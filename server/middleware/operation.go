package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/common/response"
	"github.com/jasvtfvan/oms-admin/server/model/system"
	"github.com/jasvtfvan/oms-admin/server/service"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"go.uber.org/zap"
)

var operationRecordService = service.ServiceGroupApp.System.OperationRecordService

var respPool sync.Pool
var bufferSize = 1024

func init() {
	// 当sync.Pool.Get不到可用对象，则自动执行New，点击源码可以看到该逻辑
	respPool.New = func() interface{} {
		return make([]byte, bufferSize)
	}
}

func OperationRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		v := ctx.Value("claims")
		if claims, ok := v.(*utils.CustomClaims); ok {
			if claims.LogOperation {
				userId := int(claims.BaseClaims.ID)
				handleOperationRecord(ctx, userId)
			} else {
				ctx.Next()
			}
		} else {
			response.Unauthorized(nil, "解析令牌信息失败", ctx)
			return
		}
	}
}

func handleOperationRecord(ctx *gin.Context, userId int) {
	var body []byte
	methodMap := map[string]bool{
		http.MethodPost: true,
		http.MethodPut:  true,
	}
	reqMethod := ctx.Request.Method
	if _, exists := methodMap[reqMethod]; exists {
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			global.OMS_LOG.Error("OperationRecord read request body error:", zap.Error(err))
		} else {
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}
	} else {
		query := ctx.Request.URL.RawQuery
		query, _ = url.QueryUnescape(query)
		split := strings.Split(query, "&")
		m := make(map[string]string)
		for _, v := range split {
			kv := strings.Split(v, "=")
			if len(kv) == 2 {
				m[kv[0]] = kv[1]
			}
		}
		body, _ = json.Marshal(&m)
	}
	record := system.SysOperationRecord{
		Ip:     ctx.ClientIP(),
		Method: ctx.Request.Method,
		Path:   ctx.Request.URL.Path,
		Agent:  ctx.Request.UserAgent(),
		Body:   "",
		UserID: userId,
	}

	// 上传文件时 中间件日志进行截断操作
	if strings.Contains(ctx.GetHeader("Content-Type"), "multipart/form-data") {
		record.Body = "[文件]"
	} else {
		if len(body) > bufferSize {
			record.Body = "[超出记录长度]"
		} else {
			record.Body = string(body)
		}
	}

	writer := responseBodyWriter{
		ResponseWriter: ctx.Writer,
		body:           &bytes.Buffer{},
	}
	ctx.Writer = writer
	now := time.Now() // 记录当前时间，等待response结果

	ctx.Next()

	latency := time.Since(now)
	record.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()
	record.Status = ctx.Writer.Status()
	record.Latency = latency
	record.Resp = writer.body.String()

	if strings.Contains(ctx.Writer.Header().Get("Pragma"), "public") ||
		strings.Contains(ctx.Writer.Header().Get("Expires"), "0") ||
		strings.Contains(ctx.Writer.Header().Get("Cache-Control"), "must-revalidate, post-check=0, pre-check=0") ||
		strings.Contains(ctx.Writer.Header().Get("Content-Type"), "application/force-download") ||
		strings.Contains(ctx.Writer.Header().Get("Content-Type"), "application/octet-stream") ||
		strings.Contains(ctx.Writer.Header().Get("Content-Type"), "application/vnd.ms-excel") ||
		strings.Contains(ctx.Writer.Header().Get("Content-Type"), "application/download") ||
		strings.Contains(ctx.Writer.Header().Get("Content-Disposition"), "attachment") ||
		strings.Contains(ctx.Writer.Header().Get("Content-Transfer-Encoding"), "binary") {
		if len(record.Resp) > bufferSize {
			record.Body = "超出记录长度"
		}
	}

	if err := operationRecordService.CreateSysOperationRecord(record); err != nil {
		global.OMS_LOG.Error("create operation record error:", zap.Error(err))
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

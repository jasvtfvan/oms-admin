package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jasvtfvan/oms-admin/server/global"
	"golang.org/x/sync/singleflight"
)

var (
	ErrTokenExpired     = errors.New("令牌超时")
	ErrTokenNotValidYet = errors.New("令牌尚未生效")
	ErrTokenMalformed   = errors.New("非法令牌")
	ErrTokenInvalid     = errors.New("无处理令牌")
	ErrTokenTypeError   = errors.New("令牌类型转换错误")
)

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	ID           uint
	Username     string
	LogOperation bool
	Groups       []string
	Roles        []string
}

type JWT struct {
	SigningKey        []byte
	SingleflightGroup *singleflight.Group
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.OMS_CONFIG.JWT.SigningKey),
		&singleflight.Group{},
	}
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	buf, _ := ParseDuration(global.OMS_CONFIG.JWT.BufferTime)
	exp, _ := ParseDuration(global.OMS_CONFIG.JWT.ExpiresTime)
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(buf / time.Second), // 缓冲时间转换成秒
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"OMS"},                   // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间(回拨1微秒)
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),   // 过期时间
			Issuer:    global.OMS_CONFIG.JWT.Issuer,
		},
	}
	return claims
}

// 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(j.SigningKey)
}

// 根据旧token创建新token
func (j *JWT) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
	// 多个并发操作，只有第一个操作真正执行，其他等待第一个操作结果
	v, err, _ := j.SingleflightGroup.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	if v == nil {
		return "", err
	}
	newToken, ok := v.(string)
	if !ok {
		return "", ErrTokenTypeError
	}
	return newToken, err
}

// 解析token
func (j *JWT) ParseToken(token string) (*CustomClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if jwtToken != nil {
		if claims, ok := jwtToken.Claims.(*CustomClaims); ok && jwtToken.Valid {
			return claims, nil
		}
	}
	return nil, ErrTokenInvalid
}

package response

type SysDB struct {
	Updated    bool   `json:"updated"`    // 是否已更新
	OldVersion string `json:"oldVersion"` // 老版本
	NewVersion string `json:"newVersion"` // 新版本
}

package models

import "time"

// 登录认证请求体
type AuthJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 通用响应结构体
type ResData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// 登录响应数据
type AuthData struct {
	Token string
}

// 挂载请求体
type PushData struct {
	MountPath       string `json:"mount_path"`
	Order           int    `json:"order"`
	Remark          string `json:"remark"`
	CacheExpiration int    `json:"cache_expiration"`
	WebProxy        bool   `json:"web_proxy"`
	WebdavPolicy    string `json:"webdav_policy"`
	DownProxyUrl    string `json:"down_proxy_url"`
	OrderBy         string `json:"order_by"`
	OrderDirection  string `json:"order_direction"`
	ExtractFolder   string `json:"extract_folder"`
	EnableSign      bool   `json:"enable_sign"`
	Driver          string `json:"driver"`
	Addition        string `json:"addition"`
}

// 阿里云盘挂载附加信息
type AliAddition struct {
	RefreshToken   string `json:"refresh_token"`
	ShareId        string `json:"share_id"`
	SharePwd       string `json:"share_pwd"`
	RootFolderId   string `json:"root_folder_id"`
	OrderBy        string `json:"order_by"`
	OrderDirection string `json:"order_direction"`
}

// PikPak 挂载附加信息
type PikPakAddition struct {
	RootFolderId          string `json:"root_folder_id"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	ShareId               string `json:"share_id"`
	SharePwd              string `json:"share_pwd"`
	OrderBy               string `json:"order_by"`
	OrderDirection        string `json:"order_direction"`
	Platform              string `json:"platform"`
	UseTranscodingAddress bool   `json:"use_transcoding_address"`
}

// Onedrive APP 挂载附加信息
type OnedriveAppAddition struct {
	RootFolderPath string `json:"root_folder_path"`
	Region         string `json:"region"`
	ClientId       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	TenantId       string `json:"tenant_id"`
	Email          string `json:"email"`
	ChunkSize      int    `json:"chunk_size"`
}

// 存储列表响应体
type StorageListData struct {
	Content []StorageListContent `json:"content"`
	Total   int                  `json:"total"`
}

// 存储项详细信息
type StorageListContent struct {
	Id              int       `json:"id"`
	MountPath       string    `json:"mount_path"`
	Order           int       `json:"order"`
	Driver          string    `json:"driver"`
	CacheExpiration int       `json:"cache_expiration"`
	Status          string    `json:"status"`
	Addition        string    `json:"addition"`
	Remark          string    `json:"remark"`
	Modified        time.Time `json:"modified"`
	Disabled        bool      `json:"disabled"`
	OrderBy         string    `json:"order_by"`
	OrderDirection  string    `json:"order_direction"`
	ExtractFolder   string    `json:"extract_folder"`
	WebProxy        bool      `json:"web_proxy"`
	WebdavPolicy    string    `json:"webdav_policy"`
	DownProxyURL    string    `json:"down_proxy_url"`
}

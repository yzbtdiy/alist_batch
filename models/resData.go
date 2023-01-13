package models

import "time"

type AuthJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type AuthData struct {
	Token string
}

type PushData struct {
	MountPath       string `json:"mount_path"`
	Order           int    `json:"order"`
	Remark          string `json:"remark"`
	CacheExpiration int    `json:"cache_expiration"`
	WebProxy        bool   `json:"web_proxy"`
	WebdavPolicy    string `json:"webdav_policy"`
	DownProxyUrl    string `json:"down_proxy_url"`
	ExtractFolder   string `json:"extract_folder"`
	Driver          string `json:"driver"`
	Addition        string `json:"addition"`
}

type Addition struct {
	RefreshToken   string `json:"refresh_token"`
	ShareId        string `json:"share_id"`
	SharePwd       string `json:"share_pwd"`
	RootFolderId   string `json:"root_folder_id"`
	OrderBy        string `json:"order_by"`
	OrderDirection string `json:"order_direction"`
}

type StorageListData struct {
	Content []StorageListContent `json:"content"`
	Total   int                  `json:"total"`
}

type StorageListContent struct {
	ID              int       `json:"id"`
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

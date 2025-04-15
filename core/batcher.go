package core

import (
	"encoding/json"
	"log"

	"github.com/yzbtdiy/alist_batch/models"
	"github.com/yzbtdiy/alist_batch/utils"
)

// AlistBatch 主体结构体，包含配置和API地址等
type AlistBatch struct {
	config           *models.Config
	client           *utils.HttpClient
	loginApi         string
	storageListApi   string
	addStorageApi    string
	delStorageApi    string
	updateStorageApi string
}

// NewAlistBatch 创建 AlistBatch 实例
func NewAlistBatch(config *models.Config) *AlistBatch {
	return &AlistBatch{
		config:           config,
		client:           utils.NewHttpClient(config.Token),
		loginApi:         config.Url + "/api/auth/login",
		storageListApi:   config.Url + "/api/admin/storage/list",
		addStorageApi:    config.Url + "/api/admin/storage/create",
		delStorageApi:    config.Url + "/api/admin/storage/delete",
		updateStorageApi: config.Url + "/api/admin/storage/update",
	}
}

// 检查配置文件有效性
func (a *AlistBatch) CheckConf(conf *models.Config) {
	if conf.Url == "ALIST_URL" {
		panic("url 未配置, 请检查配置文件")
	}
	if (conf.Auth.Username == "USERNAME" || conf.Auth.Password == "PASSWORD") && (conf.Token == "ALIST_TOKEN" || conf.Token == "") {
		panic("token和用户密码至少要配置一项, 请检查配置文件")
	}
	if conf.Aliyun.Enable && conf.Aliyun.RefreshToken == "ALI_YUNPAN_REFRESH_TOKEN" {
		panic("添加阿里云盘链接需要 refresh_token, 请检查配置文件")
	}
	if conf.OneDriveApp.Enable && (conf.OneDriveApp.Tenants[0].ClientId == "CLIENT_ID" || conf.OneDriveApp.Tenants[0].ClientSecret == "CLIENT_SECRET" || conf.OneDriveApp.Tenants[0].TenantId == "TENANT_ID") {
		panic("添加 OnedriveApp 需要添加 client_id, client_secret 和 tenant_id, 请检查配置文件")
	}
}

// 校验 token 是否有效
func (a *AlistBatch) VaildToken() bool {
	storageListRes := a.client.Get(a.storageListApi)
	if storageListRes.Code == 200 {
		return true
	} else {
		return false
	}
}

// 更新 token
func (a *AlistBatch) UpdateToken() {
	loginData := models.AuthJson{
		Username: a.config.Auth.Username,
		Password: a.config.Auth.Password,
	}
	authJson, _ := json.Marshal(loginData)

	resData := a.client.Post(a.loginApi, authJson)
	data, _ := json.Marshal(resData.Data)
	var tokenData models.AuthData
	json.Unmarshal(data, &tokenData)
	if resData.Code == 200 {
		utils.ModConfig("config.yaml", a.config, tokenData.Token)
		log.Println("token 已更新, 请重新运行此程序")
	} else {
		panic("token 更新失败, 请检查用户密码是否正确")
	}
}

// 释放资源
func (a *AlistBatch) Close() {
	a.client.Close()
}

package storage

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"

	"github.com/yzbtdiy/alist_batch/models"
)

// BuildAliPushData 生成添加阿里云盘分享链接挂载的JSON数据
// mountPath: 挂载路径
// aliUrl: 阿里云盘分享链接
// config: 配置信息
func BuildAliPushData(mountPath string, aliUrl string, config *models.Config) []byte {
	// 解析分享链接
	params, err := url.Parse(aliUrl)
	if err != nil {
		log.Fatalf("解析阿里云盘分享链接失败: %v", err)
	}

	// 提取分享密码
	sharePwd := ""
	if params.Query().Has("pwd") {
		sharePwd = params.Query().Get("pwd")
	}

	// 提取分享ID和文件夹ID
	pathArray := strings.Split(params.Path, "/")
	if len(pathArray) < 5 {
		log.Fatalf("无效的阿里云盘分享链接格式: %s", aliUrl)
	}

	shareId := pathArray[2]
	shareFolder := pathArray[4]

	// 构建附加信息
	addition := &models.AliAddition{
		RefreshToken:   config.Aliyun.RefreshToken,
		ShareId:        shareId,
		SharePwd:       sharePwd,
		RootFolderId:   shareFolder,
		OrderBy:        "",
		OrderDirection: "",
	}

	additionJson, err := json.Marshal(addition)
	if err != nil {
		log.Fatalf("序列化阿里云盘附加信息失败: %v", err)
	}
	additionData := string(additionJson)

	// 构建完整的挂载数据
	data := models.PushData{
		MountPath:       mountPath,
		Order:           0,
		Remark:          "",
		CacheExpiration: 30,
		WebProxy:        false,
		WebdavPolicy:    "302_redirect",
		DownProxyUrl:    "",
		OrderBy:         "",
		OrderDirection:  "",
		ExtractFolder:   "",
		EnableSign:      false,
		Driver:          "AliyundriveShare",
		Addition:        additionData,
	}

	pushJson, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("序列化阿里云盘挂载数据失败: %v", err)
	}
	return pushJson
}

// BuildUpdateAliRefreshToken 生成阿里云盘分享批量更新RefreshToken的JSON数据
// aliShareData: 存储列表内容
// refreshToken: 新的刷新令牌
func BuildUpdateAliRefreshToken(aliShareData models.StorageListContent, refreshToken string) []byte {
	var oldAddition models.AliAddition
	err := json.Unmarshal([]byte(aliShareData.Addition), &oldAddition)
	if err != nil {
		log.Fatalf("解析阿里云盘存储附加信息失败: %v", err)
	}

	oldAddition.RefreshToken = refreshToken
	aliShareData.Status = "work"

	newAddition, err := json.Marshal(oldAddition)
	if err != nil {
		log.Fatalf("序列化新的阿里云盘附加信息失败: %v", err)
	}

	aliShareData.Addition = string(newAddition)

	pushJson, err := json.Marshal(aliShareData)
	if err != nil {
		log.Fatalf("序列化更新数据失败: %v", err)
	}

	return pushJson
}

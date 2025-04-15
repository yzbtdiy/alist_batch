package storage

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/yzbtdiy/alist_batch/models"
)

// BuildOnedriverApp 生成添加OneDrive应用挂载的JSON数据
// mountPath: 挂载路径
// emailInfo: 格式为 "tid:email:path" 或 "tid:email"，其中tid为租户ID索引，email为账户邮箱，path为文件夹路径
// config: 配置信息
func BuildOnedriverApp(mountPath string, emailInfo string, config *models.Config) []byte {
	params := strings.Split(emailInfo, ":")

	// 验证参数格式
	if len(params) < 2 || len(params) > 3 {
		log.Fatalf("Invalid emailInfo format: %s. Expected format: tid:email[:path]", emailInfo)
	}

	// 解析租户ID索引
	tid, err := strconv.Atoi(params[0])
	if err != nil {
		log.Fatalf("Invalid tenant ID index: %s, error: %v", params[0], err)
	}

	// 确保索引在有效范围内
	if tid < 1 || tid > len(config.OneDriveApp.Tenants) {
		log.Fatalf("Tenant ID index out of range: %d, available range: 1-%d", tid, len(config.OneDriveApp.Tenants))
	}

	email := params[1]
	folderPath := "/"

	// 如果提供了文件夹路径，则使用提供的路径
	if len(params) == 3 {
		folderPath = params[2]
	}

	// 构建附加信息
	addition := models.OnedriveAppAddition{
		RootFolderPath: folderPath,
		Region:         config.OneDriveApp.Region,
		ClientId:       config.OneDriveApp.Tenants[tid-1].ClientId,
		ClientSecret:   config.OneDriveApp.Tenants[tid-1].ClientSecret,
		TenantId:       config.OneDriveApp.Tenants[tid-1].TenantId,
		Email:          email,
		ChunkSize:      5,
	}

	additionJson, err := json.Marshal(addition)
	if err != nil {
		log.Fatalf("Failed to marshal OneDrive APP addition: %v", err)
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
		Driver:          "OnedriveAPP",
		Addition:        additionData,
	}

	pushJson, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal push data: %v", err)
	}

	return pushJson
}

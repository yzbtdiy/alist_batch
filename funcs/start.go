package funcs

import (
	"github.com/yzbtdiy/alist_batch/models"

	"encoding/json"
	"log"
	"sync"
)

func Run() {
	confStat := CheckFile("config.yaml")
	shareStat := CheckFile("ali_share.yaml")
	if confStat {
		if shareStat {
			Start()
		} else {
			log.Println("ali_share.yaml文件不存在, 尝试生成")
			GenResFile("ali_share.yaml")
		}
	} else if shareStat {
		log.Println("config.yaml文件不存在, 尝试生成")
		GenConfFile("config.yaml")
	} else {
		log.Println("config.yaml不存在, 尝试生成")
		GenConfFile("config.yaml")
		log.Println("ali_share.yaml文件不存在, 尝试生成")
		GenResFile("ali_share.yaml")
	}
}

func Start() {
	// 读取配置文件
	conf := GetConfig("config.yaml")

	// 检查配置文件是否修改必要字段
	CheckConf(conf)

	// 拼接API
	loginApi := conf.Url + "/api/auth/login"
	storageListApi := conf.Url + "/api/admin/storage/list"
	addStorageApi := conf.Url + "/api/admin/storage/create"
	delStorageApi := conf.Url + "/api/admin/storage/delete"

	// 将用户名和密码转为json
	loginData := models.AuthJson{
		Username: conf.Auth.Username,
		Password: conf.Auth.Password,
	}
	authJson, _ := json.Marshal(loginData)

	// 检测 token 是否存在
	if conf.Token != "ALIST_TOKEN" && conf.Token != "" {
		// 携带 token 尝试读取storagelist, 若返回 200 则说明 token 有效
		storageListRes := HttpGet(storageListApi, conf.Token)
		if storageListRes.Code == 200 {
			// 如果有 -delete flag, 进行删除操作
			DeleteStorageIfHaveFlag(storageListApi, delStorageApi, conf.Token)
			// 读取阿里云盘资源链接列表
			shareListData := GetShareList("./ali_share.yaml")
			// 使用 gorouting, 感谢 nzlov: https://github.com/nzlov
			wg := &sync.WaitGroup{}
			//  遍历阿里云盘资源名和链接
			for category, shareList := range shareListData {
				for shareName, shareUrl := range shareList {
					wg.Add(1)
					go func(category, shareName, shareUrl string) {
						defer wg.Done()
						// 根据阿里云盘资源名和链接生成添加资源的 json 字符串
						pushData := BuildPushData(`/`+category+`/`+shareName, shareUrl, conf)
						// 发送请求添加资源
						pushRes := HttpPost(addStorageApi, conf.Token, pushData)
						// 返回值为 200 说明添加成功
						if pushRes.Code == 200 {
							log.Println(category + " " + shareName + " 添加完成")
						} else {
							log.Println(category + " " + shareName + " 添加失败, 请检查是否重复添加")
						}
					}(category, shareName, shareUrl)
				}
			}
			wg.Wait()
		} else {
			// 若携带 token 尝试访问 storage list 失败, 则尝试更新 token
			UpdateToken(loginApi, authJson, conf)
		}
	} else {
		// 若 token 为 ALIST_TOKEN 或空字符串, 则尝试更新 token
		UpdateToken(loginApi, authJson, conf)
	}
}

// 检查配置文件是否修改
func CheckConf(conf *models.Config) {
	if conf.Url == "ALIST_URL" {
		panic("url 未配置, 请检查配置文件")
	}
	if (conf.Auth.Username == "USERNAME" || conf.Auth.Password == "PASSWORD") && (conf.Token == "ALIST_TOKEN" || conf.Token == "") {
		panic("token和用户密码至少要配置一项, 请检查配置文件")
	}
	if conf.RefreshToken == "ALI_YUNPAN_REFRESH_TOKEN" {
		panic("refresh_token 未配置, 请检查配置文件")
	}
}

// 更新 token
func UpdateToken(loginApi string, authJson []byte, conf *models.Config) {
	resData := HttpPost(loginApi, "", authJson)
	data, _ := json.Marshal(resData.Data)
	var tokenData models.AuthData
	json.Unmarshal(data, &tokenData)
	// log.Println(tokenData.Token)
	if resData.Code == 200 {
		ModConfig("config.yaml", conf, tokenData.Token)
		log.Println("token 已更新, 请重新运行此程序")
	} else {
		panic("token 更新失败, 请检查用户密码是否正确")
	}
}

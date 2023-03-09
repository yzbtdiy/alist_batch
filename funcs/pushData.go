package funcs

import (
	"log"
	"strings"
	"sync"

	"github.com/yzbtdiy/alist_batch/models"
)

// 读取 ali_share.yaml 文件添加阿里云盘分享链接
func PushAliShares(addStorageApi string, conf *models.Config) {
	shareListData := GetShareList("./ali_share.yaml")
	// 使用 gorouting, 感谢 nzlov: https://github.com/nzlov
	wg := &sync.WaitGroup{}
	// 遍历阿里云盘资源名和链接
	for category, shareList := range shareListData {
		for shareName, shareUrl := range shareList {
			wg.Add(1)
			go func(category, shareName, shareUrl string) {
				defer wg.Done()
				// 根据阿里云盘资源名和链接生成添加资源的 json 字符串
				pushData := BuildAliPushData(`/`+category+`/`+shareName, shareUrl, conf)
				// 发送请求添加资源
				pushRes := HttpPost(addStorageApi, conf.Token, pushData)
				// 返回值为 200 说明添加成功
				if pushRes.Code == 200 {
					log.Println(category + " " + shareName + " 添加完成")
				} else if pushRes.Code == 500 {
					if strings.Split(pushRes.Message, ":")[2] == " failed to refresh token" {
						log.Println("refresh token 无效, 已尝试添加 " + category + " " + shareName)
					} else {
						log.Println("ssss" + strings.Split(pushRes.Message, ":")[2] + "sss")
						log.Println(category + " " + shareName + "添加失败, 请检查是否重复添加")
					}
				}
			}(category, shareName, shareUrl)
		}
	}
	wg.Wait()
}

// 读取 pik_share.yaml 文件添加 PikPak 分享链接
func PushPikPakShares(addStorageApi string, conf *models.Config) {
	shareListData := GetShareList("./pik_share.yaml")
	wg := &sync.WaitGroup{}
	for category, shareList := range shareListData {
		for shareName, shareUrl := range shareList {
			wg.Add(1)
			go func(category, shareName, shareUrl string) {
				defer wg.Done()
				pushData := BuildPikPakData(`/`+category+`/`+shareName, shareUrl, conf)
				pushRes := HttpPost(addStorageApi, conf.Token, pushData)
				if pushRes.Code == 200 {
					log.Println(category + " " + shareName + " 添加完成")
				} else {
					log.Println(category + " " + shareName + " 添加失败, 请检查是否重复添加")
				}
			}(category, shareName, shareUrl)
		}
	}
	wg.Wait()
}

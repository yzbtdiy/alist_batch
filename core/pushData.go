package core

import (
	"log"
	"strings"
	"sync"

	"github.com/yzbtdiy/alist_batch/utils"
)

// 批量添加阿里云盘分享链接
func (a *AlistBatch) PushAliShares() {
	shareListData := utils.GetShareList("./ali_share.yaml")
	wg := &sync.WaitGroup{}
	for category, shareList := range shareListData {
		for shareName, shareUrl := range shareList {
			wg.Add(1)
			go func(category, shareName, shareUrl string) {
				defer wg.Done()
				pushData := a.BuildAliPushData(`/`+category+`/`+shareName, shareUrl)
				pushRes := a.client.Post(a.addStorageApi, pushData)
				if pushRes.Code == 200 {
					log.Println(category + " " + shareName + " 添加完成")
				} else if pushRes.Code == 500 {
					if strings.Split(pushRes.Message, ":")[2] == " failed to refresh token" {
						log.Println("refresh token 无效, 已尝试添加 " + category + " " + shareName)
					} else {
						log.Println(category + " " + shareName + "添加失败, 请检查是否重复添加")
					}
				}
			}(category, shareName, shareUrl)
		}
	}
	wg.Wait()
}

// 批量添加 PikPak 分享链接
func (a *AlistBatch) PushPikPakShares() {
	shareListData := utils.GetShareList("./pik_share.yaml")
	wg := &sync.WaitGroup{}
	for category, shareList := range shareListData {
		for shareName, shareUrl := range shareList {
			wg.Add(1)
			go func(category, shareName, shareUrl string) {
				defer wg.Done()
				pushData := a.BuildPikPakData(`/`+category+`/`+shareName, shareUrl)
				pushRes := a.client.Post(a.addStorageApi, pushData)
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

// 批量添加 Onedrive APP 分享链接
func (a *AlistBatch) PushOnedriveApp() {
	onedriveAppList := utils.GetShareList("./onedrive_app.yaml")
	wg := &sync.WaitGroup{}
	for category, shareList := range onedriveAppList {
		for shareName, shareUrl := range shareList {
			wg.Add(1)
			go func(category, shareName, shareUrl string) {
				defer wg.Done()
				pushData := a.BuildOnedriverApp(`/`+category+`/`+shareName, shareUrl)
				pushRes := a.client.Post(a.addStorageApi, pushData)
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

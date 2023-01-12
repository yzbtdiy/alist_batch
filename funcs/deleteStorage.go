package funcs

import (
	"os"

	"github.com/yzbtdiy/alist_batch/models"

	"encoding/json"
	"flag"
	"log"
	"strconv"
)

var deleteFlag = flag.String("delete", "",
	`dis    删除已禁用存储
all    删除所有存储(慎用)`)

func DeleteStorageIfHaveFlag(storageListApi, delStorageApi, token string) {
	flag.Parse()
	if *deleteFlag != "" {
		storageData := getStorgeData(storageListApi, token)
		needDeleteIds := GetDisableStroageIds(storageData)
		DeleteStorageById(delStorageApi, needDeleteIds, token)
		os.Exit(0)
	} 
}

func GetDisableStroageIds(storageListData *models.StorageListData) []string {
	var deleteIds []string
	if *deleteFlag == "dis" {
		println("尝试删除禁用存储")
		for _, mountInfo := range storageListData.Content {
			if mountInfo.Disabled {
				deleteIds = append(deleteIds, strconv.Itoa(mountInfo.ID))
			}
		}
	} else if *deleteFlag == "all" {
		println("尝试删除所有存储")
		for _, mountInfo := range storageListData.Content {
			deleteIds = append(deleteIds, strconv.Itoa(mountInfo.ID))
		}
	}
	return deleteIds
}

func getStorgeData(storageListApi string, token string) *models.StorageListData {
	storageListRes := HttpGet(storageListApi, token)
	data, _ := json.Marshal(storageListRes.Data)
	var storageData *models.StorageListData
	json.Unmarshal(data, &storageData)
	return storageData
}

func DeleteStorageById(delStorageApi string, deleteIds []string, token string) {
	for _, id := range deleteIds {
		resData := HttpPost(delStorageApi+"?id="+id, token, []byte(""))
		if resData.Code == 200 {
			log.Println("Id为 " + id + " 的存储已删除")
		}
	}
}

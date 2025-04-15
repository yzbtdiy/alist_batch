package main

import (
	"flag"
	"log"

	"github.com/yzbtdiy/alist_batch/core"
	"github.com/yzbtdiy/alist_batch/flags"
	"github.com/yzbtdiy/alist_batch/models"
	"github.com/yzbtdiy/alist_batch/utils"
)

var conf *models.Config

func main() {
	// 读取配置文件
	conf = utils.GetConfig("config.yaml")

	// 检查配置文件是否存在
	hasConf := utils.CheckFile("config.yaml")
	if !hasConf {
		log.Println("config.yaml文件不存在, 尝试生成")
		utils.GenConfFile("config.yaml")
		return
	}

	// 检查各分享文件是否存在，不存在则生成
	if conf.Aliyun.Enable {
		if !utils.CheckFile("ali_share.yaml") {
			log.Println("ali_share.yaml文件不存在, 尝试生成")
			utils.GenAliShareFile("ali_share.yaml")
			return
		}
	}
	if conf.PikPak.Enable {
		if !utils.CheckFile("pik_share.yaml") {
			log.Println("pik_share.yaml文件不存在, 尝试生成")
			utils.GenPikShareFile("pik_share.yaml")
			return
		}
	}
	if conf.OneDriveApp.Enable {
		if !utils.CheckFile("onedrive_app.yaml") {
			log.Println("onedrive_app.yaml文件不存在, 尝试生成")
			utils.GenOnedriveAppFile("onedrive_app.yaml")
			return
		}
	}

	// 初始化批量处理器
	batcher := core.NewAlistBatch(conf)
	batcher.CheckConf(conf)

	// 检查 token 并执行批量操作
	if conf.Token != "ALIST_TOKEN" && conf.Token != "" {
		if batcher.VaildToken() {
			flag.Parse()
			if *flags.UpdateFlag != "" {
				batcher.UpdateAliStorage()
			}
			if *flags.DeleteFlag != "" {
				batcher.DeleteStorageIfHaveFlag()
			}
			if conf.Aliyun.Enable {
				batcher.PushAliShares()
			}
			if conf.PikPak.Enable {
				batcher.PushPikPakShares()
			}
			if conf.OneDriveApp.Enable {
				batcher.PushOnedriveApp()
			}
			return
		}
		batcher.UpdateToken()
	} else {
		batcher.UpdateToken()
	}
}

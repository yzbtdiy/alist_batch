package funcs

import (
	"github.com/yzbtdiy/alist_batch/models"

	"os"

	"gopkg.in/yaml.v3"
)

// 生成配置模板文件
func GenConfFile(fileName string) {
	confAuth := models.Auth{
		Username: "USERNAME",
		Password: "PASSWORD",
	}

	confAliyun := models.Aliyun{
		Enable:       true,
		RefreshToken: "ALI_YUNPAN_REFRESH_TOKEN",
	}

	confPikPak := models.PikPak{
		Enable:   true,
		Username: "PIKPAK_EMAIL",
		Password: "PIKPAK_PASSWORD",
	}

	confExample := models.Config{
		Url:    "ALIST_URL",
		Auth:   &confAuth,
		Token:  "ALIST_TOKEN",
		Aliyun: &confAliyun,
		PikPak: &confPikPak,
	}

	res, err := yaml.Marshal(confExample)
	if err != nil {
		panic(err)
	}
	os.WriteFile("./"+fileName, res, 0777)
}

// 生成阿里云盘分享链接模板文件
func GenAliShareFile(fileName string) {
	resTv := models.TvSeries{
		XiYouJi:      "https://www.aliyundrive.com/s/MmMR3zaoXLf/folder/61d259418d27bae8656f47aca23ee03b40275432",
		ShiZongZui:   "https://www.aliyundrive.com/s/hLzid1Tty6o/folder/62230dcf0c7bcec037ba4b8b80847fad1267a17a",
		FaYiQingMing: "https://www.aliyundrive.com/s/gJjg9HJtYcR/folder/61519615d363e70740ee4333a8ab1cfc0def79af",
	}
	resMv := models.Movies{
		ManWei:               "https://www.aliyundrive.com/s/rMAvoXgU6Uh/folder/621b817b0f64fa3fb67e4630ac28267a0cdc482a",
		XinHaiChengGongQiJun: "https://www.aliyundrive.com/s/FzcMCgG8YwC/folder/61ffb364be026f8c1b764182922eaeb2d3950ef4",
		LinZhengYing:         "https://www.aliyundrive.com/s/PrcaqZ2XPxM/folder/621c950a633c7c7ab8de4db1a86a1232dea530d1",
	}
	resExample := models.AliShareList{
		TvSeries: &resTv,
		Movies:   &resMv,
	}
	res, err := yaml.Marshal(resExample)
	if err != nil {
		panic(err)
	}
	os.WriteFile("./"+fileName, res, 0777)
}

// 生成 PikPak 分享链接模板文件
func GenPikShareFile(fileName string) {
	shareList := make(map[string]map[string]string)
	subList := make(map[string]string)
	subList["太空之城"] = "https://mypikpak.com/s/VNP2_7OhUCdC2aI3JSSnD--eo1"
	subList["阿飞正传"] = "https://mypikpak.com/s/VNP2d8tHvt4TVPKPacCUYRaXo1/VNP2G0YUcYmtVw025fNVqgDdo1"
	shareList["电影"] = subList
	res, err := yaml.Marshal(shareList)
	if err != nil {
		panic(err)
	}
	os.WriteFile("./"+fileName, res, 0777)
}

package flags

import "flag"

var DeleteFlag = flag.String("delete", "",
	`dis    删除已禁用存储
all    删除所有存储(慎用)`)

var UpdateFlag = flag.String("update", "",
	`ali    更新阿里refresh_token`)

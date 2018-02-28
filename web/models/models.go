package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(ClientConf))
	orm.RegisterModel(new(IpsecServerHistory))
	orm.RegisterModel(new(NoticeHistory))
}

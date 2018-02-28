package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	IPSEC_SERVER_START = iota
	IPSEC_SERVER_ONLINE
	IPSEC_SERVER_STOP
)

type IpsecServerHistory struct {
	Id    int
	Stamp int64
	State int
	Time  time.Time `orm:"-"`
}

func (this *IpsecServerHistory) AddHistory() bool {
	o := orm.NewOrm()

	his := IpsecServerHistory{}
	his.Stamp = this.Stamp
	his.State = this.State
	_, succ := o.Insert(&his)
	if succ != nil {
		return false
	}
	return true
}

func GetHistoryData(page int) []IpsecServerHistory {
	cur := page - 1
	if cur < 0 {
		cur = 0
	}
	o := orm.NewOrm()
	var data []IpsecServerHistory
	o.QueryTable("IpsecServerHistory").OrderBy("-stamp").Limit(20, 20*cur).All(&data)
	return data
}

func GetHistoryDataCount() int64 {
	var count int64 = 0
	o := orm.NewOrm()
	count, err := o.QueryTable("ipsec_server_history").Count()
	if err != nil {
		return 0
	}
	if count != 0 {
		return count
	} else {
		return 0
	}
}

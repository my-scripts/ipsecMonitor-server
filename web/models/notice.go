package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type NoticeHistory struct {
	Id    int
	Alias string
	Stamp int64
	State int
	Time  time.Time `orm:"-"`
}

func (this *NoticeHistory) AddHistory() bool {
	o := orm.NewOrm()

	his := NoticeHistory{}
	his.Alias = this.Alias
	his.Stamp = this.Stamp
	his.State = this.State
	_, succ := o.Insert(&his)
	if succ != nil {
		return false
	}
	return true
}

func GetNoticeHistoryData(page int) []NoticeHistory {
	cur := page - 1
	if cur < 0 {
		cur = 0
	}
	o := orm.NewOrm()
	var data []NoticeHistory
	o.QueryTable("NoticeHistory").OrderBy("-stamp").Limit(20, 20*cur).All(&data)
	return data
}

func GetNoticeHistoryDataCount() int64 {
	var count int64 = 0
	o := orm.NewOrm()
	count, err := o.QueryTable("NoticeHistory").Count()
	if err != nil {
		return 0
	}
	if count != 0 {
		return count
	} else {
		return 0
	}
}

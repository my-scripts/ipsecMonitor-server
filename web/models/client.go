package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ClientConf struct {
	Id    int
	Alias string `form:"alias"`
	Addr  string `form:"addr"`
}

func GetClients() []ClientConf {
	o := orm.NewOrm()
	var clients []ClientConf
	_, err := o.QueryTable("ClientConf").All(&clients)
	if err != nil {
		return nil
	}
	return clients
}

func (this *ClientConf) GetClient() bool {
	o := orm.NewOrm()
	client := ClientConf{Id: this.Id}
	err := o.Read(&client)
	if err != nil {
		return false
	} else {
		this.Id = client.Id
		this.Alias = client.Alias
		this.Addr = client.Addr
		return true
	}
	return false
}

func (this *ClientConf) SetClient() bool {
	o := orm.NewOrm()

	client := ClientConf{Id: this.Id}

	err := o.Read(&client)
	client.Id = this.Id
	client.Addr = this.Addr
	client.Alias = this.Alias
	beego.Warning(err)
	if err == orm.ErrNoRows {
		_, succ := o.Insert(&client)
		beego.Warning(succ)
		if succ != nil {
			return false
		}
		return true
	} else {
		_, succ := o.Update(&client)
		if succ != nil {
			return false
		}
		return true
	}
	return false
}

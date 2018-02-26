package models

type ClientConf struct {
	Id   int
	Addr string `form:"addr"`
}

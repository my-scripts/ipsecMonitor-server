package models

const (
	IPSEC_SERVER_ONLINE = iota
	IPSEC_SERVER_OFFLINE
)

type IpsecServerHistory struct {
	Id    int
	Stamp int64
	State int
}

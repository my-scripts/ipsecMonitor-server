package models

const (
	IPSEC_SERVER_START = iota
	IPSEC_SERVER_ONLINE
	IPSEC_SERVER_STOP
)

type IpsecServerHistory struct {
	Id    int
	Stamp int64
	State int
}

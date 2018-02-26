package client

import (
	"log"
)

type IpsecRpcClient struct {
	RpcClient
}

type Status struct {
	Status bool
}

func (this *IpsecRpcClient) Connect(host string, port int) bool {
	return this.ConnectRemote(host, port)
}

func (this *IpsecRpcClient) GetNodeHaStatus() *Status {
	replay := Status{}
	err := this.Handle.Call("Handler.RestartIpsec", 0, &replay)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &replay
}

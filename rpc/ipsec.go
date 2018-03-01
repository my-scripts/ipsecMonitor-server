package client

import (
	"log"
)

type IpsecRpcClient struct {
	RpcClient
}

type Status struct {
	Succ bool
}

type Input struct {
	Stamp int64
}

func (this *IpsecRpcClient) Connect(host string, port int) bool {
	return this.ConnectRemote(host, port)
}

func (this *IpsecRpcClient) RestartIpsec(stamp int64) *Status {
	replay := Status{}
	req := Input{Stamp: stamp}
	err := this.Handle.Call("Handler.RestartIpsec", &req, &replay)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &replay
}

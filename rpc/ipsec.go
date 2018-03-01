package client

import (
	"log"
	clientrpc "script/ipsecMonitor/client/rpc"
)

type IpsecRpcClient struct {
	RpcClient
}

func (this *IpsecRpcClient) Connect(host string, port int) bool {
	return this.ConnectRemote(host, port)
}

func (this *IpsecRpcClient) RestartIpsec(stamp int64) *clientrpc.Status {
	replay := clientrpc.Status{}
	args := clientrpc.Args{Stamp: stamp}
	err := this.Handle.Call("Handler.RestartIpsec", &args, &replay)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &replay
}

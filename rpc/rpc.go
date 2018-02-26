package client

import (
	"fmt"
	"log"
	"net/rpc"
)

type RpcClient struct {
	Handle *rpc.Client
}

func (this *RpcClient) ConnectRemote(host string, port int) bool {
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Println("failed to connect agent,", err)
		return false
	}

	this.Handle = client
	return true
}

func (this *RpcClient) Close() {
	if this.Handle != nil {
		this.Handle.Close()
		this.Handle = nil
	}
}

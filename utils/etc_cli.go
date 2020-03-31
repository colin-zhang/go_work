package utils

import (
	"github.com/coreos/etcd/clientv3"
)

type EtcdCli struct {
	clientv3.Client
}
// TODO: service discovery, distributed locker
// https://github.com/mistaker/etcdTool

func NewEtcdCli() *EtcdCli {
	clientv3.New(clientv3.Config{
		Endpoints:            nil,
		AutoSyncInterval:     0,
		DialTimeout:          0,
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
		MaxCallSendMsgSize:   0,
		MaxCallRecvMsgSize:   0,
		TLS:                  nil,
		Username:             "",
		Password:             "",
		RejectOldCluster:     false,
		DialOptions:          nil,
		LogConfig:            nil,
		Context:              nil,
		PermitWithoutStream:  false,
	})
}
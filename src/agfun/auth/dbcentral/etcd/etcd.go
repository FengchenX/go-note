package etcd

import (
	"agfun/agfun-service/dbcentral/etcddb"
)

type Client struct {
	*etcddb.Client
}

func NewClient(cli *etcddb.Client) *Client {
	client := Client{
		Client: cli,
	}
	return &client
}

var stdCli *Client

func Init() {
	etcddb.Init()
	stdCli = NewClient(etcddb.GetCli())
}

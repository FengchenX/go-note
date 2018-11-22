package etcd

import "agfun/agfun-service/dbcentral/etcddb"

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

func initCli() {
	if stdCli != nil {
		return
	}
	stdCli = NewClient(etcddb.GetCli())
}
func GetDefaultCli() *Client {
	if stdCli == nil {
		initCli()
	}
	return stdCli
}

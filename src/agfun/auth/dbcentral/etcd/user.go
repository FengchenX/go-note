package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
)

func (cli *Client) PutAccessToken(accessToken string, id uint) error {
	response, e := cli.Grant(context.TODO(), 10000)
	if e != nil {
		return e
	}
	e = cli.Put(accessToken, id, clientv3.WithLease(response.ID))
	if e != nil {
		return e
	}
	return nil
}

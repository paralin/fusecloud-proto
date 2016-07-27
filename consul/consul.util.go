package consul

import (
	"github.com/hashicorp/consul/api"
	"github.com/synrobo/proto/common"
)

func BuildConsulNode(apinode *api.Node) *ConsulNode {
	return &ConsulNode{
		Name:    apinode.Node,
		Address: common.FromIPString(apinode.Address),
	}
}

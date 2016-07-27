package consul

import (
	"github.com/hashicorp/consul/api"
	"github.com/synrobo/proto/common"
)

func ToConsulNode(apinode *api.Node) *ConsulNode {
	return &ConsulNode{
		Name:    apinode.Node,
		Address: common.FromIPString(apinode.Address),
	}
}

func ToConsulQueryMeta(apim *api.QueryMeta) *ConsulQueryMeta {
	return &ConsulQueryMeta{
		LastIndex: apim.LastIndex,
		// Convert to milliseconds
		LastContact: apim.LastContact.Nanoseconds() / 1000000,
		KnownLeader: apim.KnownLeader,
		RequestTime: apim.RequestTime.Nanoseconds() / 1000000,
	}
}

func ToConsulHealthCheck(apic *api.HealthCheck) *ConsulHealthCheck {
	return &ConsulHealthCheck{
		Node:        apic.Node,
		CheckId:     apic.CheckID,
		Name:        apic.Name,
		Status:      apic.Status,
		Notes:       apic.Notes,
		Output:      apic.Output,
		ServiceId:   apic.ServiceID,
		ServiceName: apic.ServiceName,
	}
}

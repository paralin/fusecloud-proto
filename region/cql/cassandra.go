package region_cql

import (
	"fmt"

	"github.com/fuserobotics/cassandra"
	"github.com/fuserobotics/proto/region"
)

func BuildCassandraContextForMetaRegion(reg *region.Region, cassandraIp string) (*cassandra.CassandraContext, error) {
	return cassandra.NewCassandraContextFromUri(fmt.Sprintf("cassandra://%s/meta_%s?protocol=4", cassandraIp, reg.Id))
}

package graphql

type GraphQLProto interface {
	GraphQLTypeName() string
}

type GraphQLResult struct {
	GraphQLProto

	TypeName string `json:"__typename,omitempty"`
}

func BuildGraphQLResult(proto GraphQLProto) GraphQLProto {
	if _, ok := proto.(*GraphQLResult); ok {
		return proto
	}
	return &GraphQLResult{
		GraphQLProto: proto,
		TypeName:     proto.GraphQLTypeName(),
	}
}

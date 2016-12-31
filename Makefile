PACKAGE_PATH="github.com/fuserobotics/proto/pkg/"

gengo: protogen rootgen gqlgen install-go

rootgen:
	./scripts/gen_root.bash

gqlgen:
	./scripts/gen_gqlmap.bash

protogen:
	protowrap -I $${GOPATH}/src \
		--go_out=plugins=grpc:$${GOPATH}/src \
		--proto_path $${GOPATH}/src \
		--print_structure \
		--only_specified_files \
		$$(pwd)/**/*.proto

deps:
	go get -u github.com/square/goprotowrap/cmd/protowrap

install-go:
	go install ./...

PACKAGE_PATH="github.com/fuserobotics/proto/pkg/"

gengo: protogen install-go

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
	for D in */; do go install -v github.com/fuserobotics/proto/$$D 2>/dev/null || true ; done

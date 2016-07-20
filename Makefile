PACKAGE_PATH="github.com/synrobo/proto/pkg/"
COMPILE_ROLES_DIR="./_resources/compile-roles"

gengo: compile-roles protogen
	./compile-roles ./roles/roles.compiled.go
	go install -v ./roles/

protogen:
	protowrap -I $${GOPATH}/src \
		--gogo_out=plugins=grpc:$${GOPATH}/src \
		--proto_path $${GOPATH}/src \
		--print_structure \
		--only_specified_files \
		$$(pwd)/**/*.proto

deps:
	go get -u github.com/square/goprotowrap/cmd/protowrap

compile-roles:
	go build -v $(COMPILE_ROLES_DIR)

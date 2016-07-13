PACKAGE_PATH="github.com/synrobo/proto/pkg/"

gengo:
	protowrap -I $${GOPATH}/src \
		--gogo_out=$${GOPATH}/src \
		--proto_path $${GOPATH}/src \
		--print_structure \
		--only_specified_files \
		$$(pwd)/**/*.proto

deps:
	go get -u github.com/square/goprotowrap/cmd/protowrap


PACKAGE_PATH="github.com/synrobo/proto/pkg/"

deps:
	go get -u github.com/square/goprotowrap/cmd/protowrap

gengo:
	protowrap -I $$(pwd)/src/ \
		--proto_path $$(pwd)/src \
		--gogo_out=import_prefix=$(PACKAGE_PATH):$$(pwd)/pkg/ \
		--print_structure \
		$$(pwd)/src/**/*.proto

	# Fix the import path for gogo
	sed -i -e "s#github.com/synrobo/proto/pkg/github.com/gogo/protobuf#github.com/gogo/protobuf#g" ./pkg/**/*.pb.go

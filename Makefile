gengo:
	protoc -I=$$(pwd)/src/ \
		--go_out=$$(pwd)/pkg/auth/ \
		$$(pwd)/src/auth.proto

#!/bin/bash

# Build array of auth, common, device, region...
PACKAGES=( $(find -name '*.proto' -printf '%h\n' | sort -u | cut -d/ -f2 | xargs) )

# Quote them
PACKAGES=("${PACKAGES[@]/#/\"}")
PACKAGES=("${PACKAGES[@]/%/\",}")
PACKAGES_STR=$(printf '\t%s\n' "${PACKAGES[@]}")


cat<<EOF >root.go
package proto

// List of packages with protobuf types.
var ProtoPackages []string = []string{
${PACKAGES_STR}
}
EOF

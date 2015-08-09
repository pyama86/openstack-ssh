#!/bin/bash

pkg_path="$GOPATH/src/github.com/pyama86/openstack-ssh/pkg"
XC_ARCH=${XC_ARCH:-386 amd64}
XC_OS=${XC_OS:-linux}
cd "$GOPATH/src/github.com/pyama86/openstack-ssh/cmd/openstack-ssh"
rm -rf pkg/
gox \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -output "$pkg_path/{{.OS}}_{{.Arch}}/{{.Dir}}"

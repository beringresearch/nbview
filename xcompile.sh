#!/bin/bash
set -e
# you may need to go install github.com/mitchellh/gox@v1.0.1 first

CGO_ENABLED=1 gox -ldflags "${LDFLAGS}" -output="bin/nbview_{{.OS}}_{{.Arch}}" --osarch="darwin/amd64 darwin/arm64 freebsd/386 freebsd/amd64 freebsd/arm linux/386 linux/amd64 linux/arm linux/arm64 linux/mips linux/mips64 linux/mips64le linux/mipsle linux/ppc64 linux/ppc64le linux/s390x netbsd/386 netbsd/amd64 netbsd/arm openbsd/386 openbsd/amd64 windows/386 windows/amd64"

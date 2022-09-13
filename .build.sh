#!/bin/sh

if [ -n "$1" ] && [ "$1" = "all" ]; then

    set -e
    set -x
    VERSION=${VERSION:-0.0.0}

    # disable CGO for x-compiling
    export CGO_ENABLED=0

    # build for multiple arches
    GOOS=linux   GOARCH=amd64  go build -ldflags "-X main.version=${VERSION##v}" -o release/linux/amd64/isaaksay ./cmd/isaaksay/...
    GOOS=linux   GOARCH=arm    go build -ldflags "-X main.version=${VERSION##v}" -o release/linux/arm/isaaksay ./cmd/isaaksay/...
    GOOS=linux   GOARCH=arm64  go build -ldflags "-X main.version=${VERSION##v}" -o release/linux/arm64/isaaksay ./cmd/isaaksay/...
    GOOS=windows GOARCH=amd64  go build -ldflags "-X main.version=${VERSION##v}" -o release/windows/amd64/isaaksay.exe ./cmd/isaaksay/...
    GOOS=darwin  GOARCH=amd64  go build -ldflags "-X main.version=${VERSION##v}"  -o release/darwin/amd64/isaaksay ./cmd/isaaksay/...
    GOOS=darwin  GOARCH=arm64  go build -ldflags "-X main.version=${VERSION##v}" -o release/darwin/arm64/isaaksay ./cmd/isaaksay/...

    GOOS=linux   GOARCH=amd64  go build -ldflags "-X main.version=${VERSION##v}" -o release/linux/amd64/gunersay ./cmd/gunersay/...
    GOOS=linux   GOARCH=arm    go build -ldflags "-X main.version=${VERSION##v}" -o release/linux/arm/gunersay ./cmd/gunersay/... 
    GOOS=linux   GOARCH=arm64  go build -ldflags "-X main.version=${VERSION##v}" -o release/linux/arm64/gunersay ./cmd/gunersay/...
    GOOS=windows GOARCH=amd64  go build -ldflags "-X main.version=${VERSION##v}" -o release/windows/amd64/gunersay.exe ./cmd/gunersay/...
    GOOS=darwin  GOARCH=amd64  go build -ldflags "-X main.version=${VERSION##v}" -o release/darwin/amd64/gunersay ./cmd/gunersay/...
    GOOS=darwin  GOARCH=arm64  go build -ldflags "-X main.version=${VERSION##v}" -o release/darwin/arm64/gunersay ./cmd/gunersay/...

    # tar binary files prior to upload
    tar -czvf release/isaaksay_linux_amd64.tar.gz       -C release/linux/amd64/   isaaksay
    tar -czvf release/isaaksay_linux_arm.tar.gz         -C release/linux/arm/     isaaksay
    tar -czvf release/isaaksay_linux_arm64.tar.gz       -C release/linux/arm64/   isaaksay
    tar -czvf release/isaaksay_windows_amd64.exe.tar.gz -C release/windows/amd64/ isaaksay.exe
    tar -czvf release/isaaksay_darwin_amd64.tar.gz      -C release/darwin/amd64/  isaaksay
    tar -czvf release/isaaksay_darwin_arm64.tar.gz      -C release/darwin/arm64/  isaaksay

    tar -czvf release/gunersay_linux_amd64.tar.gz       -C release/linux/amd64/   gunersay
    tar -czvf release/gunersay_linux_arm.tar.gz         -C release/linux/arm/     gunersay
    tar -czvf release/gunersay_linux_arm64.tar.gz       -C release/linux/arm64/   gunersay
    tar -czvf release/gunersay_windows_amd64.exe.tar.gz -C release/windows/amd64/ gunersay.exe
    tar -czvf release/gunersay_darwin_amd64.tar.gz      -C release/darwin/amd64/  gunersay
    tar -czvf release/gunersay_darwin_arm64.tar.gz      -C release/darwin/arm64/  gunersay

    # checksums
    if [ "$(uname -s)" = "Darwin" ]; then
        shasum -a256 release/*.tar.gz > release/himbsay_checksums.txt
    else
        sha256sum release/*.tar.gz > release/himbsay_checksums.txt
    fi
else
    set -e
    set -x
    VERSION=${VERSION:-0.0.0}

    # build native
    go build -ldflags "-X main.version=${VERSION##v}" ./cmd/isaaksay/...
    go build -ldflags "-X main.version=${VERSION##v}" ./cmd/gunersay/...
fi


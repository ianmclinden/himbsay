
<h1 align="center">himbsay - speaking terminal buddies isaak and güner</h1>

<p align="center">
    <span><img alt="isaaksay" src="https://raw.githubusercontent.com/ianmclinden/himbsay/main/.public/isaaksay.png" width="300"/><img alt="gunersay" src="https://raw.githubusercontent.com/ianmclinden/himbsay/main/.public/gunersay.png" width="300"/>
    </span>
</p>

<p align="center">
    <a href="https://github.com/ianmclinden/himbsay/actions/workflows/test.yml?query=branch%3Amain+" title="Build status">
        <img src="https://github.com/ianmclinden/himbsay/actions/workflows/test.yml/badge.svg?branch=main">
    </a>
    <a href="https://github.com/ianmclinden/himbsay/releases/latest" title="GitHub release">
        <img src="https://img.shields.io/github/release/ianmclinden/himbsay.svg">
    </a>
    <a href="https://opensource.org/licenses/MIT" title="License: MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg">
    </a>
</>

## About

A talking isaak and a talking guner. Like an overly-specific [cowsay](https://github.com/tnalpgge/rank-amateur-cowsay).


## Installation

### Linux
```sh
# Install isaaksay
curl -L https://github.com/ianmclinden/himbsay/releases/latest/download/gunersay_isaaksay_amd64.tar.gz | tar zx
sudo install -t /usr/local/bin isaaksay
```
```sh
# Install gunersay
curl -L https://github.com/ianmclinden/himbsay/releases/latest/download/gunersay_linux_amd64.tar.gz | tar zx
sudo install -t /usr/local/bin gunersay
```

Binaries for [`isaaksay/arm`](https://github.com/ianmclinden/himbsay/releases/latest/download/isaaksay_linux_arm.tar.gz), [`gunersay/arm`](https://github.com/ianmclinden/himbsay/releases/latest/download/gunersay_linux_arm.tar.gz), [`isaaksay/arm64`](https://github.com/ianmclinden/himbsay/releases/latest/download/isaaksay_linux_arm64.tar.gz), and [`gunersay/arm64`](https://github.com/ianmclinden/himbsay/releases/latest/download/gunersay_linux_arm64.tar.gz) are also available.

### macOS
```sh
# Install isaaksay
curl -L https://github.com/ianmclinden/himbsay/releases/latest/download/isaaksay_darwin_amd64.tar.gz | tar zx
sudo install -t /usr/local/bin isaaksay
```
```sh
# Install gunersay
curl -L https://github.com/ianmclinden/himbsay/releases/latest/download/gunersay_darwin_amd64.tar.gz | tar zx
sudo install -t /usr/local/bin gunersay
```

Binaries for [`isaaksay/arm64`](https://github.com/ianmclinden/himbsay/releases/latest/download/isaaksay_darwin_arm64.tar.gz) and [`gunersay/arm64`](https://github.com/ianmclinden/himbsay/releases/latest/download/gunersay_darwin_arm64.tar.gz) are also available.

### Windows

Binaries for [`isaaksay/amd64`](https://github.com/ianmclinden/himbsay/releases/latest/download/isaaksay_windows_amd64.tar.gz) and [`gunersay/amd64`](https://github.com/ianmclinden/himbsay/releases/latest/download/gunersay_windows_amd64.tar.gz) are available.


## Building

```sh
# Clone source
git clone https://github.com/ianmclinden/himbsay.git
cd himbsay
# Build
./.build.sh
# Install (optional)
go install ./...
```

## Contributors
- Ian McLinden [@ianmclinden](https://github.com/ianmclinden)
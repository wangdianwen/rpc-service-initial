#!/bin/bash
set -e

echo "=== Installing Go dependencies ==="
go mod download

echo "=== Installing Go tools ==="
go install mvdan.cc/gofumpt@latest
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/golang/mock/mockgen@latest
go install gotest.tools/gotestsum@latest
go install github.com/cweill/gotests/gotests@latest
go install github.com/maxbrunsfeld/counterfeiter/v6@latest
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
go install github.com/mdempsky/maligned@latest
go install github.com/timakin/bodyclose@latest
go install github.com/google/wire/cmd/wire@latest


echo "=== Configuring git ==="
git config --global user.name "Wang Dianwen"
git config --global user.email "wangdw2012@gmail.com"

echo "=== Setting up shell ==="
apt-get update -qq
apt-get install -y --no-install-recommends zsh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended
echo 'exec zsh' >> ~/.bashrc

 echo "=== Installing additional tools ==="
curl -fsSL https://opencode.ai/install | bash
echo 'export PATH=/root/.opencode/bin:$PATH' >> ~/.zshrc

echo "=== Installing OpenSpec ==="
npm install -g @fission-ai/openspec
export PATH="/root/.npm-global/bin:$PATH"
echo 'export PATH="/root/.npm-global/bin:$PATH"' >> ~/.zshrc

echo "=== Installation completed successfully! ==="

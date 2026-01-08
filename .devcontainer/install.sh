#!/bin/bash
set -e

echo "Installing Go dependencies..."
go mod download

echo "Installing protoc plugins..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

echo "Installing additional tools..."
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install zsh and oh-my-zsh
apt-get update
apt-get install -y zsh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended
echo 'exec zsh' >> ~/.bashrc


# Configure git
git config --global user.name "Wang Dianwen"
git config --global user.email "wangdw2012@gmail.com"

# Install OpenCode CLI
curl -fsSL https://opencode.ai/install | bash
echo 'export PATH=/root/.opencode/bin:$PATH' >> ~/.zshrc

echo "Installation completed successfully!"

#!/bin/bash
mkdir -p ~/.terraform.d/plugins/terraform-apex.com/apexprovider/apex/1.0.0/linux_amd64
go mod init github.com/richmondang/custom-tf
go fmt
go mod tidy
go build -o terraform-provider-apex
cp terraform-provider-apex ~/.terraform.d/plugins/terraform-apex.com/apexprovider/apex/1.0.0/linux_amd64
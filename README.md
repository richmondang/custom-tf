# APEX Custom Terraform Provider Example

## Initial Environment Setup

- Below steps are specific to Ubuntu 20.04 LTS - Windows Linux Subsystem

```
sudo wget https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.17.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
source ~/.bashrc
```
## Verify Go is installed properly
```
go version
```


## Install Terraform (Ubuntu)
- Reference: https://learn.hashicorp.com/tutorials/terraform/install-cli
```
sudo apt-get update && sudo apt-get install -y gnupg software-properties-common curl
curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -
sudo apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
sudo apt-get update && sudo apt-get install terraform
```

## Running Mock API Server
In a separate terminal:
```
cd api
go mod init github.com/richmondang/custom-tf/api
go mod tidy
go run .
```
This will run locally on localhost:3001


## Set Go Environment Variable to run API server, tests and build provider
```
export GO111MODULE=on
```


## Creating Custom Terraform Provider Executable
After building Go code, create terraform provider executable and copy to appropriate directory structure:
```
mkdir -p ~/.terraform.d/plugins/terraform-apex.com/apexprovider/apex/1.0.0/linux_amd64
go mod init github.com/richmondang/custom-tf
go fmt
go mod tidy
go build -o terraform-provider-apex
cp terraform-provider-apex ~/.terraform.d/plugins/terraform-apex.com/apexprovider/apex/1.0.0/linux_amd64
```
Alternatively, execute `scripts/build.sh` bash script from root folder:
```
./scripts/build.sh
```

## Note on Custom Terraform Providers - local execution
Custom Terraform Provider Executable must be located in a Terraform Plugins Directory for Terraform to recognize:
    - Linux: `~/.terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${target}`
    - Windows: `%APPDATA%\terraform.d\plugins\${host_name}/${namespace}/${type}/${version}/${target}`
Terraform searches for plugins in the format of `terraform-<TYPE>-<NAME>`, in the above example `terraform-provider-apex`, the custom plugin is of type `provider` named `apex`


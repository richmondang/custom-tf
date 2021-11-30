# Creating Custom Terraform Provider - APEX

## Initial Environment Setup

- Below steps are specific to Ubuntu 20.04 LTS - Windows Subsystem

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
## Copy Go executable to bin (global exec)
```
cd /usr/local/go/bin
sudo cp go /usr/local/bin/
sudo cp gofmt /usr/local/bin/
```

## Install Terraform
```
curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -
sudo apt-add-repository "deb [arch=$(dpkg --print-architecture)] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
sudo apt install terraform
```

## Creating Custom Terraform Provider Executable
- After building Go code, create terraform provider executable and copy to appropriate directory structure:
```
mkdir -p ~/.terraform.d/plugins/terraform-apex.com/apexprovider/apex/1.0.0/linux_amd64
go mod init
go fmt
go mod tidy
go build -o terraform-provider-apex
cp terraform-provider-apex ~/.terraform.d/plugins/terraform-apex.com/apexprovider/apex/1.0.0/linux_amd64
```
- Alternatively, execute create-binary.sh script
```
./create-binary.sh
```

## Note on Custom Terraform Providers - local execution
- Custom Terraform Provider Executable must be located in a Terraform Plugins Directory for Terraform to recognize:
    - Linux: `~/.terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${target}`
    - Windows: `%APPDATA%\terraform.d\plugins\${host_name}/${namespace}/${type}/${version}/${target}`
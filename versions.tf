terraform {
  required_providers {
    apex = {
      version = "~> 1.0.0"
      # source  = "terraform-example.com/exampleprovider/example"
      source = "terraform-apex.com/apexprovider/apex"
    }
  }
}
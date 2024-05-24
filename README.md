# Terraform Provider for Elasticsearch

This Terraform provider enables you to manage Elasticsearch resources, specifically index templates, through Terraform. It is designed for use with Elasticsearch version 7.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) v1.1.7
- [Go](https://golang.org/doc/install) 1.22 or newer

## Installation

To install this provider, you need to build it from the source and then configure Terraform to use the local binary.

### Building from Source

Clone the repository and build the provider:

```bash
git clone https://github.com/<your-username>/terraform-provider-elasticsearch.git
cd terraform-provider-elasticsearch
make install
```

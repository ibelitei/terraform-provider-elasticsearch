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

## Variables Makefile

- `HOSTNAME`: The registry hostname where the provider is hosted.
- `NAMESPACE`: The namespace of the provider.
- `NAME`: The name of the provider.
- `BINARY`: The name of the binary file that will be built.
- `VERSION`: The version of the provider.
- `OS_ARCH`: The operating system and architecture for which the provider is built.

## Default Target
The default target is `install`, which builds and installs the Terraform provider.

## Targets
### `build`
Build the Terraform provider binary for local use.
Usage:
```sh
make build
```
Build the provider and installs it in the Terraform plugins directory.
```sh
make install
```
Run an Elasticsearch instance using Docker Compose.
```sh
make run-elasticsearch
```

Run Test
```sh
make test
```


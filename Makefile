HOSTNAME=registry.terraform.io
NAMESPACE=softwaremind
NAME=elasticsearch
BINARY=terraform-provider-${NAME}
VERSION=1.0.0
OS_ARCH=darwin_arm64

default: install

build: # Builds
	@echo Build for local
	go build -o ${BINARY}

install: build
	@echo Installing provider in Terraform
	rm -rf ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

run-elasticsearch:
	@echo Running elasticsearch
	docker-compose -f docker-compose.yml up -d
	sleep 15

test: build install run-elasticsearch
	@echo Terraform apply
	rm -rf examples/.terraform*
	rm -rf examples/terraform*
	terraform -chdir=examples/ init
	terraform -chdir=examples/ apply -auto-approve


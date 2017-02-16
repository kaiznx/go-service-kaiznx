.PHONY: all

all:
	@echo "usage: make [TARGET]"
	@echo "TARGETS:"
	@echo "  install"
	@echo "  create-env"
	@echo "  deploy"
	@echo "  test"

get-godep:
	@go get github.com/tools/godep

save-deps:
	@godep save ./...

build: get-godep
	@mkdir -p bin
	@GOOS="linux" GOARCH="amd64" CGO_ENABLED=0 godep go build -o bin/application ./cmd/kaizenx/...
	@rm -rf application application.zip

# Install the app into GOBIN path
install: get-godep
	@godep go install ./cmd/...

# Create the AWS Beanstalk application
create:
	@eb init KaizenX -p go1.5 -r us-east-2

# Create the staging environment in AWS
create-env: create
	@eb create kaizenx-stg -k kaizenx_stg_key --elb-type application --sample

# Build and deploy to the staging environment
deploy: build
	@zip application.zip bin/application Procfile -r .ebextensions/*
	@git add .
	@eb use kaizenx-stg
	@eb deploy --staged
	@git rm -rf application.zip

# Run tests
test:
	@godep go test -race -v ./...
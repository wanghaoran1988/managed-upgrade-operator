SHELL := /usr/bin/env bash

OPERATOR_DOCKERFILE = ./build/Dockerfile

# Include shared Makefiles
include project.mk
include standard.mk

default: gobuild

# Extend Makefile after here

.PHONY: docker-build
docker-build: build

.PHONY: generate
generate:
	operator-sdk generate k8s
	operator-sdk generate crds
	openapi-gen --logtostderr=true \
		-i ./pkg/apis/upgrade/v1alpha1 \
		-o "" \
		-O zz_generated.openapi \
		-p ./pkg/apis/upgrade/v1alpha1 \
		-h /dev/null \
		-r "-"
	go generate pkg/cluster_upgrader/cluster_upgrader.go
	go generate pkg/maintenance/maintenance.go

.PHONY: coverage
coverage:
	hack/codecov.sh

.PHONY: run
run: 
	operator-sdk run --local --watch-namespace ""

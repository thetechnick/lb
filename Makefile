# Copyright 2019 The LB Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

export CGO_ENABLED:=0

BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
SHORT_SHA=$(shell git rev-parse --short HEAD)
VERSION?=${BRANCH}-${SHORT_SHA}
BUILD_DATE=$(shell date +%s)

MODULE=github.com/thetechnick/lb
LD_FLAGS="-w -X '$(MODULE)/pkg/version.Version=$(VERSION)' -X '$(MODULE)/pkg/version.Branch=$(BRANCH)' -X '$(MODULE)/pkg/version.Commit=$(SHORT_SHA)' -X '$(MODULE)/pkg/version.BuildDate=$(BUILD_DATE)'"

bin/linux_amd64/%: GOARGS = GOOS=linux GOARCH=amd64

bin/%: FORCE generate
	$(eval COMPONENT=$(shell basename $*))
	$(GOARGS) go build -ldflags $(LD_FLAGS) -o bin/$* cmd/$(COMPONENT)/main.go

FORCE:

clean:
	rm -rf bin/$*
.PHONEY: clean

generate: export PATH = $(shell echo "$$PWD/bin:$$PATH")
generate: tools
	@go generate ./...

tools: bin/protoc bin/protoc-gen-go

bin/protoc:
	@./hack/get-protoc.sh

bin/protoc-gen-go:
	@go build -o bin/protoc-gen-go github.com/golang/protobuf/protoc-gen-go

tidy:
	go mod tidy
.PHONEY: tidy

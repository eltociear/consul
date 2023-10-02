SHELL := /bin/bash

.PHONY: noop
noop:

##@ Build

.PHONY: tidy
tidy: ## Run go mod tidy.
	go mod tidy

##@ Checks

.PHONY: format
format: ## Format the go files.
	@for f in $$(find . -name '*.go' -print); do \
		gofmt -s -w $$f ; \
	done

.PHONY: lint
lint: ## Run the full linting rules.
	golangci-lint run -v

.PHONY: vet
vet: ## Run go vet.
	go vet ./...

##@ Testing

.PHONY: test
test: ## Run all go tests.
	go test ./...

##@ Code Generation

.PHONY: update-defaults ## Update default images versions for all relevant containers to the latest.
update-defaults: update-consul update-envoy update-dataplane

.PHONY: update-consul
update-consul: ## Update the default images for Consul CE and Enterprise containers to the latest.
	@docker pull hashicorp/consul:latest || true
	@mkdir -p tmp
	@docker run --rm hashicorp/consul:latest version -format=json | jq -r .Version > tmp/default_consul.val
	@printf "// Copyright (c) HashiCorp, Inc.\n// SPDX-License-Identifier: BUSL-1.1\n\n" > topology/default_consul.go
	@printf "package topology\n\nconst DefaultConsulImage = \"hashicorp/consul:$$(cat tmp/default_consul.val)\"\n" >> topology/default_consul.go
	@printf "const DefaultConsulEnterpriseImage = \"hashicorp/consul-enterprise:$$(cat tmp/default_consul.val)-ent\"\n" >> topology/default_consul.go
	@rm -rf tmp

.PHONY: update-envoy
update-envoy: ## Update the default image for the Envoy container to the latest.
	@docker rm -f consul-envoy-check &>/dev/null || true
	@docker pull hashicorp/consul:latest || true
	@docker run -d --name consul-envoy-check hashicorp/consul:latest
	@mkdir -p tmp
	@docker exec consul-envoy-check sh -c 'wget -q localhost:8500/v1/agent/self -O -' | jq -r '.xDS.SupportedProxies.envoy[0]' > tmp/default_envoy.val
	@docker rm -f consul-envoy-check &>/dev/null || true
	@printf "// Copyright (c) HashiCorp, Inc.\n// SPDX-License-Identifier: BUSL-1.1\n\n" > topology/default_envoy.go
	@printf "package topology\n\nconst DefaultEnvoyImage = \"envoyproxy/envoy:v$$(cat tmp/default_envoy.val)\"\n" >> topology/default_envoy.go
	@rm -rf tmp

.PHONY: update-dataplane
update-dataplane: ## Update the default image for the Consul Dataplane container to the latest.
	@docker pull hashicorp/consul-dataplane:latest || true
	@mkdir -p tmp
	@docker run --rm hashicorp/consul-dataplane:latest --version | head -n 1 | awk '{print $$3}' | cut -d'v' -f2- > tmp/default_dataplane.val
	@printf "// Copyright (c) HashiCorp, Inc.\n// SPDX-License-Identifier: BUSL-1.1\n\n" > topology/default_cdp.go
	@printf "package topology\n\nconst DefaultDataplaneImage = \"hashicorp/consul-dataplane:$$(cat tmp/default_dataplane.val)\"\n" >> topology/default_cdp.go
	@rm -rf tmp

##@ Help

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php
.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

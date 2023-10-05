#!/bin/bash

set -euo pipefail

unset CDPATH

cd "$(dirname "$0")"

###
# This script will update the default image names to the latest released versions of
# Consul, Consul Enterprise, and Consul Dataplane.
#
# For Envoy, it will interrogate the latest version of Consul for it's maximum supported
# Envoy version and use that.
###

## Update the default images for Consul CE and Enterprise containers to the latest.
function update_consul {
    local version

	docker pull hashicorp/consul:latest || true
    version="$(docker image inspect hashicorp/consul:latest | jq -r '.[0].Config.Labels."org.opencontainers.image.version"')"
    cat > topology/default_consul.go <<EOF
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package topology

const DefaultConsulImage = "hashicorp/consul:${version}"
const DefaultConsulEnterpriseImage = "hashicorp/consul-enterprise:${version}-ent"
EOF
}

## Update the default image for the Envoy container to the latest.
function update_envoy {
    local version

	docker rm -f consul-envoy-check &>/dev/null || true
	docker pull hashicorp/consul:latest || true
	docker run -d --name consul-envoy-check hashicorp/consul:latest

    while true; do
        version="$(docker exec consul-envoy-check sh -c 'wget -q localhost:8500/v1/agent/self -O -' | jq -r '.xDS.SupportedProxies.envoy[0]')"
        if [[ -n "$version" ]]; then
            break
        fi
    done

    cat > topology/default_envoy.go <<EOF
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package topology

const DefaultEnvoyImage = "envoyproxy/envoy:v${version}"
EOF
}

## Update the default image for the Consul Dataplane container to the latest.
function update_dataplane {
    local version

	docker pull hashicorp/consul-dataplane:latest || true
    version="$(docker image inspect hashicorp/consul-dataplane:latest | jq -r '.[0].Config.Labels.version')"
    cat > topology/default_cdp.go <<EOF
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package topology

const DefaultDataplaneImage = "hashicorp/consul-dataplane:${version}"
EOF
}

update_consul
update_envoy
update_dataplane

// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package peering

import "github.com/hashicorp/consul/test-integ/topoutil"

// Deprecated: topoutil.Asserter
type asserter = topoutil.Asserter

// Deprecated: topoutil.SprawlLite
type sprawlLite = topoutil.SprawlLite

// Deprecated: topoutil.NewAsserter
func newAsserter(sp sprawlLite) *asserter {
	return topoutil.NewAsserter(sp)
}

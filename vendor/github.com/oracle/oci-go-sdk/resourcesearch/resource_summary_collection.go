// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Resource Search Service
//
// Search for resources across your cloud infrastructure
//

package resourcesearch

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ResourceSummaryCollection A collection of resources that matched the search criteria.
type ResourceSummaryCollection struct {

	// A list of resources.
	Items []ResourceSummary `mandatory:"false" json:"items"`
}

func (m ResourceSummaryCollection) String() string {
	return common.PointerString(m)
}
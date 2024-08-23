/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha3

import (
	"github.com/kcp-dev/logicalcluster/v3"
)

// ResourceClaimListerExpansion allows custom methods to be added to
// ResourceClaimLister.
type ResourceClaimListerExpansion interface{}

// ResourceClaimClusterListerExpansion allows custom methods to be added to
// ResourceClaimLister.
type ResourceClaimClusterListerExpansion interface {
	// Cluster returns a lister that can list and get ResourceClaim in one workspace.
	Cluster(clusterName logicalcluster.Name) ResourceClaimLister
}

// ResourceClaimNamespaceListerExpansion allows custom methods to be added to
// ResourceClaimNamespaceLister.
type ResourceClaimNamespaceListerExpansion interface{}
package cloudservicesnetwork

import (
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/nogen/azcommonv6"
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/test/common"
)

// CloudServicesNetworkList represents a list of cloud services networks.
type CloudServicesNetworkList struct {
  // The CloudServicesNetwork items on this page
  value []CloudServicesNetwork

  // The link to the next page of items
  nextLink string
}

// Upon creation, the additional services that are provided by the platform will be allocated and
// represented in the status of this resource. All resources associated with this cloud services network will be part
// of the same layer 2 (L2) isolation domain. At least one service network must be created but may be reused across many
// virtual machines and/or Hybrid AKS clusters.
type CloudServicesNetwork struct {
  // The list of the resource properties.
  properties CloudServicesNetworkProperties

  // The name of the cloud services network.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string

  // The extended location of the cluster associated with the resource.
  extendedLocation common.ExtendedLocation

  azcommonv6.TrackedResource
}

// CloudServicesNetworkProperties represents properties of the cloud services network.
type CloudServicesNetworkProperties struct {
  // The list of egress endpoints. This allows for connection from a Hybrid AKS cluster to the specified endpoint.
  additionalEgressEndpoints []EgressEndpoint

  // The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
  associatedResourceIds []string

  // The resource ID of the Network Cloud cluster this cloud services network is associated with.
  clusterId string

  // The more detailed status of the cloud services network.
  detailedStatus CloudServicesNetworkDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // The indicator of whether the platform default endpoints are allowed for the egress traffic.
  enableDefaultEgressEndpoints CloudServicesNetworkEnableDefaultEgressEndpoints

  // The full list of additional and default egress endpoints that are currently enabled.
  enabledEgressEndpoints []EgressEndpoint

  // Field Deprecated. These fields will be empty/omitted. The list of Hybrid AKS cluster resource IDs that are associated with this cloud services network.
  hybridAksClustersAssociatedIds []string

  // The name of the interface that will be present in the virtual machine to represent this network.
  interfaceName string

  // The provisioning state of the cloud services network.
  provisioningState CloudServicesNetworkProvisioningState

  // Field Deprecated. These fields will be empty/omitted. The list of virtual machine resource IDs, excluding any Hybrid AKS virtual machines, that are currently using this cloud services network.
  virtualMachinesAssociatedIds []string
}

// EgressEndpoint represents the connection from a cloud services network to the specified endpoint for a common purpose.
type EgressEndpoint struct {
  // The descriptive category name of endpoints accessible by the AKS agent node. For example, azure-resource-management, API server, etc. The platform egress endpoints provided by default will use the category 'default'.
  category string

  // The list of endpoint dependencies.
  endpoints []EndpointDependency
}

// EndpointDependency represents the definition of an endpoint, including the domain and details.
type EndpointDependency struct {
  // The domain name of the dependency.
  domainName string

  // The port of this endpoint.
  port int64
}

// CloudServicesNetworkPatchParameters represents the body of the request to patch the cloud services network.
type CloudServicesNetworkPatchParameters struct {
  // The list of the resource properties.
  properties CloudServicesNetworkPatchProperties

  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

// CloudServicesNetworkPatchProperties represents the properties of the cloud services network that can be updated using a patch request.
type CloudServicesNetworkPatchProperties struct {
  // The list of egress endpoints. This allows for connection from a Hybrid AKS cluster to the specified endpoint.
  additionalEgressEndpoints []EgressEndpoint

  // The indicator of whether the platform default endpoints are allowed for the egress traffic.
  enableDefaultEgressEndpoints CloudServicesNetworkEnableDefaultEgressEndpoints
}

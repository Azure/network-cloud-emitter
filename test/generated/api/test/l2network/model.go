package l2network

import (
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/nogen/azcommonv6"
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/test/common"
)

// L2NetworkList represents a list of L2 networks.
type L2NetworkList struct {
  // The L2Network items on this page
  value []L2Network

  // The link to the next page of items
  nextLink string
}

// L2Network represents a network that utilizes a single isolation domain set up for layer-2 resources.
type L2Network struct {
  // The list of the resource properties.
  properties L2NetworkProperties

  // The name of the L2 network.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string

  // The extended location of the cluster associated with the resource.
  extendedLocation common.ExtendedLocation

  azcommonv6.TrackedResource
}

// L2NetworkProperties represents properties of the L2 network.
type L2NetworkProperties struct {
  // The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
  associatedResourceIds []string

  // The resource ID of the Network Cloud cluster this L2 network is associated with.
  clusterId string

  // The more detailed status of the L2 network.
  detailedStatus L2NetworkDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // Field Deprecated. These fields will be empty/omitted. The list of Hybrid AKS cluster resource ID(s) that are associated with this L2 network.
  hybridAksClustersAssociatedIds []string

  // Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
  hybridAksPluginType common.HybridAksPluginType

  // The default interface name for this L2 network in the virtual machine. This name can be overridden by the name supplied in the network attachment configuration of that virtual machine.
  interfaceName string

  // The resource ID of the Network Fabric l2IsolationDomain.
  l2IsolationDomainId string

  // The provisioning state of the L2 network.
  provisioningState L2NetworkProvisioningState

  // Field Deprecated. These fields will be empty/omitted. The list of virtual machine resource ID(s), excluding any Hybrid AKS virtual machines, that are currently using this L2 network.
  virtualMachinesAssociatedIds []string
}

// L2NetworkPatchParameters represents the body of the request to patch the L2 network.
type L2NetworkPatchParameters struct {
  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

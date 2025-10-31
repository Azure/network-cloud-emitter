package trunkednetwork

import (
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/nogen/azcommonv6"
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/test/common"
)

// TrunkedNetworkList represents a list of trunked networks.
type TrunkedNetworkList struct {
  // The TrunkedNetwork items on this page
  value []TrunkedNetwork

  // The link to the next page of items
  nextLink string
}

// TrunkedNetwork represents a network that utilizes multiple isolation domains and specified VLANs to create a trunked network.
type TrunkedNetwork struct {
  // The list of the resource properties.
  properties TrunkedNetworkProperties

  // The name of the trunked network.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string

  // The extended location of the cluster associated with the resource.
  extendedLocation common.ExtendedLocation

  azcommonv6.TrackedResource
}

// TrunkedNetworkProperties represents properties of the trunked network.
type TrunkedNetworkProperties struct {
  // The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
  associatedResourceIds []string

  // The resource ID of the Network Cloud cluster this trunked network is associated with.
  clusterId string

  // The more detailed status of the trunked network.
  detailedStatus TrunkedNetworkDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // Field Deprecated. These fields will be empty/omitted. The list of Hybrid AKS cluster resource IDs that are associated with this trunked network.
  hybridAksClustersAssociatedIds []string

  // Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
  hybridAksPluginType common.HybridAksPluginType

  // The default interface name for this trunked network in the virtual machine. This name can be overridden by the name supplied in the network attachment configuration of that virtual machine.
  interfaceName string

  // The list of resource IDs representing the Network Fabric isolation domains. It can be any combination of l2IsolationDomain and l3IsolationDomain resources.
  isolationDomainIds []string

  // The provisioning state of the trunked network.
  provisioningState TrunkedNetworkProvisioningState

  // Field Deprecated. These fields will be empty/omitted. The list of virtual machine resource IDs, excluding any Hybrid AKS virtual machines, that are currently using this trunked network.
  virtualMachinesAssociatedIds []string

  // The list of vlans that are selected from the isolation domains for trunking.
  vlans []int64
}

// TrunkedNetworkPatchParameters represents the body of the request to patch the Trunked network.
type TrunkedNetworkPatchParameters struct {
  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

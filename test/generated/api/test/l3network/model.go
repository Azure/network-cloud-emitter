package l3network

import (
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/nogen/azcommonv6"
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/test/common"
)

// L3NetworkList represents a list of L3 networks.
type L3NetworkList struct {
  // The L3Network items on this page
  value []L3Network

  // The link to the next page of items
  nextLink string
}

// L3Network represents a network that utilizes a single isolation domain set up for layer-3 resources.
type L3Network struct {
  // The list of the resource properties.
  properties L3NetworkProperties

  // The name of the L3 network.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string

  // The extended location of the cluster associated with the resource.
  extendedLocation common.ExtendedLocation

  azcommonv6.TrackedResource
}

// L3NetworkProperties represents properties of the L3 network.
type L3NetworkProperties struct {
  // The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
  associatedResourceIds []string

  // The resource ID of the Network Cloud cluster this L3 network is associated with.
  clusterId string

  // The more detailed status of the L3 network.
  detailedStatus L3NetworkDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // Field Deprecated. These fields will be empty/omitted. The list of Hybrid AKS cluster resource IDs that are associated with this L3 network.
  hybridAksClustersAssociatedIds []string

  // Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The indicator of whether or not to disable IPAM allocation on the network attachment definition injected into the Hybrid AKS Cluster.
  hybridAksIpamEnabled HybridAksIpamEnabled

  // Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
  hybridAksPluginType common.HybridAksPluginType

  // The default interface name for this L3 network in the virtual machine. This name can be overridden by the name supplied in the network attachment configuration of that virtual machine.
  interfaceName string

  // The type of the IP address allocation, defaulted to "DualStack".
  ipAllocationType IpAllocationType

  // The IPV4 prefix (CIDR) assigned to this L3 network. Required when the IP allocation type
  // is IPV4 or DualStack.
  ipv4ConnectedPrefix string

  // The IPV6 prefix (CIDR) assigned to this L3 network. Required when the IP allocation type
  // is IPV6 or DualStack.
  ipv6ConnectedPrefix string

  // The resource ID of the Network Fabric l3IsolationDomain.
  l3IsolationDomainId string

  // The provisioning state of the L3 network.
  provisioningState L3NetworkProvisioningState

  // Field Deprecated. These fields will be empty/omitted. The list of virtual machine resource IDs, excluding any Hybrid AKS virtual machines, that are currently using this L3 network.
  virtualMachinesAssociatedIds []string

  // The VLAN from the l3IsolationDomain that is used for this network.
  vlan int64
}

// L3NetworkPatchParameters represents the body of the request to patch the cloud services network.
type L3NetworkPatchParameters struct {
  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

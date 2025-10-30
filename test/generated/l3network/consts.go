package l3network

// The more detailed status of the L3 network.
type L3NetworkDetailedStatus string

const (
  L3NetworkDetailedStatusError L3NetworkDetailedStatus = "Error"

  L3NetworkDetailedStatusAvailable L3NetworkDetailedStatus = "Available"

  L3NetworkDetailedStatusProvisioning L3NetworkDetailedStatus = "Provisioning"
)

// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The indicator of whether or not to disable IPAM allocation on the network attachment definition injected into the Hybrid AKS Cluster.
type HybridAksIpamEnabled string

const (
  HybridAksIpamEnabledTrue HybridAksIpamEnabled = "True"

  HybridAksIpamEnabledFalse HybridAksIpamEnabled = "False"
)

// The type of the IP address allocation, defaulted to "DualStack".
type IpAllocationType string

const (
  IpAllocationTypeIPV4 IpAllocationType = "IPV4"

  IpAllocationTypeIPV6 IpAllocationType = "IPV6"

  IpAllocationTypeDualStack IpAllocationType = "DualStack"
)

// The provisioning state of the L3 network.
type L3NetworkProvisioningState string

const (
  L3NetworkProvisioningStateSucceeded L3NetworkProvisioningState = "Succeeded"

  L3NetworkProvisioningStateFailed L3NetworkProvisioningState = "Failed"

  L3NetworkProvisioningStateCanceled L3NetworkProvisioningState = "Canceled"

  L3NetworkProvisioningStateProvisioning L3NetworkProvisioningState = "Provisioning"

  L3NetworkProvisioningStateAccepted L3NetworkProvisioningState = "Accepted"
)

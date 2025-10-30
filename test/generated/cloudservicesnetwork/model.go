package cloudservicesnetwork

// CloudServicesNetworkList represents a list of cloud services networks.
type CloudServicesNetworkList struct{}

// Upon creation, the additional services that are provided by the platform will be allocated and
// represented in the status of this resource. All resources associated with this cloud services network will be part
// of the same layer 2 (L2) isolation domain. At least one service network must be created but may be reused across many
// virtual machines and/or Hybrid AKS clusters.
type CloudServicesNetwork struct{}

// CloudServicesNetworkProperties represents properties of the cloud services network.
type CloudServicesNetworkProperties struct{}

// CloudServicesNetworkPatchParameters represents the body of the request to patch the cloud services network.
type CloudServicesNetworkPatchParameters struct{}

// CloudServicesNetworkPatchProperties represents the properties of the cloud services network that can be updated using a patch request.
type CloudServicesNetworkPatchProperties struct{}

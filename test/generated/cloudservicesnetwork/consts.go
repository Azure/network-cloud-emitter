package cloudservicesnetwork

// The more detailed status of the cloud services network.
type CloudServicesNetworkDetailedStatus string

const (
  CloudServicesNetworkDetailedStatusError CloudServicesNetworkDetailedStatus = "Error"

  CloudServicesNetworkDetailedStatusAvailable CloudServicesNetworkDetailedStatus = "Available"

  CloudServicesNetworkDetailedStatusProvisioning CloudServicesNetworkDetailedStatus = "Provisioning"
)

// The indicator of whether the platform default endpoints are allowed for the egress traffic.
type CloudServicesNetworkEnableDefaultEgressEndpoints string

const (
  CloudServicesNetworkEnableDefaultEgressEndpointsTrue CloudServicesNetworkEnableDefaultEgressEndpoints = "True"

  CloudServicesNetworkEnableDefaultEgressEndpointsFalse CloudServicesNetworkEnableDefaultEgressEndpoints = "False"
)

// The provisioning state of the cloud services network.
type CloudServicesNetworkProvisioningState string

const (
  CloudServicesNetworkProvisioningStateSucceeded CloudServicesNetworkProvisioningState = "Succeeded"

  CloudServicesNetworkProvisioningStateFailed CloudServicesNetworkProvisioningState = "Failed"

  CloudServicesNetworkProvisioningStateCanceled CloudServicesNetworkProvisioningState = "Canceled"

  CloudServicesNetworkProvisioningStateProvisioning CloudServicesNetworkProvisioningState = "Provisioning"

  CloudServicesNetworkProvisioningStateAccepted CloudServicesNetworkProvisioningState = "Accepted"
)

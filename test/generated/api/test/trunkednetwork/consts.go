package trunkednetwork

// The more detailed status of the trunked network.
type TrunkedNetworkDetailedStatus string

const (
  TrunkedNetworkDetailedStatusError TrunkedNetworkDetailedStatus = "Error"

  TrunkedNetworkDetailedStatusAvailable TrunkedNetworkDetailedStatus = "Available"

  TrunkedNetworkDetailedStatusProvisioning TrunkedNetworkDetailedStatus = "Provisioning"
)

// The provisioning state of the trunked network.
type TrunkedNetworkProvisioningState string

const (
  TrunkedNetworkProvisioningStateSucceeded TrunkedNetworkProvisioningState = "Succeeded"

  TrunkedNetworkProvisioningStateFailed TrunkedNetworkProvisioningState = "Failed"

  TrunkedNetworkProvisioningStateCanceled TrunkedNetworkProvisioningState = "Canceled"

  TrunkedNetworkProvisioningStateProvisioning TrunkedNetworkProvisioningState = "Provisioning"

  TrunkedNetworkProvisioningStateAccepted TrunkedNetworkProvisioningState = "Accepted"
)

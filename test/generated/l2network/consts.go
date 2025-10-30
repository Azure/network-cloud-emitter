package l2network

// The more detailed status of the L2 network.
type L2NetworkDetailedStatus string

const (
  L2NetworkDetailedStatusError L2NetworkDetailedStatus = "Error"

  L2NetworkDetailedStatusAvailable L2NetworkDetailedStatus = "Available"

  L2NetworkDetailedStatusProvisioning L2NetworkDetailedStatus = "Provisioning"
)

// The provisioning state of the L2 network.
type L2NetworkProvisioningState string

const (
  L2NetworkProvisioningStateSucceeded L2NetworkProvisioningState = "Succeeded"

  L2NetworkProvisioningStateFailed L2NetworkProvisioningState = "Failed"

  L2NetworkProvisioningStateCanceled L2NetworkProvisioningState = "Canceled"

  L2NetworkProvisioningStateProvisioning L2NetworkProvisioningState = "Provisioning"

  L2NetworkProvisioningStateAccepted L2NetworkProvisioningState = "Accepted"
)

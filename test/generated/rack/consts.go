package rack

// The more detailed status of the rack.
type RackDetailedStatus string

const (
  RackDetailedStatusError RackDetailedStatus = "Error"

  RackDetailedStatusAvailable RackDetailedStatus = "Available"

  RackDetailedStatusProvisioning RackDetailedStatus = "Provisioning"
)

// The provisioning state of the rack resource.
type RackProvisioningState string

const (
  RackProvisioningStateSucceeded RackProvisioningState = "Succeeded"

  RackProvisioningStateFailed RackProvisioningState = "Failed"

  RackProvisioningStateCanceled RackProvisioningState = "Canceled"

  RackProvisioningStateProvisioning RackProvisioningState = "Provisioning"

  RackProvisioningStateAccepted RackProvisioningState = "Accepted"
)

package volume

// The more detailed status of the volume.
type VolumeDetailedStatus string

const (
  VolumeDetailedStatusError VolumeDetailedStatus = "Error"

  VolumeDetailedStatusActive VolumeDetailedStatus = "Active"

  VolumeDetailedStatusProvisioning VolumeDetailedStatus = "Provisioning"
)

// The provisioning state of the volume.
type VolumeProvisioningState string

const (
  VolumeProvisioningStateSucceeded VolumeProvisioningState = "Succeeded"

  VolumeProvisioningStateFailed VolumeProvisioningState = "Failed"

  VolumeProvisioningStateCanceled VolumeProvisioningState = "Canceled"

  VolumeProvisioningStateProvisioning VolumeProvisioningState = "Provisioning"

  VolumeProvisioningStateAccepted VolumeProvisioningState = "Accepted"
)

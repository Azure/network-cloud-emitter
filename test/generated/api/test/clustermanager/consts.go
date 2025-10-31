package clustermanager

// The detailed status that provides additional information about the cluster manager.
type ClusterManagerDetailedStatus string

const (
  ClusterManagerDetailedStatusError ClusterManagerDetailedStatus = "Error"

  ClusterManagerDetailedStatusAvailable ClusterManagerDetailedStatus = "Available"

  ClusterManagerDetailedStatusProvisioning ClusterManagerDetailedStatus = "Provisioning"

  ClusterManagerDetailedStatusProvisioningFailed ClusterManagerDetailedStatus = "ProvisioningFailed"

  ClusterManagerDetailedStatusUpdating ClusterManagerDetailedStatus = "Updating"

  ClusterManagerDetailedStatusUpdateFailed ClusterManagerDetailedStatus = "UpdateFailed"
)

// The provisioning state of the cluster manager.
type ClusterManagerProvisioningState string

const (
  ClusterManagerProvisioningStateSucceeded ClusterManagerProvisioningState = "Succeeded"

  ClusterManagerProvisioningStateFailed ClusterManagerProvisioningState = "Failed"

  ClusterManagerProvisioningStateCanceled ClusterManagerProvisioningState = "Canceled"

  ClusterManagerProvisioningStateProvisioning ClusterManagerProvisioningState = "Provisioning"

  ClusterManagerProvisioningStateAccepted ClusterManagerProvisioningState = "Accepted"

  ClusterManagerProvisioningStateUpdating ClusterManagerProvisioningState = "Updating"
)

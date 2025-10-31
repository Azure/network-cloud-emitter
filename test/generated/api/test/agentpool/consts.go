package agentpool

// The current status of the agent pool.
type AgentPoolDetailedStatus string

const (
  AgentPoolDetailedStatusAvailable AgentPoolDetailedStatus = "Available"

  AgentPoolDetailedStatusError AgentPoolDetailedStatus = "Error"

  AgentPoolDetailedStatusProvisioning AgentPoolDetailedStatus = "Provisioning"
)

// The provisioning state of the agent pool.
type AgentPoolProvisioningState string

const (
  AgentPoolProvisioningStateAccepted AgentPoolProvisioningState = "Accepted"

  AgentPoolProvisioningStateCanceled AgentPoolProvisioningState = "Canceled"

  AgentPoolProvisioningStateDeleting AgentPoolProvisioningState = "Deleting"

  AgentPoolProvisioningStateFailed AgentPoolProvisioningState = "Failed"

  AgentPoolProvisioningStateInProgress AgentPoolProvisioningState = "InProgress"

  AgentPoolProvisioningStateSucceeded AgentPoolProvisioningState = "Succeeded"

  AgentPoolProvisioningStateUpdating AgentPoolProvisioningState = "Updating"
)

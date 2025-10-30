package agentpool

// The size of the hugepages to allocate.
type HugepagesSize string

const (
  HugepagesSize2M HugepagesSize = "2M"

  HugepagesSize1G HugepagesSize = "1G"
)

// The current status of the agent pool.
type AgentPoolDetailedStatus string

const (
  AgentPoolDetailedStatusAvailable AgentPoolDetailedStatus = "Available"

  AgentPoolDetailedStatusError AgentPoolDetailedStatus = "Error"

  AgentPoolDetailedStatusProvisioning AgentPoolDetailedStatus = "Provisioning"
)

// The selection of how this agent pool is utilized, either as a system pool or a user pool. System pools run the features and critical services for the Kubernetes Cluster, while user pools are dedicated to user workloads. Every Kubernetes cluster must contain at least one system node pool with at least one node.
type AgentPoolMode string

const (
  AgentPoolModeSystem AgentPoolMode = "System"

  AgentPoolModeUser AgentPoolMode = "User"

  AgentPoolModeNotApplicable AgentPoolMode = "NotApplicable"
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

package agentpool

// AgentPool represents the agent pool of Kubernetes cluster.
type AgentPool struct{}

// AgentPoolProperties represents the properties of the Kubernetes cluster agent pool.
type AgentPoolProperties struct{}

// AgentOptions are configurations that will be applied to each agent in an agent pool.
type AgentOptions struct{}

// AgentPoolUpgradeSettings specifies the upgrade settings for an agent pool.
type AgentPoolUpgradeSettings struct{}

// AgentPoolPatchParameters represents the body of the request to patch the Kubernetes cluster agent pool.
type AgentPoolPatchParameters struct{}

// AgentPoolPatchProperties represents the properties of an agent pool that can be modified.
type AgentPoolPatchProperties struct{}

// NodePoolAdministratorConfigurationPatch represents the patching capabilities for the administrator configuration.
type NodePoolAdministratorConfigurationPatch struct{}

// AgentPoolList represents a list of Kubernetes cluster agent pools.
type AgentPoolList struct{}

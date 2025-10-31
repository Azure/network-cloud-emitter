package agentpool

import (
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/nogen/azcommonv6"
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/test/common"
)

// AgentPool represents the agent pool of Kubernetes cluster.
type AgentPool struct {
  // The list of the resource properties.
  properties AgentPoolProperties

  // The name of the Kubernetes cluster agent pool.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string

  // The extended location of the cluster associated with the resource.
  extendedLocation common.ExtendedLocation

  azcommonv6.TrackedResource
}

// AgentPoolProperties represents the properties of the Kubernetes cluster agent pool.
type AgentPoolProperties struct {
  // The administrator credentials to be used for the nodes in this agent pool.
  administratorConfiguration common.AdministratorConfiguration

  // The configurations that will be applied to each agent in this agent pool.
  agentOptions common.AgentOptions

  // The configuration of networks being attached to the agent pool for use by the workloads that run on this Kubernetes cluster.
  attachedNetworkConfiguration common.AttachedNetworkConfiguration

  // The list of availability zones of the Network Cloud cluster used for the provisioning of nodes in this agent pool. If not specified, all availability zones will be used.
  availabilityZones []string

  // The number of virtual machines that use this configuration.
  count int64

  // The current status of the agent pool.
  detailedStatus AgentPoolDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // The Kubernetes version running in this agent pool.
  kubernetesVersion string

  // The labels applied to the nodes in this agent pool.
  labels []common.KubernetesLabel

  // The selection of how this agent pool is utilized, either as a system pool or a user pool. System pools run the features and critical services for the Kubernetes Cluster, while user pools are dedicated to user workloads. Every Kubernetes cluster must contain at least one system node pool with at least one node.
  mode common.AgentPoolMode

  // The provisioning state of the agent pool.
  provisioningState AgentPoolProvisioningState

  // The taints applied to the nodes in this agent pool.
  taints []common.KubernetesLabel

  // The configuration of the agent pool.
  upgradeSettings common.AgentPoolUpgradeSettings

  // The name of the VM SKU that determines the size of resources allocated for node VMs.
  vmSkuName string
}

// AgentPoolPatchParameters represents the body of the request to patch the Kubernetes cluster agent pool.
type AgentPoolPatchParameters struct {
  // The list of the resource properties.
  properties AgentPoolPatchProperties

  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

// AgentPoolPatchProperties represents the properties of an agent pool that can be modified.
type AgentPoolPatchProperties struct {
  // The configuration of administrator credentials for the control plane nodes.
  administratorConfiguration NodePoolAdministratorConfigurationPatch

  // The number of virtual machines that use this configuration.
  count int64

  // The configuration of the agent pool.
  upgradeSettings common.AgentPoolUpgradeSettings
}

// NodePoolAdministratorConfigurationPatch represents the patching capabilities for the administrator configuration.
type NodePoolAdministratorConfigurationPatch struct {
  // SshPublicKey represents the public key used to authenticate with a resource through SSH.
  sshPublicKeys []common.SshPublicKey
}

// AgentPoolList represents a list of Kubernetes cluster agent pools.
type AgentPoolList struct {
  // The AgentPool items on this page
  value []AgentPool

  // The link to the next page of items
  nextLink string
}

package agentpool

import (
  "networkcloud/agentpool"
  "networkcloud/common"
  "networkcloud/console"
  console_2 "networkcloud/console"
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
}

// AgentPoolProperties represents the properties of the Kubernetes cluster agent pool.
type AgentPoolProperties struct {
  // The administrator credentials to be used for the nodes in this agent pool.
  administratorConfiguration common.AdministratorConfiguration

  // The configurations that will be applied to each agent in this agent pool.
  agentOptions AgentOptions

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
  labels []<Unresolved Symbol>

  // The selection of how this agent pool is utilized, either as a system pool or a user pool. System pools run the features and critical services for the Kubernetes Cluster, while user pools are dedicated to user workloads. Every Kubernetes cluster must contain at least one system node pool with at least one node.
  mode AgentPoolMode

  // The provisioning state of the agent pool.
  provisioningState AgentPoolProvisioningState

  // The taints applied to the nodes in this agent pool.
  taints []<Unresolved Symbol>

  // The configuration of the agent pool.
  upgradeSettings AgentPoolUpgradeSettings

  // The name of the VM SKU that determines the size of resources allocated for node VMs.
  vmSkuName string
}

// AgentOptions are configurations that will be applied to each agent in an agent pool.
type AgentOptions struct {
  // The number of hugepages to allocate.
  hugepagesCount int64

  // The size of the hugepages to allocate.
  hugepagesSize HugepagesSize
}

// AgentPoolUpgradeSettings specifies the upgrade settings for an agent pool.
type AgentPoolUpgradeSettings struct {
  // The maximum time in seconds that is allowed for a node drain to complete before proceeding with the upgrade of the agent pool. If not specified during creation, a value of 1800 seconds is used.
  drainTimeout int64

  // The maximum number or percentage of nodes that are surged during upgrade. This can either be set to an integer (e.g. '5') or a percentage (e.g. '50%'). If a percentage is specified, it is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded up. If not specified during creation, a value of 1 is used. One of MaxSurge and MaxUnavailable must be greater than 0.
  maxSurge string

  // The maximum number or percentage of nodes that can be unavailable during upgrade. This can either be set to an integer (e.g. '5') or a percentage (e.g. '50%'). If a percentage is specified, it is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded up. If not specified during creation, a value of 0 is used. One of MaxSurge and MaxUnavailable must be greater than 0.
  maxUnavailable string
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
  upgradeSettings AgentPoolUpgradeSettings
}

// NodePoolAdministratorConfigurationPatch represents the patching capabilities for the administrator configuration.
type NodePoolAdministratorConfigurationPatch struct {
  // SshPublicKey represents the public key used to authenticate with a resource through SSH.
  sshPublicKeys []console_2.SshPublicKey
}

// AgentPoolList represents a list of Kubernetes cluster agent pools.
type AgentPoolList struct {
  // The AgentPool items on this page
  value []AgentPool

  // The link to the next page of items
  nextLink string
}

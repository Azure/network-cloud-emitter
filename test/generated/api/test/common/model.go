package common

import "time"

// ExtendedLocation represents the Azure custom location where the resource will be created.
type ExtendedLocation struct {
  // The resource ID of the extended location on which the resource will be created.
  name string

  // The extended location type, for example, CustomLocation.
  type string
}

// AdministrativeCredentials represents the admin credentials for the device requiring password-based authentication.
type AdministrativeCredentials struct {
  // The password of the administrator of the device used during initialization.
  password string

  // The username of the administrator of the device used during initialization.
  username string
}

// ManagedResourceGroupConfiguration represents the configuration of the resource group managed by Azure.
type ManagedResourceGroupConfiguration struct {
  // The location of the managed resource group. If not specified, the location of the parent resource is chosen.
  location string

  // The name for the managed resource group. If not specified, the unique name is automatically generated.
  name string
}

// SecretRotationStatus represents the status of a secret rotation.
type SecretRotationStatus struct {
  // The maximum number of days the secret may be used before it must be changed.
  expirePeriodDays int64

  // The date and time when the secret was last changed.
  lastRotationTime time.Time

  // The number of days a secret exists before rotations will be attempted.
  rotationPeriodDays int64

  // The reference to the secret in a key vault.
  secretArchiveReference SecretArchiveReference

  // The type name used to identify the purpose of the secret.
  secretType string
}

// SecretArchiveReference represents the reference to a secret in a key vault.
type SecretArchiveReference struct {
  // The resource ID of the key vault containing the secret.
  keyVaultId string

  // The name of the secret in the key vault.
  secretName string

  // The version of the secret in the key vault.
  secretVersion string
}

// NetworkAttachment represents the single network attachment.
type NetworkAttachment struct {
  // The resource ID of the associated network attached to the virtual machine.
  // It can be one of cloudServicesNetwork, l3Network, l2Network or trunkedNetwork resources.
  attachedNetworkId string

  // The indicator of whether this is the default gateway.
  // Only one of the attached networks (including the CloudServicesNetwork attachment) for a single machine may be specified as True.
  defaultGateway DefaultGateway

  // The IP allocation mechanism for the virtual machine.
  // Dynamic and Static are only valid for l3Network which may also specify Disabled.
  // Otherwise, Disabled is the only permitted value.
  ipAllocationMethod VirtualMachineIPAllocationMethod

  // The IPv4 address of the virtual machine.
  //
  // This field is used only if the attached network has IPAllocationType of IPV4 or DualStack.
  //
  // If IPAllocationMethod is:
  // Static - this field must contain a user specified IPv4 address from within the subnet specified in the attached network.
  // Dynamic - this field is read-only, but will be populated with an address from within the subnet specified in the attached network.
  // Disabled - this field will be empty.
  ipv4Address string

  // The IPv6 address of the virtual machine.
  //
  // This field is used only if the attached network has IPAllocationType of IPV6 or DualStack.
  //
  // If IPAllocationMethod is:
  // Static - this field must contain an IPv6 address range from within the range specified in the attached network.
  // Dynamic - this field is read-only, but will be populated with an range from within the subnet specified in the attached network.
  // Disabled - this field will be empty.
  ipv6Address string

  // The MAC address of the interface for the virtual machine that corresponds to this network attachment.
  macAddress string

  // The associated network's interface name.
  // If specified, the network attachment name has a maximum length of 15 characters and must be unique to this virtual machine.
  // If the user doesn’t specify this value, the default interface name of the network resource will be used.
  // For a CloudServicesNetwork resource, this name will be ignored.
  networkAttachmentName string
}

// SshPublicKey represents the public key used to authenticate with a resource through SSH.
type SshPublicKey struct {
  // The SSH public key data.
  keyData string
}

// KeySetUser represents the properties of the user in the key set.
type KeySetUser struct {
  // The user name that will be used for access.
  azureUserName string

  // The free-form description for this user.
  description string

  // The SSH public key that will be provisioned for user access. The user is expected to have the corresponding SSH private key for logging in.
  sshPublicKey SshPublicKey

  // The user principal name (email format) used to validate this user's group membership.
  userPrincipalName string
}

// KeySetUserStatus represents the status of the key set user.
type KeySetUserStatus struct {
  // The user name that will be used for access.
  azureUserName string

  // The indicator of whether the user is currently deployed for access.
  status BareMetalMachineKeySetUserSetupStatus

  // The additional information describing the current status of this user, if any available.
  statusMessage string
}

// AdministratorConfiguration represents the administrative credentials that will be applied to the control plane and agent pool nodes in Kubernetes clusters.
type AdministratorConfiguration struct {
  // The user name for the administrator that will be applied to the operating systems that run Kubernetes nodes. If not supplied, a user name will be chosen by the service.
  adminUsername string

  // The SSH configuration for the operating systems that run the nodes in the Kubernetes cluster. In some cases, specification of public keys may be required to produce a working environment.
  sshPublicKeys []SshPublicKey
}

// AgentOptions are configurations that will be applied to each agent in an agent pool.
type AgentOptions struct {
  // The number of hugepages to allocate.
  hugepagesCount int64

  // The size of the hugepages to allocate.
  hugepagesSize HugepagesSize
}

// AttachedNetworkConfiguration represents the set of workload networks to attach to a resource.
type AttachedNetworkConfiguration struct {
  // The list of Layer 2 Networks and related configuration for attachment.
  l2Networks []L2NetworkAttachmentConfiguration

  // The list of Layer 3 Networks and related configuration for attachment.
  l3Networks []L3NetworkAttachmentConfiguration

  // The list of Trunked Networks and related configuration for attachment.
  trunkedNetworks []TrunkedNetworkAttachmentConfiguration
}

// L2NetworkAttachmentConfiguration represents the configuration of the attachment of a Layer 2 network.
type L2NetworkAttachmentConfiguration struct {
  // The resource ID of the network that is being configured for attachment.
  networkId string

  // The indicator of how this network will be utilized by the Kubernetes cluster.
  pluginType KubernetesPluginType
}

// L3NetworkAttachmentConfiguration represents the configuration of the attachment of a Layer 3 network.
type L3NetworkAttachmentConfiguration struct {
  // The indication of whether this network will or will not perform IP address management and allocate IP addresses when attached.
  ipamEnabled L3NetworkConfigurationIpamEnabled

  // The resource ID of the network that is being configured for attachment.
  networkId string

  // The indicator of how this network will be utilized by the Kubernetes cluster.
  pluginType KubernetesPluginType
}

// TrunkedNetworkAttachmentConfiguration represents the configuration of the attachment of a trunked network.
type TrunkedNetworkAttachmentConfiguration struct {
  // The resource ID of the network that is being configured for attachment.
  networkId string

  // The indicator of how this network will be utilized by the Kubernetes cluster.
  pluginType KubernetesPluginType
}

// KubernetesLabel represents a single entry for a Kubernetes label or taint such as those used on a node or pod.
type KubernetesLabel struct {
  // The name of the label or taint.
  key string

  // The value of the label or taint.
  value string
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

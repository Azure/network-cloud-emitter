package common

// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
type HybridAksPluginType string

const (
  HybridAksPluginTypeDPDK HybridAksPluginType = "DPDK"

  HybridAksPluginTypeSRIOV HybridAksPluginType = "SRIOV"

  HybridAksPluginTypeOSDevice HybridAksPluginType = "OSDevice"
)

// The indicator of whether this is the default gateway.
// Only one of the attached networks (including the CloudServicesNetwork attachment) for a single machine may be specified as True.
type DefaultGateway string

const (
  DefaultGatewayTrue DefaultGateway = "True"

  DefaultGatewayFalse DefaultGateway = "False"
)

// The IP allocation mechanism for the virtual machine.
// Dynamic and Static are only valid for l3Network which may also specify Disabled.
// Otherwise, Disabled is the only permitted value.
type VirtualMachineIPAllocationMethod string

const (
  VirtualMachineIPAllocationMethodDynamic VirtualMachineIPAllocationMethod = "Dynamic"

  VirtualMachineIPAllocationMethodStatic VirtualMachineIPAllocationMethod = "Static"

  VirtualMachineIPAllocationMethodDisabled VirtualMachineIPAllocationMethod = "Disabled"
)

// The indicator of whether the user is currently deployed for access.
type BareMetalMachineKeySetUserSetupStatus string

const (
  BareMetalMachineKeySetUserSetupStatusActive BareMetalMachineKeySetUserSetupStatus = "Active"

  BareMetalMachineKeySetUserSetupStatusInvalid BareMetalMachineKeySetUserSetupStatus = "Invalid"
)

// The size of the hugepages to allocate.
type HugepagesSize string

const (
  HugepagesSize2M HugepagesSize = "2M"

  HugepagesSize1G HugepagesSize = "1G"
)

// The indicator of how this network will be utilized by the Kubernetes cluster.
type KubernetesPluginType string

const (
  KubernetesPluginTypeDPDK KubernetesPluginType = "DPDK"

  KubernetesPluginTypeSRIOV KubernetesPluginType = "SRIOV"

  KubernetesPluginTypeOSDevice KubernetesPluginType = "OSDevice"

  KubernetesPluginTypeMACVLAN KubernetesPluginType = "MACVLAN"

  KubernetesPluginTypeIPVLAN KubernetesPluginType = "IPVLAN"
)

// The indication of whether this network will or will not perform IP address management and allocate IP addresses when attached.
type L3NetworkConfigurationIpamEnabled string

const (
  L3NetworkConfigurationIpamEnabledTrue L3NetworkConfigurationIpamEnabled = "True"

  L3NetworkConfigurationIpamEnabledFalse L3NetworkConfigurationIpamEnabled = "False"
)

// The selection of how this agent pool is utilized, either as a system pool or a user pool. System pools run the features and critical services for the Kubernetes Cluster, while user pools are dedicated to user workloads. Every Kubernetes cluster must contain at least one system node pool with at least one node.
type AgentPoolMode string

const (
  AgentPoolModeSystem AgentPoolMode = "System"

  AgentPoolModeUser AgentPoolMode = "User"

  AgentPoolModeNotApplicable AgentPoolMode = "NotApplicable"
)

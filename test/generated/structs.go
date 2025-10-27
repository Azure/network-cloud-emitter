package main

// Localized display information for this particular operation.
type OperationDisplay struct{}

// BareMetalMachineProperties represents the properties of a bare metal machine.
type BareMetalMachineProperties struct{}

// AdministrativeCredentials represents the admin credentials for the device requiring password-based authentication.
type AdministrativeCredentials struct{}

// HardwareInventory represents the hardware configuration of this machine as exposed to the customer, including information acquired from the model/sku information and from the ironic inspector.
type HardwareInventory struct{}

// HardwareInventoryNetworkInterface represents the network interface details as part of a hardware inventory.
type HardwareInventoryNetworkInterface struct{}

// Type Deprecated. Will be removed in an upcoming version. Nic represents the network interface card details.
type Nic struct{}

// Type Deprecated. Will be removed in an upcoming version. LldpNeighbor represents the details about the device connected to the NIC.
type LldpNeighbor struct{}

// HardwareValidationStatus represents the latest hardware validation details performed for this bare metal machine.
type HardwareValidationStatus struct{}

// RuntimeProtectionStatus represents the runtime protection status of the bare metal machine.
type RuntimeProtectionStatus struct{}

// SecretRotationStatus represents the status of a secret rotation.
type SecretRotationStatus struct{}

// SecretArchiveReference represents the reference to a secret in a key vault.
type SecretArchiveReference struct{}

// BareMetalMachine represents the physical machine in the rack.
type BareMetalMachine struct{}

// ExtendedLocation represents the Azure custom location where the resource will be created.
type ExtendedLocation struct{}

// BareMetalMachineList represents a list of bare metal machines.
type BareMetalMachineList struct{}

// CloudServicesNetworkProperties represents properties of the cloud services network.
type CloudServicesNetworkProperties struct{}

// EgressEndpoint represents the connection from a cloud services network to the specified endpoint for a common purpose.
type EgressEndpoint struct{}

// EndpointDependency represents the definition of an endpoint, including the domain and details.
type EndpointDependency struct{}

// Upon creation, the additional services that are provided by the platform will be allocated and
// represented in the status of this resource. All resources associated with this cloud services network will be part
// of the same layer 2 (L2) isolation domain. At least one service network must be created but may be reused across many
// virtual machines and/or Hybrid AKS clusters.
type CloudServicesNetwork struct{}

// CloudServicesNetworkList represents a list of cloud services networks.
type CloudServicesNetworkList struct{}

// ClusterManagerProperties represents the properties of a cluster manager.
type ClusterManagerProperties struct{}

// ClusterAvailableVersion represents the cluster version that the cluster manager can be asked to create and manage.
type ClusterAvailableVersion struct{}

// ManagedResourceGroupConfiguration represents the configuration of the resource group managed by Azure.
type ManagedResourceGroupConfiguration struct{}

// ClusterManager represents a control-plane to manage one or more on-premises clusters.
type ClusterManager struct{}

// ClusterManagerList represents a list of cluster manager objects.
type ClusterManagerList struct{}

// ClusterProperties represents the properties of a cluster.
type ClusterProperties struct{}

// RackDefinition represents details regarding the rack.
type RackDefinition struct{}

// BareMetalMachineConfigurationData represents configuration for the bare metal machine.
type BareMetalMachineConfigurationData struct{}

// StorageApplianceConfigurationData represents configuration for the storage application.
type StorageApplianceConfigurationData struct{}

// AnalyticsOutputSettings represents the settings for the log analytics workspace used for output of logs from this cluster.
type AnalyticsOutputSettings struct{}

// IdentitySelector represents the selection of a managed identity for use.
type IdentitySelector struct{}

// ClusterAvailableUpgradeVersion represents the various cluster upgrade parameters.
type ClusterAvailableUpgradeVersion struct{}

// ClusterCapacity represents various details regarding compute capacity.
type ClusterCapacity struct{}

// ServicePrincipalInformation represents the details of the service principal to be used by the cluster during Arc Appliance installation.
type ServicePrincipalInformation struct{}

// CommandOutputSettings represents the settings for commands run within the cluster such as bare metal machine run read-only commands.
type CommandOutputSettings struct{}

// ValidationThreshold indicates allowed machine and node hardware and deployment failures.
type ValidationThreshold struct{}

// RuntimeProtectionConfiguration represents the runtime protection configuration for the cluster.
type RuntimeProtectionConfiguration struct{}

// ClusterSecretArchive configures the key vault to archive the secrets of the cluster for later retrieval.
type ClusterSecretArchive struct{}

// SecretArchiveSettings represents the settings for the secret archive used to hold credentials for the cluster.
type SecretArchiveSettings struct{}

// ClusterUpdateStrategy represents the strategy for updating the cluster.
type ClusterUpdateStrategy struct{}

// VulnerabilityScanningSettings represents the settings for how security vulnerability scanning is applied to the cluster.
type VulnerabilityScanningSettings struct{}

// Cluster represents the on-premises Network Cloud cluster.
type Cluster struct{}

// ClusterList represents a list of clusters.
type ClusterList struct{}

// KubernetesClusterProperties represents the properties of Kubernetes cluster resource.
type KubernetesClusterProperties struct{}

// AadConfiguration represents the Azure Active Directory Integration properties.
type AadConfiguration struct{}

// AdministratorConfiguration represents the administrative credentials that will be applied to the control plane and agent pool nodes in Kubernetes clusters.
type AdministratorConfiguration struct{}

// SshPublicKey represents the public key used to authenticate with a resource through SSH.
type SshPublicKey struct{}

// AvailableUpgrade represents an upgrade available for a Kubernetes cluster.
type AvailableUpgrade struct{}

// ControlPlaneNodeConfiguration represents the selection of virtual machines and size of the control plane for a Kubernetes cluster.
type ControlPlaneNodeConfiguration struct{}

// FeatureStatus contains information regarding a Kubernetes cluster feature.
type FeatureStatus struct{}

// InitialAgentPoolConfiguration specifies the configuration of a pool of virtual machines that are initially defined with a Kubernetes cluster.
type InitialAgentPoolConfiguration struct{}

// AgentOptions are configurations that will be applied to each agent in an agent pool.
type AgentOptions struct{}

// AttachedNetworkConfiguration represents the set of workload networks to attach to a resource.
type AttachedNetworkConfiguration struct{}

// L2NetworkAttachmentConfiguration represents the configuration of the attachment of a Layer 2 network.
type L2NetworkAttachmentConfiguration struct{}

// L3NetworkAttachmentConfiguration represents the configuration of the attachment of a Layer 3 network.
type L3NetworkAttachmentConfiguration struct{}

// TrunkedNetworkAttachmentConfiguration represents the configuration of the attachment of a trunked network.
type TrunkedNetworkAttachmentConfiguration struct{}

// KubernetesLabel represents a single entry for a Kubernetes label or taint such as those used on a node or pod.
type KubernetesLabel struct{}

// AgentPoolUpgradeSettings specifies the upgrade settings for an agent pool.
type AgentPoolUpgradeSettings struct{}

// NetworkConfiguration specifies the Kubernetes cluster network related configuration.
type NetworkConfiguration struct{}

// BgpServiceLoadBalancerConfiguration represents the configuration of a BGP service load balancer.
type BgpServiceLoadBalancerConfiguration struct{}

// BgpAdvertisement represents the association of IP address pools to the communities and peers.
type BgpAdvertisement struct{}

// ServiceLoadBalancerBgpPeer represents the configuration of the BGP service load balancer for the Kubernetes cluster.
type ServiceLoadBalancerBgpPeer struct{}

// IpAddressPool represents a pool of IP addresses that can be allocated to a service.
type IpAddressPool struct{}

// L2ServiceLoadBalancerConfiguration represents the configuration of a layer 2 service load balancer.
type L2ServiceLoadBalancerConfiguration struct{}

// KubernetesClusterNode represents the details of a node in a Kubernetes cluster.
type KubernetesClusterNode struct{}

// NetworkAttachment represents the single network attachment.
type NetworkAttachment struct{}

// KubernetesCluster represents the Kubernetes cluster hosted on Network Cloud.
type KubernetesCluster struct{}

// KubernetesClusterList represents a list of Kubernetes clusters.
type KubernetesClusterList struct{}

// L2NetworkProperties represents properties of the L2 network.
type L2NetworkProperties struct{}

// L2Network represents a network that utilizes a single isolation domain set up for layer-2 resources.
type L2Network struct{}

// L2NetworkList represents a list of L2 networks.
type L2NetworkList struct{}

// L3NetworkProperties represents properties of the L3 network.
type L3NetworkProperties struct{}

// L3Network represents a network that utilizes a single isolation domain set up for layer-3 resources.
type L3Network struct{}

// L3NetworkList represents a list of L3 networks.
type L3NetworkList struct{}

// RackSkuProperties represents the properties of compute-related hardware for a rack. This supports both aggregator and compute racks.
type RackSkuProperties struct{}

// MachineSkuSlot represents a single SKU and rack slot associated with the machine.
type MachineSkuSlot struct{}

// MachineSkuProperties represents the properties of the machine SKU.
type MachineSkuProperties struct{}

// Disk represents the properties of the disk.
type MachineDisk struct{}

// NetworkInterface represents properties of the network interface.
type NetworkInterface struct{}

// StorageApplianceSkuSlot represents the single SKU and rack slot associated with the storage appliance.
type StorageApplianceSkuSlot struct{}

// StorageApplianceSkuProperties represents the properties of the storage appliance SKU.
type StorageApplianceSkuProperties struct{}

// RackSku represents the SKU information of the rack.
type RackSku struct{}

// RackSkuList represents a list of rack SKUs.
type RackSkuList struct{}

// RackProperties represents the properties of the rack.
type RackProperties struct{}

// Rack represents the hardware of the rack and is dependent upon the cluster for lifecycle.
type Rack struct{}

// RackList represents a list of racks.
type RackList struct{}

// StorageApplianceProperties represents the properties of the storage appliance.
type StorageApplianceProperties struct{}

// StorageAppliance represents on-premises Network Cloud storage appliance.
type StorageAppliance struct{}

// StorageApplianceList represents a list of storage appliances.
type StorageApplianceList struct{}

// TrunkedNetworkProperties represents properties of the trunked network.
type TrunkedNetworkProperties struct{}

// TrunkedNetwork represents a network that utilizes multiple isolation domains and specified VLANs to create a trunked network.
type TrunkedNetwork struct{}

// TrunkedNetworkList represents a list of trunked networks.
type TrunkedNetworkList struct{}

// VirtualMachineProperties represents the properties of the virtual machine.
type VirtualMachineProperties struct{}

// VirtualMachinePlacementHint represents a single scheduling hint of the virtual machine.
type VirtualMachinePlacementHint struct{}

// StorageProfile represents information about a disk.
type StorageProfile struct{}

// OsDisk represents configuration of the boot disk.
type OsDisk struct{}

// ImageRepositoryCredentials represents the credentials used to login to the image repository.
type ImageRepositoryCredentials struct{}

// VirtualMachine represents the on-premises Network Cloud virtual machine.
type VirtualMachine struct{}

// VirtualMachineList represents a list of virtual machines.
type VirtualMachineList struct{}

// VolumeProperties represents properties of the volume resource.
type VolumeProperties struct{}

// Volume represents storage made available for use by resources running on the cluster.
type Volume struct{}

// VolumeList represents a list of volumes.
type VolumeList struct{}

// The current status of an async operation.
type OperationStatusResult struct{}

// OperationStatusResultProperties represents additional properties of the operation status result.
type OperationStatusResultProperties struct{}

// BareMetalMachinePatchParameters represents the body of the request to patch bare metal machine properties.
type BareMetalMachinePatchParameters struct{}

// BareMetalMachinePatchProperties represents the properties of the bare metal machine that can be patched.
type BareMetalMachinePatchProperties struct{}

// BareMetalMachineCordonParameters represents the body of the request to evacuate workloads from node on a bare metal machine.
type BareMetalMachineCordonParameters struct{}

// BareMetalMachinePowerOffParameters represents the body of the request to power off bare metal machine.
type BareMetalMachinePowerOffParameters struct{}

// BareMetalMachineReplaceParameters represents the body of the request to physically swap a bare metal machine for another.
type BareMetalMachineReplaceParameters struct{}

// BareMetalMachineRunCommandParameters represents the body of the request to execute a script on the bare metal machine.
type BareMetalMachineRunCommandParameters struct{}

// BareMetalMachineRunDataExtractsParameters represents the body of request containing list of curated data extraction commands to run on the bare metal machine.
type BareMetalMachineRunDataExtractsParameters struct{}

// BareMetalMachineCommandSpecification represents the command and optional arguments to exercise against the bare metal machine.
type BareMetalMachineCommandSpecification struct{}

// BareMetalMachineRunReadCommandsParameters represents the body of request containing list of read-only commands to run on the bare metal machine.
type BareMetalMachineRunReadCommandsParameters struct{}

// CloudServicesNetworkPatchParameters represents the body of the request to patch the cloud services network.
type CloudServicesNetworkPatchParameters struct{}

// CloudServicesNetworkPatchProperties represents the properties of the cloud services network that can be updated using a patch request.
type CloudServicesNetworkPatchProperties struct{}

// ClusterManagerPatchParameters represents the body of the request to patch the cluster properties.
type ClusterManagerPatchParameters struct{}

// ClusterPatchParameters represents the body of the request to patch the cluster properties.
type ClusterPatchParameters struct{}

// ClusterPatchProperties represents the properties of the cluster for patching.
type ClusterPatchProperties struct{}

// VulnerabilityScanningSettingsPatch represents the settings for how security vulnerability scanning is applied to the cluster.
type VulnerabilityScanningSettingsPatch struct{}

// BareMetalMachineKeySetProperties represents the properties of bare metal machine key set.
type BareMetalMachineKeySetProperties struct{}

// KeySetUser represents the properties of the user in the key set.
type KeySetUser struct{}

// KeySetUserStatus represents the status of the key set user.
type KeySetUserStatus struct{}

// BareMetalMachineKeySet represents the bare metal machine key set.
type BareMetalMachineKeySet struct{}

// BareMetalMachineKeySetList represents a list of bare metal machine key sets.
type BareMetalMachineKeySetList struct{}

// BareMetalMachineKeySetPatchParameters represents the body of the request to patch the bare metal machine key set.
type BareMetalMachineKeySetPatchParameters struct{}

// BareMetalMachineKeySetPatchProperties represents the properties of bare metal machine key set that can be patched.
type BareMetalMachineKeySetPatchProperties struct{}

// BmcKeySetProperties represents the properties of baseboard management controller key set.
type BmcKeySetProperties struct{}

// BmcKeySet represents the baseboard management controller key set.
type BmcKeySet struct{}

// BmcKeySetList represents a list of baseboard management controller key sets.
type BmcKeySetList struct{}

// BmcKeySetPatchParameters represents the body of the request to patch the baseboard management controller key set.
type BmcKeySetPatchParameters struct{}

// BmcKeySetPatchProperties represents the properties of baseboard management controller key set that are patchable.
type BmcKeySetPatchProperties struct{}

// ClusterContinueUpdateVersionParameters represents the body of the request to continue the update of a cluster version.
type ClusterContinueUpdateVersionParameters struct{}

// ClusterDeployParameters represents the body of the request to deploy cluster.
type ClusterDeployParameters struct{}

// ClusterMetricsConfigurationProperties represents the properties of metrics configuration for the cluster.
type ClusterMetricsConfigurationProperties struct{}

// ClusterMetricsConfiguration represents the metrics configuration of an on-premises Network Cloud cluster.
type ClusterMetricsConfiguration struct{}

// ClusterMetricsConfigurationList represents a list of metrics configuration of the cluster.
type ClusterMetricsConfigurationList struct{}

// ClusterMetricsConfigurationPatchParameters represents the body of the request to patch the metrics configuration of cluster.
type ClusterMetricsConfigurationPatchParameters struct{}

// ClusterMetricsConfigurationPatchProperties represents the properties of metrics configuration for the cluster for patching.
type ClusterMetricsConfigurationPatchProperties struct{}

// ClusterScanRuntimeParameters defines the parameters for the cluster scan runtime operation.
type ClusterScanRuntimeParameters struct{}

// ClusterUpdateVersionParameters represents the body of the request to update cluster version.
type ClusterUpdateVersionParameters struct{}

// KubernetesClusterPatchParameters represents the body of the request to patch the Hybrid AKS cluster.
type KubernetesClusterPatchParameters struct{}

// KubernetesClusterPatchProperties represents the properties of the Kubernetes cluster that can be patched.
type KubernetesClusterPatchProperties struct{}

// AdministratorConfigurationPatch represents the patching capabilities for the administrator configuration.
type AdministratorConfigurationPatch struct{}

// ControlPlaneNodePatchConfiguration represents the properties of the control plane that can be patched for this Kubernetes cluster.
type ControlPlaneNodePatchConfiguration struct{}

// AgentPoolProperties represents the properties of the Kubernetes cluster agent pool.
type AgentPoolProperties struct{}

// AgentPool represents the agent pool of Kubernetes cluster.
type AgentPool struct{}

// AgentPoolList represents a list of Kubernetes cluster agent pools.
type AgentPoolList struct{}

// AgentPoolPatchParameters represents the body of the request to patch the Kubernetes cluster agent pool.
type AgentPoolPatchParameters struct{}

// AgentPoolPatchProperties represents the properties of an agent pool that can be modified.
type AgentPoolPatchProperties struct{}

// NodePoolAdministratorConfigurationPatch represents the patching capabilities for the administrator configuration.
type NodePoolAdministratorConfigurationPatch struct{}

// KubernetesClusterFeatureProperties represents the properties of a Kubernetes cluster feature.
type KubernetesClusterFeatureProperties struct{}

// StringKeyValuePair represents a single entry in a mapping of keys to values.
type StringKeyValuePair struct{}

// KubernetesClusterFeature represents the feature of a Kubernetes cluster.
type KubernetesClusterFeature struct{}

// KubernetesClusterFeatureList represents the list of Kubernetes cluster feature resources.
type KubernetesClusterFeatureList struct{}

// KubernetesClusterFeaturePatchParameters represents the body of the request to patch the Kubernetes cluster feature.
type KubernetesClusterFeaturePatchParameters struct{}

// KubernetesClusterFeaturePatchProperties represents the Kubernetes cluster feature properties for patching.
type KubernetesClusterFeaturePatchProperties struct{}

// KubernetesClusterRestartNodeParameters represents the body of the request to restart the node of a Kubernetes cluster.
type KubernetesClusterRestartNodeParameters struct{}

// L2NetworkPatchParameters represents the body of the request to patch the L2 network.
type L2NetworkPatchParameters struct{}

// L3NetworkPatchParameters represents the body of the request to patch the cloud services network.
type L3NetworkPatchParameters struct{}

// RackPatchParameters represents the body of the request to patch the rack properties.
type RackPatchParameters struct{}

// RacksPatchProperties represents the properties of the rack during patching.
type RacksPatchProperties struct{}

// StorageAppliancePatchParameters represents the body of the request to patch storage appliance properties.
type StorageAppliancePatchParameters struct{}

// StorageAppliancePatchProperties represents the properties of the storage appliance that can be patched.
type StorageAppliancePatchProperties struct{}

// StorageApplianceEnableRemoteVendorManagementParameters represents the body of the request to enable remote vendor management of a storage appliance.
type StorageApplianceEnableRemoteVendorManagementParameters struct{}

// TrunkedNetworkPatchParameters represents the body of the request to patch the Trunked network.
type TrunkedNetworkPatchParameters struct{}

// VirtualMachinePatchParameters represents the body of the request to patch the virtual machine.
type VirtualMachinePatchParameters struct{}

// VirtualMachinePatchProperties represents the properties of the virtual machine that can be patched.
type VirtualMachinePatchProperties struct{}

// ConsoleProperties represents the properties of the virtual machine console.
type ConsoleProperties struct{}

// Console represents the console of an on-premises Network Cloud virtual machine.
type Console struct{}

// ConsoleList represents a list of virtual machine consoles.
type ConsoleList struct{}

// ConsolePatchParameters represents the body of the request to patch the virtual machine console.
type ConsolePatchParameters struct{}

// ConsolePatchProperties represents the properties of the virtual machine console that can be patched.
type ConsolePatchProperties struct{}

// VirtualMachinePowerOffParameters represents the body of the request to power off virtual machine.
type VirtualMachinePowerOffParameters struct{}

// VolumePatchParameters represents the body of the request to patch the volume resource.
type VolumePatchParameters struct{}

// AgentPoolConfiguration specifies the configuration of a pool of nodes.
type AgentPoolConfiguration struct{}

// TagsParameter represents the resource tags.
type TagsParameter struct{}

package kubernetescluster

import (
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/nogen/azcommonv6"
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/test/common"
)

// KubernetesCluster represents the Kubernetes cluster hosted on Network Cloud.
type KubernetesCluster struct {
  // The list of the resource properties.
  properties KubernetesClusterProperties

  // The name of the Kubernetes cluster.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string

  // The extended location of the cluster associated with the resource.
  extendedLocation common.ExtendedLocation

  azcommonv6.TrackedResource
}

// KubernetesClusterProperties represents the properties of Kubernetes cluster resource.
type KubernetesClusterProperties struct {
  // The Azure Active Directory Integration properties.
  aadConfiguration AadConfiguration

  // The administrative credentials that will be applied to the control plane and agent pool nodes that do not specify their own values.
  administratorConfiguration common.AdministratorConfiguration

  // The full list of network resource IDs that are attached to this cluster, including those attached only to specific agent pools.
  attachedNetworkIds []string

  // The list of versions that this Kubernetes cluster can be upgraded to.
  availableUpgrades []AvailableUpgrade

  // The resource ID of the Network Cloud cluster.
  clusterId string

  // The resource ID of the connected cluster set up when this Kubernetes cluster is created.
  connectedClusterId string

  // The current running version of Kubernetes on the control plane.
  controlPlaneKubernetesVersion string

  // The defining characteristics of the control plane for this Kubernetes Cluster.
  controlPlaneNodeConfiguration ControlPlaneNodeConfiguration

  // The current status of the Kubernetes cluster.
  detailedStatus KubernetesClusterDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // The current feature settings.
  featureStatuses []FeatureStatus

  // The agent pools that are created with this Kubernetes cluster for running critical system services and workloads. This data in this field is only used during creation, and the field will be empty following the creation of the Kubernetes Cluster. After creation, the management of agent pools is done using the agentPools sub-resource.
  initialAgentPoolConfigurations []InitialAgentPoolConfiguration

  // The Kubernetes version for this cluster.
  kubernetesVersion string

  // The configuration of the managed resource group associated with the resource.
  managedResourceGroupConfiguration common.ManagedResourceGroupConfiguration

  // The configuration of the Kubernetes cluster networking, including the attachment of networks that span the cluster.
  networkConfiguration NetworkConfiguration

  // The details of the nodes in this cluster.
  nodes []KubernetesClusterNode

  // The provisioning state of the Kubernetes cluster resource.
  provisioningState KubernetesClusterProvisioningState
}

// AadConfiguration represents the Azure Active Directory Integration properties.
type AadConfiguration struct {
  // The list of Azure Active Directory group object IDs that will have an administrative role on the Kubernetes cluster.
  adminGroupObjectIds []string
}

// AvailableUpgrade represents an upgrade available for a Kubernetes cluster.
type AvailableUpgrade struct {
  // The version lifecycle indicator.
  availabilityLifecycle AvailabilityLifecycle

  // The version available for upgrading.
  version string
}

// ControlPlaneNodeConfiguration represents the selection of virtual machines and size of the control plane for a Kubernetes cluster.
type ControlPlaneNodeConfiguration struct {
  // The administrator credentials to be used for the nodes in the control plane.
  administratorConfiguration common.AdministratorConfiguration

  // The list of availability zones of the Network Cloud cluster to be used for the provisioning of nodes in the control plane. If not specified, all availability zones will be used.
  availabilityZones []string

  // The number of virtual machines that use this configuration.
  count int64

  // The name of the VM SKU supplied during creation.
  vmSkuName string
}

// FeatureStatus contains information regarding a Kubernetes cluster feature.
type FeatureStatus struct {
  // The status representing the state of this feature.
  detailedStatus FeatureDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // The name of the feature.
  name string

  // The version of the feature.
  version string
}

// InitialAgentPoolConfiguration specifies the configuration of a pool of virtual machines that are initially defined with a Kubernetes cluster.
type InitialAgentPoolConfiguration struct {
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

  // The labels applied to the nodes in this agent pool.
  labels []common.KubernetesLabel

  // The selection of how this agent pool is utilized, either as a system pool or a user pool. System pools run the features and critical services for the Kubernetes Cluster, while user pools are dedicated to user workloads. Every Kubernetes cluster must contain at least one system node pool with at least one node.
  mode common.AgentPoolMode

  // The name that will be used for the agent pool resource representing this agent pool.
  name string

  // The taints applied to the nodes in this agent pool.
  taints []common.KubernetesLabel

  // The configuration of the agent pool.
  upgradeSettings common.AgentPoolUpgradeSettings

  // The name of the VM SKU that determines the size of resources allocated for node VMs.
  vmSkuName string
}

// NetworkConfiguration specifies the Kubernetes cluster network related configuration.
type NetworkConfiguration struct {
  // The configuration of networks being attached to the cluster for use by the workloads that run on this Kubernetes cluster.
  attachedNetworkConfiguration common.AttachedNetworkConfiguration

  // The configuration of the BGP service load balancer for this Kubernetes cluster. A maximum of one service load balancer may be specified, either Layer 2 or BGP.
  bgpServiceLoadBalancerConfiguration BgpServiceLoadBalancerConfiguration

  // The resource ID of the associated Cloud Services network.
  cloudServicesNetworkId string

  // The resource ID of the Layer 3 network that is used for creation of the Container Networking Interface network.
  cniNetworkId string

  // The IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in service CIDR.
  dnsServiceIp string

  // The configuration of the Layer 2 service load balancer for this Kubernetes cluster. A maximum of one service load balancer may be specified, either Layer 2 or BGP.
  l2ServiceLoadBalancerConfiguration L2ServiceLoadBalancerConfiguration

  // The CIDR notation IP ranges from which to assign pod IPs. One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is expected for dual-stack networking.
  podCidrs []string

  // The CIDR notation IP ranges from which to assign service IPs. One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is expected for dual-stack networking.
  serviceCidrs []string
}

// BgpServiceLoadBalancerConfiguration represents the configuration of a BGP service load balancer.
type BgpServiceLoadBalancerConfiguration struct {
  // The association of IP address pools to the communities and peers, allowing for announcement of IPs.
  bgpAdvertisements []BgpAdvertisement

  // The list of additional BgpPeer entities that the Kubernetes cluster will peer with. All peering must be explicitly defined.
  bgpPeers []ServiceLoadBalancerBgpPeer

  // The indicator to specify if the load balancer peers with the network fabric.
  fabricPeeringEnabled FabricPeeringEnabled

  // The list of pools of IP addresses that can be allocated to load balancer services.
  ipAddressPools []IpAddressPool
}

// BgpAdvertisement represents the association of IP address pools to the communities and peers.
type BgpAdvertisement struct {
  // The indicator of if this advertisement is also made to the network fabric associated with the Network Cloud Cluster. This field is ignored if fabricPeeringEnabled is set to False.
  advertiseToFabric AdvertiseToFabric

  // The names of the BGP communities to be associated with the announcement, utilizing a BGP community string in 1234:1234 format.
  communities []string

  // The names of the IP address pools associated with this announcement.
  ipAddressPools []string

  // The names of the BGP peers to limit this advertisement to. If no values are specified, all BGP peers will receive this advertisement.
  peers []string
}

// ServiceLoadBalancerBgpPeer represents the configuration of the BGP service load balancer for the Kubernetes cluster.
type ServiceLoadBalancerBgpPeer struct {
  // The indicator of BFD enablement for this BgpPeer.
  bfdEnabled BfdEnabled

  // The indicator to enable multi-hop peering support.
  bgpMultiHop BgpMultiHop

  // Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The requested BGP hold time value. This field uses ISO 8601 duration format, for example P1H.
  holdTime string

  // Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The requested BGP keepalive time value. This field uses ISO 8601 duration format, for example P1H.
  keepAliveTime string

  // The autonomous system number used for the local end of the BGP session.
  myAsn int64

  // The name used to identify this BGP peer for association with a BGP advertisement.
  name string

  // The authentication password for routers enforcing TCP MD5 authenticated sessions.
  password string

  // The IPv4 or IPv6 address used to connect this BGP session.
  peerAddress string

  // The autonomous system number expected from the remote end of the BGP session.
  peerAsn int64

  // The port used to connect this BGP session.
  peerPort int64
}

// IpAddressPool represents a pool of IP addresses that can be allocated to a service.
type IpAddressPool struct {
  // The list of IP address ranges. Each range can be a either a subnet in CIDR format or an explicit start-end range of IP addresses. For a BGP service load balancer configuration, only CIDR format is supported and excludes /32 (IPv4) and /128 (IPv6) prefixes.
  addresses []string

  // The indicator to determine if automatic allocation from the pool should occur.
  autoAssign BfdEnabled

  // The name used to identify this IP address pool for association with a BGP advertisement.
  name string

  // The indicator to prevent the use of IP addresses ending with .0 and .255 for this pool. Enabling this option will only use IP addresses between .1 and .254 inclusive.
  onlyUseHostIps BfdEnabled
}

// L2ServiceLoadBalancerConfiguration represents the configuration of a layer 2 service load balancer.
type L2ServiceLoadBalancerConfiguration struct {
  // The list of pools of IP addresses that can be allocated to load balancer services.
  ipAddressPools []IpAddressPool
}

// KubernetesClusterNode represents the details of a node in a Kubernetes cluster.
type KubernetesClusterNode struct {
  // The resource ID of the agent pool that this node belongs to. This value is not represented on control plane nodes.
  agentPoolId string

  // The availability zone this node is running within.
  availabilityZone string

  // The resource ID of the bare metal machine that hosts this node.
  bareMetalMachineId string

  // The number of CPU cores configured for this node, derived from the VM SKU specified.
  cpuCores int64

  // The detailed state of this node.
  detailedStatus KubernetesClusterNodeDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // The size of the disk configured for this node. Allocations are measured in gibibytes.
  diskSizeGB int64

  // The machine image used to deploy this node.
  image string

  // The currently running version of Kubernetes and bundled features running on this node.
  kubernetesVersion string

  // The list of labels on this node that have been assigned to the agent pool containing this node.
  labels []common.KubernetesLabel

  // The amount of memory configured for this node, derived from the vm SKU specified. Allocations are measured in gibibytes.
  memorySizeGB int64

  // The mode of the agent pool containing this node. Not applicable for control plane nodes.
  mode common.AgentPoolMode

  // The name of this node, as realized in the Kubernetes cluster.
  name string

  // The NetworkAttachments made to this node.
  networkAttachments []common.NetworkAttachment

  // The power state of this node.
  powerState KubernetesNodePowerState

  // The role of this node in the cluster.
  role KubernetesNodeRole

  // The list of taints that have been assigned to the agent pool containing this node.
  taints []common.KubernetesLabel

  // The VM SKU name that was used to create this cluster node.
  vmSkuName string
}

// KubernetesClusterPatchParameters represents the body of the request to patch the Hybrid AKS cluster.
type KubernetesClusterPatchParameters struct {
  // The list of the resource properties.
  properties KubernetesClusterPatchProperties

  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

// KubernetesClusterPatchProperties represents the properties of the Kubernetes cluster that can be patched.
type KubernetesClusterPatchProperties struct {
  // The configuration of the default administrator credentials.
  administratorConfiguration AdministratorConfigurationPatch

  // The defining characteristics of the control plane that can be patched for this Kubernetes cluster.
  controlPlaneNodeConfiguration ControlPlaneNodePatchConfiguration

  // The Kubernetes version for this cluster.
  kubernetesVersion string
}

// AdministratorConfigurationPatch represents the patching capabilities for the administrator configuration.
type AdministratorConfigurationPatch struct {
  // SshPublicKey represents the public key used to authenticate with a resource through SSH.
  sshPublicKeys []common.SshPublicKey
}

// ControlPlaneNodePatchConfiguration represents the properties of the control plane that can be patched for this Kubernetes cluster.
type ControlPlaneNodePatchConfiguration struct {
  // The configuration of administrator credentials for the control plane nodes.
  administratorConfiguration AdministratorConfigurationPatch

  // The number of virtual machines that use this configuration.
  count int64
}

// KubernetesClusterRestartNodeParameters represents the body of the request to restart the node of a Kubernetes cluster.
type KubernetesClusterRestartNodeParameters struct {
  // The name of the node to restart.
  nodeName string
}

// KubernetesClusterList represents a list of Kubernetes clusters.
type KubernetesClusterList struct {
  // The KubernetesCluster items on this page
  value []KubernetesCluster

  // The link to the next page of items
  nextLink string
}

package kubernetescluster

import (
  "networkcloud/common"
  "networkcloud/console"
  console_2 "networkcloud/console"
  "networkcloud/kubernetescluster"
)

// KubernetesClusterList represents a list of Kubernetes clusters.
type KubernetesClusterList struct {
  // The KubernetesCluster items on this page
  value []KubernetesCluster

  // The link to the next page of items
  nextLink string
}

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
  availableUpgrades []<Unresolved Symbol>

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
  featureStatuses []<Unresolved Symbol>

  // The agent pools that are created with this Kubernetes cluster for running critical system services and workloads. This data in this field is only used during creation, and the field will be empty following the creation of the Kubernetes Cluster. After creation, the management of agent pools is done using the agentPools sub-resource.
  initialAgentPoolConfigurations []<Unresolved Symbol>

  // The Kubernetes version for this cluster.
  kubernetesVersion string

  // The configuration of the managed resource group associated with the resource.
  managedResourceGroupConfiguration common.ManagedResourceGroupConfiguration

  // The configuration of the Kubernetes cluster networking, including the attachment of networks that span the cluster.
  networkConfiguration NetworkConfiguration

  // The details of the nodes in this cluster.
  nodes []<Unresolved Symbol>

  // The provisioning state of the Kubernetes cluster resource.
  provisioningState KubernetesClusterProvisioningState
}

// AadConfiguration represents the Azure Active Directory Integration properties.
type AadConfiguration struct {
  // The list of Azure Active Directory group object IDs that will have an administrative role on the Kubernetes cluster.
  adminGroupObjectIds []string
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
  bgpAdvertisements []<Unresolved Symbol>

  // The list of additional BgpPeer entities that the Kubernetes cluster will peer with. All peering must be explicitly defined.
  bgpPeers []<Unresolved Symbol>

  // The indicator to specify if the load balancer peers with the network fabric.
  fabricPeeringEnabled FabricPeeringEnabled

  // The list of pools of IP addresses that can be allocated to load balancer services.
  ipAddressPools []<Unresolved Symbol>
}

// L2ServiceLoadBalancerConfiguration represents the configuration of a layer 2 service load balancer.
type L2ServiceLoadBalancerConfiguration struct {
  // The list of pools of IP addresses that can be allocated to load balancer services.
  ipAddressPools []<Unresolved Symbol>
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
  sshPublicKeys []console_2.SshPublicKey
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

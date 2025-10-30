package kubernetescluster

// KubernetesClusterList represents a list of Kubernetes clusters.
type KubernetesClusterList struct{}

// KubernetesCluster represents the Kubernetes cluster hosted on Network Cloud.
type KubernetesCluster struct{}

// KubernetesClusterProperties represents the properties of Kubernetes cluster resource.
type KubernetesClusterProperties struct{}

// AadConfiguration represents the Azure Active Directory Integration properties.
type AadConfiguration struct{}

// ControlPlaneNodeConfiguration represents the selection of virtual machines and size of the control plane for a Kubernetes cluster.
type ControlPlaneNodeConfiguration struct{}

// NetworkConfiguration specifies the Kubernetes cluster network related configuration.
type NetworkConfiguration struct{}

// BgpServiceLoadBalancerConfiguration represents the configuration of a BGP service load balancer.
type BgpServiceLoadBalancerConfiguration struct{}

// L2ServiceLoadBalancerConfiguration represents the configuration of a layer 2 service load balancer.
type L2ServiceLoadBalancerConfiguration struct{}

// KubernetesClusterPatchParameters represents the body of the request to patch the Hybrid AKS cluster.
type KubernetesClusterPatchParameters struct{}

// KubernetesClusterPatchProperties represents the properties of the Kubernetes cluster that can be patched.
type KubernetesClusterPatchProperties struct{}

// AdministratorConfigurationPatch represents the patching capabilities for the administrator configuration.
type AdministratorConfigurationPatch struct{}

// ControlPlaneNodePatchConfiguration represents the properties of the control plane that can be patched for this Kubernetes cluster.
type ControlPlaneNodePatchConfiguration struct{}

// KubernetesClusterRestartNodeParameters represents the body of the request to restart the node of a Kubernetes cluster.
type KubernetesClusterRestartNodeParameters struct{}

package kubernetescluster

// The current status of the Kubernetes cluster.
type KubernetesClusterDetailedStatus string

const (
  KubernetesClusterDetailedStatusAvailable KubernetesClusterDetailedStatus = "Available"

  KubernetesClusterDetailedStatusError KubernetesClusterDetailedStatus = "Error"

  KubernetesClusterDetailedStatusProvisioning KubernetesClusterDetailedStatus = "Provisioning"
)

// The indicator to specify if the load balancer peers with the network fabric.
type FabricPeeringEnabled string

const (
  FabricPeeringEnabledTrue FabricPeeringEnabled = "True"

  FabricPeeringEnabledFalse FabricPeeringEnabled = "False"
)

// The provisioning state of the Kubernetes cluster resource.
type KubernetesClusterProvisioningState string

const (
  KubernetesClusterProvisioningStateSucceeded KubernetesClusterProvisioningState = "Succeeded"

  KubernetesClusterProvisioningStateFailed KubernetesClusterProvisioningState = "Failed"

  KubernetesClusterProvisioningStateCanceled KubernetesClusterProvisioningState = "Canceled"

  KubernetesClusterProvisioningStateAccepted KubernetesClusterProvisioningState = "Accepted"

  KubernetesClusterProvisioningStateInProgress KubernetesClusterProvisioningState = "InProgress"

  KubernetesClusterProvisioningStateCreated KubernetesClusterProvisioningState = "Created"

  KubernetesClusterProvisioningStateUpdating KubernetesClusterProvisioningState = "Updating"

  KubernetesClusterProvisioningStateDeleting KubernetesClusterProvisioningState = "Deleting"
)

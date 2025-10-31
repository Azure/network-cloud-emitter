package kubernetescluster

// The version lifecycle indicator.
type AvailabilityLifecycle string

const (
  AvailabilityLifecyclePreview AvailabilityLifecycle = "Preview"

  AvailabilityLifecycleGenerallyAvailable AvailabilityLifecycle = "GenerallyAvailable"
)

// The current status of the Kubernetes cluster.
type KubernetesClusterDetailedStatus string

const (
  KubernetesClusterDetailedStatusAvailable KubernetesClusterDetailedStatus = "Available"

  KubernetesClusterDetailedStatusError KubernetesClusterDetailedStatus = "Error"

  KubernetesClusterDetailedStatusProvisioning KubernetesClusterDetailedStatus = "Provisioning"
)

// The status representing the state of this feature.
type FeatureDetailedStatus string

const (
  FeatureDetailedStatusRunning FeatureDetailedStatus = "Running"

  FeatureDetailedStatusFailed FeatureDetailedStatus = "Failed"

  FeatureDetailedStatusUnknown FeatureDetailedStatus = "Unknown"
)

// The indicator of if this advertisement is also made to the network fabric associated with the Network Cloud Cluster. This field is ignored if fabricPeeringEnabled is set to False.
type AdvertiseToFabric string

const (
  AdvertiseToFabricTrue AdvertiseToFabric = "True"

  AdvertiseToFabricFalse AdvertiseToFabric = "False"
)

// The indicator to determine if automatic allocation from the pool should occur.
type BfdEnabled string

const (
  BfdEnabledTrue BfdEnabled = "True"

  BfdEnabledFalse BfdEnabled = "False"
)

// The indicator to enable multi-hop peering support.
type BgpMultiHop string

const (
  BgpMultiHopTrue BgpMultiHop = "True"

  BgpMultiHopFalse BgpMultiHop = "False"
)

// The indicator to specify if the load balancer peers with the network fabric.
type FabricPeeringEnabled string

const (
  FabricPeeringEnabledTrue FabricPeeringEnabled = "True"

  FabricPeeringEnabledFalse FabricPeeringEnabled = "False"
)

// The detailed state of this node.
type KubernetesClusterNodeDetailedStatus string

const (
  KubernetesClusterNodeDetailedStatusAvailable KubernetesClusterNodeDetailedStatus = "Available"

  KubernetesClusterNodeDetailedStatusError KubernetesClusterNodeDetailedStatus = "Error"

  KubernetesClusterNodeDetailedStatusProvisioning KubernetesClusterNodeDetailedStatus = "Provisioning"

  KubernetesClusterNodeDetailedStatusRunning KubernetesClusterNodeDetailedStatus = "Running"

  KubernetesClusterNodeDetailedStatusScheduling KubernetesClusterNodeDetailedStatus = "Scheduling"

  KubernetesClusterNodeDetailedStatusStopped KubernetesClusterNodeDetailedStatus = "Stopped"

  KubernetesClusterNodeDetailedStatusTerminating KubernetesClusterNodeDetailedStatus = "Terminating"

  KubernetesClusterNodeDetailedStatusUnknown KubernetesClusterNodeDetailedStatus = "Unknown"
)

// The power state of this node.
type KubernetesNodePowerState string

const (
  KubernetesNodePowerStateOn KubernetesNodePowerState = "On"

  KubernetesNodePowerStateOff KubernetesNodePowerState = "Off"

  KubernetesNodePowerStateUnknown KubernetesNodePowerState = "Unknown"
)

// The role of this node in the cluster.
type KubernetesNodeRole string

const (
  KubernetesNodeRoleControlPlane KubernetesNodeRole = "ControlPlane"

  KubernetesNodeRoleWorker KubernetesNodeRole = "Worker"
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

package kubernetesclusterfeature

// The lifecycle indicator of the feature.
type KubernetesClusterFeatureAvailabilityLifecycle string

const (
  KubernetesClusterFeatureAvailabilityLifecyclePreview KubernetesClusterFeatureAvailabilityLifecycle = "Preview"

  KubernetesClusterFeatureAvailabilityLifecycleGenerallyAvailable KubernetesClusterFeatureAvailabilityLifecycle = "GenerallyAvailable"
)

// The detailed status of the feature.
type KubernetesClusterFeatureDetailedStatus string

const (
  KubernetesClusterFeatureDetailedStatusError KubernetesClusterFeatureDetailedStatus = "Error"

  KubernetesClusterFeatureDetailedStatusProvisioning KubernetesClusterFeatureDetailedStatus = "Provisioning"

  KubernetesClusterFeatureDetailedStatusInstalled KubernetesClusterFeatureDetailedStatus = "Installed"
)

// The provisioning state of the Kubernetes cluster feature.
type KubernetesClusterFeatureProvisioningState string

const (
  KubernetesClusterFeatureProvisioningStateAccepted KubernetesClusterFeatureProvisioningState = "Accepted"

  KubernetesClusterFeatureProvisioningStateCanceled KubernetesClusterFeatureProvisioningState = "Canceled"

  KubernetesClusterFeatureProvisioningStateDeleting KubernetesClusterFeatureProvisioningState = "Deleting"

  KubernetesClusterFeatureProvisioningStateFailed KubernetesClusterFeatureProvisioningState = "Failed"

  KubernetesClusterFeatureProvisioningStateSucceeded KubernetesClusterFeatureProvisioningState = "Succeeded"

  KubernetesClusterFeatureProvisioningStateUpdating KubernetesClusterFeatureProvisioningState = "Updating"
)

// The indicator of if the feature is required or optional. Optional features may be deleted by the user, while required features are managed with the kubernetes cluster lifecycle.
type KubernetesClusterFeatureRequired string

const (
  KubernetesClusterFeatureRequiredTrue KubernetesClusterFeatureRequired = "True"

  KubernetesClusterFeatureRequiredFalse KubernetesClusterFeatureRequired = "False"
)

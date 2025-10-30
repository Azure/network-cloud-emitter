package kubernetesclusterfeature

import "networkcloud/kubernetesclusterfeature"

// KubernetesClusterFeature represents the feature of a Kubernetes cluster.
type KubernetesClusterFeature struct {
  // The list of the resource properties.
  properties KubernetesClusterFeatureProperties

  // The name of the feature.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string
}

// KubernetesClusterFeatureProperties represents the properties of a Kubernetes cluster feature.
type KubernetesClusterFeatureProperties struct {
  // The lifecycle indicator of the feature.
  availabilityLifecycle KubernetesClusterFeatureAvailabilityLifecycle

  // The detailed status of the feature.
  detailedStatus KubernetesClusterFeatureDetailedStatus

  // The descriptive message for the detailed status of the feature.
  detailedStatusMessage string

  // The configured options for the feature.
  options []<Unresolved Symbol>

  // The provisioning state of the Kubernetes cluster feature.
  provisioningState KubernetesClusterFeatureProvisioningState

  // The indicator of if the feature is required or optional. Optional features may be deleted by the user, while required features are managed with the kubernetes cluster lifecycle.
  required KubernetesClusterFeatureRequired

  // The version of the feature.
  version string
}

// KubernetesClusterFeaturePatchParameters represents the body of the request to patch the Kubernetes cluster feature.
type KubernetesClusterFeaturePatchParameters struct {
  // The list of the resource properties.
  properties KubernetesClusterFeaturePatchProperties

  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

// KubernetesClusterFeaturePatchProperties represents the Kubernetes cluster feature properties for patching.
type KubernetesClusterFeaturePatchProperties struct {
  // The configured options for the feature.
  options []<Unresolved Symbol>
}

// KubernetesClusterFeatureList represents the list of Kubernetes cluster feature resources.
type KubernetesClusterFeatureList struct {
  // The KubernetesClusterFeature items on this page
  value []KubernetesClusterFeature

  // The link to the next page of items
  nextLink string
}

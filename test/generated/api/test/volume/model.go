package volume

import (
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/nogen/azcommonv6"
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/test/common"
)

// VolumeList represents a list of volumes.
type VolumeList struct {
  // The Volume items on this page
  value []Volume

  // The link to the next page of items
  nextLink string
}

// Volume represents storage made available for use by resources running on the cluster.
type Volume struct {
  // The list of the resource properties.
  properties VolumeProperties

  // The name of the volume.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string

  // The extended location of the cluster associated with the resource.
  extendedLocation common.ExtendedLocation

  azcommonv6.TrackedResource
}

// VolumeProperties represents properties of the volume resource.
type VolumeProperties struct {
  // The list of resource IDs that attach the volume. It may include virtual machines and Hybrid AKS clusters.
  attachedTo []string

  // The more detailed status of the volume.
  detailedStatus VolumeDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // The provisioning state of the volume.
  provisioningState VolumeProvisioningState

  // The unique identifier of the volume.
  serialNumber string

  // The size of the allocation for this volume in Mebibytes.
  sizeMiB int64
}

// VolumePatchParameters represents the body of the request to patch the volume resource.
type VolumePatchParameters struct {
  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

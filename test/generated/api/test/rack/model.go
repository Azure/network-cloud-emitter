package rack

import (
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/nogen/azcommonv6"
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/test/common"
)

// RackList represents a list of racks.
type RackList struct {
  // The Rack items on this page
  value []Rack

  // The link to the next page of items
  nextLink string
}

// Rack represents the hardware of the rack and is dependent upon the cluster for lifecycle.
type Rack struct {
  // The list of the resource properties.
  properties RackProperties

  // The name of the rack.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string

  // The extended location of the cluster associated with the resource.
  extendedLocation common.ExtendedLocation

  azcommonv6.TrackedResource
}

// RackProperties represents the properties of the rack.
type RackProperties struct {
  // The value that will be used for machines in this rack to represent the availability zones that can be referenced by Hybrid AKS Clusters for node arrangement.
  availabilityZone string

  // The resource ID of the cluster the rack is created for. This value is set when the rack is created by the cluster.
  clusterId string

  // The more detailed status of the rack.
  detailedStatus RackDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // The provisioning state of the rack resource.
  provisioningState RackProvisioningState

  // The free-form description of the rack location. (e.g. “DTN Datacenter, Floor 3, Isle 9, Rack 2B”)
  rackLocation string

  // The unique identifier for the rack within Network Cloud cluster. An alternate unique alphanumeric value other than a serial number may be provided if desired.
  rackSerialNumber string

  // The SKU for the rack.
  rackSkuId string
}

// RackPatchParameters represents the body of the request to patch the rack properties.
type RackPatchParameters struct {
  // The list of the resource properties.
  properties RacksPatchProperties

  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

// RacksPatchProperties represents the properties of the rack during patching.
type RacksPatchProperties struct {
  // The free-form description of the rack location. (e.g. “DTN Datacenter, Floor 3, Isle 9, Rack 2B”)
  rackLocation string

  // The globally unique identifier for the rack.
  rackSerialNumber string
}

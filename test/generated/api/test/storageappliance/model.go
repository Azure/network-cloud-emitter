package storageappliance

import (
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/nogen/azcommonv6"
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/test/common"
)

// StorageAppliance represents on-premises Network Cloud storage appliance.
type StorageAppliance struct {
  // The list of the resource properties.
  properties StorageApplianceProperties

  // The name of the storage appliance.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string

  // The extended location of the cluster associated with the resource.
  extendedLocation common.ExtendedLocation

  azcommonv6.TrackedResource
}

// StorageApplianceProperties represents the properties of the storage appliance.
type StorageApplianceProperties struct {
  // The credentials of the administrative interface on this storage appliance.
  administratorCredentials common.AdministrativeCredentials

  // The total capacity of the storage appliance. Measured in GiB.
  capacity int64

  // The amount of storage consumed.
  capacityUsed int64

  // The resource ID of the cluster this storage appliance is associated with. Measured in GiB.
  clusterId string

  // The detailed status of the storage appliance.
  detailedStatus StorageApplianceDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // The endpoint for the management interface of the storage appliance.
  managementIpv4Address string

  // The manufacturer of the storage appliance.
  manufacturer string

  // The model of the storage appliance.
  model string

  // The provisioning state of the storage appliance.
  provisioningState StorageApplianceProvisioningState

  // The resource ID of the rack where this storage appliance resides.
  rackId string

  // The slot the storage appliance is in the rack based on the BOM configuration.
  rackSlot int64

  // The indicator of whether the storage appliance supports remote vendor management.
  remoteVendorManagementFeature RemoteVendorManagementFeature

  // The indicator of whether the remote vendor management feature is enabled or disabled, or unsupported if it is an unsupported feature.
  remoteVendorManagementStatus RemoteVendorManagementStatus

  // The list of statuses that represent secret rotation activity.
  secretRotationStatus []common.SecretRotationStatus

  // The serial number for the storage appliance.
  serialNumber string

  // The SKU for the storage appliance.
  storageApplianceSkuId string

  // The version of the storage appliance.
  version string
}

// StorageAppliancePatchParameters represents the body of the request to patch storage appliance properties.
type StorageAppliancePatchParameters struct {
  // The list of the resource properties.
  properties StorageAppliancePatchProperties

  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

// StorageAppliancePatchProperties represents the properties of the storage appliance that can be patched.
type StorageAppliancePatchProperties struct {
  // The serial number for the storage appliance.
  serialNumber string
}

// StorageApplianceEnableRemoteVendorManagementParameters represents the body of the request to enable remote vendor management of a storage appliance.
type StorageApplianceEnableRemoteVendorManagementParameters struct {
  // Field Deprecated. This field is not used and will be rejected if provided. The list of IPv4 subnets (in CIDR format), IPv6 subnets (in CIDR format), or hostnames that the storage appliance needs accessible in order to turn on the remote vendor management.
  supportEndpoints []string
}

// StorageApplianceList represents a list of storage appliances.
type StorageApplianceList struct {
  // The StorageAppliance items on this page
  value []StorageAppliance

  // The link to the next page of items
  nextLink string
}

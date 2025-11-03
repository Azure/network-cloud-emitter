package clustermanager

import (
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/nogen/azcommonv6"
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/test/common"
)

// ClusterManager represents a control-plane to manage one or more on-premises clusters.
type ClusterManager struct {
  // The list of the resource properties.
  properties ClusterManagerProperties

  // The name of the cluster manager.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string

  // The managed service identities assigned to this resource.
  identity

  azcommonv6.TrackedResource
}

// ClusterManagerProperties represents the properties of a cluster manager.
type ClusterManagerProperties struct {
  // The resource ID of the Log Analytics workspace that is used for the logs collection.
  analyticsWorkspaceId string

  // Field deprecated, this value will no longer influence the cluster manager allocation process and will be removed in a future version. The Azure availability zones within the region that will be used to support the cluster manager resource.
  availabilityZones []string

  // The list of the cluster versions the manager supports. It is used as input in clusterVersion property of a cluster resource.
  clusterVersions []ClusterAvailableVersion

  // The detailed status that provides additional information about the cluster manager.
  detailedStatus ClusterManagerDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // The resource ID of the fabric controller that has one to one mapping with the cluster manager.
  fabricControllerId string

  // The configuration of the managed resource group associated with the resource.
  managedResourceGroupConfiguration common.ManagedResourceGroupConfiguration

  // The extended location (custom location) that represents the cluster manager's control plane location. This extended location is used when creating cluster and rack manifest resources.
  managerExtendedLocation common.ExtendedLocation

  // The provisioning state of the cluster manager.
  provisioningState ClusterManagerProvisioningState

  // Field deprecated, this value will no longer influence the cluster manager allocation process and will be removed in a future version. The size of the Azure virtual machines to use for hosting the cluster manager resource.
  vmSize string
}

// ClusterAvailableVersion represents the cluster version that the cluster manager can be asked to create and manage.
type ClusterAvailableVersion struct {
  // The last date the version of the platform is supported.
  supportExpiryDate string

  // The version of the cluster to be deployed.
  targetClusterVersion string
}

// ClusterManagerPatchParameters represents the body of the request to patch the cluster properties.
type ClusterManagerPatchParameters struct {
  // The identity for the resource.
  identity

  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

// ClusterManagerList represents a list of cluster manager objects.
type ClusterManagerList struct {
  // The ClusterManager items on this page
  value []ClusterManager

  // The link to the next page of items
  nextLink string
}

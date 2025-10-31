package clustermetricsconfiguration

import (
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/nogen/azcommonv6"
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/test/common"
)

// ClusterMetricsConfiguration represents the metrics configuration of an on-premises Network Cloud cluster.
type ClusterMetricsConfiguration struct {
  // The list of the resource properties.
  properties ClusterMetricsConfigurationProperties

  // The name of the metrics configuration for the cluster.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string

  // The extended location of the cluster associated with the resource.
  extendedLocation common.ExtendedLocation

  azcommonv6.TrackedResource
}

// ClusterMetricsConfigurationProperties represents the properties of metrics configuration for the cluster.
type ClusterMetricsConfigurationProperties struct {
  // The interval in minutes by which metrics will be collected.
  collectionInterval int64

  // The more detailed status of the metrics configuration.
  detailedStatus ClusterMetricsConfigurationDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // The list of metrics that are available for the cluster but disabled at the moment.
  disabledMetrics []string

  // The list of metric names that have been chosen to be enabled in addition to the core set of enabled metrics.
  enabledMetrics []string

  // The provisioning state of the metrics configuration.
  provisioningState ClusterMetricsConfigurationProvisioningState
}

// ClusterMetricsConfigurationPatchParameters represents the body of the request to patch the metrics configuration of cluster.
type ClusterMetricsConfigurationPatchParameters struct {
  // The list of the resource properties.
  properties ClusterMetricsConfigurationPatchProperties

  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

// ClusterMetricsConfigurationPatchProperties represents the properties of metrics configuration for the cluster for patching.
type ClusterMetricsConfigurationPatchProperties struct {
  // The interval in minutes by which metrics will be collected.
  collectionInterval int64

  // The list of metric names that have been chosen to be enabled in addition to the core set of enabled metrics.
  enabledMetrics []string
}

// ClusterMetricsConfigurationList represents a list of metrics configuration of the cluster.
type ClusterMetricsConfigurationList struct {
  // The ClusterMetricsConfiguration items on this page
  value []ClusterMetricsConfiguration

  // The link to the next page of items
  nextLink string
}

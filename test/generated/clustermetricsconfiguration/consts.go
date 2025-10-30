package clustermetricsconfiguration

// The more detailed status of the metrics configuration.
type ClusterMetricsConfigurationDetailedStatus string

const (
  ClusterMetricsConfigurationDetailedStatusProcessing ClusterMetricsConfigurationDetailedStatus = "Processing"

  ClusterMetricsConfigurationDetailedStatusApplied ClusterMetricsConfigurationDetailedStatus = "Applied"

  ClusterMetricsConfigurationDetailedStatusError ClusterMetricsConfigurationDetailedStatus = "Error"
)

// The provisioning state of the metrics configuration.
type ClusterMetricsConfigurationProvisioningState string

const (
  ClusterMetricsConfigurationProvisioningStateSucceeded ClusterMetricsConfigurationProvisioningState = "Succeeded"

  ClusterMetricsConfigurationProvisioningStateFailed ClusterMetricsConfigurationProvisioningState = "Failed"

  ClusterMetricsConfigurationProvisioningStateCanceled ClusterMetricsConfigurationProvisioningState = "Canceled"

  ClusterMetricsConfigurationProvisioningStateAccepted ClusterMetricsConfigurationProvisioningState = "Accepted"

  ClusterMetricsConfigurationProvisioningStateProvisioning ClusterMetricsConfigurationProvisioningState = "Provisioning"
)

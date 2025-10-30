package cluster

// ClusterList represents a list of clusters.
type ClusterList struct{}

// Cluster represents the on-premises Network Cloud cluster.
type Cluster struct{}

// ClusterProperties represents the properties of a cluster.
type ClusterProperties struct{}

// RackDefinition represents details regarding the rack.
type RackDefinition struct{}

// AnalyticsOutputSettings represents the settings for the log analytics workspace used for output of logs from this cluster.
type AnalyticsOutputSettings struct{}

// IdentitySelector represents the selection of a managed identity for use.
type IdentitySelector struct{}

// ClusterCapacity represents various details regarding compute capacity.
type ClusterCapacity struct{}

// ServicePrincipalInformation represents the details of the service principal to be used by the cluster during Arc Appliance installation.
type ServicePrincipalInformation struct{}

// CommandOutputSettings represents the settings for commands run within the cluster such as bare metal machine run read-only commands.
type CommandOutputSettings struct{}

// ValidationThreshold indicates allowed machine and node hardware and deployment failures.
type ValidationThreshold struct{}

// RuntimeProtectionConfiguration represents the runtime protection configuration for the cluster.
type RuntimeProtectionConfiguration struct{}

// ClusterSecretArchive configures the key vault to archive the secrets of the cluster for later retrieval.
type ClusterSecretArchive struct{}

// SecretArchiveSettings represents the settings for the secret archive used to hold credentials for the cluster.
type SecretArchiveSettings struct{}

// ClusterUpdateStrategy represents the strategy for updating the cluster.
type ClusterUpdateStrategy struct{}

// VulnerabilityScanningSettings represents the settings for how security vulnerability scanning is applied to the cluster.
type VulnerabilityScanningSettings struct{}

// ClusterPatchParameters represents the body of the request to patch the cluster properties.
type ClusterPatchParameters struct{}

// ClusterPatchProperties represents the properties of the cluster for patching.
type ClusterPatchProperties struct{}

// VulnerabilityScanningSettingsPatch represents the settings for how security vulnerability scanning is applied to the cluster.
type VulnerabilityScanningSettingsPatch struct{}

// ClusterContinueUpdateVersionParameters represents the body of the request to continue the update of a cluster version.
type ClusterContinueUpdateVersionParameters struct{}

// ClusterDeployParameters represents the body of the request to deploy cluster.
type ClusterDeployParameters struct{}

// ClusterScanRuntimeParameters defines the parameters for the cluster scan runtime operation.
type ClusterScanRuntimeParameters struct{}

// ClusterUpdateVersionParameters represents the body of the request to update cluster version.
type ClusterUpdateVersionParameters struct{}

package cluster

import (
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/nogen/azcommonv6"
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/test/common"
)

// ClusterList represents a list of clusters.
type ClusterList struct {
  // The Cluster items on this page
  value []Cluster

  // The link to the next page of items
  nextLink string
}

// Cluster represents the on-premises Network Cloud cluster.
type Cluster struct {
  // The list of the resource properties.
  properties ClusterProperties

  // The name of the cluster.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string

  // The extended location of the cluster manager associated with the cluster.
  extendedLocation common.ExtendedLocation

  // The managed service identities assigned to this resource.
  identity

  azcommonv6.TrackedResource
}

// ClusterProperties represents the properties of a cluster.
type ClusterProperties struct {
  // The rack definition that is intended to reflect only a single rack in a single rack cluster, or an aggregator rack in a multi-rack cluster.
  aggregatorOrSingleRackDefinition RackDefinition

  // The settings for the log analytics workspace used for output of logs from this cluster.
  analyticsOutputSettings AnalyticsOutputSettings

  // Field Deprecated. The resource ID of the Log Analytics Workspace that will be used for storing relevant logs.
  analyticsWorkspaceId string

  // The list of cluster runtime version upgrades available for this cluster.
  availableUpgradeVersions []ClusterAvailableUpgradeVersion

  // The capacity supported by this cluster.
  clusterCapacity ClusterCapacity

  // The latest heartbeat status between the cluster manager and the cluster.
  clusterConnectionStatus ClusterConnectionStatus

  // The extended location (custom location) that represents the cluster's control plane location. This extended location is used to route the requests of child objects of the cluster that are handled by the platform operator.
  clusterExtendedLocation common.ExtendedLocation

  // The customer-provided location information to identify where the cluster resides.
  clusterLocation string

  // The latest connectivity status between cluster manager and the cluster.
  clusterManagerConnectionStatus ClusterManagerConnectionStatus

  // The resource ID of the cluster manager that manages this cluster. This is set by the Cluster Manager when the cluster is created.
  clusterManagerId string

  // The service principal to be used by the cluster during Arc Appliance installation.
  clusterServicePrincipal ServicePrincipalInformation

  // The type of rack configuration for the cluster.
  clusterType ClusterType

  // The current runtime version of the cluster.
  clusterVersion string

  // The settings for commands run in this cluster, such as bare metal machine run read only commands and data extracts.
  commandOutputSettings CommandOutputSettings

  // The validation threshold indicating the allowable failures of compute machines during environment validation and deployment.
  computeDeploymentThreshold ValidationThreshold

  // The list of rack definitions for the compute racks in a multi-rack
  // cluster, or an empty list in a single-rack cluster.
  computeRackDefinitions []RackDefinition

  // The current detailed status of the cluster.
  detailedStatus ClusterDetailedStatus

  // The descriptive message about the detailed status.
  detailedStatusMessage string

  // Field Deprecated. This field will not be populated in an upcoming version. The extended location (custom location) that represents the Hybrid AKS control plane location. This extended location is used when creating provisioned clusters (Hybrid AKS clusters).
  hybridAksExtendedLocation common.ExtendedLocation

  // The configuration of the managed resource group associated with the resource.
  managedResourceGroupConfiguration common.ManagedResourceGroupConfiguration

  // The count of Manual Action Taken (MAT) events that have not been validated.
  manualActionCount int64

  // The resource ID of the Network Fabric associated with the cluster.
  networkFabricId string

  // The provisioning state of the cluster.
  provisioningState ClusterProvisioningState

  // The settings for cluster runtime protection.
  runtimeProtectionConfiguration RuntimeProtectionConfiguration

  // The configuration for use of a key vault to store secrets for later retrieval by the operator.
  secretArchive ClusterSecretArchive

  // The settings for the secret archive used to hold credentials for the cluster.
  secretArchiveSettings SecretArchiveSettings

  // The support end date of the runtime version of the cluster.
  supportExpiryDate string

  // The strategy for updating the cluster.
  updateStrategy ClusterUpdateStrategy

  // The settings for how security vulnerability scanning is applied to the cluster.
  vulnerabilityScanningSettings VulnerabilityScanningSettings

  // The list of workload resource IDs that are hosted within this cluster.
  workloadResourceIds []string
}

// RackDefinition represents details regarding the rack.
type RackDefinition struct {
  // The zone name used for this rack when created. Availability zones are used for workload placement.
  availabilityZone string

  // The unordered list of bare metal machine configuration.
  bareMetalMachineConfigurationData []BareMetalMachineConfigurationData

  // The resource ID of the network rack that matches this rack definition.
  networkRackId string

  // The free-form description of the rack's location.
  rackLocation string

  // The unique identifier for the rack within Network Cloud cluster. An alternate unique alphanumeric value other than a serial number may be provided if desired.
  rackSerialNumber string

  // The resource ID of the sku for the rack being added.
  rackSkuId string

  // The list of storage appliance configuration data for this rack.
  storageApplianceConfigurationData []StorageApplianceConfigurationData
}

// BareMetalMachineConfigurationData represents configuration for the bare metal machine.
type BareMetalMachineConfigurationData struct {
  // The connection string for the baseboard management controller including IP address and protocol.
  bmcConnectionString string

  // The credentials of the baseboard management controller on this bare metal machine. The password field is expected to be an Azure Key Vault key URL. Until the cluster is converted to utilize managed identity by setting the secret archive settings, the actual password value should be provided instead.
  bmcCredentials common.AdministrativeCredentials

  // The MAC address of the BMC for this machine.
  bmcMacAddress string

  // The MAC address associated with the PXE NIC card.
  bootMacAddress string

  // The free-form additional information about the machine, e.g. an asset tag.
  machineDetails string

  // The user-provided name for the bare metal machine created from this specification.
  // If not provided, the machine name will be generated programmatically.
  machineName string

  // The slot the physical machine is in the rack based on the BOM configuration.
  rackSlot int64

  // The serial number of the machine. Hardware suppliers may use an alternate value. For example, service tag.
  serialNumber string
}

// StorageApplianceConfigurationData represents configuration for the storage application.
type StorageApplianceConfigurationData struct {
  // The credentials of the administrative interface on this storage appliance. The password field is expected to be an Azure Key Vault key URL. Until the cluster is converted to utilize managed identity by setting the secret archive settings, the actual password value should be provided instead.
  adminCredentials common.AdministrativeCredentials

  // The slot that storage appliance is in the rack based on the BOM configuration.
  rackSlot int64

  // The serial number of the appliance.
  serialNumber string

  // The user-provided name for the storage appliance that will be created from this specification.
  storageApplianceName string
}

// AnalyticsOutputSettings represents the settings for the log analytics workspace used for output of logs from this cluster.
type AnalyticsOutputSettings struct {
  // The resource ID of the analytics workspace that is to be used by the specified identity.
  analyticsWorkspaceId string

  // The selection of the managed identity to use with this analytics workspace. The identity type must be either system assigned or user assigned.
  associatedIdentity IdentitySelector
}

// IdentitySelector represents the selection of a managed identity for use.
type IdentitySelector struct {
  // The type of managed identity that is being selected.
  identityType ManagedServiceIdentitySelectorType

  // The user assigned managed identity resource ID to use. Mutually exclusive with a system assigned identity type.
  userAssignedIdentityResourceId *string
}

// ClusterAvailableUpgradeVersion represents the various cluster upgrade parameters.
type ClusterAvailableUpgradeVersion struct {
  // The indicator of whether the control plane will be impacted during the upgrade.
  controlImpact ControlImpact

  // The expected duration needed for this upgrade.
  expectedDuration string

  // The impact description including the specific details and release notes.
  impactDescription string

  // The last date the version of the platform is supported.
  supportExpiryDate string

  // The target version this cluster will be upgraded to.
  targetClusterVersion string

  // The indicator of whether the workload will be impacted during the upgrade.
  workloadImpact WorkloadImpact
}

// ClusterCapacity represents various details regarding compute capacity.
type ClusterCapacity struct {
  // The remaining appliance-based storage in GB available for workload use. Measured in gibibytes.
  availableApplianceStorageGB int64

  // The remaining number of cores that are available in this cluster for workload use.
  availableCoreCount int64

  // The remaining machine or host-based storage in GB available for workload use. Measured in gibibytes.
  availableHostStorageGB int64

  // The remaining memory in GB that are available in this cluster for workload use. Measured in gibibytes.
  availableMemoryGB int64

  // The total appliance-based storage in GB supported by this cluster for workload use. Measured in gibibytes.
  totalApplianceStorageGB int64

  // The total number of cores that are supported by this cluster for workload use.
  totalCoreCount int64

  // The total machine or host-based storage in GB supported by this cluster for workload use. Measured in gibibytes.
  totalHostStorageGB int64

  // The total memory supported by this cluster for workload use. Measured in gibibytes.
  totalMemoryGB int64
}

// ServicePrincipalInformation represents the details of the service principal to be used by the cluster during Arc Appliance installation.
type ServicePrincipalInformation struct {
  // The application ID, also known as client ID, of the service principal.
  applicationId string

  // The password of the service principal.
  password string

  // The principal ID, also known as the object ID, of the service principal.
  principalId string

  // The tenant ID, also known as the directory ID, of the tenant in which the service principal is created.
  tenantId string
}

// CommandOutputSettings represents the settings for commands run within the cluster such as bare metal machine run read-only commands.
type CommandOutputSettings struct {
  // The selection of the managed identity to use with this storage account container. The identity type must be either system assigned or user assigned.
  associatedIdentity IdentitySelector

  // The URL of the storage account container that is to be used by the specified identities.
  containerUrl string
}

// ValidationThreshold indicates allowed machine and node hardware and deployment failures.
type ValidationThreshold struct {
  // Selection of how the type evaluation is applied to the cluster calculation.
  grouping ValidationThresholdGrouping

  // Selection of how the threshold should be evaluated.
  type ValidationThresholdType

  // The numeric threshold value.
  value int64
}

// RuntimeProtectionConfiguration represents the runtime protection configuration for the cluster.
type RuntimeProtectionConfiguration struct {
  // The mode of operation for runtime protection.
  enforcementLevel RuntimeProtectionEnforcementLevel
}

// ClusterSecretArchive configures the key vault to archive the secrets of the cluster for later retrieval.
type ClusterSecretArchive struct {
  // The resource ID of the key vault to archive the secrets of the cluster.
  keyVaultId string

  // The indicator if the specified key vault should be used to archive the secrets of the cluster.
  useKeyVault ClusterSecretArchiveEnabled
}

// SecretArchiveSettings represents the settings for the secret archive used to hold credentials for the cluster.
type SecretArchiveSettings struct {
  // The selection of the managed identity to use with this vault URI. The identity type must be either system assigned or user assigned.
  associatedIdentity IdentitySelector

  // The URI for the key vault used as the secret archive.
  vaultUri string
}

// ClusterUpdateStrategy represents the strategy for updating the cluster.
type ClusterUpdateStrategy struct {
  // The maximum number of worker nodes that can be offline within the increment of update, e.g., rack-by-rack.
  // Limited by the maximum number of machines in the increment. Defaults to the whole increment size.
  maxUnavailable int64

  // The mode of operation for runtime protection.
  strategyType ClusterUpdateStrategyType

  // Selection of how the threshold should be evaluated.
  thresholdType ValidationThresholdType

  // The numeric threshold value.
  thresholdValue int64

  // The time to wait between the increments of update defined by the strategy.
  waitTimeMinutes int64
}

// VulnerabilityScanningSettings represents the settings for how security vulnerability scanning is applied to the cluster.
type VulnerabilityScanningSettings struct {
  // The mode selection for container vulnerability scanning.
  containerScan VulnerabilityScanningSettingsContainerScan
}

// ClusterPatchParameters represents the body of the request to patch the cluster properties.
type ClusterPatchParameters struct {
  // The identity for the resource.
  identity

  // The list of the resource properties.
  properties ClusterPatchProperties

  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

// ClusterPatchProperties represents the properties of the cluster for patching.
type ClusterPatchProperties struct {
  // The rack definition that is intended to reflect only a single rack in a single rack cluster, or an aggregator rack in a multi-rack cluster.
  aggregatorOrSingleRackDefinition RackDefinition

  // The settings for the log analytics workspace used for output of logs from this cluster.
  analyticsOutputSettings *AnalyticsOutputSettings

  // The customer-provided location information to identify where the cluster resides.
  clusterLocation string

  // The service principal to be used by the cluster during Arc Appliance installation.
  clusterServicePrincipal *ServicePrincipalInformation

  // The settings for commands run in this cluster, such as bare metal machine run read only commands and data extracts.
  commandOutputSettings *CommandOutputSettings

  // The validation threshold indicating the allowable failures of compute machines during environment validation and deployment.
  computeDeploymentThreshold *ValidationThreshold

  // The list of rack definitions for the compute racks in a multi-rack
  // cluster, or an empty list in a single-rack cluster.
  computeRackDefinitions []RackDefinition

  // The settings for cluster runtime protection.
  runtimeProtectionConfiguration *RuntimeProtectionConfiguration

  // The configuration for use of a key vault to store secrets for later retrieval by the operator.
  secretArchive *ClusterSecretArchive

  // The settings for the secret archive used to hold credentials for the cluster.
  secretArchiveSettings *SecretArchiveSettings

  // The strategy for updating the cluster.
  updateStrategy *ClusterUpdateStrategy

  // The settings for how security vulnerability scanning is applied to the cluster.
  vulnerabilityScanningSettings *VulnerabilityScanningSettingsPatch
}

// VulnerabilityScanningSettingsPatch represents the settings for how security vulnerability scanning is applied to the cluster.
type VulnerabilityScanningSettingsPatch struct {
  // The mode selection for container vulnerability scanning.
  containerScan VulnerabilityScanningSettingsContainerScan
}

// ClusterContinueUpdateVersionParameters represents the body of the request to continue the update of a cluster version.
type ClusterContinueUpdateVersionParameters struct {
  // The mode by which the cluster will target the next grouping of servers to continue the update.
  machineGroupTargetingMode ClusterContinueUpdateVersionMachineGroupTargetingMode
}

// ClusterDeployParameters represents the body of the request to deploy cluster.
type ClusterDeployParameters struct {
  // The names of bare metal machines in the cluster that should be skipped during environment validation.
  skipValidationsForMachines []string
}

// ClusterScanRuntimeParameters defines the parameters for the cluster scan runtime operation.
type ClusterScanRuntimeParameters struct {
  // The choice of if the scan operation should run the scan.
  scanActivity ClusterScanRuntimeParametersScanActivity
}

// ClusterUpdateVersionParameters represents the body of the request to update cluster version.
type ClusterUpdateVersionParameters struct {
  // The version to be applied to the cluster during update.
  targetClusterVersion string
}

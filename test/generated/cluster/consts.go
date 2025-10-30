package cluster

// The type of managed identity that is being selected.
type ManagedServiceIdentitySelectorType string

const (
  ManagedServiceIdentitySelectorTypeSystemAssignedIdentity ManagedServiceIdentitySelectorType = "SystemAssignedIdentity"

  ManagedServiceIdentitySelectorTypeUserAssignedIdentity ManagedServiceIdentitySelectorType = "UserAssignedIdentity"
)

// The latest heartbeat status between the cluster manager and the cluster.
type ClusterConnectionStatus string

const (
  ClusterConnectionStatusConnected ClusterConnectionStatus = "Connected"

  ClusterConnectionStatusDisconnected ClusterConnectionStatus = "Disconnected"

  ClusterConnectionStatusTimeout ClusterConnectionStatus = "Timeout"

  ClusterConnectionStatusUndefined ClusterConnectionStatus = "Undefined"
)

// The latest connectivity status between cluster manager and the cluster.
type ClusterManagerConnectionStatus string

const (
  ClusterManagerConnectionStatusConnected ClusterManagerConnectionStatus = "Connected"

  ClusterManagerConnectionStatusUnreachable ClusterManagerConnectionStatus = "Unreachable"
)

// The type of rack configuration for the cluster.
type ClusterType string

const (
  ClusterTypeSingleRack ClusterType = "SingleRack"

  ClusterTypeMultiRack ClusterType = "MultiRack"
)

// Selection of how the type evaluation is applied to the cluster calculation.
type ValidationThresholdGrouping string

const (
  ValidationThresholdGroupingPerCluster ValidationThresholdGrouping = "PerCluster"

  ValidationThresholdGroupingPerRack ValidationThresholdGrouping = "PerRack"
)

// Selection of how the threshold should be evaluated.
type ValidationThresholdType string

const (
  ValidationThresholdTypeCountSuccess ValidationThresholdType = "CountSuccess"

  ValidationThresholdTypePercentSuccess ValidationThresholdType = "PercentSuccess"
)

// The current detailed status of the cluster.
type ClusterDetailedStatus string

const (
  ClusterDetailedStatusPendingDeployment ClusterDetailedStatus = "PendingDeployment"

  ClusterDetailedStatusDeploying ClusterDetailedStatus = "Deploying"

  ClusterDetailedStatusRunning ClusterDetailedStatus = "Running"

  ClusterDetailedStatusUpdating ClusterDetailedStatus = "Updating"

  ClusterDetailedStatusUpdatePaused ClusterDetailedStatus = "UpdatePaused"

  ClusterDetailedStatusDegraded ClusterDetailedStatus = "Degraded"

  ClusterDetailedStatusDeleting ClusterDetailedStatus = "Deleting"

  ClusterDetailedStatusDisconnected ClusterDetailedStatus = "Disconnected"

  ClusterDetailedStatusFailed ClusterDetailedStatus = "Failed"
)

// The provisioning state of the cluster.
type ClusterProvisioningState string

const (
  ClusterProvisioningStateSucceeded ClusterProvisioningState = "Succeeded"

  ClusterProvisioningStateFailed ClusterProvisioningState = "Failed"

  ClusterProvisioningStateCanceled ClusterProvisioningState = "Canceled"

  ClusterProvisioningStateAccepted ClusterProvisioningState = "Accepted"

  ClusterProvisioningStateValidating ClusterProvisioningState = "Validating"

  ClusterProvisioningStateUpdating ClusterProvisioningState = "Updating"
)

// The mode of operation for runtime protection.
type RuntimeProtectionEnforcementLevel string

const (
  RuntimeProtectionEnforcementLevelAudit RuntimeProtectionEnforcementLevel = "Audit"

  RuntimeProtectionEnforcementLevelDisabled RuntimeProtectionEnforcementLevel = "Disabled"

  RuntimeProtectionEnforcementLevelOnDemand RuntimeProtectionEnforcementLevel = "OnDemand"

  RuntimeProtectionEnforcementLevelPassive RuntimeProtectionEnforcementLevel = "Passive"

  RuntimeProtectionEnforcementLevelRealTime RuntimeProtectionEnforcementLevel = "RealTime"
)

// The indicator if the specified key vault should be used to archive the secrets of the cluster.
type ClusterSecretArchiveEnabled string

const (
  ClusterSecretArchiveEnabledTrue ClusterSecretArchiveEnabled = "True"

  ClusterSecretArchiveEnabledFalse ClusterSecretArchiveEnabled = "False"
)

// The mode of operation for runtime protection.
type ClusterUpdateStrategyType string

const (
  ClusterUpdateStrategyTypeRack ClusterUpdateStrategyType = "Rack"

  ClusterUpdateStrategyTypePauseAfterRack ClusterUpdateStrategyType = "PauseAfterRack"
)

// The mode selection for container vulnerability scanning.
type VulnerabilityScanningSettingsContainerScan string

const (
  VulnerabilityScanningSettingsContainerScanDisabled VulnerabilityScanningSettingsContainerScan = "Disabled"

  VulnerabilityScanningSettingsContainerScanEnabled VulnerabilityScanningSettingsContainerScan = "Enabled"
)

// The mode by which the cluster will target the next grouping of servers to continue the update.
type ClusterContinueUpdateVersionMachineGroupTargetingMode string

const (
  ClusterContinueUpdateVersionMachineGroupTargetingModeAlphaByRack ClusterContinueUpdateVersionMachineGroupTargetingMode = "AlphaByRack"
)

// The choice of if the scan operation should run the scan.
type ClusterScanRuntimeParametersScanActivity string

const (
  ClusterScanRuntimeParametersScanActivityScan ClusterScanRuntimeParametersScanActivity = "Scan"

  ClusterScanRuntimeParametersScanActivitySkip ClusterScanRuntimeParametersScanActivity = "Skip"
)

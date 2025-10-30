package storageappliance

// The detailed status of the storage appliance.
type StorageApplianceDetailedStatus string

const (
  StorageApplianceDetailedStatusAvailable StorageApplianceDetailedStatus = "Available"

  StorageApplianceDetailedStatusDegraded StorageApplianceDetailedStatus = "Degraded"

  StorageApplianceDetailedStatusError StorageApplianceDetailedStatus = "Error"

  StorageApplianceDetailedStatusProvisioning StorageApplianceDetailedStatus = "Provisioning"
)

// The provisioning state of the storage appliance.
type StorageApplianceProvisioningState string

const (
  StorageApplianceProvisioningStateSucceeded StorageApplianceProvisioningState = "Succeeded"

  StorageApplianceProvisioningStateFailed StorageApplianceProvisioningState = "Failed"

  StorageApplianceProvisioningStateCanceled StorageApplianceProvisioningState = "Canceled"

  StorageApplianceProvisioningStateProvisioning StorageApplianceProvisioningState = "Provisioning"

  StorageApplianceProvisioningStateAccepted StorageApplianceProvisioningState = "Accepted"
)

// The indicator of whether the storage appliance supports remote vendor management.
type RemoteVendorManagementFeature string

const (
  RemoteVendorManagementFeatureSupported RemoteVendorManagementFeature = "Supported"

  RemoteVendorManagementFeatureUnsupported RemoteVendorManagementFeature = "Unsupported"
)

// The indicator of whether the remote vendor management feature is enabled or disabled, or unsupported if it is an unsupported feature.
type RemoteVendorManagementStatus string

const (
  RemoteVendorManagementStatusEnabled RemoteVendorManagementStatus = "Enabled"

  RemoteVendorManagementStatusDisabled RemoteVendorManagementStatus = "Disabled"

  RemoteVendorManagementStatusUnsupported RemoteVendorManagementStatus = "Unsupported"
)

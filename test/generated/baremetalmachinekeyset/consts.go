package baremetalmachinekeyset

// The more detailed status of the key set.
type BareMetalMachineKeySetDetailedStatus string

const (
  BareMetalMachineKeySetDetailedStatusAllActive BareMetalMachineKeySetDetailedStatus = "AllActive"

  BareMetalMachineKeySetDetailedStatusSomeInvalid BareMetalMachineKeySetDetailedStatus = "SomeInvalid"

  BareMetalMachineKeySetDetailedStatusAllInvalid BareMetalMachineKeySetDetailedStatus = "AllInvalid"

  BareMetalMachineKeySetDetailedStatusValidating BareMetalMachineKeySetDetailedStatus = "Validating"
)

// The access level allowed for the users in this key set.
type BareMetalMachineKeySetPrivilegeLevel string

const (
  BareMetalMachineKeySetPrivilegeLevelStandard BareMetalMachineKeySetPrivilegeLevel = "Standard"

  BareMetalMachineKeySetPrivilegeLevelSuperuser BareMetalMachineKeySetPrivilegeLevel = "Superuser"
)

// The provisioning state of the bare metal machine key set.
type BareMetalMachineKeySetProvisioningState string

const (
  BareMetalMachineKeySetProvisioningStateSucceeded BareMetalMachineKeySetProvisioningState = "Succeeded"

  BareMetalMachineKeySetProvisioningStateFailed BareMetalMachineKeySetProvisioningState = "Failed"

  BareMetalMachineKeySetProvisioningStateCanceled BareMetalMachineKeySetProvisioningState = "Canceled"

  BareMetalMachineKeySetProvisioningStateAccepted BareMetalMachineKeySetProvisioningState = "Accepted"

  BareMetalMachineKeySetProvisioningStateProvisioning BareMetalMachineKeySetProvisioningState = "Provisioning"
)

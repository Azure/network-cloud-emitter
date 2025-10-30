package bmckeyset

// The more detailed status of the key set.
type BmcKeySetDetailedStatus string

const (
  BmcKeySetDetailedStatusAllActive BmcKeySetDetailedStatus = "AllActive"

  BmcKeySetDetailedStatusSomeInvalid BmcKeySetDetailedStatus = "SomeInvalid"

  BmcKeySetDetailedStatusAllInvalid BmcKeySetDetailedStatus = "AllInvalid"

  BmcKeySetDetailedStatusValidating BmcKeySetDetailedStatus = "Validating"
)

// The access level allowed for the users in this key set.
type BmcKeySetPrivilegeLevel string

const (
  BmcKeySetPrivilegeLevelReadOnly BmcKeySetPrivilegeLevel = "ReadOnly"

  BmcKeySetPrivilegeLevelAdministrator BmcKeySetPrivilegeLevel = "Administrator"
)

// The provisioning state of the baseboard management controller key set.
type BmcKeySetProvisioningState string

const (
  BmcKeySetProvisioningStateSucceeded BmcKeySetProvisioningState = "Succeeded"

  BmcKeySetProvisioningStateFailed BmcKeySetProvisioningState = "Failed"

  BmcKeySetProvisioningStateCanceled BmcKeySetProvisioningState = "Canceled"

  BmcKeySetProvisioningStateAccepted BmcKeySetProvisioningState = "Accepted"

  BmcKeySetProvisioningStateProvisioning BmcKeySetProvisioningState = "Provisioning"
)

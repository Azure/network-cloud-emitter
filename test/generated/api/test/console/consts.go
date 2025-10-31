package console

// The more detailed status of the console.
type ConsoleDetailedStatus string

const (
  ConsoleDetailedStatusReady ConsoleDetailedStatus = "Ready"

  ConsoleDetailedStatusError ConsoleDetailedStatus = "Error"
)

// The indicator of whether the console access is enabled.
type ConsoleEnabled string

const (
  ConsoleEnabledTrue ConsoleEnabled = "True"

  ConsoleEnabledFalse ConsoleEnabled = "False"
)

// The provisioning state of the virtual machine console.
type ConsoleProvisioningState string

const (
  ConsoleProvisioningStateSucceeded ConsoleProvisioningState = "Succeeded"

  ConsoleProvisioningStateFailed ConsoleProvisioningState = "Failed"

  ConsoleProvisioningStateCanceled ConsoleProvisioningState = "Canceled"

  ConsoleProvisioningStateProvisioning ConsoleProvisioningState = "Provisioning"

  ConsoleProvisioningStateAccepted ConsoleProvisioningState = "Accepted"
)

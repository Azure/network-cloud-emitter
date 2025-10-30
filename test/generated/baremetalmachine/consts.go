package baremetalmachine

// The cordon status of the bare metal machine.
type BareMetalMachineCordonStatus string

const (
  BareMetalMachineCordonStatusCordoned BareMetalMachineCordonStatus = "Cordoned"

  BareMetalMachineCordonStatusUncordoned BareMetalMachineCordonStatus = "Uncordoned"
)

// The more detailed status of the bare metal machine.
type BareMetalMachineDetailedStatus string

const (
  BareMetalMachineDetailedStatusPreparing BareMetalMachineDetailedStatus = "Preparing"

  BareMetalMachineDetailedStatusError BareMetalMachineDetailedStatus = "Error"

  BareMetalMachineDetailedStatusAvailable BareMetalMachineDetailedStatus = "Available"

  BareMetalMachineDetailedStatusProvisioning BareMetalMachineDetailedStatus = "Provisioning"

  BareMetalMachineDetailedStatusProvisioned BareMetalMachineDetailedStatus = "Provisioned"

  BareMetalMachineDetailedStatusDeprovisioning BareMetalMachineDetailedStatus = "Deprovisioning"
)

// The outcome of the hardware validation.
type BareMetalMachineHardwareValidationResult string

const (
  BareMetalMachineHardwareValidationResultPass BareMetalMachineHardwareValidationResult = "Pass"

  BareMetalMachineHardwareValidationResultFail BareMetalMachineHardwareValidationResult = "Fail"
)

// The power state derived from the baseboard management controller.
type BareMetalMachinePowerState string

const (
  BareMetalMachinePowerStateOn BareMetalMachinePowerState = "On"

  BareMetalMachinePowerStateOff BareMetalMachinePowerState = "Off"
)

// The provisioning state of the bare metal machine.
type BareMetalMachineProvisioningState string

const (
  BareMetalMachineProvisioningStateSucceeded BareMetalMachineProvisioningState = "Succeeded"

  BareMetalMachineProvisioningStateFailed BareMetalMachineProvisioningState = "Failed"

  BareMetalMachineProvisioningStateCanceled BareMetalMachineProvisioningState = "Canceled"

  BareMetalMachineProvisioningStateProvisioning BareMetalMachineProvisioningState = "Provisioning"

  BareMetalMachineProvisioningStateAccepted BareMetalMachineProvisioningState = "Accepted"
)

// The indicator of whether the bare metal machine is ready to receive workloads.
type BareMetalMachineReadyState string

const (
  BareMetalMachineReadyStateTrue BareMetalMachineReadyState = "True"

  BareMetalMachineReadyStateFalse BareMetalMachineReadyState = "False"
)

// The indicator of whether to evacuate the node workload when the bare metal machine is cordoned.
type BareMetalMachineEvacuate string

const (
  BareMetalMachineEvacuateTrue BareMetalMachineEvacuate = "True"

  BareMetalMachineEvacuateFalse BareMetalMachineEvacuate = "False"
)

// The indicator of whether to skip the graceful OS shutdown and power off the bare metal machine immediately.
type BareMetalMachineSkipShutdown string

const (
  BareMetalMachineSkipShutdownTrue BareMetalMachineSkipShutdown = "True"

  BareMetalMachineSkipShutdownFalse BareMetalMachineSkipShutdown = "False"
)

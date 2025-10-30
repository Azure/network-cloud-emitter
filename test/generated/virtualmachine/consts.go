package virtualmachine

// Selects the boot method for the virtual machine.
type VirtualMachineBootMethod string

const (
  VirtualMachineBootMethodUEFI VirtualMachineBootMethod = "UEFI"

  VirtualMachineBootMethodBIOS VirtualMachineBootMethod = "BIOS"
)

// The indicator of whether this is the default gateway.
// Only one of the attached networks (including the CloudServicesNetwork attachment) for a single machine may be specified as True.
type DefaultGateway string

const (
  DefaultGatewayTrue DefaultGateway = "True"

  DefaultGatewayFalse DefaultGateway = "False"
)

// The IP allocation mechanism for the virtual machine.
// Dynamic and Static are only valid for l3Network which may also specify Disabled.
// Otherwise, Disabled is the only permitted value.
type VirtualMachineIPAllocationMethod string

const (
  VirtualMachineIPAllocationMethodDynamic VirtualMachineIPAllocationMethod = "Dynamic"

  VirtualMachineIPAllocationMethodStatic VirtualMachineIPAllocationMethod = "Static"

  VirtualMachineIPAllocationMethodDisabled VirtualMachineIPAllocationMethod = "Disabled"
)

// The more detailed status of the virtual machine.
type VirtualMachineDetailedStatus string

const (
  VirtualMachineDetailedStatusAvailable VirtualMachineDetailedStatus = "Available"

  VirtualMachineDetailedStatusError VirtualMachineDetailedStatus = "Error"

  VirtualMachineDetailedStatusProvisioning VirtualMachineDetailedStatus = "Provisioning"

  VirtualMachineDetailedStatusRunning VirtualMachineDetailedStatus = "Running"

  VirtualMachineDetailedStatusScheduling VirtualMachineDetailedStatus = "Scheduling"

  VirtualMachineDetailedStatusStopped VirtualMachineDetailedStatus = "Stopped"

  VirtualMachineDetailedStatusTerminating VirtualMachineDetailedStatus = "Terminating"

  VirtualMachineDetailedStatusUnknown VirtualMachineDetailedStatus = "Unknown"
)

// Field Deprecated, the value will be ignored if provided. The indicator of whether one of the specified CPU cores is isolated to run the emulator thread for this virtual machine.
type VirtualMachineIsolateEmulatorThread string

const (
  VirtualMachineIsolateEmulatorThreadTrue VirtualMachineIsolateEmulatorThread = "True"

  VirtualMachineIsolateEmulatorThreadFalse VirtualMachineIsolateEmulatorThread = "False"
)

// The power state of the virtual machine.
type VirtualMachinePowerState string

const (
  VirtualMachinePowerStateOn VirtualMachinePowerState = "On"

  VirtualMachinePowerStateOff VirtualMachinePowerState = "Off"

  VirtualMachinePowerStateUnknown VirtualMachinePowerState = "Unknown"
)

// The provisioning state of the virtual machine.
type VirtualMachineProvisioningState string

const (
  VirtualMachineProvisioningStateSucceeded VirtualMachineProvisioningState = "Succeeded"

  VirtualMachineProvisioningStateFailed VirtualMachineProvisioningState = "Failed"

  VirtualMachineProvisioningStateCanceled VirtualMachineProvisioningState = "Canceled"

  VirtualMachineProvisioningStateProvisioning VirtualMachineProvisioningState = "Provisioning"

  VirtualMachineProvisioningStateAccepted VirtualMachineProvisioningState = "Accepted"
)

// The strategy for creating the OS disk.
type OsDiskCreateOption string

const (
  // Utilize the local storage of the host machine.
  OsDiskCreateOptionEphemeral OsDiskCreateOption = "Ephemeral"

  // Utilize a storage appliance backed volume to host the disk.
  OsDiskCreateOptionPersistent OsDiskCreateOption = "Persistent"
)

// The strategy for deleting the OS disk.
type OsDiskDeleteOption string

const (
  OsDiskDeleteOptionDelete OsDiskDeleteOption = "Delete"
)

// Field Deprecated, use virtualizationModel instead. The type of the virtio interface.
type VirtualMachineVirtioInterfaceType string

const (
  VirtualMachineVirtioInterfaceTypeModern VirtualMachineVirtioInterfaceType = "Modern"

  VirtualMachineVirtioInterfaceTypeTransitional VirtualMachineVirtioInterfaceType = "Transitional"
)

// The type of the device model to use.
type VirtualMachineDeviceModelType string

const (
  // Traditional and most compatible device virtualization interface.
  VirtualMachineDeviceModelTypeT1 VirtualMachineDeviceModelType = "T1"

  // Modern and enhanced device virtualization interface.
  VirtualMachineDeviceModelTypeT2 VirtualMachineDeviceModelType = "T2"

  // Improved security and functionality (including TPM and secure boot support). Required for windows 11 and server 2025.
  VirtualMachineDeviceModelTypeT3 VirtualMachineDeviceModelType = "T3"
)

// The indicator of whether to skip the graceful OS shutdown and power off the virtual machine immediately.
type SkipShutdown string

const (
  SkipShutdownTrue SkipShutdown = "True"

  SkipShutdownFalse SkipShutdown = "False"
)

package virtualmachine

// VirtualMachineList represents a list of virtual machines.
type VirtualMachineList struct{}

// VirtualMachine represents the on-premises Network Cloud virtual machine.
type VirtualMachine struct{}

// VirtualMachineProperties represents the properties of the virtual machine.
type VirtualMachineProperties struct{}

// NetworkAttachment represents the single network attachment.
type NetworkAttachment struct{}

// StorageProfile represents information about a disk.
type StorageProfile struct{}

// OsDisk represents configuration of the boot disk.
type OsDisk struct{}

// ImageRepositoryCredentials represents the credentials used to login to the image repository.
type ImageRepositoryCredentials struct{}

// VirtualMachinePatchParameters represents the body of the request to patch the virtual machine.
type VirtualMachinePatchParameters struct{}

// VirtualMachinePatchProperties represents the properties of the virtual machine that can be patched.
type VirtualMachinePatchProperties struct{}

// VirtualMachinePowerOffParameters represents the body of the request to power off virtual machine.
type VirtualMachinePowerOffParameters struct{}

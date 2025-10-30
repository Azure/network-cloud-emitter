package baremetalmachine

// BareMetalMachineList represents a list of bare metal machines.
type BareMetalMachineList struct{}

// BareMetalMachine represents the physical machine in the rack.
type BareMetalMachine struct{}

// BareMetalMachineProperties represents the properties of a bare metal machine.
type BareMetalMachineProperties struct{}

// HardwareInventory represents the hardware configuration of this machine as exposed to the customer, including information acquired from the model/sku information and from the ironic inspector.
type HardwareInventory struct{}

// HardwareValidationStatus represents the latest hardware validation details performed for this bare metal machine.
type HardwareValidationStatus struct{}

// RuntimeProtectionStatus represents the runtime protection status of the bare metal machine.
type RuntimeProtectionStatus struct{}

// BareMetalMachinePatchParameters represents the body of the request to patch bare metal machine properties.
type BareMetalMachinePatchParameters struct{}

// BareMetalMachinePatchProperties represents the properties of the bare metal machine that can be patched.
type BareMetalMachinePatchProperties struct{}

// BareMetalMachineCordonParameters represents the body of the request to evacuate workloads from node on a bare metal machine.
type BareMetalMachineCordonParameters struct{}

// BareMetalMachinePowerOffParameters represents the body of the request to power off bare metal machine.
type BareMetalMachinePowerOffParameters struct{}

// BareMetalMachineReplaceParameters represents the body of the request to physically swap a bare metal machine for another.
type BareMetalMachineReplaceParameters struct{}

// BareMetalMachineRunCommandParameters represents the body of the request to execute a script on the bare metal machine.
type BareMetalMachineRunCommandParameters struct{}

// BareMetalMachineRunDataExtractsParameters represents the body of request containing list of curated data extraction commands to run on the bare metal machine.
type BareMetalMachineRunDataExtractsParameters struct{}

// BareMetalMachineRunReadCommandsParameters represents the body of request containing list of read-only commands to run on the bare metal machine.
type BareMetalMachineRunReadCommandsParameters struct{}

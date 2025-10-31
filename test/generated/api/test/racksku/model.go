package racksku

import "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/nogen/azcommonv6"

// RackSku represents the SKU information of the rack.
type RackSku struct {
  // The list of the resource properties.
  properties RackSkuProperties

  // The name of the rack SKU.
  name string

  azcommonv6.ProxyResource
}

// RackSkuProperties represents the properties of compute-related hardware for a rack. This supports both aggregator and compute racks.
type RackSkuProperties struct {
  // The list of machine SKUs and associated rack slot for the compute-dedicated machines in this rack model.
  computeMachines []MachineSkuSlot

  // The list of machine SKUs and associated rack slot for the control-plane dedicated machines in this rack model.
  controllerMachines []MachineSkuSlot

  // The free-form text describing the rack.
  description string

  // The maximum number of compute racks supported by an aggregator rack. 0 if this is a compute rack or a rack for a single rack cluster(rackType="Single").
  maxClusterSlots int64

  // The provisioning state of the rack SKU resource.
  provisioningState RackSkuProvisioningState

  // The type of the rack.
  rackType RackSkuType

  // The list of appliance SKUs and associated rack slot for the storage appliance(s) in this rack model.
  storageAppliances []StorageApplianceSkuSlot

  // The list of supported SKUs if the rack is an aggregator.
  supportedRackSkuIds []string
}

// MachineSkuSlot represents a single SKU and rack slot associated with the machine.
type MachineSkuSlot struct {
  // The list of the resource properties.
  properties MachineSkuProperties

  // The position in the rack for the machine.
  rackSlot int64
}

// MachineSkuProperties represents the properties of the machine SKU.
type MachineSkuProperties struct {
  // The type of bootstrap protocol used.
  bootstrapProtocol BootstrapProtocol

  // The count of CPU cores for this machine.
  cpuCores int64

  // The count of CPU sockets for this machine.
  cpuSockets int64

  // The list of disks.
  disks []MachineDisk

  // The generation of the architecture.
  generation string

  // The hardware version of the machine.
  hardwareVersion string

  // The maximum amount of memory. Measured in gibibytes.
  memoryCapacityGB int64

  // The model of the machine.
  model string

  // The list of network interfaces.
  networkInterfaces []NetworkInterface

  // The count of SMT and physical core threads for this machine.
  totalThreads int64

  // The make of the machine.
  vendor string
}

// Disk represents the properties of the disk.
type MachineDisk struct {
  // The maximum amount of storage. Measured in gibibytes.
  capacityGB int64

  // The connection type of the rack SKU resource.
  connection MachineSkuDiskConnectionType

  // The disk type of rack SKU resource.
  type DiskType
}

// NetworkInterface represents properties of the network interface.
type NetworkInterface struct {
  // The partial address of Peripheral Component Interconnect (PCI).
  address string

  // The connection type of the device.
  deviceConnectionType DeviceConnectionType

  // The model name of the device.
  model string

  // The physical slot for this device.
  physicalSlot int64

  // The number of ports on the device.
  portCount int64

  // The maximum amount of data in gigabits that the line card transmits through a port at any given second.
  portSpeed int64

  // The vendor name of the device.
  vendor string
}

// StorageApplianceSkuSlot represents the single SKU and rack slot associated with the storage appliance.
type StorageApplianceSkuSlot struct {
  // The list of the resource properties.
  properties StorageApplianceSkuProperties

  // The position in the rack for the storage appliance.
  rackSlot int64
}

// StorageApplianceSkuProperties represents the properties of the storage appliance SKU.
type StorageApplianceSkuProperties struct {
  // The maximum capacity of the storage appliance. Measured in gibibytes.
  capacityGB int64

  // The model of the storage appliance.
  model string
}

// RackSkuList represents a list of rack SKUs.
type RackSkuList struct {
  // The RackSku items on this page
  value []RackSku

  // The link to the next page of items
  nextLink string
}

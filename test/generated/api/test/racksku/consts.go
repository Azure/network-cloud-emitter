package racksku

// The type of bootstrap protocol used.
type BootstrapProtocol string

const (
  BootstrapProtocolPXE BootstrapProtocol = "PXE"
)

// The connection type of the rack SKU resource.
type MachineSkuDiskConnectionType string

const (
  MachineSkuDiskConnectionTypePCIE MachineSkuDiskConnectionType = "PCIE"

  MachineSkuDiskConnectionTypeSATA MachineSkuDiskConnectionType = "SATA"

  MachineSkuDiskConnectionTypeRAID MachineSkuDiskConnectionType = "RAID"

  MachineSkuDiskConnectionTypeSAS MachineSkuDiskConnectionType = "SAS"
)

// The disk type of rack SKU resource.
type DiskType string

const (
  DiskTypeHDD DiskType = "HDD"

  DiskTypeSSD DiskType = "SSD"
)

// The connection type of the device.
type DeviceConnectionType string

const (
  DeviceConnectionTypePCI DeviceConnectionType = "PCI"
)

// The provisioning state of the rack SKU resource.
type RackSkuProvisioningState string

const (
  RackSkuProvisioningStateCanceled RackSkuProvisioningState = "Canceled"

  RackSkuProvisioningStateFailed RackSkuProvisioningState = "Failed"

  RackSkuProvisioningStateSucceeded RackSkuProvisioningState = "Succeeded"
)

// The type of the rack.
type RackSkuType string

const (
  RackSkuTypeAggregator RackSkuType = "Aggregator"

  RackSkuTypeCompute RackSkuType = "Compute"

  RackSkuTypeSingle RackSkuType = "Single"
)

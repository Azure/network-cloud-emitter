package baremetalmachine

import (
  "networkcloud/baremetalmachine"
  "networkcloud/common"
)

// BareMetalMachineList represents a list of bare metal machines.
type BareMetalMachineList struct {
  // The BareMetalMachine items on this page
  value []BareMetalMachine

  // The link to the next page of items
  nextLink string
}

// BareMetalMachine represents the physical machine in the rack.
type BareMetalMachine struct {
  // The list of the resource properties.
  properties BareMetalMachineProperties

  // The name of the bare metal machine.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string

  // The extended location of the cluster associated with the resource.
  extendedLocation common.ExtendedLocation
}

// BareMetalMachineProperties represents the properties of a bare metal machine.
type BareMetalMachineProperties struct {
  // The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
  associatedResourceIds []string

  // The connection string for the baseboard management controller including IP address and protocol.
  bmcConnectionString string

  // The credentials of the baseboard management controller on this bare metal machine.
  bmcCredentials common.AdministrativeCredentials

  // The MAC address of the BMC device.
  bmcMacAddress string

  // The MAC address of a NIC connected to the PXE network.
  bootMacAddress string

  // The resource ID of the cluster this bare metal machine is associated with.
  clusterId string

  // The cordon status of the bare metal machine.
  cordonStatus BareMetalMachineCordonStatus

  // The more detailed status of the bare metal machine.
  detailedStatus BareMetalMachineDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // The hardware inventory, including information acquired from the model/sku information and from the ironic inspector.
  hardwareInventory HardwareInventory

  // The details of the latest hardware validation performed for this bare metal machine.
  hardwareValidationStatus HardwareValidationStatus

  // Field Deprecated. These fields will be empty/omitted. The list of the resource IDs for the HybridAksClusters that have nodes hosted on this bare metal machine.
  hybridAksClustersAssociatedIds []string

  // The name of this machine represented by the host object in the Cluster's Kubernetes control plane.
  kubernetesNodeName string

  // The version of Kubernetes running on this machine.
  kubernetesVersion string

  // The cluster version that has been applied to this machine during deployment or a version update.
  machineClusterVersion string

  // The custom details provided by the customer.
  machineDetails string

  // The OS-level hostname assigned to this machine.
  machineName string

  // The list of roles that are assigned to the cluster node running on this machine.
  machineRoles []string

  // The unique internal identifier of the bare metal machine SKU.
  machineSkuId string

  // The IPv4 address that is assigned to the bare metal machine during the cluster deployment.
  oamIpv4Address string

  // The IPv6 address that is assigned to the bare metal machine during the cluster deployment.
  oamIpv6Address string

  // The image that is currently provisioned to the OS disk.
  osImage string

  // The power state derived from the baseboard management controller.
  powerState BareMetalMachinePowerState

  // The provisioning state of the bare metal machine.
  provisioningState BareMetalMachineProvisioningState

  // The resource ID of the rack where this bare metal machine resides.
  rackId string

  // The rack slot in which this bare metal machine is located, ordered from the bottom up i.e. the lowest slot is 1.
  rackSlot int64

  // The indicator of whether the bare metal machine is ready to receive workloads.
  readyState BareMetalMachineReadyState

  // The runtime protection status of the bare metal machine.
  runtimeProtectionStatus RuntimeProtectionStatus

  // The list of statuses that represent secret rotation activity.
  secretRotationStatus []<Unresolved Symbol>

  // The serial number of the bare metal machine.
  serialNumber string

  // The discovered value of the machine's service tag.
  serviceTag string

  // Field Deprecated. These fields will be empty/omitted. The list of the resource IDs for the VirtualMachines that are hosted on this bare metal machine.
  virtualMachinesAssociatedIds []string
}

// HardwareInventory represents the hardware configuration of this machine as exposed to the customer, including information acquired from the model/sku information and from the ironic inspector.
type HardwareInventory struct {
  // Freeform data extracted from the environment about this machine. This information varies depending on the specific hardware and configuration.
  additionalHostInformation string

  // The list of network interfaces and associated details for the bare metal machine.
  interfaces []<Unresolved Symbol>

  // Field Deprecated. Will be removed in an upcoming version. The list of network interface cards and associated details for the bare metal machine.
  nics []<Unresolved Symbol>
}

// HardwareValidationStatus represents the latest hardware validation details performed for this bare metal machine.
type HardwareValidationStatus struct {
  // The timestamp of the hardware validation execution.
  lastValidationTime time.Time

  // The outcome of the hardware validation.
  result BareMetalMachineHardwareValidationResult
}

// RuntimeProtectionStatus represents the runtime protection status of the bare metal machine.
type RuntimeProtectionStatus struct {
  // The timestamp when the malware definitions were last updated.
  definitionsLastUpdated time.Time

  // The version of the malware definitions.
  definitionsVersion string

  // The timestamp of the most recently completed scan, or empty if there has never been a scan.
  scanCompletedTime time.Time

  // The timestamp of the most recently scheduled scan, or empty if no scan has been scheduled.
  scanScheduledTime time.Time

  // The timestamp of the most recently started scan, or empty if there has never been a scan.
  scanStartedTime time.Time
}

// BareMetalMachinePatchParameters represents the body of the request to patch bare metal machine properties.
type BareMetalMachinePatchParameters struct {
  // The list of the resource properties.
  properties BareMetalMachinePatchProperties

  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

// BareMetalMachinePatchProperties represents the properties of the bare metal machine that can be patched.
type BareMetalMachinePatchProperties struct {
  // The details provided by the customer during the creation of rack manifests
  // that allows for custom data to be associated with this machine.
  machineDetails string
}

// BareMetalMachineCordonParameters represents the body of the request to evacuate workloads from node on a bare metal machine.
type BareMetalMachineCordonParameters struct {
  // The indicator of whether to evacuate the node workload when the bare metal machine is cordoned.
  evacuate BareMetalMachineEvacuate
}

// BareMetalMachinePowerOffParameters represents the body of the request to power off bare metal machine.
type BareMetalMachinePowerOffParameters struct {
  // The indicator of whether to skip the graceful OS shutdown and power off the bare metal machine immediately.
  skipShutdown BareMetalMachineSkipShutdown
}

// BareMetalMachineReplaceParameters represents the body of the request to physically swap a bare metal machine for another.
type BareMetalMachineReplaceParameters struct {
  // The credentials of the baseboard management controller on this bare metal machine. The password field is expected to be an Azure Key Vault key URL. Until the cluster is converted to utilize managed identity by setting the secret archive settings, the actual password value should be provided instead.
  bmcCredentials common.AdministrativeCredentials

  // The MAC address of the BMC device.
  bmcMacAddress string

  // The MAC address of a NIC connected to the PXE network.
  bootMacAddress string

  // The OS-level hostname assigned to this machine.
  machineName string

  // The serial number of the bare metal machine.
  serialNumber string
}

// BareMetalMachineRunCommandParameters represents the body of the request to execute a script on the bare metal machine.
type BareMetalMachineRunCommandParameters struct {
  // The list of string arguments that will be passed to the script in order as separate arguments.
  arguments []string

  // The maximum time the script is allowed to run.
  // If the execution time exceeds the maximum, the script will be stopped, any output produced until then will be captured, and the exit code matching a timeout will be returned (252).
  limitTimeSeconds int64

  // The base64 encoded script to execute on the bare metal machine.
  script string
}

// BareMetalMachineRunDataExtractsParameters represents the body of request containing list of curated data extraction commands to run on the bare metal machine.
type BareMetalMachineRunDataExtractsParameters struct {
  // The list of curated data extraction commands to be executed directly against the target machine.
  commands []<Unresolved Symbol>

  // The maximum time the commands are allowed to run.
  // If the execution time exceeds the maximum, the script will be stopped, any output produced until then will be captured, and the exit code matching a timeout will be returned (252).
  limitTimeSeconds int64
}

// BareMetalMachineRunReadCommandsParameters represents the body of request containing list of read-only commands to run on the bare metal machine.
type BareMetalMachineRunReadCommandsParameters struct {
  // The list of read-only commands to be executed directly against the target machine.
  commands []<Unresolved Symbol>

  // The maximum time the commands are allowed to run.
  // If the execution time exceeds the maximum, the script will be stopped, any output produced until then will be captured, and the exit code matching a timeout will be returned (252).
  limitTimeSeconds int64
}

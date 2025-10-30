package virtualmachine

import (
  "networkcloud/common"
  "networkcloud/console"
  console_2 "networkcloud/console"
  "networkcloud/virtualmachine"
)

// VirtualMachineList represents a list of virtual machines.
type VirtualMachineList struct {
  // The VirtualMachine items on this page
  value []VirtualMachine

  // The link to the next page of items
  nextLink string
}

// VirtualMachine represents the on-premises Network Cloud virtual machine.
type VirtualMachine struct {
  // The list of the resource properties.
  properties VirtualMachineProperties

  // The name of the virtual machine.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string

  // The extended location of the cluster associated with the resource.
  extendedLocation common.ExtendedLocation
}

// VirtualMachineProperties represents the properties of the virtual machine.
type VirtualMachineProperties struct {
  // The name of the administrator to which the ssh public keys will be added into the authorized keys.
  adminUsername string

  // The cluster availability zone containing this virtual machine.
  availabilityZone string

  // The resource ID of the bare metal machine that hosts the virtual machine.
  bareMetalMachineId string

  // Selects the boot method for the virtual machine.
  bootMethod VirtualMachineBootMethod

  // The cloud service network that provides platform-level services for the virtual machine.
  cloudServicesNetworkAttachment NetworkAttachment

  // The resource ID of the cluster the virtual machine is created for.
  clusterId string

  // The extended location to use for creation of a VM console resource.
  consoleExtendedLocation common.ExtendedLocation

  // The number of CPU cores in the virtual machine.
  cpuCores int64

  // The more detailed status of the virtual machine.
  detailedStatus VirtualMachineDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // Field Deprecated, the value will be ignored if provided. The indicator of whether one of the specified CPU cores is isolated to run the emulator thread for this virtual machine.
  isolateEmulatorThread VirtualMachineIsolateEmulatorThread

  // The memory size of the virtual machine. Allocations are measured in gibibytes.
  memorySizeGB int64

  // The list of network attachments to the virtual machine.
  networkAttachments []NetworkAttachment

  // The Base64 encoded cloud-init network data.
  networkData string

  // The scheduling hints for the virtual machine.
  placementHints []<Unresolved Symbol>

  // The power state of the virtual machine.
  powerState VirtualMachinePowerState

  // The provisioning state of the virtual machine.
  provisioningState VirtualMachineProvisioningState

  // The list of ssh public keys. Each key will be added to the virtual machine using the cloud-init ssh_authorized_keys mechanism for the adminUsername.
  sshPublicKeys []console_2.SshPublicKey

  // The storage profile that specifies size and other parameters about the disks related to the virtual machine.
  storageProfile StorageProfile

  // The Base64 encoded cloud-init user data.
  userData string

  // Field Deprecated, use virtualizationModel instead. The type of the virtio interface.
  virtioInterface VirtualMachineVirtioInterfaceType

  // The type of the device model to use.
  vmDeviceModel VirtualMachineDeviceModelType

  // The virtual machine image that is currently provisioned to the OS disk, using the full url and tag notation used to pull the image.
  vmImage string

  // The credentials used to login to the image repository that has access to the specified image.
  vmImageRepositoryCredentials ImageRepositoryCredentials

  // The resource IDs of volumes that are attached to the virtual machine.
  volumes []string
}

// NetworkAttachment represents the single network attachment.
type NetworkAttachment struct {
  // The resource ID of the associated network attached to the virtual machine.
  // It can be one of cloudServicesNetwork, l3Network, l2Network or trunkedNetwork resources.
  attachedNetworkId string

  // The indicator of whether this is the default gateway.
  // Only one of the attached networks (including the CloudServicesNetwork attachment) for a single machine may be specified as True.
  defaultGateway DefaultGateway

  // The IP allocation mechanism for the virtual machine.
  // Dynamic and Static are only valid for l3Network which may also specify Disabled.
  // Otherwise, Disabled is the only permitted value.
  ipAllocationMethod VirtualMachineIPAllocationMethod

  // The IPv4 address of the virtual machine.
  //
  // This field is used only if the attached network has IPAllocationType of IPV4 or DualStack.
  //
  // If IPAllocationMethod is:
  // Static - this field must contain a user specified IPv4 address from within the subnet specified in the attached network.
  // Dynamic - this field is read-only, but will be populated with an address from within the subnet specified in the attached network.
  // Disabled - this field will be empty.
  ipv4Address string

  // The IPv6 address of the virtual machine.
  //
  // This field is used only if the attached network has IPAllocationType of IPV6 or DualStack.
  //
  // If IPAllocationMethod is:
  // Static - this field must contain an IPv6 address range from within the range specified in the attached network.
  // Dynamic - this field is read-only, but will be populated with an range from within the subnet specified in the attached network.
  // Disabled - this field will be empty.
  ipv6Address string

  // The MAC address of the interface for the virtual machine that corresponds to this network attachment.
  macAddress string

  // The associated network's interface name.
  // If specified, the network attachment name has a maximum length of 15 characters and must be unique to this virtual machine.
  // If the user doesn’t specify this value, the default interface name of the network resource will be used.
  // For a CloudServicesNetwork resource, this name will be ignored.
  networkAttachmentName string
}

// StorageProfile represents information about a disk.
type StorageProfile struct {
  // The disk to use with this virtual machine.
  osDisk OsDisk

  // The resource IDs of volumes that are requested to be attached to the virtual machine.
  volumeAttachments []string
}

// OsDisk represents configuration of the boot disk.
type OsDisk struct {
  // The strategy for creating the OS disk.
  createOption OsDiskCreateOption

  // The strategy for deleting the OS disk.
  deleteOption OsDiskDeleteOption

  // The size of the disk. Required if the createOption is Ephemeral. Allocations are measured in gibibytes.
  diskSizeGB int64
}

// ImageRepositoryCredentials represents the credentials used to login to the image repository.
type ImageRepositoryCredentials struct {
  // The password or token used to access an image in the target repository.
  password string

  // The URL of the authentication server used to validate the repository credentials.
  registryUrl string

  // The username used to access an image in the target repository.
  username string
}

// VirtualMachinePatchParameters represents the body of the request to patch the virtual machine.
type VirtualMachinePatchParameters struct {
  // The list of the resource properties.
  properties VirtualMachinePatchProperties

  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

// VirtualMachinePatchProperties represents the properties of the virtual machine that can be patched.
type VirtualMachinePatchProperties struct {
  // The credentials used to login to the image repository that has access to the specified image.
  vmImageRepositoryCredentials ImageRepositoryCredentials
}

// VirtualMachinePowerOffParameters represents the body of the request to power off virtual machine.
type VirtualMachinePowerOffParameters struct {
  // The indicator of whether to skip the graceful OS shutdown and power off the virtual machine immediately.
  skipShutdown SkipShutdown
}

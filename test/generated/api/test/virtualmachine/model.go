package virtualmachine

import (
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/nogen/azcommonv6"
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/test/common"
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

  azcommonv6.TrackedResource
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
  cloudServicesNetworkAttachment common.NetworkAttachment

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
  networkAttachments []common.NetworkAttachment

  // The Base64 encoded cloud-init network data.
  networkData string

  // The scheduling hints for the virtual machine.
  placementHints []VirtualMachinePlacementHint

  // The power state of the virtual machine.
  powerState VirtualMachinePowerState

  // The provisioning state of the virtual machine.
  provisioningState VirtualMachineProvisioningState

  // The list of ssh public keys. Each key will be added to the virtual machine using the cloud-init ssh_authorized_keys mechanism for the adminUsername.
  sshPublicKeys []common.SshPublicKey

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

// VirtualMachinePlacementHint represents a single scheduling hint of the virtual machine.
type VirtualMachinePlacementHint struct {
  // The specification of whether this hint supports affinity or anti-affinity with the referenced resources.
  hintType VirtualMachinePlacementHintType

  // The resource ID of the target object that the placement hints will be checked against, e.g., the bare metal node to host the virtual machine.
  resourceId string

  // The indicator of whether the hint is a hard or soft requirement during scheduling.
  schedulingExecution VirtualMachineSchedulingExecution

  // The scope for the virtual machine affinity or anti-affinity placement hint. It should always be "Machine" in the case of node affinity.
  scope VirtualMachinePlacementHintPodAffinityScope
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

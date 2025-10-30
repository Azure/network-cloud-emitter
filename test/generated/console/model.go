package console

import (
  "networkcloud/common"
  "networkcloud/console"
)

// Console represents the console of an on-premises Network Cloud virtual machine.
type Console struct {
  // The list of the resource properties.
  properties ConsoleProperties

  // The name of the virtual machine console.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string

  // The extended location of the cluster manager associated with the cluster this virtual machine is created on.
  extendedLocation common.ExtendedLocation
}

// ConsoleProperties represents the properties of the virtual machine console.
type ConsoleProperties struct {
  // The more detailed status of the console.
  detailedStatus ConsoleDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // The indicator of whether the console access is enabled.
  enabled ConsoleEnabled

  // The date and time after which the key will be disallowed access.
  expiration time.Time

  // The resource ID of the private link service that is used to provide virtual machine console access.
  privateLinkServiceId string

  // The provisioning state of the virtual machine console.
  provisioningState ConsoleProvisioningState

  // The SSH public key that will be provisioned for user access. The user is expected to have the corresponding SSH private key for logging in.
  sshPublicKey SshPublicKey

  // The unique identifier for the virtual machine that is used to access the console.
  virtualMachineAccessId string
}

// SshPublicKey represents the public key used to authenticate with a resource through SSH.
type SshPublicKey struct {
  // The SSH public key data.
  keyData string
}

// ConsolePatchParameters represents the body of the request to patch the virtual machine console.
type ConsolePatchParameters struct {
  // The list of the resource properties.
  properties ConsolePatchProperties

  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

// ConsolePatchProperties represents the properties of the virtual machine console that can be patched.
type ConsolePatchProperties struct {
  // The indicator of whether the console access is enabled.
  enabled ConsoleEnabled

  // The date and time after which the key will be disallowed access.
  expiration time.Time

  // The SSH public key that will be provisioned for user access. The user is expected to have the corresponding SSH private key for logging in.
  sshPublicKey SshPublicKey
}

// ConsoleList represents a list of virtual machine consoles.
type ConsoleList struct {
  // The Console items on this page
  value []Console

  // The link to the next page of items
  nextLink string
}

package baremetalmachinekeyset

import (
  "time"

  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/nogen/azcommonv6"
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/test/common"
)

// BareMetalMachineKeySet represents the bare metal machine key set.
type BareMetalMachineKeySet struct {
  // The list of the resource properties.
  properties BareMetalMachineKeySetProperties

  // The name of the bare metal machine key set.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string

  // The extended location of the cluster associated with the resource.
  extendedLocation common.ExtendedLocation

  azcommonv6.TrackedResource
}

// BareMetalMachineKeySetProperties represents the properties of bare metal machine key set.
type BareMetalMachineKeySetProperties struct {
  // The object ID of Azure Active Directory group that all users in the list must be in for access to be granted. Users that are not in the group will not have access.
  azureGroupId string

  // The more detailed status of the key set.
  detailedStatus BareMetalMachineKeySetDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // The date and time after which the users in this key set will be removed from the bare metal machines.
  expiration time.Time

  // The list of IP addresses of jump hosts with management network access from which a login will be allowed for the users.
  jumpHostsAllowed []string

  // The last time this key set was validated.
  lastValidation time.Time

  // The name of the group that users will be assigned to on the operating system of the machines.
  osGroupName string

  // The access level allowed for the users in this key set.
  privilegeLevel BareMetalMachineKeySetPrivilegeLevel

  // The provisioning state of the bare metal machine key set.
  provisioningState BareMetalMachineKeySetProvisioningState

  // The unique list of permitted users.
  userList []common.KeySetUser

  // The status evaluation of each user.
  userListStatus []common.KeySetUserStatus
}

// BareMetalMachineKeySetPatchParameters represents the body of the request to patch the bare metal machine key set.
type BareMetalMachineKeySetPatchParameters struct {
  // The list of the resource properties.
  properties BareMetalMachineKeySetPatchProperties

  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

// BareMetalMachineKeySetPatchProperties represents the properties of bare metal machine key set that can be patched.
type BareMetalMachineKeySetPatchProperties struct {
  // The date and time after which the users in this key set will be removed from the bare metal machines.
  expiration time.Time

  // The list of IP addresses of jump hosts with management network access from which a login will be allowed for the users.
  jumpHostsAllowed []string

  // The unique list of permitted users.
  userList []common.KeySetUser
}

// BareMetalMachineKeySetList represents a list of bare metal machine key sets.
type BareMetalMachineKeySetList struct {
  // The BareMetalMachineKeySet items on this page
  value []BareMetalMachineKeySet

  // The link to the next page of items
  nextLink string
}

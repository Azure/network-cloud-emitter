package bmckeyset

import (
  "networkcloud/bmckeyset"
  "networkcloud/common"
)

// BmcKeySet represents the baseboard management controller key set.
type BmcKeySet struct {
  // The list of the resource properties.
  properties BmcKeySetProperties

  // The name of the baseboard management controller key set.
  name string

  // If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
  eTag string

  // The extended location of the cluster associated with the resource.
  extendedLocation common.ExtendedLocation
}

// BmcKeySetProperties represents the properties of baseboard management controller key set.
type BmcKeySetProperties struct {
  // The object ID of Azure Active Directory group that all users in the list must be in for access to be granted. Users that are not in the group will not have access.
  azureGroupId string

  // The more detailed status of the key set.
  detailedStatus BmcKeySetDetailedStatus

  // The descriptive message about the current detailed status.
  detailedStatusMessage string

  // The date and time after which the users in this key set will be removed from the baseboard management controllers.
  expiration time.Time

  // The last time this key set was validated.
  lastValidation time.Time

  // The access level allowed for the users in this key set.
  privilegeLevel BmcKeySetPrivilegeLevel

  // The provisioning state of the baseboard management controller key set.
  provisioningState BmcKeySetProvisioningState

  // The unique list of permitted users.
  userList []<Unresolved Symbol>

  // The status evaluation of each user.
  userListStatus []<Unresolved Symbol>
}

// BmcKeySetPatchParameters represents the body of the request to patch the baseboard management controller key set.
type BmcKeySetPatchParameters struct {
  // The list of the resource properties.
  properties BmcKeySetPatchProperties

  // The Azure resource tags that will replace the existing ones.
  tags map[string]string
}

// BmcKeySetPatchProperties represents the properties of baseboard management controller key set that are patchable.
type BmcKeySetPatchProperties struct {
  // The date and time after which the users in this key set will be removed from the baseboard management controllers.
  expiration time.Time

  // The unique list of permitted users.
  userList []<Unresolved Symbol>
}

// BmcKeySetList represents a list of baseboard management controller key sets.
type BmcKeySetList struct {
  // The BmcKeySet items on this page
  value []BmcKeySet

  // The link to the next page of items
  nextLink string
}

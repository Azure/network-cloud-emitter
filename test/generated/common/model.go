package common

import "networkcloud/console"

// ExtendedLocation represents the Azure custom location where the resource will be created.
type ExtendedLocation struct {
  // The resource ID of the extended location on which the resource will be created.
  name string

  // The extended location type, for example, CustomLocation.
  type string
}

// ManagedResourceGroupConfiguration represents the configuration of the resource group managed by Azure.
type ManagedResourceGroupConfiguration struct {
  // The location of the managed resource group. If not specified, the location of the parent resource is chosen.
  location string

  // The name for the managed resource group. If not specified, the unique name is automatically generated.
  name string
}

// AdministrativeCredentials represents the admin credentials for the device requiring password-based authentication.
type AdministrativeCredentials struct {
  // The password of the administrator of the device used during initialization.
  password string

  // The username of the administrator of the device used during initialization.
  username string
}

// AdministratorConfiguration represents the administrative credentials that will be applied to the control plane and agent pool nodes in Kubernetes clusters.
type AdministratorConfiguration struct {
  // The user name for the administrator that will be applied to the operating systems that run Kubernetes nodes. If not supplied, a user name will be chosen by the service.
  adminUsername string

  // The SSH configuration for the operating systems that run the nodes in the Kubernetes cluster. In some cases, specification of public keys may be required to produce a working environment.
  sshPublicKeys []console.SshPublicKey
}

// AttachedNetworkConfiguration represents the set of workload networks to attach to a resource.
type AttachedNetworkConfiguration struct {
  // The list of Layer 2 Networks and related configuration for attachment.
  l2Networks []<Unresolved Symbol>

  // The list of Layer 3 Networks and related configuration for attachment.
  l3Networks []<Unresolved Symbol>

  // The list of Trunked Networks and related configuration for attachment.
  trunkedNetworks []<Unresolved Symbol>
}

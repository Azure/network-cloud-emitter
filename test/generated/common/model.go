package common

// ExtendedLocation represents the Azure custom location where the resource will be created.
type ExtendedLocation struct{}

// ManagedResourceGroupConfiguration represents the configuration of the resource group managed by Azure.
type ManagedResourceGroupConfiguration struct{}

// AdministrativeCredentials represents the admin credentials for the device requiring password-based authentication.
type AdministrativeCredentials struct{}

// AdministratorConfiguration represents the administrative credentials that will be applied to the control plane and agent pool nodes in Kubernetes clusters.
type AdministratorConfiguration struct{}

// AttachedNetworkConfiguration represents the set of workload networks to attach to a resource.
type AttachedNetworkConfiguration struct{}

package racksku

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

package racksku

import "networkcloud/racksku"

// RackSku represents the SKU information of the rack.
type RackSku struct {
  // The list of the resource properties.
  properties RackSkuProperties

  // The name of the rack SKU.
  name string
}

// RackSkuProperties represents the properties of compute-related hardware for a rack. This supports both aggregator and compute racks.
type RackSkuProperties struct {
  // The list of machine SKUs and associated rack slot for the compute-dedicated machines in this rack model.
  computeMachines []<Unresolved Symbol>

  // The list of machine SKUs and associated rack slot for the control-plane dedicated machines in this rack model.
  controllerMachines []<Unresolved Symbol>

  // The free-form text describing the rack.
  description string

  // The maximum number of compute racks supported by an aggregator rack. 0 if this is a compute rack or a rack for a single rack cluster(rackType="Single").
  maxClusterSlots int64

  // The provisioning state of the rack SKU resource.
  provisioningState RackSkuProvisioningState

  // The type of the rack.
  rackType RackSkuType

  // The list of appliance SKUs and associated rack slot for the storage appliance(s) in this rack model.
  storageAppliances []<Unresolved Symbol>

  // The list of supported SKUs if the rack is an aggregator.
  supportedRackSkuIds []string
}

// RackSkuList represents a list of rack SKUs.
type RackSkuList struct {
  // The RackSku items on this page
  value []RackSku

  // The link to the next page of items
  nextLink string
}

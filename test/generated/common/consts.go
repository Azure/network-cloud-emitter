package common

// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
type HybridAksPluginType string

const (
  HybridAksPluginTypeDPDK HybridAksPluginType = "DPDK"

  HybridAksPluginTypeSRIOV HybridAksPluginType = "SRIOV"

  HybridAksPluginTypeOSDevice HybridAksPluginType = "OSDevice"
)

import { createModule, StrictDescriptor } from "./create-module.js";

export const armCommon = createModule(
  "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/nogen/azcommonv6",
  {
    kind: "package",
    path: "dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git/api/nogen/azcommonv6",
    members: {
      ProxyResource: { kind: "var" },
      TrackedResource: { kind: "var" },
    },
  } satisfies StrictDescriptor,
);

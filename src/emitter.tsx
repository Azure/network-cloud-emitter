import { Output, renderAsync, writeOutput } from "@alloy-js/core";
import * as go from "@alloy-js/go";
import { EmitContext } from "@typespec/compiler";
import { Resources } from "./components/resources.jsx";
import { ResourceManager, ResourceManagerContext } from "./context/resource-manager.js";
import { TspContext } from "./context/tsp-context.js";

export async function $onEmit(context: EmitContext) {
  const output = await renderAsync(
    <TspContext.Provider value={{ program: context.program }}>
      <ResourceManagerContext.Provider value={{ resourceManager: new ResourceManager() }}>
        <Output>
          <go.ModuleDirectory name="dev.azure.com/msazuredev/AzureForOperatorsIndustry/_git/nc-rp-api.git">
            <Resources />
          </go.ModuleDirectory>
        </Output>
      </ResourceManagerContext.Provider>
    </TspContext.Provider>,
  );

  writeOutput(output, context.emitterOutputDir);
}

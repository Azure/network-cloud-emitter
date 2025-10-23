import { Output, renderAsync, writeOutput } from "@alloy-js/core";
import { EmitContext, listServices, Model } from "@typespec/compiler";
import { Package } from "./components/package.jsx";
import { Structs } from "./components/structs.jsx";
import { TspContext } from "./context/tsp-context.js";

export async function $onEmit(context: EmitContext) {
  const models: Model[] = [];
  const services = listServices(context.program);
  for (const service of services) {
    models.push(...service.type.models.values());
  }

  const output = await renderAsync(
    <TspContext.Provider value={{ program: context.program }}>
      <Output>
        <Package name="main">
          <Structs types={models} />
        </Package>
      </Output>
    </TspContext.Provider>,
  );

  writeOutput(output, context.emitterOutputDir);
}

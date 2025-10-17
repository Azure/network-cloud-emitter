import { EmitContext } from "@typespec/compiler";
import { Output, writeOutput } from "@typespec/emitter-framework";

export async function $onEmit(context: EmitContext) {
  await writeOutput(context.program, <Output program={context.program} />, context.emitterOutputDir);
}

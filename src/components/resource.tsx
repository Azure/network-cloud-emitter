import { For } from "@alloy-js/core";
import * as go from "@alloy-js/go";
import { useResourceManager as useResourceManagerContext } from "../context/resource-manager.js";
import { StructDeclaration } from "./struct-declaration.js";

export interface ResourceProps {
  name: string;
}

export function Resource(props: ResourceProps) {
  const resourceManager = useResourceManagerContext().resourceManager;

  const models = [];
  const consts = [];

  for (const type of resourceManager.resources.get(props.name) || []) {
    if (type.kind === "Model") {
      models.push(type);
    } else if (type.kind === "Union") {
      consts.push(type);
    }
  }

  return (
    <go.SourceDirectory path={props.name.toLowerCase()}>
      <go.SourceFile path="model.go">
        <For each={models} doubleHardline>
          {(model) => <StructDeclaration type={model} />}
        </For>
      </go.SourceFile>

      <go.SourceFile path="consts.go"></go.SourceFile>
    </go.SourceDirectory>
  );
}

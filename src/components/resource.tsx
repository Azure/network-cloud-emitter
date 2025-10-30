import { For } from "@alloy-js/core";
import * as go from "@alloy-js/go";
import { useResourceManager as useResourceManagerContext } from "../context/resource-manager.js";
import { OpenEnumDeclaration } from "./open-enum-declaration.jsx";
import { StructDeclaration } from "./struct-type-declaration.jsx";

export interface ResourceProps {
  name: string;
}

export function Resource(props: ResourceProps) {
  const resourceManager = useResourceManagerContext().resourceManager;

  const models = [];
  const enums = [];

  for (const type of resourceManager.resources.get(props.name) || []) {
    if (type.kind === "Model") {
      models.push(type);
    } else if (type.kind === "Union") {
      enums.push(type);
    }
  }

  return (
    <go.SourceDirectory path={props.name.toLowerCase()}>
      <go.SourceFile path="model.go" package={props.name.toLowerCase()}>
        <For each={models} doubleHardline>
          {(m) => <StructDeclaration type={m} />}
        </For>
      </go.SourceFile>

      <go.SourceFile path="consts.go">
        <For each={enums} doubleHardline>
          {(e) => <OpenEnumDeclaration type={e} />}
        </For>
      </go.SourceFile>
    </go.SourceDirectory>
  );
}

import { For } from "@alloy-js/core";
import * as go from "@alloy-js/go";
import { Model } from "@typespec/compiler";
import { useTsp } from "../context/tsp-context.js";

export interface StructDeclarationProps extends Omit<go.StructTypeDeclarationProps, "name"> {
  type: Model;
}

export function StructDeclaration(props: StructDeclarationProps) {
  const { $ } = useTsp();
  return <go.StructTypeDeclaration name={props.type.name} doc={$.type.getDoc(props.type)}></go.StructTypeDeclaration>;
}

export interface StructsProps {
  path?: string;
  types: Model[];
}

export function Structs(props: StructsProps) {
  return (
    <go.SourceFile path={props.path ?? "structs.go"}>
      <For each={props.types} doubleHardline>
        {(type) => <StructDeclaration type={type} />}
      </For>
    </go.SourceFile>
  );
}

import { For, refkey, Show } from "@alloy-js/core";
import * as go from "@alloy-js/go";
import { Model } from "@typespec/compiler";
import { useTsp } from "../context/tsp-context.js";
import { TypeExpression } from "./type-expression.jsx";

export interface StructDeclarationProps extends Omit<go.StructTypeDeclarationProps, "name"> {
  name?: string;
  type: Model;
}

export function StructDeclaration(props: StructDeclarationProps) {
  const { $ } = useTsp();

  return (
    <go.StructTypeDeclaration
      name={props.name ?? props.type.name}
      refkey={props.refkey ?? refkey(props.type)}
      doc={$.type.getDoc(props.type)}
    >
      <For each={props.type.properties.entries()} doubleHardline>
        {([propName, prop]) => (
          <go.StructMember
            name={propName}
            refkey={refkey(prop)}
            doc={$.type.getDoc(prop)}
            type={<TypeExpression type={prop.type} />}
          />
        )}
      </For>

      <Show when={props.type.baseModel && props.type.baseModel.kind === "Model"}>
        <hbr />
        <hbr />
        <go.StructEmbed>
          <TypeExpression type={props.type.baseModel!} />
        </go.StructEmbed>
      </Show>
    </go.StructTypeDeclaration>
  );
}

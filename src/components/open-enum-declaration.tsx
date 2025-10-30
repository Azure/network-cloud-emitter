import { For, refkey, Show } from "@alloy-js/core";
import * as go from "@alloy-js/go";
import { getUnionAsEnum } from "@azure-tools/typespec-azure-core";
import { ignoreDiagnostics, Scalar, type Union } from "@typespec/compiler";
import { useTsp } from "../context/tsp-context.js";
import { reportDiagnostic } from "../lib.js";
import { TypeExpression } from "./type-expression.jsx";

export interface OpenEnumDeclarationProps extends Omit<go.TypeDeclarationProps, "name"> {
  name?: string;
  type: Union;
}

export function OpenEnumDeclaration(props: OpenEnumDeclarationProps) {
  const { $ } = useTsp();
  if (ignoreDiagnostics(getUnionAsEnum(props.type)) === undefined) {
    throw new Error("The provided union type cannot be represented as an open enum");
  }

  const type = ignoreDiagnostics(getUnionAsEnum(props.type))!;

  if (!props.type.name || props.type.name === "") {
    reportDiagnostic($.program, { code: "type-declaration-missing-name", target: props.type });
  }

  const name = props.name ?? go.useGoNamePolicy().getName(props.type.name!, "type");
  const doc = props.doc ?? $.type.getDoc(props.type);

  return (
    <>
      <go.TypeDeclaration name={name} doc={doc} refkey={refkey(props.type)}>
        <TypeExpression type={getUnionAsEnumValueType(props.type)!} />
      </go.TypeDeclaration>
      <hbr />
      <hbr />
      <go.VariableDeclarationGroup const={true}>
        <For each={type.flattenedMembers.keys()} doubleHardline>
          {(key) => {
            return (
              <go.VariableDeclaration
                doc={$.type.getDoc(type.flattenedMembers.get(key)!.type)}
                name={`${name}${key.toString()}`}
                type={refkey(props.type)}
                const={true}
                refkey={refkey(type.flattenedMembers.get(key)!.type)}
              >
                <Show when={type.flattenedMembers.get(key)!.value !== undefined}>
                  {typeof type.flattenedMembers.get(key)!.value === "string"
                    ? `"${type.flattenedMembers.get(key)!.value}"`
                    : type.flattenedMembers.get(key)!.value}
                </Show>
              </go.VariableDeclaration>
            );
          }}
        </For>
      </go.VariableDeclarationGroup>
    </>
  );
}

function getUnionAsEnumValueType(union: Union): Scalar | undefined {
  for (const variant of union.variants.values()) {
    if (variant.type.kind === "Union") {
      const ret = getUnionAsEnumValueType(variant.type);
      if (ret) return ret;
    } else if (variant.type.kind === "Scalar") {
      return variant.type;
    }
  }
  return undefined;
}

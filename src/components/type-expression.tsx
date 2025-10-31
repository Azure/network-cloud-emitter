import { code, refkey, type Children } from "@alloy-js/core";
import { Pointer } from "@alloy-js/go";
import {
  getNamespaceFullName,
  getTypeName,
  isVoidType,
  type IntrinsicType,
  type Scalar,
  type Type,
} from "@typespec/compiler";
import type { Typekit } from "@typespec/compiler/typekit";
import { useTsp } from "../context/tsp-context.js";
import { armCommon } from "../external/arm-common.js";
import { time } from "../external/time.js";
import { reportDiagnostic } from "../lib.js";

export interface TypeExpressionProps {
  type: Type;
}

export function TypeExpression(props: TypeExpressionProps): Children {
  if (props.type.kind === "Union") {
    const nullableType = getNullableUnionInnerType(props.type);
    if (nullableType) {
      return (
        <Pointer>
          <TypeExpression type={nullableType} />
        </Pointer>
      );
    }
  }

  const { $ } = useTsp();

  if (isDeclaration($, props.type)) {
    if (
      props.type.kind === "Model" &&
      props.type.namespace &&
      getNamespaceFullName(props.type.namespace) === "Azure.ResourceManager.CommonTypes"
    ) {
      return <>{(armCommon as any)[props.type.name]}</>;
    }
    return code`${refkey(props.type)}`;
  }

  if ($.scalar.is(props.type)) {
    return getScalarIntrinsicExpression($, props.type);
  } else if ($.array.is(props.type)) {
    return code`[]${(<TypeExpression type={props.type.indexer.value} />)}`;
  } else if ($.record.is(props.type)) {
    return code`map[string]${(<TypeExpression type={props.type.indexer.value} />)}`;
  } else if ($.literal.isString(props.type)) {
    // Go doesn't have literal types, so we map them to their corresponding Go types
    return code`string`;
  } else if ($.literal.isNumeric(props.type)) {
    return Number.isInteger(props.type.value) ? code`int` : code`float64`;
  } else if ($.literal.isBoolean(props.type)) {
    return code`bool`;
  } else if (isVoidType(props.type)) {
    return code`interface{}`;
  }

  throw new Error(`Unsupported type for TypeExpression: ${props.type.kind} (${getTypeName(props.type)})`);
}

// Go type mappings for TypeSpec intrinsic types
const intrinsicNameToGoType = new Map<string, Children | null>([
  // Core types
  ["unknown", code`any`], // Matches Go's any
  ["string", code`string`], // Matches Go's string
  ["boolean", code`bool`], // Matches Go's bool
  ["null", code`nil`], // Matches Go's nil
  ["void", code`nil`], // No direct equivalent, use nil
  ["never", null], // No direct equivalent in Go
  ["bytes", code`[]byte`], // Matches Go's byte slice

  // Numeric types
  ["numeric", code`float64`], // Parent type for all numeric types, use most precise
  ["integer", code`int`], // Broad integer category, maps to int
  ["float", code`float64`], // Broad float category, maps to float64
  ["decimal", code`float64`], // Decimal numbers as float64
  ["decimal128", code`float64`], // 128-bit decimal as float64
  ["int64", code`int64`], // 64-bit signed integer
  ["int32", code`int32`], // 32-bit signed integer
  ["int16", code`int16`], // 16-bit signed integer
  ["int8", code`int8`], // 8-bit signed integer
  ["safeint", code`int`], // Safe integer, use int as default
  ["uint64", code`uint64`], // 64-bit unsigned integer
  ["uint32", code`uint32`], // 32-bit unsigned integer
  ["uint16", code`uint16`], // 16-bit unsigned integer
  ["uint8", code`uint8`], // 8-bit unsigned integer (byte)
  ["float32", code`float32`], // 32-bit floating point
  ["float64", code`float64`], // 64-bit floating point

  // Date and time types
  ["plainDate", <>{time.Time}</>], // Use time.Time for dates
  ["plainTime", <>{time.Time}</>], // Use time.Time for times
  ["utcDateTime", <>{time.Time}</>], // Use time.Time for UTC date-times
  ["offsetDateTime", <>{time.Time}</>], // Use time.Time for timezone-specific date-times
  ["duration", <>{time.Duration}</>], // Duration as time.Duration

  // String types
  ["url", code`string`], // URLs as strings in Go (or could use url.URL)
]);

export function getScalarIntrinsicExpression($: Typekit, type: Scalar | IntrinsicType): Children | null {
  let intrinsicName: string;

  if ($.scalar.isUtcDateTime(type) || $.scalar.extendsUtcDateTime(type)) {
    return <>{time.Time}</>;
  }
  if ($.scalar.is(type)) {
    intrinsicName = $.scalar.getStdBase(type)?.name ?? "";
  } else {
    intrinsicName = type.name;
  }

  const goType = intrinsicNameToGoType.get(intrinsicName);

  if (!goType) {
    reportDiagnostic($.program, { code: "go-unsupported-scalar", target: type });
    return code`any`; // Fallback to any if unsupported
  }

  return goType;
}

function isDeclaration($: Typekit, type: Type): boolean {
  switch (type.kind) {
    case "Namespace":
    case "Interface":
    case "Enum":
    case "Operation":
    case "EnumMember":
      return true;
    case "UnionVariant":
      return false;

    case "Model":
      if ($.array.is(type) || $.record.is(type)) {
        return false;
      }
      return true;
    case "Union":
      return Boolean(type.name);
    default:
      return false;
  }
}

function getNullableUnionInnerType(union: Type): Type | null {
  if (union.kind !== "Union") {
    return null;
  }

  const variants = [...union.variants.values()];

  // Check if this is a nullable union (type | null)
  if (variants.length === 2) {
    const nullVariant = variants.find((v) => v.type.kind === "Intrinsic" && v.type.name === "null");
    if (nullVariant) {
      const nonNullVariant = variants.find((v) => v !== nullVariant);
      return nonNullVariant?.type || null;
    }
  }

  return null;
}

export { intrinsicNameToGoType };

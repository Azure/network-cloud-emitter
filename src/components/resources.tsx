import { For } from "@alloy-js/core";
import * as go from "@alloy-js/go";
import { getUnionAsEnum } from "@azure-tools/typespec-azure-core";
import { resolveArmResources, ResolvedResource } from "@azure-tools/typespec-azure-resource-manager";
import { ignoreDiagnostics, Model, Type, Union } from "@typespec/compiler";
import { unsafe_Realm } from "@typespec/compiler/experimental";
import { HttpOperation } from "@typespec/http";
import { useResourceManager as useResourceManagerContext } from "../context/resource-manager.js";
import { useTsp } from "../context/tsp-context.js";
import { Resource } from "./resource.jsx";

export interface ResourcesProps {}

export function Resources(props: ResourcesProps) {
  const tspContext = useTsp();
  const { program } = tspContext;
  const provider = resolveArmResources(program);

  for (const armResource of provider.resources?.values() ?? []) {
    extractTypesForResource(armResource);
  }

  return (
    <>
      <go.SourceDirectory path={"api/test"}>
        <For each={provider.resources ?? []}>{(armResource) => <Resource name={armResource.type.name} />}</For>
        <Resource name="common" />
      </go.SourceDirectory>
    </>
  );
}

function extractTypesForResource(resourceModel: ResolvedResource) {
  const models = new Set<Model>();
  const unions = new Set<Union>();

  for (const resource of resourceModel.operations ?? []) {
    if (resource.operations.lifecycle.read) {
      resource.operations.lifecycle.read.map((op) => handleOperationTypes(op.httpOperation, models, unions));
    }

    if (resource.operations.lifecycle.createOrUpdate) {
      resource.operations.lifecycle.createOrUpdate.map((op) => handleOperationTypes(op.httpOperation, models, unions));
    }

    if (resource.operations.lifecycle.delete) {
      resource.operations.lifecycle.delete.map((op) => handleOperationTypes(op.httpOperation, models, unions));
    }

    if (resource.operations.lifecycle.update) {
      resource.operations.lifecycle.update.map((op) => handleOperationTypes(op.httpOperation, models, unions));
    }

    if (resource.operations.actions) {
      resource.operations.actions.map((action) => handleOperationTypes(action.httpOperation, models, unions));
    }

    if (resource.operations.lists) {
      resource.operations.lists.map((list) => handleOperationTypes(list.httpOperation, models, unions));
    }
  }

  const resourceManager = useResourceManagerContext().resourceManager;
  for (const model of models) {
    resourceManager.addType(resourceModel.type.name, model);
  }
  for (const union of unions) {
    resourceManager.addType(resourceModel.type.name, union);
  }
}

function handleOperationTypes(operation: HttpOperation, models: Set<Model>, unions: Set<Union>) {
  if (operation.parameters.body?.bodyKind === "single") {
    handleTypes(operation.parameters.body.type, models, unions);
  }
  for (const responses of Object.values(operation.responses)) {
    for (const response of responses.responses) {
      if (response.body?.bodyKind === "single") {
        handleTypes(response.body.type, models, unions);
      }
    }
  }
}

function handleTypes(type: Type, models: Set<Model>, unions: Set<Union>) {
  const tspContext = useTsp();
  const { $ } = tspContext;

  if (unsafe_Realm.realmForType.has(type)) {
    return;
  }

  if ($.model.is(type)) {
    type = $.model.getEffectiveModel(type);
    if (models.has(type)) {
      return;
    } else if ($.array.is(type) || $.record.is(type)) {
      handleTypes(type.indexer.value, models, unions);
    } else if (type.namespace?.name !== "NetworkCloud") {
      return;
    } else {
      models.add(type);
      for (const prop of type.properties.values()) {
        handleTypes(prop.type, models, unions);
      }
      if (type.baseModel) {
        handleTypes(type.baseModel, models, unions);
      }
      const additionalProperties = $.model.getAdditionalPropertiesRecord(type);
      if (additionalProperties) {
        handleTypes(additionalProperties, models, unions);
      }
    }
  } else if ($.union.is(type)) {
    if (unions.has(type)) {
      return;
    } else if (type.namespace?.name !== "NetworkCloud") {
      return;
    } else if ($.union.isValidEnum(type) || ignoreDiagnostics(getUnionAsEnum(type))) {
      unions.add(type);
    } else {
      type.variants.forEach((variant) => {
        handleTypes(variant.type, models, unions);
      });
    }
  }
}

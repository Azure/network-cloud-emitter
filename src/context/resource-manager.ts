import { ComponentContext, createNamedContext, refkey, Refkey, shallowReactive, useContext } from "@alloy-js/core";
import { Type } from "@typespec/compiler";

export class ResourceManager {
  knownTypes: WeakMap<Type, string> = new WeakMap();
  refkeys: WeakMap<Type, Refkey> = new WeakMap();
  resources: Map<string, Set<Type>> = shallowReactive(new Map());

  constructor() {
    this.resources.set("common", shallowReactive(new Set()));
  }

  addType(resourceName: string, type: Type) {
    if (this.knownTypes.has(type)) {
      const existingResource = this.knownTypes.get(type);
      // same resource or common, skip
      if (existingResource === resourceName || existingResource === "common") {
        return;
      }
      // already exists in other resource, move to common
      this.knownTypes.set(type, "common");
      this.resources.get(existingResource!)!.delete(type);
      this.resources.get("common")!.add(type);
      return;
    } else {
      this.knownTypes.set(type, resourceName);
      const key = refkey();
      this.refkeys.set(type, key);
      if (!this.resources.has(resourceName)) {
        this.resources.set(resourceName, shallowReactive(new Set()));
      }
      this.resources.get(resourceName)!.add(type);
    }
  }

  getRefkey(type: Type): Refkey {
    return this.refkeys.get(type)!;
  }
}

export type ResourceManagerContext = {
  resourceManager: ResourceManager;
};

export const ResourceManagerContext: ComponentContext<{ resourceManager: ResourceManager }> = createNamedContext<{
  resourceManager: ResourceManager;
}>("ResourceManagerContext");

export function useResourceManager() {
  const context = useContext(ResourceManagerContext)!;

  if (!context) {
    throw new Error(
      "ResourceManagerContext is not set. Make sure the component is wrapped in ResourceManagerContext.Provider or the emitter framework Output component.",
    );
  }

  return context as ResourceManagerContext;
}

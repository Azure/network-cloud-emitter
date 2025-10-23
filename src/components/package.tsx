import { Children, Scope, useScope } from "@alloy-js/core";
import * as go from "@alloy-js/go";

export interface PackageProps {
  name: string;
  children?: Children;
}

export function Package(props: PackageProps) {
  const currentScope = useScope();

  // Check if we need to create a module scope first
  const needsModule =
    !currentScope || (!(currentScope instanceof go.GoPackageScope) && !(currentScope instanceof go.GoModuleScope));

  if (needsModule) {
    // Create a module scope first, then a package scope within it
    const moduleScope = new go.GoModuleScope("main");
    const packageSymbol = new go.PackageSymbol(props.name);
    const packageScope = new go.GoPackageScope(packageSymbol, moduleScope);

    return (
      <Scope value={moduleScope}>
        <Scope value={packageScope}>{props.children}</Scope>
      </Scope>
    );
  } else {
    // Create package scope within existing compatible scope
    const packageSymbol = new go.PackageSymbol(props.name);
    const packageScope = go.createGoPackageScope(packageSymbol);

    return <Scope value={packageScope}>{props.children}</Scope>;
  }
}

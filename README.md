# Network Cloud Emitter

A proof-of-concept (POC) project demonstrating how to use the [Alloy framework](https://alloy-framework.github.io/alloy/) to generate Go code for Azure Network Cloud service from TypeSpec definitions.

## Purpose

This project showcases the power of the Alloy framework for code generation by:

- **TypeSpec to Go Code Generation**: Converting Network Cloud service's TypeSpec definitions into Go model structs and constants.
- **Automated Code Generation Pipeline & Framework Integration**: Demonstrating a complete workflow from TypeSpec specifications to generated Go code while showing how to integrate Alloy with TypeSpec compiler and leverage React-like components for structured code generation.

## Quick Start

### Prerequisites

- Node.js 18+
- PNPM 10.13+
- TypeSpec Compiler installed globally:

```bash
npm install -g @typespec/compiler
```

### Installation

```bash
# Clone the repository
git clone <repository-url>
cd network-cloud-emitter

# Install dependencies
pnpm install

# Build the emitter
pnpm run build
```

### Generate Code

```bash
# Generate Go code from TypeSpec definitions
pnpm run regenerate
```

This will:

1. Clean any existing generated code
2. Build the emitter
3. Run TypeSpec compilation with the custom emitter
4. Generate Go code in `test/generated/` directory

### Key Components

- **`src/emitter.tsx`**: The main emitter that orchestrates the code generation process using React-like JSX syntax
- **`src/components/`**: Reusable components for generating different types of Go code structures
- **`src/context/`**: Alloy contexts for sharing data between components during generation
- **`test/spec/`**: TypeSpec definitions for Azure Network Cloud service resources
- **`test/generated/`**: Output directory for generated Go code

## Debugging

When working with this emitter, you have several debugging options available to help troubleshoot code generation issues:

### 1. JavaScript Debugger Mode with Conditional Breakpoints

Use VS Code's JavaScript Debug Terminal with conditional breakpoints for precise debugging:

1. **Open JavaScript Debug Terminal** - In VS Code:
   - Press `Ctrl+Shift+P` (Windows) or `Cmd+Shift+P` (Mac)
   - Type "Debug: JavaScript Debug Terminal" and select it
   - This opens a terminal with Node.js debugging capabilities enabled

2. **Set conditional breakpoints** - In VS Code, click in the gutter next to line numbers and add conditions:
   - Break when processing specific TypeSpec types: `model.name === "AgentPool"`
   - Break when specific properties exist: `symbol.properties?.length > 5`

3. **Debug the generation process** - In the JavaScript Debug Terminal:
   ```bash
   # Run the generation process - debugger will automatically attach
   pnpm run regenerate
   ```

### 2. Console.log Debugging

Add strategic console.log statements throughout your emitter code:

```tsx
// In your emitter components
function ResourceComponent({ model }: { model: Model }) {
  console.log("Processing model:", model.name);
  console.log(
    "Model properties:",
    model.properties?.map((p) => p.name),
  );

  // Log the generated content before rendering
  const generatedCode = `type ${model.name} struct { /* ... */ }`;
  console.log("Generated Go code:", generatedCode);

  return <GoStruct name={model.name}>{/* component content */}</GoStruct>;
}
```

### 3. Alloy Framework Tracing with ALLOY_TRACE

The Alloy framework provides built-in tracing capabilities using the `ALLOY_TRACE` environment variable:

**Basic Usage:**

```bash
# Trace all scope operations
$env:ALLOY_TRACE = "scope"
pnpm run regenerate

# Trace symbol creation and updates
$env:ALLOY_TRACE = "symbol.create,symbol.update"
pnpm run regenerate

# Trace scope updates and symbol creation
$env:ALLOY_TRACE = "scope.update,symbol.create"
pnpm run regenerate
```

**Available trace phases:**

- `scope.create` - Trace when new scopes are created
- `scope.update` - Trace scope updates and modifications
- `scope.copySymbols` - Trace symbol copying between scopes
- `scope.moveSymbols` - Trace symbol movement operations
- `symbol.create` - Trace symbol creation
- `symbol.update` - Trace symbol updates
- `symbol.bind` - Trace symbol binding operations
- `render.component` - Trace component rendering
- `render.effect` - Trace render effects

The trace output will show detailed information about Alloy's internal operations, including:

- Scope creation and hierarchy
- Symbol binding and resolution
- Component rendering phases
- Code generation pipeline steps

**Note:** On Unix-based systems (macOS/Linux), use `export` instead of `$env:`:

```bash
export ALLOY_TRACE="scope.update,symbol.create"
pnpm run regenerate
```

## Contributing

This is a POC project demonstrating Alloy framework capabilities. Feel free to explore, modify, and extend the code generation logic to suit your needs.

## License

MIT

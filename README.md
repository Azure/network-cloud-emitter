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

## Contributing

This is a POC project demonstrating Alloy framework capabilities. Feel free to explore, modify, and extend the code generation logic to suit your needs.

## License

MIT

#!/usr/bin/env node

import { execSync } from "child_process";
import fs from "fs";
import path from "path";
import { fileURLToPath } from "url";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const projectRoot = path.resolve(__dirname, "..");
console.log("Project root:", projectRoot);

console.log("Cleaning generated directory...");
const generatedDir = path.join(projectRoot, "test", "generated");
if (fs.existsSync(generatedDir)) {
  fs.rmSync(generatedDir, { recursive: true, force: true });
}

const command = `tsp compile --emit=${projectRoot} --option="network-cloud-emitter.emitter-output-dir=${path.join(projectRoot, "test", "generated")}" "${path.join(projectRoot, "test", "spec", "NetworkCloud.Management", "main.tsp")}"`;
console.log("Regenerating:", command);
execSync(command, { stdio: "inherit", cwd: projectRoot });

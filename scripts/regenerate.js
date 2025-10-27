#!/usr/bin/env node

import { execSync } from "child_process";
import path from "path";
import { fileURLToPath } from "url";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const projectRoot = path.resolve(__dirname, "..");
const command = `tsp compile --emit=${projectRoot} --option="network-cloud-emitter.emitter-output-dir=${path.join(projectRoot, "test", "generated")}" "${path.join(projectRoot, "test", "spec", "NetworkCloud.Management", "main.tsp")}"`;

console.log("Running:", command);
console.log("Project root:", projectRoot);
execSync(command, { stdio: "inherit", cwd: projectRoot });

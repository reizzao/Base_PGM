import { RunTester } from "../../global/helpers.js";
import { createTypes } from "./index.js";
import { seedTypesGolang } from "./seed.types.golang.js";

RunTester(createTypes(seedTypesGolang));

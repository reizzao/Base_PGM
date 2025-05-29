import { RunTester } from "../../global/helpers.ts";
import { typesFactory } from "./index.ts";
import { seedTypesGolang } from "./seed.types.golang.ts";

RunTester(typesFactory.create(seedTypesGolang));

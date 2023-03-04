export * from './modules.service';
import { ModulesService } from './modules.service';
export * from './terraform.service';
import { TerraformService } from './terraform.service';
export const APIS = [ModulesService, TerraformService];

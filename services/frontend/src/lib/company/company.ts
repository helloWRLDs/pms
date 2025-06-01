import { Project } from "../project/project";
import { ListItems, Pagination } from "../utils/list";

export type Company = {
  id: string;
  name: string;
  codename: string;
  people_count: number;
  created_at: {
    seconds: number;
  };
  updated_at: {
    seconds: number;
  };
  projects?: ListItems<Project>;
  bin?: string;
  address?: string;
  description?: string;
};

export type CompanyCreation = Pick<Company, "name" | "codename" | "bin" | "address" | "description">;

export type CompanyFilters = Pagination & {
  user_id?: string;
  company_id?: string;
  company_codename?: string;
  company_name?: string;
};

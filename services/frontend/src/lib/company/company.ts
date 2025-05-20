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
};

export type CompanyCreation = Pick<Company, "name" | "codename">;

export type CompanyFilters = Pagination & {
  user_id?: string;
  company_id?: string;
  company_codename?: string;
  company_name?: string;
};

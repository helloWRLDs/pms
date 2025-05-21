import { Pagination } from "../utils/list";

export type Project = {
  id: string;
  title: string;
  code_name: string;
  description: string;
  status: string;
  company_id: string;
  created_at: {
    seconds: number;
  };
  updated_at: {
    seconds: number;
  };
  total_tasks: number;
  done_tasks: number;
};

export type ProjectCreation = Pick<
  Project,
  "title" | "description" | "company_id" | "code_name"
>;

export type ProjectFilters = Pagination & {
  company_id?: string;
  title?: string;
  status?: string;
};

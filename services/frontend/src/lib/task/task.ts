import { Pagination } from "../utils/list";

export type Task = {
  id: string;
  title: string;
  body: string;
  status: string;
  sprint_id?: string;
  project_id?: string;
  priority: number;
  code: string;
  created_at: {
    seconds: number;
  };
  updated_at: {
    seconds: number;
  };
  assignee_id?: string;
  due_date: {
    seconds: number;
  };
};

export type TaskCreation = Omit<
  Task,
  "id" | "created_at" | "updated_at" | "code"
>;

export type TaskFilter = Pagination & {
  sprint_id?: string;
  sprint_name?: string;
  status?: string;
  assignee_id?: string;
  priority?: number;
  title?: string;
  project_id?: string;
  project_name?: string;
  code?: string;
};

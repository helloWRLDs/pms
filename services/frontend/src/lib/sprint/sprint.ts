import { Task } from "../task/task";
import { Pagination } from "../utils/list";

export type Sprint = {
  id: string;
  title: string;
  description: string;
  start_date: {
    seconds: number;
  };
  end_date: {
    seconds: number;
  };
  project_id: string;
  created_at: {
    seconds: number;
  };
  updated_at: {
    seconds: number;
  };
  tasks: Array<Task>;
};

export type SprintCreation = Omit<Sprint, "id" | "created_at" | "updated_at">;

export type SprintFilter = Pagination & {
  project_id?: string;
  title?: string;
};

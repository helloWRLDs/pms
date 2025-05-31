import { DeepPartial } from "../utils/generics";
import { Pagination } from "../utils/list";
import { TaskType } from "./tasktype";

export type Task = {
  id: string;
  title: string;
  body: string;
  status: string;
  task_type?: TaskType;
  type?: TaskType;
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

export type TaskOptional = DeepPartial<Task>;

export type TaskFilter = Pagination & {
  sprint_id?: string;
  sprint_name?: string;
  status?: string;
  type?: TaskType;
  assignee_id?: string;
  priority?: number;
  title?: string;
  project_id?: string;
  project_name?: string;
  code?: string;
};

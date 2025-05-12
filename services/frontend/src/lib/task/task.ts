import { TaskStatus } from "./status";

export type Task = {
  id: string;
  title: string;
  body: string;
  status: string;
  sprint_id?: string;
  project_id?: string;
  priority: number;
  created_at: {
    seconds: number;
  };
  updated_at: {
    seconds: number;
  };
  assignee_id: string;
  due_date: {
    seconds: number;
  };
};

export type TaskCreation = Omit<Task, "id" | "created_at" | "updated_at">;

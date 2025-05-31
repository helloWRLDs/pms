export type TaskStats = {
  total_tasks: number;
  total_points: number;
  in_progress_tasks: number;
  to_do_tasks: number;
  done_tasks: number;
};

export type UserTaskStats = {
  user_id: string;
  first_name: string;
  last_name: string;
  stats: Record<string, TaskStats>;
};

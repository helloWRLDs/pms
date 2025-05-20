import { User } from "../user/user";
import { Pagination } from "../utils/list";

export type TaskComment = {
  id: string;
  user: User;
  task_id: string;
  body: string;
  created_at: {
    seconds: number;
  };
};

export type TaskCommentCreation = Omit<
  TaskComment,
  "id" | "created_at" | "user"
> & {
  user_id: string;
};

export type TaskCommentFilter = Pagination &
  Pick<TaskComment, "task_id"> & { user_id: string };

import { API } from "./api";
import { Task, TaskCreation, TaskFilter } from "../lib/task/task";
import { buildQuery, ListItems } from "../lib/utils/list";
import {
  TaskComment,
  TaskCommentCreation,
  TaskCommentFilter,
} from "../lib/task/comment";

class TaskAPI extends API {
  async unassign(taskID: string, userID: string): Promise<void> {
    try {
      await this.req.delete(
        `${this.baseURL}/tasks/${taskID}/assignments/${userID}`
      );
    } catch (err) {
      console.error("Failed to unassign task:", err);
      throw err;
    }
  }

  async assign(taskID: string, userID: string): Promise<void> {
    try {
      await this.req.post(
        `${this.baseURL}/tasks/${taskID}/assignments/${userID}`
      );
    } catch (err) {
      console.error("Failed to assign task:", err);
      throw err;
    }
  }

  async list(filter: TaskFilter): Promise<ListItems<Task>> {
    try {
      const res = await this.req.get(
        buildQuery(`${this.baseURL}/tasks`, filter)
      );
      return res.data;
    } catch (err) {
      console.error("Failed to list tasks:", err);
      throw err;
    }
  }

  async create(task: TaskCreation): Promise<void> {
    try {
      await this.req.post(`${this.baseURL}/tasks`, task);
    } catch (err) {
      console.error("Failed to create task:", err);
      throw err;
    }
  }

  async get(id: string): Promise<Task> {
    try {
      const res = await this.req.get(`${this.baseURL}/tasks/${id}`);
      return res.data;
    } catch (err) {
      console.error("Failed to get task:", err);
      throw err;
    }
  }

  async update(id: string, task: Task): Promise<void> {
    try {
      await this.req.put(`${this.baseURL}/tasks/${id}`, task);
    } catch (err) {
      console.error("Failed to update task:", err);
      throw err;
    }
  }

  async delete(id: string): Promise<void> {
    try {
      await this.req.delete(`${this.baseURL}/tasks/${id}`);
    } catch (err) {
      console.error("Failed to delete task:", err);
      throw err;
    }
  }

  async createComment(comment: TaskCommentCreation): Promise<TaskComment> {
    try {
      const res = await this.req.post(
        `${this.baseURL}/tasks/${comment.task_id}/comments`,
        comment
      );
      return res.data;
    } catch (err) {
      console.error("Failed to create comment:", err);
      throw err;
    }
  }

  async listComments(
    taskID: string,
    filter: TaskCommentFilter
  ): Promise<ListItems<TaskComment>> {
    try {
      const res = await this.req.get(
        buildQuery(`${this.baseURL}/tasks/${taskID}/comments`, filter)
      );
      return res.data;
    } catch (err) {
      console.error("Failed to list comments:", err);
      throw err;
    }
  }
}

export const taskAPI = new TaskAPI();

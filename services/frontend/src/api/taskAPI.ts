import { API } from "./api";
import { useCompanyStore } from "../store/selectedCompanyStore";
import { Task, TaskCreation, TaskFilter } from "../lib/task/task";
import { buildQuery, ListItems } from "../lib/utils/list";
import { useProjectStore } from "../store/selectedProjectStore";
import {
  TaskComment,
  TaskCommentCreation,
  TaskCommentFilter,
} from "../lib/task/comment";

class TaskAPI extends API {
  private projectID: string;

  setupProjectID() {
    const selectedCompany = useProjectStore.getState().project;
    this.projectID = selectedCompany?.id ?? "";
  }

  constructor() {
    super();
    this.projectID = useProjectStore.getState().project?.id ?? "";
  }

  async unassign(taskID: string, userID: string): Promise<void> {
    this.setupProjectID();
    try {
      await this.req.delete(
        `${this.baseURL}/projects/${this.projectID}/tasks/${taskID}/assignment/${userID}`
      );
      return;
    } catch (e) {
      console.error(e);
      throw e;
    }
  }

  async assign(taskID: string, userID: string): Promise<void> {
    this.setupProjectID();
    console.log(`assign called: task=${taskID} user=${userID}`);
    try {
      await this.req.post(
        `${this.baseURL}/projects/${this.projectID}/tasks/${taskID}/assignment/${userID}`
      );
      return;
    } catch (err) {
      console.log(err);
      throw err;
    }
  }

  async list(filter: TaskFilter): Promise<ListItems<Task>> {
    this.setupProjectID();
    try {
      const res = await this.req.get(
        buildQuery(`${this.baseURL}/projects/${this.projectID}/tasks`, filter)
      );
      return res.data;
    } catch (err) {
      console.log(err);
      throw err;
    }
  }

  async create(task: TaskCreation) {
    this.setupProjectID();
    try {
      const res = await this.req.post(
        `${this.baseURL}/projects/${this.projectID}/tasks`,
        task
      );
      console.log(res.data);
    } catch (err) {
      console.log(err);
    }
  }

  async get(id: string): Promise<Task> {
    this.setupProjectID();
    try {
      const res = await this.req.get(
        `${this.baseURL}/projects/${this.projectID}/tasks/${id}`
      );
      return res.data;
    } catch (err) {
      console.log(err);
      throw err;
    }
  }

  async update(id: string, task: Task): Promise<void> {
    this.setupProjectID();
    try {
      await this.req.put(
        `${this.baseURL}/projects/${this.projectID}/tasks/${id}`,
        task
      );
    } catch (err) {
      console.log(err);
    }
  }

  async delete(id: string): Promise<void> {
    this.setupProjectID();
    try {
      await this.req.delete(
        `${this.baseURL}/projects/${this.projectID}/tasks/${id}`
      );
    } catch (err) {
      console.log(err);
    }
  }

  async createComment(comment: TaskCommentCreation): Promise<TaskComment> {
    this.setupProjectID();
    try {
      const res = await this.req.post(
        `${this.baseURL}/projects/${this.projectID}/tasks/${comment.task_id}/comments`,
        comment
      );
      return res.data;
    } catch (err) {
      console.log(err);
    }
    return {} as TaskComment;
  }

  async listComments(
    taskID: string,
    filter: TaskCommentFilter
  ): Promise<ListItems<TaskComment>> {
    this.setupProjectID();
    try {
      const res = await this.req.get(
        buildQuery(
          `${this.baseURL}/projects/${this.projectID}/tasks/${taskID}/comments`,
          filter
        )
      );
      return res.data;
    } catch (err) {
      console.log(err);
    }
    return {} as ListItems<TaskComment>;
  }
}

export const taskAPI = new TaskAPI();

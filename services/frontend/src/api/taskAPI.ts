import { API } from "./api";
import { useCompanyStore } from "../store/selectedCompanyStore";
import { Task, TaskCreation } from "../lib/task/task";
import { ListItems } from "../lib/utils/list";

class TaskAPI extends API {
  private projectID: string;
  constructor() {
    super();
    const selectedCompany = useCompanyStore.getState().selectedCompany;
    if (selectedCompany) {
      this.projectID = selectedCompany.id;
    } else {
      this.projectID = "";
    }
  }

  async listTasks(): Promise<ListItems<Task>> {
    try {
      const res = await this.req.get(
        `${this.baseURL}/projects/${this.projectID}/tasks`
      );
      return res.data;
    } catch (err) {
      console.log(err);
    }
    return {} as ListItems<Task>;
  }

  async createTask(task: TaskCreation) {
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

  async getTask(id: string) {
    try {
      const res = await this.req.get(
        `${this.baseURL}/projects/${this.projectID}/tasks/${id}`
      );
      return res.data;
    } catch (err) {
      console.log(err);
    }
  }
}

export const taskAPI = new TaskAPI();

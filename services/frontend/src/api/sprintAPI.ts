import { Sprint, SprintCreation, SprintFilter } from "../lib/sprint/sprint";
import { ListItems } from "../lib/utils/list";
import { API } from "./api";

class SprintAPI extends API {
  async get(id: string): Promise<Sprint> {
    try {
      const res = await this.req.get(`${this.baseURL}/sprints/${id}`);
      return res.data;
    } catch (err) {
      console.error("Failed to get sprint:", err);
      throw err;
    }
  }

  async list(filter: SprintFilter): Promise<ListItems<Sprint>> {
    try {
      const res = await this.req.get(`${this.baseURL}/sprints`, {
        params: filter,
      });
      console.log("Sprint API - Response:", res.data);
      return res.data;
    } catch (err) {
      console.error("Failed to list sprints:", err);
      throw err;
    }
  }

  async create(creation: SprintCreation): Promise<Sprint> {
    try {
      const res = await this.req.post(`${this.baseURL}/sprints`, creation);
      return res.data;
    } catch (err) {
      console.error("Failed to create sprint:", err);
      throw err;
    }
  }

  async update(id: string, sprint: Sprint): Promise<Sprint> {
    try {
      const res = await this.req.put(`${this.baseURL}/sprints/${id}`, sprint);
      return res.data;
    } catch (err) {
      console.error("Failed to update sprint:", err);
      throw err;
    }
  }
}

const sprintAPI = new SprintAPI();

export default sprintAPI;

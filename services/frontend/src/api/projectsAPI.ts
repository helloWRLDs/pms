import {
  Project,
  ProjectCreation,
  ProjectFilters,
} from "../lib/project/project";
import { buildQuery, ListItems } from "../lib/utils/list";
import { API } from "./api";

class ProjectAPI extends API {
  async list(filter: ProjectFilters): Promise<ListItems<Project>> {
    try {
      const res = await this.req.get(
        buildQuery(`${this.baseURL}/projects`, filter)
      );
      return res.data;
    } catch (err) {
      console.log(err);
    }
    return {} as ListItems<Project>;
  }

  async get(id: string): Promise<Project> {
    try {
      const res = await this.req.get(`${this.baseURL}/projects/${id}`);
      return res.data;
    } catch (err) {
      console.log(err);
    }
    return {} as Project;
  }

  async create(creation: ProjectCreation): Promise<void> {
    try {
      await this.req.post(`${this.baseURL}/projects`, creation);
    } catch (err) {
      console.log(err);
    }
  }
}

const projectAPI = new ProjectAPI();

export default projectAPI;

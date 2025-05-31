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
        buildQuery(`/projects`, {
          ...filter,
          page: filter.page || 1,
          per_page: filter.per_page || 10,
        })
      );
      return res.data;
    } catch (err) {
      console.error("Failed to list projects:", err);
      throw err;
    }
  }

  async get(id: string): Promise<Project> {
    try {
      const res = await this.req.get(`/projects/${id}`);
      return res.data;
    } catch (err) {
      console.error("Failed to get project:", err);
      throw err;
    }
  }

  async create(creation: ProjectCreation): Promise<void> {
    try {
      await this.req.post(`/projects`, creation);
    } catch (err) {
      console.error("Failed to create project:", err);
      throw err;
    }
  }
}

const projectAPI = new ProjectAPI();

export default projectAPI;

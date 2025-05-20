import { Sprint, SprintCreation, SprintFilter } from "../lib/sprint/sprint";
import { buildQuery, ListItems } from "../lib/utils/list";
import { useCompanyStore } from "../store/selectedCompanyStore";
import { useProjectStore } from "../store/selectedProjectStore";
import { API } from "./api";

class SprintAPI extends API {
  private projectID: string;

  setupProjectID() {
    const selectedProject = useProjectStore.getState().project;
    this.projectID = selectedProject?.id ?? "";
  }

  constructor() {
    super();
    const selectedCompany = useCompanyStore.getState().selectedCompany;
    if (selectedCompany) {
      this.projectID = selectedCompany.id;
    } else {
      this.projectID = "";
    }
  }

  async get(id: string): Promise<Sprint> {
    this.setupProjectID();
    try {
      const res = await this.req.get(
        `${this.baseURL}/projects/${this.projectID}/sprints/${id}`
      );
      return res.data;
    } catch (err) {
      console.log(err);
    }
    return {} as Sprint;
  }

  async list(filter: SprintFilter): Promise<ListItems<Sprint>> {
    this.setupProjectID();
    try {
      const res = await this.req.get(
        buildQuery(`${this.baseURL}/projects/${this.projectID}/sprints`, filter)
      );
      return res.data;
    } catch (err) {
      console.log(err);
    }
    return {} as ListItems<Sprint>;
  }

  async create(creation: SprintCreation): Promise<void> {
    this.setupProjectID();
    try {
      await this.req.post(
        `${this.baseURL}/projects/${this.projectID}/sprints`,
        creation
      );
    } catch (err) {
      console.log(err);
    }
  }
}

const sprintAPI = new SprintAPI();

export default sprintAPI;

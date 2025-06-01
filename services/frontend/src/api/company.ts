import {
  Company,
  CompanyCreation,
  CompanyFilters,
} from "../lib/company/company";
import { ListItems } from "../lib/utils/list";
import { API } from "./api";

class CompanyAPI extends API {
  async get(id: string): Promise<Company> {
    try {
      const res = await this.req.get<Company>(
        `${this.baseURL}/companies/${id}`
      );
      return res.data;
    } catch (err) {
      console.log(err);
    }
    return {} as Company;
  }

  async addParticipant(companyID: string, userID: string): Promise<void> {
    await this.req.post(
      `${this.baseURL}/companies/${companyID}/participants/${userID}`
    );
  }

  async create(creation: CompanyCreation): Promise<void> {
    try {
      await this.req.post(`${this.baseURL}/companies`, creation);
    } catch (e) {
      console.error(e);
      throw e;
    }
  }

  async list(filters: CompanyFilters): Promise<ListItems<Company>> {
    try {
      const res = await this.req.get<ListItems<Company>>(
        `${this.baseURL}/companies`,
        { params: filters }
      );
      return res.data;
    } catch (err) {
      console.log(err);
      throw err;
    }
  }
}

const companyAPI = new CompanyAPI();

export default companyAPI;

import { Company, CompanyFilters } from "../lib/company/company";
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

  async list(filters: CompanyFilters): Promise<ListItems<Company>> {
    try {
      const res = await this.req.get<ListItems<Company>>(
        `${this.baseURL}/companies`,
        { params: filters }
      );
      return res.data;
    } catch (err) {
      console.log(err);
    }
    return {} as ListItems<Company>;
  }
}

const companyAPI = new CompanyAPI();

export default companyAPI;

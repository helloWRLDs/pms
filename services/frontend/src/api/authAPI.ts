import { AuthData, Session } from "../lib/user/session";
import { User, UserCredentials, UserNew } from "../lib/user/user";
import { Company } from "../lib/company/company";
import { API } from "./api";
import { ListItems, Pagination } from "../lib/utils/list";

class AuthAPI extends API {
  async login(creds: UserCredentials) {
    console.log(this.baseURL);
    try {
      const res = await this.req.post<AuthData>(
        `${this.baseURL}/auth/login`,
        creds
      );
      console.log(res);
      return res.data;
    } catch (err) {
      console.log(err);
    }
  }

  async register(creation: UserNew) {
    try {
      const res = await this.req.post(
        `${this.baseURL}/auth/register`,
        creation
      );
      console.log(res);
    } catch (e) {
      console.log(e);
    }
  }

  async getSession(): Promise<Session> {
    try {
      const res = await this.req.get<Session>(`${this.baseURL}/session`);
      return res.data;
    } catch (err) {
      console.log(err);
    }
    return {} as Session;
  }

  async updateSession(session: Session): Promise<void> {
    try {
      await this.req.put(`${this.baseURL}/session`, session);
    } catch (err) {
      console.log(err);
    }
  }

  async getUser(id: string): Promise<User> {
    try {
      const res = await this.req.get<User>(`${this.baseURL}/users/${id}`);
      return res.data;
    } catch (err) {
      console.log(err);
    }
    return {} as User;
  }

  async getCompany(id: string): Promise<Company> {
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

  async listCompanies(filters: Pagination): Promise<ListItems<Company>> {
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

export const authAPI = new AuthAPI();

import { AuthData, Session } from "../lib/user/session";
import { User, UserCredentials, UserFilter, UserNew } from "../lib/user/user";
import { buildQuery, ListItems } from "../lib/utils/list";
import { API } from "./api";

class AuthAPI extends API {
  async initiateGoogleOAuth(): Promise<{ auth_url: string }> {
    const endpoint = `${this.baseURL}/auth/oauth2/google`;
    try {
      const res = await this.req.get<{ auth_url: string }>(endpoint);
      return res.data;
    } catch (err) {
      throw err;
    }
  }

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

  async listUsers(filter: UserFilter): Promise<ListItems<User>> {
    try {
      const res = await this.req.get(
        buildQuery(`${this.baseURL}/users`, filter)
      );
      return res.data;
    } catch (e) {
      console.error("failed fetching users: ", e);
    }
    return {} as ListItems<User>;
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

  async updateUser(userId: string, userData: Partial<User>): Promise<User> {
    const response = await this.req.put(
      `${this.baseURL}/users/${userId}`,
      userData
    );
    return response.data;
  }
}

const authAPI = new AuthAPI();

export default authAPI;

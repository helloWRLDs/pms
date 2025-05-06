import { Company } from "../lib/company/company";
import { User, UserCredentials } from "../lib/user/new";
import { AuthData, Session } from "../lib/user/session";
import { List } from "../lib/utils";
import { buildQuery, Pagination } from "../lib/utils/pagination";
import { request } from "./api";
import { APIConfig } from "./client";

const authAPI = (access_token?: string) => {
  const config: APIConfig = {
    baseURL: "http://localhost:8080/api/v1",
    headers: {
      "Content-Type": "application/json",
    },
  };

  if (access_token) {
    config.headers.Authorization = `Bearer ${access_token}`;
  }

  const req = request(config);

  const login = (creds: UserCredentials): Promise<AuthData> => {
    return req.post<AuthData>("/auth/login", creds);
  };

  const register = (newUser: UserCredentials): Promise<void> => {
    return req.post<void>("/auth/register", newUser);
  };

  const getUser = (id: string): Promise<User> => {
    return req.get<User>(`/users/${id}`);
  };

  const updateUser = (id: string, updated_user: User): Promise<void> => {
    return req.put<void>(`/users/${id}`, updated_user);
  };

  const listCompanies = (pagination: Pagination): Promise<List<Company>> => {
    return req.get<List<Company>>(buildQuery("/companies", pagination));
  };

  const getCompany = (companyID: string): Promise<Company> => {
    return req.get<Company>(`/companies/${companyID}`);
  };

  const getSession = (): Promise<Session> => {
    return req.get<Session>("/session");
  };

  const updateSession = (session: Session): Promise<void> => {
    return req.put<void>("/session", session);
  };

  return {
    getSession,
    updateSession,
    getCompany,
    listCompanies,
    login,
    register,
    getUser,
    updateUser,
  };
};

export default authAPI;

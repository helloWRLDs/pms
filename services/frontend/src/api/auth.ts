import { AuthData, User, UserCredentials } from "../lib/user";
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

  const register = (newUser: User): Promise<void> => {
    return req.post<void>("/auth/register", newUser);
  };

  return { login, register };
};

export default authAPI;

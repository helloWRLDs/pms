import { AuthData, User, UserCredentials } from "../lib/user";
import { UserCreation } from "../lib/user/new";
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

  const register = (newUser: UserCreation): Promise<void> => {
    return req.post<void>("/auth/register", newUser);
  };

  const getUser = (id: string): Promise<User> => {
    return req.get<User>(`/users/${id}`);
  };

  const updateUser = (id: string, updated_user: User): Promise<void> => {
    return req.put<void>(`/users/${id}`, updated_user);
  };

  return { login, register, getUser, updateUser };
};

export default authAPI;

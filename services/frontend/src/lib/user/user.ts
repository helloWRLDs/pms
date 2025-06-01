import { DeepPartial } from "../utils/generics";
import { Pagination } from "../utils/list";

export type User = {
  id: string;
  first_name: string;
  last_name: string;
  email: string;
  password: string;
  phone: string;
  avatar_img: string;
  avatar_url: string;
  bio: string;
  created_at: {
    seconds: number;
    nanos: number;
  };
  updated_at: {
    seconds: number;
    nanos: number;
  };
};

export type UserCredentials = Pick<User, "email" | "password">;
export type UserNew = Pick<
  User,
  "email" | "password" | "first_name" | "last_name"
>;
export type UserOptional = DeepPartial<User>;
export type UserFilter = Pagination & {
  company_id: string;
  email?: string;
};

export interface AuthData {
  session_id?: string;
  access_token?: string;
  user?: User;
  exp?: number;
  selected_company_id?: string;
}

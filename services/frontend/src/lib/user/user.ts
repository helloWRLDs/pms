import { DeepPartial } from "../utils/generics";

export type User = {
  id: string;
  name: string;
  email: string;
  password: string;
  phone: string;
  avatar_img: string;
  bio: string;
  created_at: number;
  updated_at: number;
};

export type UserCredentials = Pick<User, "email" | "password">;
export type UserNew = Pick<User, "email" | "password" | "name">;
export type UserOptional = DeepPartial<User>;

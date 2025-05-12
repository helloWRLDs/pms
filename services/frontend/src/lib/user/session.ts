import { UserOptional } from "./user";

export type AuthData = {
  session_id: string;
  access_token: string;
  exp: number;
  user: UserOptional;
};

export type Session = {
  session_id: string;
  user_id: string;
  access_token: string;
  exp: number;
  selected_company_id: string;
};

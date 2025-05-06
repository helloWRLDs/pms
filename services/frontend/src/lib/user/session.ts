import { User } from "./new";

export type AuthData = {
  session_id?: string;
  access_token?: string;
  exp?: number;
  user?: User;
};

export type Session = {
  session_id: string;
  user_id: string;
  access_token: string;
  exp: number;
  selected_company_id: string;
};

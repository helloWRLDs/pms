export interface UserCredentials {
  email: string;
  password: string;
}

export interface User {
  id: string;
  email: string;
  phone: string;
  name: string;
  bio: string;
  avatar_img: string;
  created_at: number;
  updated_at: number;
}

export interface AuthData {
  session_id?: string;
  access_token?: string;
  user?: User;
  exp?: number;
  selected_company_id?: string;
}

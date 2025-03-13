export interface UserCredentials {
  email: string;
  password: string;
}

export interface User {
  id: string;
  email: string;
  name: string;
  created_at: number;
  updated_at: number;
}

export interface AuthData {
  session_id?: string | null;
  access_token?: string | null;
  user?: User | null;
  exp?: number | null;
}

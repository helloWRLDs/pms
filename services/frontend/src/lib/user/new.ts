export type UserCredentials = {
  name?: string;
  email: string;
  password?: string;
};

export type User = UserCredentials & {
  id: string;
  phone?: string;
  avatar_img?: string;
  bio?: string;
  created_at?: number;
  updated_at?: number;
};

import { Permission } from "./permission";

export type Role = {
  name: string;
  company_id?: string;
  permissions: Permission[];
  created_at: {
    seconds: number;
    nanos: number;
  };
};

export type RoleFilter = {
  company_id?: string;
  with_default?: boolean;
  page?: number;
  per_page?: number;
};

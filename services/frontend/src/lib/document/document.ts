import { Pagination } from "../utils/list";

export type Document = {
  id: string;
  title: string;
  body: string;
  project_id: string;
  created_at: {
    seconds: number;
  };
};

export type DocumentCreation = Pick<Document, "project_id" | "title"> & {
  sprint_id?: string;
};

export type DocumentFilter = Pagination & {
  project_id?: string;
  title?: string;
  date_from?: string;
  date_to?: string;
  order_by?: string;
  order_direction?: string;
};

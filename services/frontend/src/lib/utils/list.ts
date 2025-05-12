export type List = {
  page: number;
  per_page: number;
  total_pages: number;
  total_items: number;
};

export type ListItems<T> = List & {
  items: T[];
};

export type Pagination = {
  page: number;
  per_page: number;
};

export type ListFilters = Pagination & {
  query?: string;
};

export const buildQuery = (baseURL: string, pagination?: {}): string => {
  const params = new URLSearchParams();
  if (pagination) {
    Object.entries(pagination).forEach(([k, v]) => {
      if (v !== undefined && v !== null) {
        params.append(k, `${v}`);
      }
    });
  }
  const q = params.toString();
  return q ? `${baseURL}?${q}` : baseURL;
};

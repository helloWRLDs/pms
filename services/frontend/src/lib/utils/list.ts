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

export const buildQuery = (baseURL: string, query: {}): string => {
  const params = new URLSearchParams();
  if (query) {
    Object.entries(query).forEach(([k, v]) => {
      if (v !== undefined && v !== null) {
        params.append(k, `${v}`);
      }
    });
  }
  const q = params.toString();
  return q ? `${baseURL}?${q}` : baseURL;
};

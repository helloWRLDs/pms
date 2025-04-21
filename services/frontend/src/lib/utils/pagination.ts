export interface Pagination {
  page?: number;
  per_page?: number;
}

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

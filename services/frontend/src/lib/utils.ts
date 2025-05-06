export type List<T> = {
  page: number;
  per_page: number;
  total_items: number;
  total_pages: number;
  items: Array<T>;
};

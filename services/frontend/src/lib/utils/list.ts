export interface List<T> {
  page: number;
  per_page: number;
  total_pages: number;
  total_items: number;
  items: T[];
}

export interface Company {
  id: string;
  name: string;
  codename: string;
  people_count: number;
  created_at: {
    seconds: number;
  };
  updated_at: {
    seconds: number;
  };
}

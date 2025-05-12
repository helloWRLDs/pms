export type Project = {
  id: string;
  title: string;
  description: string;
  status: string;
  company_id: string;
  created_at: {
    seconds: number;
  };
  updated_at: {
    seconds: number;
  };
  progress: number;
};

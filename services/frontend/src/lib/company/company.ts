import { Project } from "../project/project";
import { List } from "../utils";

export type Company = {
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
  projects?: List<Project>;
};

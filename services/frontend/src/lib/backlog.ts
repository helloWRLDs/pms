export default class Backlog {
  id: string;
  name: string;
  description: string;
  created_at: number;
  updated_at: number;

  constructor(
    id: string,
    name: string,
    description: string,
    created_at: number,
    updated_at: number
  ) {
    this.id = id;
    this.name = name;
    this.description = description;
    this.created_at = created_at;
    this.updated_at = updated_at;
  }
}

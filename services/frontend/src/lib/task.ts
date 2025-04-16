export class Task {
  id: string;
  title: string;
  body: string;
  status: string;
  executor: string;
  backlog_id: string;
  created_at: string;
  updated_at: string;

  constructor(data: any) {
    if (typeof data == "string") {
      data = JSON.parse(data);
    }
    this.id = data.id;
    this.title = data.title;
    this.body = data.body;
    this.status = data.status;
    this.executor = data.executor;
    this.backlog_id = data.backlog_id;
    this.created_at = data.created_at;
    this.updated_at = data.updated_at;
  }
}

export namespace Task {
  export enum TaskStatus {
    CREATED,
    IN_PROGRESS,
    PENDING,
    DONE,
  }

  export const STATUSES: TaskStatus[] = [
    TaskStatus.CREATED,
    TaskStatus.IN_PROGRESS,
    TaskStatus.PENDING,
    TaskStatus.DONE,
  ];

  export const toString = (task: TaskStatus): string => {
    switch (task) {
      case TaskStatus.CREATED:
        return "created";
      case TaskStatus.IN_PROGRESS:
        return "in progress";
      case TaskStatus.DONE:
        return "done";
      case TaskStatus.PENDING:
        return "pending";
      default:
        return "unknown";
    }
  };

  export const fromString = (task: string): TaskStatus => {
    switch (task) {
      case "created":
        return TaskStatus.CREATED;
      case "in progress":
        return TaskStatus.IN_PROGRESS;
      case "done":
        return TaskStatus.DONE;
      case "pending":
        return TaskStatus.PENDING;
      default:
        throw new Error("Unknown task status");
    }
  };
}

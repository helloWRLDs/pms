const Enum = {
  CREATED: "CREATED",
  IN_PROGRESS: "IN_PROGRESS",
  PENDING: "PENDING",
  DONE: "DONE",
  ARCHIVED: "ARCHIVED",
} as const;

export type TaskStatus = (typeof Enum)[keyof typeof Enum];
export const getTaskStatuses: TaskStatus[] = Object.values(Enum);

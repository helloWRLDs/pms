import { capitalize } from "../utils/string";
import { ObjectWrapper } from "../utils/wrapper";

export const Enum = {
  CREATED: "CREATED",
  IN_PROGRESS: "IN_PROGRESS",
  PENDING: "PENDING",
  DONE: "DONE",
  ARCHIVED: "ARCHIVED",
} as const;

export type TaskStatus = (typeof Enum)[keyof typeof Enum];
export const getTaskStatuses: TaskStatus[] = Object.values(Enum);

export const StatusFilterValues: ObjectWrapper[] = [
  "All",
  ...getTaskStatuses,
].map((status) => {
  if (status === "All") {
    return { label: "All", value: "" };
  }
  return {
    label: capitalize(status.toString().replace("_", "")),
    value: status,
  };
});

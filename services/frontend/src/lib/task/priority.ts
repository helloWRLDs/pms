import { ObjectWrapper } from "../utils/wrapper";

export const AVAILABLE_PRIORITIES = [1, 2, 3, 4, 5];

export class Priority {
  priority: number;
  static availablePriorities = AVAILABLE_PRIORITIES;

  constructor(priority: number) {
    this.priority = priority;
  }

  toString(): string {
    switch (this.priority) {
      case 1:
        return "Lowest";
      case 2:
        return "Low";
      case 3:
        return "Medium";
      case 4:
        return "High";
      case 5:
        return "Critical";
      default:
        return "All";
    }
  }

  getColor(): string {
    switch (this.priority) {
      case 1:
        return "red";
      case 2:
        return "orange";
      case 3:
        return "yellow";
      case 4:
        return "green";
      case 5:
        return "blue";
      default:
        return "gray";
    }
  }
}

export const PriorityFilterValues: ObjectWrapper[] = [
  0,
  ...AVAILABLE_PRIORITIES,
].map((priority) => {
  if (priority === 0) {
    return { label: "All", value: priority };
  }
  return {
    label: new Priority(priority).toString(),
    value: priority,
  };
});

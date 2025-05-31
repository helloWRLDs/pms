export const TaskTypes = {
  BUG: "bug",
  STORY: "story",
  SUB_TASK: "sub_task",
  FEATURE: "feature",
  CHORE: "chore",
  REFACTOR: "refactor",
  TEST: "test",
  DOCUMENTATION: "documentation",
} as const;

export type TaskType = (typeof TaskTypes)[keyof typeof TaskTypes];

export const getTaskTypes = Object.values(TaskTypes);

export interface TaskTypeConfig {
  label: string;
  color: string;
  icon: string;
  description: string;
}

export const taskTypeConfigs: Record<TaskType, TaskTypeConfig> = {
  [TaskTypes.BUG]: {
    label: "Bug",
    color: "#ef4444", // red
    icon: "🐛",
    description: "Something isn't working correctly",
  },
  [TaskTypes.STORY]: {
    label: "Story",
    color: "#22c55e", // green
    icon: "📖",
    description: "A user story or requirement",
  },
  [TaskTypes.SUB_TASK]: {
    label: "Sub Task",
    color: "#8b5cf6", // purple
    icon: "📋",
    description: "A smaller task that's part of a larger one",
  },
  [TaskTypes.FEATURE]: {
    label: "Feature",
    color: "#3b82f6", // blue
    icon: "✨",
    description: "A new feature or enhancement",
  },
  [TaskTypes.CHORE]: {
    label: "Chore",
    color: "#6b7280", // gray
    icon: "🔧",
    description: "Maintenance or housekeeping task",
  },
  [TaskTypes.REFACTOR]: {
    label: "Refactor",
    color: "#f59e0b", // amber
    icon: "♻️",
    description: "Code refactoring or restructuring",
  },
  [TaskTypes.TEST]: {
    label: "Test",
    color: "#10b981", // emerald
    icon: "🧪",
    description: "Testing related task",
  },
  [TaskTypes.DOCUMENTATION]: {
    label: "Documentation",
    color: "#06b6d4", // cyan
    icon: "📚",
    description: "Documentation writing or updates",
  },
};

export const getTaskTypeConfig = (type: TaskType): TaskTypeConfig => {
  return taskTypeConfigs[type];
};

export const getTaskTypeLabel = (type: TaskType): string => {
  return taskTypeConfigs[type]?.label || type;
};

export const getTaskTypeColor = (type: TaskType): string => {
  return taskTypeConfigs[type]?.color || "#6b7280";
};

export const getTaskTypeIcon = (type: TaskType): string => {
  return taskTypeConfigs[type]?.icon || "📌";
};

export const isValidTaskType = (type: string): type is TaskType => {
  return Object.values(TaskTypes).includes(type as TaskType);
};

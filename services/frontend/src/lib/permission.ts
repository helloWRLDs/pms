const Permissions = {
  // User permissions
  USER_READ_PERMISSION: "user:read",
  USER_WRITE_PERMISSION: "user:write",
  USER_DELETE_PERMISSION: "user:delete",
  USER_INVITE_PERMISSION: "user:invite",

  // Company permissions (company and organization are the same entity)
  COMPANY_READ_PERMISSION: "company:read",
  COMPANY_WRITE_PERMISSION: "company:write",
  COMPANY_DELETE_PERMISSION: "company:delete",
  COMPANY_INVITE_PERMISSION: "company:invite",

  // Project permissions
  PROJECT_READ_PERMISSION: "project:read",
  PROJECT_WRITE_PERMISSION: "project:write",
  PROJECT_DELETE_PERMISSION: "project:delete",
  PROJECT_INVITE_PERMISSION: "project:invite",

  // Task permissions
  TASK_READ_PERMISSION: "task:read",
  TASK_WRITE_PERMISSION: "task:write",
  TASK_DELETE_PERMISSION: "task:delete",
  TASK_INVITE_PERMISSION: "task:invite",

  // Role permissions
  ROLE_READ_PERMISSION: "role:read",
  ROLE_WRITE_PERMISSION: "role:write",
  ROLE_DELETE_PERMISSION: "role:delete",
  ROLE_INVITE_PERMISSION: "role:invite",

  // Sprint permissions
  SPRINT_READ_PERMISSION: "sprint:read",
  SPRINT_WRITE_PERMISSION: "sprint:write",
  SPRINT_DELETE_PERMISSION: "sprint:delete",
  SPRINT_INVITE_PERMISSION: "sprint:invite",
} as const;

export type Permission = (typeof Permissions)[keyof typeof Permissions];

export { Permissions };

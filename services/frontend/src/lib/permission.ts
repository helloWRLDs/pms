const Permissions = {
  ORG_READ_PERMISSION: "org:read",
  ORG_WRITE_PERMISSION: "org:write",
  USER_READ_PERMISSION: "user:read",
  USER_WRITE_PERMISSION: "user:write",
  USER_DELETE_PERMISSION: "user:delete",
  USER_ADD_PERMISSION: "user:add",
  PROJECT_READ_PERMISSION: "project:read",
  PROJECT_WRITE_PERMISSION: "project:write",
  PROJECT_DELETE_PERMISSION: "project:delete",
  PROJECT_ADD_PERMISSION: "project:add",
  TASK_READ_PERMISSION: "task:read",
  TASK_WRITE_PERMISSION: "task:write",
  TASK_DELETE_PERMISSION: "task:delete",
  TASK_ADD_PERMISSION: "task:add",
  ROLE_READ_PERMISSION: "role:read",
  ROLE_WRITE_PERMISSION: "role:write",
  ROLE_DELETE_PERMISSION: "role:delete",
  ROLE_ADD_PERMISSION: "role:add",
  SPRINT_READ_PERMISSION: "sprint:read",
  SPRINT_WRITE_PERMISSION: "sprint:write",
  SPRINT_DELETE_PERMISSION: "sprint:delete",
  SPRINT_ADD_PERMISSION: "sprint:add",
} as const;

export type Permission = (typeof Permissions)[keyof typeof Permissions];

export { Permissions };

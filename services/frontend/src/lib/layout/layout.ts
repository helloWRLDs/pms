export const Layouts = {
  Projects: "Projects",
  Companies: "Companies",
  Login: "Login",
} as const;

export type Layout = (typeof Layouts)[keyof typeof Layouts];

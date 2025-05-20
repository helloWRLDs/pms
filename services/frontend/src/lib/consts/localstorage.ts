import { ValueOf } from "../utils/generics";

export const LocalStorageKeysMap = {
  AUTH: "auth",
  TEST: "test",
  PROJECT: "project",
  COMPANY: "company",
  SPRINT: "sprint",
} as const;

export type LocaStorageKey = ValueOf<typeof LocalStorageKeysMap>;

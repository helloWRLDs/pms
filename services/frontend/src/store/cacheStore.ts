import { create } from "zustand";
import { Project } from "../lib/project/project";
import { Sprint } from "../lib/sprint/sprint";
import { Company } from "../lib/company/company";
import { persist } from "zustand/middleware";
import { User } from "../lib/user/user";

interface CacheStore {
  projects: Record<string, Project> | null;
  sprints: Record<string, Sprint> | null;
  companies: Record<string, Company> | null;
  assignees: Record<string, User> | null;
  setProjects: (projects: Record<string, Project>) => void;
  getProjects: () => Record<string, Project> | null;
  getProject: (id: string) => Project | null;
  setSprints: (sprints: Record<string, Sprint>) => void;
  getSprints: () => Record<string, Sprint> | null;
  getSprint: (id: string) => Sprint | null;
  setCompanies: (companies: Record<string, Company>) => void;
  getCompanies: () => Record<string, Company> | null;
  getCompany: (id: string) => Company | null;
  setAssignees: (user: Record<string, User>) => void;
  getAssignees: () => Record<string, User> | null;
  getAssignee: (id: string) => User | null;
}

export const useCacheStore = create<CacheStore>()(
  persist(
    (set, get) => ({
      projects: null,
      sprints: null,
      companies: null,
      assignees: null,
      setProjects: (projects) => set({ ...get(), projects: projects }),
      getProjects: () => get().projects,
      getProject: (id) => get().projects?.[id] ?? null,
      setSprints: (sprints) => set({ ...get(), sprints: sprints }),
      getSprints: () => get().sprints,
      getSprint: (id) => get().sprints?.[id] ?? null,
      setCompanies: (companies) => set({ ...get(), companies: companies }),
      getCompanies: () => get().companies,
      getCompany: (id) => get().companies?.[id] ?? null,
      setAssignees: (assignees) => set({ ...get(), assignees: assignees }),
      getAssignees: () => get().assignees,
      getAssignee: (id) => get().assignees?.[id] ?? null,
    }),
    {
      name: "cache-data",
      storage: localStorage,
      // storage: createCookieStorage<CacheStore>("cache-data"),
    }
  )
);

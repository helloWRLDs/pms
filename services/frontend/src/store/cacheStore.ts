import { create } from "zustand";
import { Project } from "../lib/project/project";
import { Sprint } from "../lib/sprint/sprint";
import { Company } from "../lib/company/company";
import { persist } from "zustand/middleware";
import { User } from "../lib/user/user";

// Generic cache item wrapper
interface CacheItem<T> {
  data: T | null;
  lastUpdated: number;
  isLoading: boolean;
  error: string | null;
}

// Entity collections
interface EntityCollections {
  projects: Record<string, Project>;
  sprints: Record<string, Sprint>;
  companies: Record<string, Company>;
  assignees: Record<string, User>;
}

// Cache metadata
interface CacheMetadata {
  lastGlobalUpdate: number;
  version: string;
  selectedCompanyId: string | null;
  selectedProjectId: string | null;
}

// Cache state
interface CacheState {
  collections: {
    projects: CacheItem<Record<string, Project>>;
    sprints: CacheItem<Record<string, Sprint>>;
    companies: CacheItem<Record<string, Company>>;
    assignees: CacheItem<Record<string, User>>;
  };
  metadata: CacheMetadata;
}

// Cache actions
interface CacheActions {
  // Entity actions
  setProjects: (projects: Record<string, Project>) => void;
  setSprints: (sprints: Record<string, Sprint>) => void;
  setCompanies: (companies: Record<string, Company>) => void;
  setAssignees: (assignees: Record<string, User>) => void;

  // Loading state actions
  setLoading: (entity: keyof EntityCollections, isLoading: boolean) => void;
  setError: (entity: keyof EntityCollections, error: string | null) => void;

  // Getters
  getProjects: () => CacheItem<Record<string, Project>>;
  getSprints: () => CacheItem<Record<string, Sprint>>;
  getCompanies: () => CacheItem<Record<string, Company>>;
  getAssignees: () => CacheItem<Record<string, User>>;

  // Individual item getters
  getProject: (id: string) => Project | null;
  getSprint: (id: string) => Sprint | null;
  getCompany: (id: string) => Company | null;
  getAssignee: (id: string) => User | null;

  // Metadata actions
  setSelectedCompany: (companyId: string | null) => void;
  setSelectedProject: (projectId: string | null) => void;
}

type CacheStore = CacheState & CacheActions;

const initialState: CacheState = {
  collections: {
    projects: { data: null, lastUpdated: 0, isLoading: false, error: null },
    sprints: { data: null, lastUpdated: 0, isLoading: false, error: null },
    companies: { data: null, lastUpdated: 0, isLoading: false, error: null },
    assignees: { data: null, lastUpdated: 0, isLoading: false, error: null },
  },
  metadata: {
    lastGlobalUpdate: Date.now(),
    version: "1.0.0",
    selectedCompanyId: null,
    selectedProjectId: null,
  },
};

export const useCacheStore = create<CacheStore>()(
  persist(
    (set, get) => ({
      ...initialState,

      // Entity setters
      setProjects: (projects) =>
        set((state) => ({
          collections: {
            ...state.collections,
            projects: {
              data: projects,
              lastUpdated: Date.now(),
              isLoading: false,
              error: null,
            },
          },
        })),

      setSprints: (sprints) =>
        set((state) => ({
          collections: {
            ...state.collections,
            sprints: {
              data: sprints,
              lastUpdated: Date.now(),
              isLoading: false,
              error: null,
            },
          },
        })),

      setCompanies: (companies) =>
        set((state) => ({
          collections: {
            ...state.collections,
            companies: {
              data: companies,
              lastUpdated: Date.now(),
              isLoading: false,
              error: null,
            },
          },
        })),

      setAssignees: (assignees) =>
        set((state) => ({
          collections: {
            ...state.collections,
            assignees: {
              data: assignees,
              lastUpdated: Date.now(),
              isLoading: false,
              error: null,
            },
          },
        })),

      // Loading state actions
      setLoading: (entity, isLoading) =>
        set((state) => ({
          collections: {
            ...state.collections,
            [entity]: {
              ...state.collections[entity],
              isLoading,
            },
          },
        })),

      setError: (entity, error) =>
        set((state) => ({
          collections: {
            ...state.collections,
            [entity]: {
              ...state.collections[entity],
              error,
              isLoading: false,
            },
          },
        })),

      // Getters
      getProjects: () => get().collections.projects,
      getSprints: () => get().collections.sprints,
      getCompanies: () => get().collections.companies,
      getAssignees: () => get().collections.assignees,

      // Individual item getters
      getProject: (id) => get().collections.projects.data?.[id] ?? null,
      getSprint: (id) => get().collections.sprints.data?.[id] ?? null,
      getCompany: (id) => get().collections.companies.data?.[id] ?? null,
      getAssignee: (id) => get().collections.assignees.data?.[id] ?? null,

      // Metadata actions
      setSelectedCompany: (companyId) =>
        set((state) => ({
          metadata: {
            ...state.metadata,
            selectedCompanyId: companyId,
          },
        })),

      setSelectedProject: (projectId) =>
        set((state) => ({
          metadata: {
            ...state.metadata,
            selectedProjectId: projectId,
          },
        })),
    }),
    {
      name: "cache-data",
      storage: {
        getItem: (name) => {
          const str = localStorage.getItem(name);
          if (!str) return null;
          try {
            return JSON.parse(str);
          } catch {
            return null;
          }
        },
        setItem: (name, value) => {
          localStorage.setItem(name, JSON.stringify(value));
        },
        removeItem: (name) => {
          localStorage.removeItem(name);
        },
      },
    }
  )
);

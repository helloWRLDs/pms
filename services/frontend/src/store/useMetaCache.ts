import { create } from "zustand";
import { Company } from "../lib/company/company";
import { Project } from "../lib/project/project";
import { Sprint } from "../lib/sprint/sprint";
import { User } from "../lib/user/user";
import { persist } from "zustand/middleware";

interface CacheState {
  metadata: {
    selectedCompany: Company | null;
    selectedProject: Project | null;
  };
}

interface CacheActions {
  setSelectedCompany: (companyId: Company | null) => void;
  setSelectedProject: (projectId: Project | null) => void;
}

type CacheStore = CacheState & CacheActions;

const useMetaCache = create<CacheStore>()(
  persist(
    (set, get) => ({
      metadata: {
        selectedCompany: null,
        selectedProject: null,
      },
      setSelectedCompany: (company: Company | null) =>
        set((state: CacheState) => ({
          metadata: {
            ...state.metadata,
            selectedCompany: company,
            selectedProject: null,
          },
        })),
      setSelectedProject: (project: Project | null) =>
        set((state: CacheState) => ({
          metadata: { ...state.metadata, selectedProject: project },
        })),
    }),
    {
      name: "meta-cache",
      storage: {
        getItem: (name: string) => {
          const str = localStorage.getItem(name);
          if (!str) return null;
          try {
            return JSON.parse(str);
          } catch {
            return null;
          }
        },
        setItem: (name: string, value: any) => {
          localStorage.setItem(name, JSON.stringify(value));
        },
        removeItem: (name: string) => {
          localStorage.removeItem(name);
        },
      },
    }
  )
);

export default useMetaCache;

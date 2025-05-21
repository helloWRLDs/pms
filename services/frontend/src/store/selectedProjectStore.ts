import { create } from "zustand";
import { Project } from "../lib/project/project";
import { createJSONStorage, persist } from "zustand/middleware";
import { LocalStorageKeysMap } from "../lib/consts/localstorage";

interface ProjectStore {
  project: Project | null;
  selectProject: (project: Project) => void;
  resetProject: () => void;
  getProject: () => Project | null;
}

export const useProjectStore = create<ProjectStore>()(
  persist(
    (set, get) => ({
      project: null,
      selectProject: (project: Project) => set({ project }),
      resetProject: () => set({ project: null }),
      getProject: () => get().project,
    }),
    {
      name: LocalStorageKeysMap.PROJECT,
      storage: createJSONStorage(() => sessionStorage),
    }
  )
);

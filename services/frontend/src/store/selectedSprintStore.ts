import { createJSONStorage, persist } from "zustand/middleware";
import { Sprint } from "../lib/sprint/sprint";
import { create } from "zustand";
import { LocalStorageKeysMap } from "../lib/consts/localstorage";

interface SprintStore {
  sprint: Sprint | null;
  selectSprint: (sprint: Sprint) => void;
  resetSprint: () => void;
  getSprint: () => Sprint | null;
}

export const useSprintStore = create<SprintStore>()(
  persist(
    (set, get) => ({
      sprint: null,
      selectSprint: (sprint: Sprint) => set({ sprint: sprint }),
      resetSprint: () => set({ sprint: null }),
      getSprint: () => get().sprint,
    }),
    {
      name: LocalStorageKeysMap.SPRINT,
      storage: createJSONStorage(() => sessionStorage),
    }
  )
);

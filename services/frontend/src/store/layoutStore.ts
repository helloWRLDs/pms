import { create } from "zustand";
import { Layout } from "../lib/layout/layout";

interface LayoutStore {
  activeLayout: Layout | null;
  setLayout: (layout: Layout) => void;
}

export const useLayoutStore = create<LayoutStore>()((set, get) => ({
  activeLayout: null,
  setLayout: (layout: Layout) => set({ activeLayout: layout }),
  get: () => get().activeLayout,
}));

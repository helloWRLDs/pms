import { create } from "zustand";
import { Company } from "../lib/company/company";
import { createJSONStorage, persist } from "zustand/middleware";
import { LocalStorageKeysMap } from "../lib/consts/localstorage";

interface CompanyStore {
  selectedCompany: Company | null;
  selectCompany: (selectedCompany: Company) => void;
  resetCompany: () => void;
  getCompany: () => Company | null;
}

export const useCompanyStore = create<CompanyStore>()(
  persist(
    (set, get) => ({
      selectedCompany: null,
      selectCompany: (selectedCompany: Company) => set({ selectedCompany }),
      resetCompany: () => set({ selectedCompany: null }),
      getCompany: () => get().selectedCompany,
    }),
    {
      name: LocalStorageKeysMap.COMPANY,
      storage: createJSONStorage(() => sessionStorage),
    }
  )
);

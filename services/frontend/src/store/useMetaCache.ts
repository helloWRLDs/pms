import { create } from "zustand";
import { Company } from "../lib/company/company";
import { Project } from "../lib/project/project";

import { persist } from "zustand/middleware";

interface CacheState {
  metadata: {
    selectedCompany: Company | null;
    selectedProject: Project | null;
    currentUserId: string | null; // Track which user this cache belongs to
  };
}

interface CacheActions {
  setSelectedCompany: (companyId: Company | null) => void;
  setSelectedProject: (projectId: Project | null) => void;
  setCurrentUser: (userId: string | null) => void;
  clearCache: () => void;
  checkUserAndClearIfDifferent: (currentUserId: string | null) => void;
}

type CacheStore = CacheState & CacheActions;

const useMetaCache = create<CacheStore>()(
  persist(
    (set, get) => ({
      metadata: {
        selectedCompany: null,
        selectedProject: null,
        currentUserId: null,
      },
      setSelectedCompany: (company: Company | null) =>
        set((state: CacheState) => ({
          metadata: {
            ...state.metadata,
            selectedCompany: company,
            selectedProject: null, // Clear project when company changes
          },
        })),
      setSelectedProject: (project: Project | null) =>
        set((state: CacheState) => ({
          metadata: { ...state.metadata, selectedProject: project },
        })),
      setCurrentUser: (userId: string | null) =>
        set((state: CacheState) => ({
          metadata: { ...state.metadata, currentUserId: userId },
        })),
      clearCache: () =>
        set(() => ({
          metadata: {
            selectedCompany: null,
            selectedProject: null,
            currentUserId: null,
          },
        })),
      checkUserAndClearIfDifferent: (currentUserId: string | null) => {
        const state = get();
        const existingUserId = state.metadata.currentUserId;

        // Only update if the user actually changed
        if (existingUserId !== currentUserId) {
          if (existingUserId && existingUserId !== currentUserId) {
            // Different user detected, clear the cache
            console.log("🔄 Different user detected, clearing meta cache", {
              previousUser: existingUserId,
              newUser: currentUserId,
            });
            set(() => ({
              metadata: {
                selectedCompany: null,
                selectedProject: null,
                currentUserId: currentUserId,
              },
            }));
          } else {
            // First time setting user or going from null to a user
            console.log("👤 Setting current user:", currentUserId);
            set((state: CacheState) => ({
              metadata: { ...state.metadata, currentUserId: currentUserId },
            }));
          }
        }
        // If user hasn't changed, do nothing to prevent unnecessary updates
      },
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

import { create } from "zustand";
import { AuthData } from "../lib/user/session";
import { persist } from "zustand/middleware";
import { LocalStorageKeysMap } from "../lib/consts/localstorage";
import { jwtDecode } from "jwt-decode";

interface AuthStore {
  auth: AuthData | null;
  isAuthenticated: () => boolean;
  setAuth: (auth: AuthData) => void;
  clearAuth: () => void;
  getAuth: () => AuthData | null;
}

export const useAuthStore = create<AuthStore>()(
  persist(
    (set, get) => ({
      auth: null,
      isAuthenticated: () => {
        const auth = get().auth;
        if (!auth) {
          return false;
        }
        if (!auth.access_token) {
          set({ auth: null });
          return false;
        }
        const decoded = jwtDecode(auth.access_token);
        if (!decoded.exp || decoded.exp * 1000 < Date.now()) {
          set({ auth: null });
          return false;
        }
        return true;
      },
      setAuth: (auth: AuthData) => set({ auth }),
      clearAuth: () => set({ auth: null }),
      getAuth: () => get().auth,
    }),
    {
      name: LocalStorageKeysMap.AUTH,
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

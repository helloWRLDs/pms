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
        if (!decoded.exp || decoded.exp > Date.now()) {
          set({ auth: null });
          return false;
        }
        return true;
      },
      setAuth: (auth: AuthData) => set({ auth }),
      clearAuth: () => set({ auth: null }),
    }),
    {
      name: LocalStorageKeysMap.AUTH,
    }
  )
);

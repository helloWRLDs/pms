import Cookies from "js-cookie";
import type { PersistStorage } from "zustand/middleware";

const createCookieStorage = <T>(): PersistStorage<T> => ({
  getItem: (name) => {
    const cookie = Cookies.get(name);
    if (!cookie) return null;
    try {
      return JSON.parse(cookie);
    } catch (e) {
      console.error("Failed to parse cookie: ", e);
      return null;
    }
  },
  setItem: (name, value) => {
    Cookies.set(name, JSON.stringify(value), { expires: 7 });
  },
  removeItem: (name) => {
    Cookies.remove(name);
  },
});

export default createCookieStorage;

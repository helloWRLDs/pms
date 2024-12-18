export const useLocalStorage = <T>(key: string) => {
  const set = (value: T) => {
    try {
      window.localStorage.setItem(key, JSON.stringify(value));
    } catch (e) {
      console.log(`error setting value with key=${key}`);
    }
  };

  const get = (): T | undefined => {
    try {
      const value = window.localStorage.getItem(key);
      return value ? JSON.parse(value) : undefined;
    } catch (e) {
      console.log(`error getting value with key=${key}`);
    }
  };

  const remove = () => {
    try {
      window.localStorage.removeItem(key);
    } catch (e) {
      console.log(`error removing value with key=${key}`);
    }
  };

  return [set, get, remove] as const;
};

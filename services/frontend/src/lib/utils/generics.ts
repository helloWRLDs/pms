/**
 * For enums
 */
export type ValueOf<T> = T[keyof T];

/**
 * Make all fields optional
 */
export type DeepPartial<T> = T extends object
  ? {
      [K in keyof T]?: DeepPartial<T[K]>;
    }
  : T;

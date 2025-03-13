import { configureStore } from "@reduxjs/toolkit";
import reducer from "./authSlice";

export const authStore = configureStore({
  reducer: reducer,
});

export type RootState = ReturnType<typeof authStore.getState>;
export type AppDispath = typeof authStore.dispatch;

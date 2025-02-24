import { configureStore } from "@reduxjs/toolkit";
import reducer from "./authSlice";
// import { persist } from "zustand/middleware";
// import { create } from "zustand";
// import { jwtDecode } from "jwt-decode";

export const authStore = configureStore({
  reducer: reducer,
});

export type RootState = ReturnType<typeof authStore.getState>;
export type AppDispath = typeof authStore.dispatch;

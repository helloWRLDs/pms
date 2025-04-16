import { configureStore, createSlice, PayloadAction } from "@reduxjs/toolkit";
import { isTokenValid } from "../utils/jwt";
import { AuthData } from "../lib/user";

// Local Storage Key
const AUTH_KEY = "auth";

// Load initial state from localStorage
const loadInitialState = (): AuthData => {
  const nullState: AuthData = {};
  try {
    const data = localStorage.getItem(AUTH_KEY);
    if (!data) {
      return nullState;
    }
    const state: AuthData = JSON.parse(data);
    if (!isTokenValid(state.access_token || "")) {
      return nullState;
    }
    return state;
  } catch (e) {
    console.log(`Failed loading initial auth state: ${e}`);
    return nullState;
  }
};

const initialState = loadInitialState();

// Slice
const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    login: (state: AuthData, action: PayloadAction<AuthData>) => {
      state.access_token = action.payload.access_token;
      state.user = action.payload.user;
      localStorage.setItem(AUTH_KEY, JSON.stringify(state));
    },
    logout: () => {
      localStorage.removeItem(AUTH_KEY);
      return {};
    },
  },
});

// Actions
export const { login, logout } = authSlice.actions;

// Store
export const authStore = configureStore({
  reducer: {
    auth: authSlice.reducer,
  },
});

// Types
export type RootState = ReturnType<typeof authStore.getState>;
export type AppDispatch = typeof authStore.dispatch;

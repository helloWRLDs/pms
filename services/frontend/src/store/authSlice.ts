import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { isTokenValid } from "../utils/jwt";
import { AuthData } from "../lib/user";

const AUTH_KEY = "auth";

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
    console.log(`failed finding initial state: ${e}`);
    return nullState;
  }
};

const initialState = loadInitialState();

const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    login: (state: AuthData, action: PayloadAction<AuthData>) => {
      state.access_token = action.payload.access_token;
      state.user = action.payload.user;
      localStorage.setItem(AUTH_KEY, JSON.stringify(state));
    },
    logout: (state: AuthData) => {
      state.access_token = null;
      state.user = null;
      localStorage.removeItem(AUTH_KEY);
    },
  },
});

export const { login, logout } = authSlice.actions;
export default authSlice.reducer;

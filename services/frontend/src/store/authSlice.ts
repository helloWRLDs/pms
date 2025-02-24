import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { isTokenValid } from "../utils/jwt";

const AUTH_KEY = "auth";

interface UserData {
  id: string;
  email: string;
}

interface AuthState {
  token: string | null;
  user: UserData | null;
}

const loadInitialState = (): AuthState => {
  const nullState: AuthState = { token: null, user: null };
  try {
    const data = localStorage.getItem(AUTH_KEY);
    if (!data) {
      return nullState;
    }
    const state: AuthState = JSON.parse(data);
    if (!isTokenValid(state.token || "")) {
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
    login: (state: AuthState, action: PayloadAction<AuthState>) => {
      state.token = action.payload.token;
      state.user = action.payload.user;
      localStorage.setItem(AUTH_KEY, JSON.stringify(state));
    },
    logout: (state: AuthState) => {
      state.token = null;
      state.user = null;
      localStorage.removeItem(AUTH_KEY);
    },
  },
});

export const { login, logout } = authSlice.actions;
export default authSlice.reducer;

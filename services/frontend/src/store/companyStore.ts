import { configureStore, createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Company } from "../lib/company";

// Initial state (null means nothing selected yet)
const initialState: Company | null = null;

const selectedCompanySlice = createSlice({
  name: "selectedCompany",
  initialState,
  reducers: {
    setCompany(state: Company | null, action: PayloadAction<Company>) {
      if (state) {
        state.id = action.payload.id;
        state.code_name = action.payload.code_name;
      }
    },
    resetCompany() {
      return null;
    },
  },
});

const selectedCompanyStore = configureStore({
  reducer: selectedCompanySlice.reducer,
});

export type RootState = ReturnType<typeof selectedCompanyStore.getState>;
export type AppDispath = typeof selectedCompanyStore.dispatch;
export const { setCompany, resetCompany } = selectedCompanySlice.actions;

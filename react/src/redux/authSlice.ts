import { createSlice } from "@reduxjs/toolkit";

interface AuthState {
  isLogged: boolean;
}

const initialState: AuthState = {
  isLogged: false,
};

export const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    setLoginState: (state, action: { payload: boolean }) => {
      state.isLogged = action.payload;
    },
  },
});

export const { setLoginState } = authSlice.actions;

export default authSlice.reducer;

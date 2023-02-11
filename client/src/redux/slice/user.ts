import { createSlice, PayloadAction } from '@reduxjs/toolkit';

export interface UserState {
    isAuthenticated: boolean;
}

const initialState: UserState = {
    isAuthenticated: localStorage.getItem("apiToken") !== null && localStorage.getItem("apiToken") !== ""
}

export const userSlice = createSlice({
    name: "user",
    initialState,
    reducers: {
        changeUserAuthentication: (state, action: PayloadAction<boolean>) => {
            state.isAuthenticated = action.payload;
        }
    },
});

export const { changeUserAuthentication } = userSlice.actions;

export default userSlice.reducer;


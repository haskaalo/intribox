import { configureStore } from "@reduxjs/toolkit";
import UserReducer from "./slice/user";
import MediaGridReducer from "./slice/mediagrid";

const store = configureStore({
    reducer: {
        user: UserReducer,
        mediagrid: MediaGridReducer,
    }
});

export default store;


// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch


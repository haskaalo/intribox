import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface Album {
    id: string;
    title: string;
    description: string;
    created_at: number;
}

export interface AlbumListState {
    albumlist: Album[];
    albumID: Record<string, number> // Hashmap <ID, index> to avoid O(n**2) search time
};

const initialState: AlbumListState = {
    albumlist: [],
    albumID: {},
};

export const albumListSlice = createSlice({
    name: "albumlist",
    initialState,
    reducers: {
        addAlbums: (state, action: PayloadAction<Album[]>) => {
            // Remove albums that are already in loadedMedias
            const albumToAdd = action.payload.filter(album => !(album.id in state.albumID));

            const initialLength = state.albumlist.length;

            state.albumlist = state.albumlist.concat(albumToAdd);

            // Add IDs and index to albumID (hashset)
            for (let i = 0; i < albumToAdd.length; i++) {
                state.albumID[albumToAdd[i].id] = initialLength+i;
            }
        }
    }
});

export const { addAlbums } = albumListSlice.actions;

export default albumListSlice.reducer;

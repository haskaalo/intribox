import { createSlice, PayloadAction } from '@reduxjs/toolkit';


export interface Media {
    id: string; // uuid
    name: string; // name of media
    uploaded_time: number; // unix timestamp
    size: number; // as byte
    download_url: string // url to get get media from cdn
    downloaded: boolean // has the file been loaded into the browser?
}

export interface MediaGridState {
    loadedMedias: Media[],
    loadedMediaID: Record<string, number>, // Hashmap <ID, index> to avoid O(n**2) search time
}

const initialState: MediaGridState = {
    loadedMedias: [],
    loadedMediaID: {},
}

export const mediaGridSlice = createSlice({
    name: "mediagrid",
    initialState,
    reducers: {
        addMedia: (state, action: PayloadAction<Media[]>) => {
            // Remove medias that are already in loadedMedias
            const mediaToAdd = action.payload.filter(media => !(media.id in state.loadedMediaID));

            const initialLength = state.loadedMedias.length;

            state.loadedMedias = mediaToAdd.concat(state.loadedMedias);

            // Add IDs and index to loadedMediaID (hashset)
            for (let i = 0; i < mediaToAdd.length; i++) {
                state.loadedMediaID[mediaToAdd[i].id] = initialLength+i;
            }
        },
        mediaBeenDownloaded: (state, action: PayloadAction<string>) => { // payload is ID of media
            const index = state.loadedMediaID[action.payload];
            state.loadedMedias[index].downloaded = true;
        },
        reset: (state) => {
            state.loadedMediaID = initialState.loadedMediaID;
            state.loadedMedias = initialState.loadedMedias;
        }
     },
});

export const { addMedia, mediaBeenDownloaded, reset } = mediaGridSlice.actions;

export default mediaGridSlice.reducer;

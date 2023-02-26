import { giveErrorFromStatusCode, handleKnownError, KnownError } from "./error";

interface CreateNewEmptyAlbumParams {
    name: string;
}

export const CreateNewEmptyAlbum = async (params: CreateNewEmptyAlbumParams) => {
    const response = await fetch(`${BUILDCONFIG.apiUrl}/album/new`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json; charset=utf-8",
            "X-Intribox-Token": localStorage.getItem("apiToken")
        },
        body: JSON.stringify({
            title: params.name,
            description: "",
            media_ids: [],
        }),
    }).catch(() => {
        throw new Error(KnownError.NETWORK_ERROR);
    });

    const errorVal = giveErrorFromStatusCode(response.status);
    if (errorVal != null) {
        handleKnownError(errorVal);
        throw new Error(errorVal)
    }

    const responseJSON: {id: string} = await response.json();

    return responseJSON.id;
};

interface GetListAlbumResponse {
    id: string;
    title: string;
    description: string;
    created_at: number;
}

export const GetListAlbum = async () => {
    const response = await fetch(`${BUILDCONFIG.apiUrl}/album/list`, {
        method: "GET",
        headers: {
            "X-Intribox-Token": localStorage.getItem("apiToken"),
        }
    }).catch(() => {
        throw new Error(KnownError.NETWORK_ERROR);
    })

    const errorVal = giveErrorFromStatusCode(response.status);
    if (errorVal != null) {
        handleKnownError(errorVal);
        throw new Error(errorVal);
    }

    const responseJSON: GetListAlbumResponse[] = await response.json();

    return responseJSON;
}
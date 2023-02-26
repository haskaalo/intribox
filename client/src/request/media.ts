import { giveErrorFromStatusCode, handleKnownError, KnownError } from "./error";

interface UploadMediaParams {
    file: File;
}

interface UploadMediaResponse {
    id: string;
    name: string;
    uploaded_time: number;
    size: number;
    download_url: string;
}

export const UploadMedia = async (params: UploadMediaParams): Promise<UploadMediaResponse> => {
    const formData = new FormData();
    formData.append("file", params.file, params.file.name);
    formData.append("content-type", params.file.type);

    const response = await fetch(`${BUILDCONFIG.apiUrl}/media/new`, {
        method: "POST",
        body: formData,
        redirect: "follow",
        headers: {
            "X-Intribox-Token": localStorage.getItem("apiToken"),
        }
    }).catch(() => {
        throw new Error(KnownError.NETWORK_ERROR);
    });

    const errorVal = giveErrorFromStatusCode(response.status);

    if (errorVal !== null) {
        handleKnownError(errorVal);
        throw new Error(errorVal);
    }

    const responseJSON: UploadMediaResponse = await response.json();

    return responseJSON;
}

interface GetMediaListResponse {
    id: string;
    name: string;
    uploaded_time: number;
    size: number;
    download_url: string;
}

export const GetListMedia = async () => {
    const response = await fetch(`${BUILDCONFIG.apiUrl}/media/list`, {
        method: "GET",
        redirect: "follow",
        headers: {
            "X-Intribox-Token": localStorage.getItem("apiToken"),
        }
    }).catch(() => {
        throw new Error(KnownError.NETWORK_ERROR);
    });

    const errorVal = giveErrorFromStatusCode(response.status);

    if (errorVal != null) {
        handleKnownError(errorVal);
        throw new Error(errorVal);
    }

    const responseJSON: GetMediaListResponse[] = await response.json();

    return responseJSON;
}

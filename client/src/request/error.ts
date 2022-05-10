export enum KnownError {
    "NETWORK_ERROR" = "A network error happened while requesting data",
    "NOT_FOUND" = "Not Found",
    "INTERNAL_ERROR" = "Internal Error",
}

export const giveErrorFromStatusCode = (status: number) => {
    switch (status) {
        case 404: {
            return KnownError.NOT_FOUND;
        }
        case 500: {
            return KnownError.INTERNAL_ERROR;
        }
        default: {
            return null;
        }
    }
};
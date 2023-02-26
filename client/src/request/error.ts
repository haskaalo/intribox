import { changeUserAuthentication } from "@home/redux/slice/user";
import store from "@home/redux/store";

export enum KnownError {
    "NETWORK_ERROR" = "A network error happened while requesting data",
    "NOT_FOUND" = "Not Found",
    "INTERNAL_ERROR" = "Internal Error",
    "UNAUTHORIZED" = "Unauthorized",
    "CONFLICT" = "Conflict"
}

export const giveErrorFromStatusCode = (status: number) => {
    switch (status) {
        case 404: {
            return KnownError.NOT_FOUND;
        }
        case 500: {
            return KnownError.INTERNAL_ERROR;
        }
        case 401: {
            return KnownError.UNAUTHORIZED;
        }
        case 409: {
            return KnownError.CONFLICT;
        }
        default: {
            return null;
        }
    }
};

export const handleKnownError = (s: KnownError) => {
    // eslint-disable-next-line default-case
    switch (s) {
        case KnownError.UNAUTHORIZED: {
            store.dispatch(changeUserAuthentication(false));
            break;
        }
    }
}
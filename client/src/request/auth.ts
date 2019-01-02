import { KnownError, giveErrorFromStatusCode } from "./index";

export const LoginUser = async (email: string, password: string): Promise<string> => {
    const response = await fetch(`${BUILDCONFIG.apiUrl}/auth/login`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json; charset=utf-8",
        },
        redirect: "follow",
        body: JSON.stringify({email, password}),
    }).catch((err) => {
        throw new Error(KnownError.NETWORK_ERROR);
    });

    const ErrorVal = giveErrorFromStatusCode(response.status);

    if (ErrorVal !== null) {
        throw new Error(ErrorVal);
    }

    const responseJSON: {apiToken: string} = await response.json();

    return responseJSON.apiToken;
};

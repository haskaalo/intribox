import { KnownError, giveErrorFromStatusCode } from "./error";

export const LoginUser = async (email: string, password: string): Promise<string> => {
    const response = await fetch(`${BUILDCONFIG.apiUrl}/auth/login`, {
        method: "POST",
        headers: {
            // eslint-disable-next-line @typescript-eslint/naming-convention
            "Content-Type": "application/json; charset=utf-8",
        },
        redirect: "follow",
        body: JSON.stringify({email, password}),
    }).catch(() => {
        throw new Error(KnownError.NETWORK_ERROR);
    });

    const errorVal = giveErrorFromStatusCode(response.status);

    if (errorVal !== null) {
        throw new Error(errorVal);
    }

    const responseJSON: {apiToken: string} = await response.json();

    return responseJSON.apiToken;
};

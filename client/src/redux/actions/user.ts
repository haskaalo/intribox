import { ActionTypes, ActionsUnion } from "./types";

export interface IUser {
    isAuthenticated: boolean;
}

export const ChangeUserAuth = (authenticated: boolean) => ({
    type: ActionTypes.AUTH_USER as ActionTypes.AUTH_USER,
    payload: {
        isAuthenticated: authenticated,
    },
});
export type ChangeUserAuthAction = ReturnType<typeof ChangeUserAuth>;

const UserActions = { ChangeUserAuth };

export type UserActions = ActionsUnion<typeof UserActions>;

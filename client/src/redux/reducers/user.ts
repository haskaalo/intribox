import { IUser, UserActions } from "../actions/user";
import { ActionTypes } from "../actions/types";

const initialState: IUser = {
    isAuthenticated: localStorage.getItem("apiToken") !== null && localStorage.getItem("apiToken") !== "",
};

function user(state = initialState, action: UserActions): IUser {
    switch (action.type) {
        case ActionTypes.AUTH_USER:
            return {...state, isAuthenticated: action.payload.isAuthenticated};
        default:
            return state;
    }
}

export default user;

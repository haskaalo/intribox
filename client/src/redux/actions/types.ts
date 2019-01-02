export enum ActionTypes {
    "AUTH_USER" = "[user] authenticate user",
}

type FunctionType = (...args: any[]) => any;
type ActionsObjects = {[actionKey: string]: FunctionType};
export type ActionsUnion<A extends ActionsObjects> = ReturnType<A[keyof A]>;

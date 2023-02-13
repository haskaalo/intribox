import * as React from "react";
import { useSelector } from "react-redux";
import { Navigate } from "react-router-dom";
import { RootState } from "@home/redux/store";

function RequireSession({children}: {children: React.ReactNode}) {

    const isAuthenticated = useSelector((state: RootState) => state.user.isAuthenticated);
    
    if (!isAuthenticated) {
        return <Navigate to="/auth/sign_in" replace />
    }

    // eslint-disable-next-line react/jsx-no-useless-fragment
    return <>{children}</>;
}

export default RequireSession;

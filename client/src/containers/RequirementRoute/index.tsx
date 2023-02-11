import * as React from "react";
import { connect } from "react-redux";
import {UserState} from "@home/redux/slice/user";
import { RouteProps, Navigate } from "react-router-dom";


interface IPropsRedux {
    isAuthenticated: boolean;
}

class RequirementRoute extends React.Component<RouteProps & IPropsRedux> {
    render() {
        const {children, isAuthenticated} = this.props;

        if (!isAuthenticated) {
            return <Navigate to="/auth/sign_in" replace />;
        }

        return children;
    }
}

const mapStateToProps = ({user}: {user: UserState}): IPropsRedux => ({
    isAuthenticated: user.isAuthenticated,
});

export default connect(mapStateToProps)(RequirementRoute);

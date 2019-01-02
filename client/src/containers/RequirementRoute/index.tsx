import * as React from "react";
import { connect } from "react-redux";
import {IUser} from "@home/redux/actions/user";
import { Route, RouteProps, Redirect } from "react-router";

export interface IProps {
    redirectHomeAuth: boolean;
}

interface IPropsRedux {
    isAuthenticated: boolean;
}

// tslint:disable:no-shadowed-variable
class RequirementRoute extends React.Component<RouteProps & IProps & IPropsRedux> {
     render() {
        const {redirectHomeAuth, isAuthenticated, component: Component, ...props } = this.props;

        if (redirectHomeAuth === true) {
            return <Route {...props} render={(props) => isAuthenticated ? <Redirect to="/home" /> : <Component {...props} />} />;
        } else {
            return <Route {...props} render={(props) => isAuthenticated ? <Component {...props} /> : <Redirect to="/auth/sign_in" />} />;
        }
    }
}

const mapStateToProps = ({user}: {user: IUser}, ownProps: IProps): IPropsRedux => {
    return {
        isAuthenticated: user.isAuthenticated,
    };
};

export default connect(mapStateToProps)(RequirementRoute);

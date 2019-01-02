import * as React from "react";
import {BrowserRouter, Switch, Route, Redirect } from "react-router-dom";
import * as Loadable from "react-loadable";
import RequirementRoute from "./containers/RequirementRoute";
import "./styles/base.scss";
import { Provider } from "react-redux";
import store from "./redux/store";
import Loading from "./components/Loading";
import Home from "./components/pages/home";

const SignIn = Loadable({
    loader: () => import("./components/pages/auth/SignIn"),
    loading: Loading,
});

const App = () => (
    <Provider store={store}>
    <BrowserRouter>
        <Switch>
            <Route exact path="/" render={() => <Redirect to="/home" />} />
            <RequirementRoute redirectHomeAuth={false} path="/home" component={Home}/>
            <RequirementRoute redirectHomeAuth={true} path="/auth/sign_in" component={SignIn} />
        </Switch>
    </BrowserRouter>
    </Provider>
);

export default App;

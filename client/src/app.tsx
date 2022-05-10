import * as React from "react";
import {BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import * as Loadable from "react-loadable";
import { Provider } from "react-redux";
import RequirementRoute from "./containers/RequirementRoute";
import "./styles/base.scss";
import store from "./redux/store";
import Loading from "./components/Loading";
import Home from "./components/pages/home";


const SignIn = Loadable({
    loader: () => import("./components/pages/auth/SignIn"),
    loading: Loading,
});

class App extends React.Component {
    render() {
        return <Provider store={store}>
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Navigate to="/home" />}/>
                <Route path="/auth/sign_in" element={<SignIn />} />
                <Route path="/home" element={<RequirementRoute><Home /></RequirementRoute>} />
            </Routes>
        </BrowserRouter>
    </Provider>
    }
}
export default App;

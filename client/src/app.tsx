import * as React from "react";
import {BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import * as Loadable from "react-loadable";
import { Provider } from "react-redux";
import RequireSession from "./containers/RequireSession";
import "./styles/base.scss";
import store from "./redux/store";
import Loading from "./components/Loading";


const SignIn = Loadable({
    loader: () => import("./components/pages/auth/SignIn"),
    loading: Loading,
});

const Home = Loadable({
    loader: () => import("./components/pages/home"),
    loading: Loading
});

class App extends React.Component {
    render() {
        return <Provider store={store}>
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Navigate to="/home" />}/>
                <Route path="/auth/sign_in" element={<SignIn />} />
                <Route path="/home" element={<RequireSession><Home /></RequireSession>} />
            </Routes>
        </BrowserRouter>
    </Provider>
    }
}

export default App;

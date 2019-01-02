import * as React from "react";
import { Container } from "reactstrap";

import NavbarHeader from "./navbarheader";

class Home extends React.Component {
    render() {
        return <Container fluid={true}>
            <NavbarHeader/>
        </Container>;
    }
}

export default Home;

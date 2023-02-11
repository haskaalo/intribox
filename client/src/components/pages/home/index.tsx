import * as React from "react";
import { Container } from "reactstrap";
import MediaGrid from "./MediaGrid";

import NavbarHeader from "./navbarheader";

class Home extends React.Component {
    render() {
        return <Container fluid>
            <NavbarHeader/>
            <MediaGrid />
        </Container>;
    }
}

export default Home;

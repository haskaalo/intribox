import * as React from "react";
import { Container, Row } from "reactstrap";
import MediaGrid from "./MediaGrid";

import NavbarHeader from "./navbarheader";

class Home extends React.Component {
    render() {
        return <Container fluid>
            <NavbarHeader/>
            <Row>
                <MediaGrid />
            </Row>
        </Container>;
    }
}

export default Home;

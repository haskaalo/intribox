import * as React from "react";
import { Navbar, NavbarBrand, Form, Input, Col } from "reactstrap";

class NavbarHeader extends React.Component {
    render() {
        return <Navbar className="home-navbar">
            <NavbarBrand className="hidden-sm-block">IntriBox</NavbarBrand>
            <Col md={6} className="mx-auto">
                <Form>
                    <Input bsSize="lg" placeholder="Search albums, locations, etc.." />
                </Form>
            </Col>
        </Navbar>;
    }
}

export default NavbarHeader;

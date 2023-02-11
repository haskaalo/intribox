import { UploadMedia } from "@home/request/media";
import * as React from "react";
import { Navbar, NavbarBrand, Form, Input, Col, NavItem, Nav, Label } from "reactstrap";

function NavbarHeader() {
    async function doMediaUpload(file: File) {
        try {
            const response = await UploadMedia({file});
            alert(`Successfull ID: ${response.id}`);
        } catch (err) {
            alert(`Error while uploading: ${  err.message}`);
        }
    }

    function handleUpload(f: FileList) {
        for (let i = 0; i < f.length; i++) {
            doMediaUpload(f[i]);
        }
    }
    return <Navbar className="home-navbar">
    <NavbarBrand className="hidden-sm-block">IntriBox</NavbarBrand>
    <Col md={6} className="mx-auto">
        <Form>
            <Input bsSize="lg" placeholder="Search albums, locations, etc.." />
        </Form>
    </Col>
    <Nav>
        <NavItem>
            <Label >
                Upload
                <Input type="file" style={{display: "none"}} onChange={(e) => handleUpload(e.target.files)}/>
            </Label>
        </NavItem>
    </Nav>
</Navbar>;
}

export default NavbarHeader;

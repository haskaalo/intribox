import { addMedia } from "@home/redux/slice/mediagrid";
import { AppDispatch } from "@home/redux/store";
import { UploadMedia } from "@home/request/media";
import * as React from "react";
import { useDispatch } from "react-redux";
import { Navbar, NavbarBrand, Form, Input, Col, NavItem, Nav, Label } from "reactstrap";

function NavbarHeader() {
    const dispatch: AppDispatch = useDispatch();
    
    async function doMediaUpload(file: File) {
        try {
            const response = await UploadMedia({file});
            dispatch(addMedia([{...response, downloaded: false}]));
        } catch (err) {
            alert(`Error while uploading: ${  err.message}`); // TODO: Change this!
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
            <Label style={{cursor: "pointer"}}>
                Upload
                <Input type="file"  multiple style={{display: "none"}} onChange={(e) => handleUpload(e.target.files)}/>
            </Label>
        </NavItem>
    </Nav>
</Navbar>;
}

export default NavbarHeader;

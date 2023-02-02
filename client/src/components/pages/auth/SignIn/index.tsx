import React, {useState } from "react";
import { ChangeUserAuth } from "@home/redux/actions/user";
import { Container, Row, Form, FormGroup, Button, Input } from "reactstrap";
import { connect } from "react-redux";
import {useNavigate} from "react-router-dom";
import { LoginUser, KnownError } from "@home/request";

export interface IProps {
    UserAuth: typeof ChangeUserAuth;
}

function SignIn(props: IProps) {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const navigate = useNavigate();

    async function handleFormSubmit(event: React.FormEvent) {
        event.preventDefault();
        const { UserAuth } = props;

        const response = await LoginUser(email, password).catch((err: Error) => {
            if (err.message === KnownError.NOT_FOUND) {
                alert("User not found or password doesn't exist");
                
            } else {
                alert(err.message);
            }
        });

        if (response) {
            localStorage.setItem("apiToken", response);
            UserAuth(true);
            navigate("/home");
        }
    }

    return <Container fluid>
        <Row>
            <Form className="mx-auto text-center" onSubmit={e => handleFormSubmit(e)}>
                <h1>IntriBox</h1>
                <FormGroup>
                    <Input type="email" name="email" placeholder="E-mail" bsSize="lg" value={email} onChange={e => setEmail(e.target.value)} />
                </FormGroup>
                <FormGroup>
                    <Input type="password" name="password" placeholder="Password" bsSize="lg" value={password} onChange={e => setPassword(e.target.value)} />
                </FormGroup>
                <Button type="submit" color="primary" size="lg">Sign In</Button>
            </Form>
        </Row>
    </Container>;
}

export default connect(null, { UserAuth: ChangeUserAuth })(SignIn);
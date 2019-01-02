import * as React from "react";
import { Container, Row, Form, FormGroup, Button, Input } from "reactstrap";
import { LoginUser, KnownError } from "@home/request";
import { ChangeUserAuth } from "@home/redux/actions/user";
import { connect } from "react-redux";

export interface IState {
    input: {
        email: string;
        password: string;
    };
}

export interface IProps {
    UserAuth: typeof ChangeUserAuth;
}

class SignIn extends React.Component<IProps, IState> {
    constructor(props: IProps) {
        super(props);
        this.state = {
            input: {
                email: "",
                password: "",
            },
        };

        this.handleFormSubmit = this.handleFormSubmit.bind(this);
        this.handleInputChange = this.handleInputChange.bind(this);
    }

    render() {
        return <Container fluid={true}>
            <Row>
                <Form className="mx-auto text-center" onSubmit={this.handleFormSubmit}>
                    <h1>IntriBox</h1>
                    <FormGroup>
                        <Input type="email" name="email" placeholder="E-mail"  bsSize="lg" value={this.state.input.email} onChange={this.handleInputChange} />
                    </FormGroup>
                    <FormGroup>
                        <Input type="password" name="password" placeholder="Password" bsSize="lg" value={this.state.input.password} onChange={this.handleInputChange} />
                    </FormGroup>
                    <Button type="submit" color="primary" size="lg">Sign In</Button>
                </Form>
            </Row>
        </Container>;
    }

    private handleInputChange(event: React.ChangeEvent<HTMLInputElement>) {
        const currentState = this.state.input;
        currentState[event.target.name as keyof IState["input"]] = event.target.value;

        this.setState({
            input: currentState,
        });
    }

    private async handleFormSubmit(event: React.FormEvent) {
        event.preventDefault();

        const response = await LoginUser(this.state.input.email, this.state.input.password).catch((err: Error) => {
            if (err.message === KnownError.NOT_FOUND) {
                alert("User not found or password doesn't exist");
                return;
            } else {
                alert(err.message);
                return;
            }
        });

        if (response) {
            localStorage.setItem("apiToken", response);
            this.props.UserAuth(true);
        }
    }
}

export default connect(null, { UserAuth: ChangeUserAuth })(SignIn);

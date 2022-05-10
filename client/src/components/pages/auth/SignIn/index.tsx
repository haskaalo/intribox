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

    private handleInputChange(event: React.ChangeEvent<HTMLInputElement>) {
        const { input } = this.state;

        input[event.target.name as keyof IState["input"]] = event.target.value;
        this.setState({input});
    }

    private async handleFormSubmit(event: React.FormEvent) {
        event.preventDefault();
        const { input } = this.state;
        const { UserAuth } = this.props;

        const response = await LoginUser(input.email, input.password).catch((err: Error) => {
            if (err.message === KnownError.NOT_FOUND) {
                alert("User not found or password doesn't exist");
                
            } else {
                alert(err.message);
                
            }
        });

        if (response) {
            localStorage.setItem("apiToken", response);
            UserAuth(true);
        }
    }

    render() {
        const { input } = this.state;
        return <Container fluid>
            <Row>
                <Form className="mx-auto text-center" onSubmit={this.handleFormSubmit}>
                    <h1>IntriBox</h1>
                    <FormGroup>
                        <Input type="email" name="email" placeholder="E-mail"  bsSize="lg" value={input.email} onChange={this.handleInputChange} />
                    </FormGroup>
                    <FormGroup>
                        <Input type="password" name="password" placeholder="Password" bsSize="lg" value={input.password} onChange={this.handleInputChange} />
                    </FormGroup>
                    <Button type="submit" color="primary" size="lg">Sign In</Button>
                </Form>
            </Row>
        </Container>;
    }
}

export default connect(null, { UserAuth: ChangeUserAuth })(SignIn);

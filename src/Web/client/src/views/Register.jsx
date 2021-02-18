import React from "react";
import {CONSTANTS} from "../env";
// reactstrap components
import {
  Button,
  Card,
  CardHeader,
  CardBody,
  CardFooter,
  CardText,
  FormGroup,
  Form,
  Input,
  Row,
  Col,
  Alert
} from "reactstrap";
import { Link, Redirect } from "react-router-dom";
import * as axios from "axios";




class Register extends React.Component {
  state = {
    name: null,
    email: null,
    organization: null,
    password: null,
    notification: {
      status: null,
      message: null
    },
    loading: false

  }

  componentDidMount(){
    var isUserLoggedIn = sessionStorage.getItem("MOTI_USER")
    var tokenAvail = sessionStorage.getItem("MOTI_TOKEN")

    if(isUserLoggedIn && tokenAvail){
      return this.props.history.push("/admin/dashboard")
    }
    sessionStorage.clear()
  }

  async RegisterUser(){
      if(!this.state.email || !this.state.password, !this.state.name, !this.state.organization){
          return false
      }
      this.setState({notification:{status: null, message: null}, loading: true})
      await this.submitRegister()
  }


  async submitRegister(){

    try {
        var data = this.state;
        delete data.notification
        delete data.loading
        
        let response = await axios.post(`${CONSTANTS.baseUrl}/api/v1/Auth/Register`, data);
        this.setState({notification:{status: "success", message: response.data.message}});

    } catch(error){
      if(error.response){
        return this.setState({notification:{status: "danger", message: error.response.data.message}, loading: false});
      }
      return this.setState({notification:{status: "danger", message: "A network problem has occured"}, loading: false});
    }

    this.setState({loading: false})
  }

  renderNotifications(){
    if(this.state.notification.status){
      return (<Alert color={this.state.notification.status}>{this.state.notification.message}</Alert>)
    }
    return null
  }
  render() {
    return (
      <>
          <Row className="text-center">
            <Col md="4" sm="8" xl="4" style={{margin: '10% auto'}}>
                
            <Form autoComplete="false" onSubmit={(e)=>{e.preventDefault();this.RegisterUser()}}>
              <Card>
              <h3 className="text-center" style={{margin: '20px 0', padding: 0}}>TROLL•Tower</h3>
                <CardBody>
                  {this.renderNotifications()}
                  <Row>
                      <Col md="12">
                        <FormGroup>
                          <Input placeholder="Your Name" type="text" id="name" required defaultValue={this.state.name} style={{marginTop: 10, marginBottom: 10}}
                            onChange={(e)=>{
                              this.setState({name: e.target.value})
                          }}/>
                        </FormGroup>
                      </Col>
                    </Row>
                    <Row>
                      <Col md="12">
                        <FormGroup>
                          <Input placeholder="Your Organization Name" type="text" id="organization" required defaultValue={this.state.organization} style={{marginTop: 10, marginBottom: 10}}
                            onChange={(e)=>{
                              this.setState({organization: e.target.value})
                          }}/>
                        </FormGroup>
                      </Col>
                    </Row>
                    <Row>
                      <Col md="12">
                        <FormGroup>
                          <Input placeholder="Email" type="email" id="email" required defaultValue={this.state.email} style={{marginTop: 10, marginBottom: 10}}
                            onChange={(e)=>{
                              this.setState({email: e.target.value})
                          }}/>
                        </FormGroup>
                      </Col>
                    </Row>
                    <Row>
                      <Col md="12">
                        <FormGroup>
                          <Input
                            placeholder="Your Password"
                            type="password"
                            id="password"
                            defaultValue={this.state.password}
                            style={{marginTop: 10, marginBottom: 10}}
                            required
                            onChange={(e)=>{
                              this.setState({password: e.target.value})
                            }}
                          />
                        </FormGroup>
                      </Col>
                    </Row>
                  
                </CardBody>
                <CardFooter>
                  <Button className="btn-fill" color="primary" type="submit" disabled={this.state.loading==true}>
                    Register
                  </Button>
                  <p className="text-center" style={{margin: "30px 0 0"}}>
                    <Link to="/auth/login">
                    Already have an account ? Login
                  </Link>
                  </p>
                </CardFooter>
              </Card>
              </Form>
            </Col>
          </Row>
      </>
    );
  }
}

export default Register;

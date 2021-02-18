import React from "react";
import {CONSTANTS} from "../env";
// reactstrap components
import {
  Button,
  Card,
  CardHeader,
  CardBody,
  CardFooter,
  FormGroup,
  Form,
  Input,
  Row,
  Col,
  Nav,
  Alert
} from "reactstrap";
import SettingsNavbar from "../components/Sidebar/SettingsNavbar"
import * as axios from "axios";
import auth from "../common/auth";
class UserProfile extends React.Component {

  state = {
    email: null,
    username: null,
    name: null,
    surname: null,
    organization: null,
    notification: {
      status: null,
      message: null
    },
    user: null,
    loading: false,
    token: null
  }

  async UpdateUser(){
      this.setState({notification:{status: null, message: null}, loading: true})
      await this.submitUpdate()
  }

  async submitUpdate(){

    var cid = this.state.user.id
    var tkn = this.state.token
    try {
        var data = {
          name: this.state.name,
          surname: this.state.surname,
          organization: this.state.organization
        }
        
        let response = await axios.put(`${CONSTANTS.baseUrl}/api/v1/User/${cid}`, data, { headers: { Authorization: `Bearer ${tkn}` } });
        this.setState({notification:{status: "success", message: response.data.message}});
        await this.getUser()
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

  async getUser(){
    this.setState({notification:{status: null, message: null}, loading: true})
    
    console.log(this.state.user)
    try {
      let response = await axios.get(`${CONSTANTS.baseUrl}/api/v1/User/${this.state.user.id}`, { headers: { Authorization: `Bearer ${this.state.token}` } });
      if(response.data.status){
        this.setState({
          email: response.data.data.email,
          name: response.data.data.name,
          surname: response.data.data.surname
        });
      }
      console.log(response)
      this.setState({notification:{status: null, message: null}, loading: false})
      
    } catch(error){
      if(error.response){
        return this.setState({notification:{status: "danger", message: error.response.data.message}, loading: false});
      }
      return this.setState({notification:{status: "danger", message: "A network problem has occured"}, loading: false});
    }
  }

  async componentWillMount(){
    this.setState({user: auth.getUser(), token: auth.getToken()});
    await this.getUser();
  }

  render() {
    return (
      <>
        <div className="content">
          <Row>
            <Col md="11" lg="11" sm="11">
            <Form onSubmit={async (e)=>{e.preventDefault();await this.UpdateUser()}}>
              <Card>
                <CardHeader>
                  <h5 className="title">Edit Profile</h5>
                </CardHeader>
                <CardBody>
                  
                    <Row>
                      <Col className="pr-md-1" md="6" lg="6">
                        <FormGroup>
                          <label>Name</label>
                          <Input
                            defaultValue={this.state.name}
                            placeholder="Name"
                            type="text"
                            required
                            onChange={(e)=>{
                              this.setState({name: e.target.value.trim()})
                            }}
                          />
                        </FormGroup>
                      </Col>
                    </Row>
                    <Row>
                      <Col className="pr-md-1" md="6" lg="6">
                        <FormGroup>
                          <label htmlFor="exampleInputEmail1">
                            Email address
                          </label>
                          <Input placeholder="mike@email.com" type="email" onChange={(e)=>{
                              this.setState({name: e.target.value.trim()})
                            }}
                            disabled
                            defaultValue={this.state.email}
                          />
                        </FormGroup>
                      </Col>
                      <Col className="pl-md-1" md="6" lg="6">
                        <FormGroup>
                          <label>Organization</label>
                          <Input
                            defaultValue={this.state.organization}
                            placeholder="Organization"
                            type="text"
                            required
                            disabled
                            onChange={(e)=>{
                              this.setState({name: e.target.value.trim()})
                            }}
                            defaultValue={this.state.organization}
                          />
                        </FormGroup>
                      </Col>
                    </Row>
                  
                </CardBody>
                <CardFooter>
                  <Button className="btn-fill" color="primary" type="submit" disabled={this.state.loading}>
                    Save
                  </Button>
                </CardFooter>
              </Card>
              </Form>
            </Col>
            <Col md="1" sm="1" lg="1" style={{padding: 0}}>
              <SettingsNavbar />
            </Col>
          </Row>
        </div>
      </>
    );
  }
}

export default UserProfile;

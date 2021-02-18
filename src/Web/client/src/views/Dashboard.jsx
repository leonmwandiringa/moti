import React from "react";
// nodejs library that concatenates classes

// reactstrap components
import {
  Button,
  ButtonGroup,
  Card,
  CardHeader,
  CardBody,
  CardTitle,
  DropdownToggle,
  DropdownMenu,
  DropdownItem,
  UncontrolledDropdown,
  Label,
  FormGroup,
  Alert,
  Table,
  Row,
  Col,
  UncontrolledTooltip
} from "reactstrap";
import * as axios from "axios";
import auth from "../common/auth";
import {CONSTANTS} from "../env";
import { Link } from "react-router-dom";

class Dashboard extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      data: [],
      notification: {
        status: null,
        message: null
      },
      loading: false
    };
  }

  async getUserTunnels(){
    try {
        var data = []
        var token = auth.getToken();
        this.setState({loading: true})

        let response = await axios.get(`${CONSTANTS.baseUrl}/api/v1/pipeline/tunnel`, {headers: { Authorization: `Bearer ${token}` }});
        response.data.status || response.data.data ? this.setState({data: response.data.data}) : this.setState({data: []}) 

    } catch(error){
      if(error.response){
        return this.setState({notification:{status: "danger", message: error.response.data.message || "Resource was not found"}, loading: false});
      }
      return this.setState({notification:{status: "danger", message: "A network problem has occured"}, loading: false});
    }
    this.setState({loading: false})
  }

  componentWillMount(){
      this.getUserTunnels()
  }

  renderNotifications(){
    if(this.state.notification.status){
      return (<Alert style={{width: '50%', margin: 'auto', textAlign: 'center', position: 'absolute', left: '25%', zIndex: '99999999'}} color={this.state.notification.status}>{this.state.notification.message}</Alert>)
    }
    return null
  }

  rendersTunnels(){
    if(this.state.data.length == 0){
      return (<h4>No Tunnels were found</h4>)
    }
    let tunnelVals = this.state.data
    return (tunnelVals.map((val, index)=>{
      return (<tr key={val.server_name}>
          <td className="td-actions text-right">{val.domain_name}</td>
          <td className="td-actions text-right">{val.domain_plan}</td>
          <td className="td-actions text-right">{val.max_quota}</td>
          <td className="td-actions text-right">{val.quota}</td>
          <td className="td-actions text-right">{val.status}</td>
          <td className="td-actions text-right">{val.created_at}</td>
        </tr>)
      }))
  }

  render() {
    return (
      <>
      {this.renderNotifications()}
        <div className="content">
          <Row>
            <Col lg="12" md="12">
              <Card className="card-tasks">
                {/* <CardHeader>
                  <h6 className="title d-inline">Server/Instances</h6>
                </CardHeader> */}
                <CardBody>
                  <div className="table-full-width">
                    <Table>
                      <thead>
                        <th>Domain</th>
                        <th>Domain Plan</th>
                        <th>Max Hits</th>
                        <th>Hits</th>
                        <th>Status</th>
                        <th>Created At</th>
                      </thead>
                      <tbody>
                        {this.rendersTunnels()}
                      </tbody>
                    </Table>
                  </div>
                </CardBody>
              </Card>
            </Col>
            </Row>
        </div>
      </>
    );
  }
}

export default Dashboard;

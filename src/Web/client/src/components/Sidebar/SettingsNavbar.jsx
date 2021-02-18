import React from "react";
import {
    Card,
    CardBody,
    Nav,
  } from "reactstrap";
// reactstrap components

class SettingsSidebar extends React.Component {
  constructor(props) {
    super(props);
    this.activeRoute.bind(this);
  }
  // verifies if routeName is the one active (in browser input)
  activeRoute(routeName) {
    return window.location.pathname == routeName ? "active" : "";
  }

  render() {
    const { bgColor, routes, logo } = this.props;
    return (
        <Card className="card-user">
            <CardBody>
                <Nav>
                    <li
                        className={this.activeRoute("/admin/profile")}
                            key="1"
                        >
                            <a href="/admin/profile"
                                className="nav-link"
                                style={{opacity: this.activeRoute("/admin/profile") ? 1 : 0.6}}
                            >
                            <i className="tim-icons icon-single-02" />
                        </a>
                    </li>
                </Nav>
                <Nav>
                    <li
                        className={this.activeRoute("/admin/organization")}
                            key="1"
                        >
                            <a href="/admin/organization"
                                className="nav-link"
                                style={{opacity: this.activeRoute("/admin/organization") ? 1 : 0.6}}
                            >
                            <i className="tim-icons icon-bank" />
                        </a>
                    </li>
                </Nav>
            </CardBody>
      </Card>
    );
  }
}

export default SettingsSidebar;

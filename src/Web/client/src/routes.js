import Dashboard from "views/Dashboard.jsx";
import UserProfile from "views/UserProfile.jsx";
import Login from "views/Login.jsx";
import Register from "views/Register.jsx";

var routes = [
  {
    path: "/dashboard",
    name: "Dashboard",
    icon: "tim-icons icon-chart-pie-36",
    component: Dashboard,
    Allowed: ["DEVELOPER", "ADMIN", "SUPERADMIN"],
    layout: "/admin",
    sidebar: true
  },
  {
    path: "/dashboard",
    name: "Domains",
    icon: "tim-icons icon-planet",
    component: Dashboard,
    Allowed: ["ADMIN", "SUPERADMIN"],
    layout: "/admin",
    sidebar: true
  },
  {
    path: "/dashboard",
    name: "Sessions",
    icon: "tim-icons icon-planet",
    component: Dashboard,
    Allowed: ["DEVELOPER", "ADMIN", "SUPERADMIN"],
    layout: "/admin",
    sidebar: true
  },
  {
    path: "/dashboard",
    name: "Metrics",
    icon: "tim-icons icon-chart-bar-32",
    component: Dashboard,
    Allowed: ["ADMIN", "SUPERADMIN", "DEVELOPER"],
    layout: "/admin",
    sidebar: true
  },
  {
    path: "/profile",
    name: "User Profile",
    icon: "tim-icons icon-single-02",
    component: UserProfile,
    Allowed: ["ADMIN", "SUPERADMIN", "DEVELOPER"],
    layout: "/admin",
    sidebar: false
  },
  {
    path: "/login",
    name: "Login",
    icon: "tim-icons icon-single-02",
    component: Login,
    layout: "/auth",
    sidebar: false
  },
  {
    path: "/register",
    name: "Register",
    icon: "tim-icons icon-single-02",
    component: Register,
    layout: "/auth",
    sidebar: false
  }
];
export default routes;

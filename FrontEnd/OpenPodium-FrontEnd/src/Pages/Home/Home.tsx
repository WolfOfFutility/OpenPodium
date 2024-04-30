/* eslint-disable @typescript-eslint/no-explicit-any */

import "./Home.css"

// Import Modules
import { useState, useEffect } from 'react';
import { Link, NavigateFunction, useLocation, useNavigate } from 'react-router-dom';

// Models
import AuthObject from '../../Models/AuthObject';

// Redirect to the login page if there is no auth.
const checkForAuth = (state: any, navigate: NavigateFunction, setAuth: React.Dispatch<React.SetStateAction<AuthObject>>) => {
  if(state == null) {
    navigate("/login")
  }
  else {
    setAuth(state.auth)
  }
}

const menuConfig = [
  {
    name: "Home",
    url: "/home"
  },
  {
    name: "Repos",
    url: "/repos"
  },
  {
    name: "Pipelines",
    url: "/pipelines"
  },
  {
    name: "Environments",
    url: "/environments"
  },
  {
    name: "Applications",
    url: "/applications"
  },
  {
    name: "Settings",
    url: "/settings"
  },
]

const Home = () => {
    const navigate = useNavigate();
    const [auth, setAuth] = useState(new AuthObject());
    const state = useLocation().state;
    const path = useLocation().pathname;

    // Render the page, look for auth
    useEffect(() => {
        checkForAuth(state, navigate, setAuth);
    }, [state, navigate, setAuth])

    return (
        <div id="home-wrapper">
          <div id="menu-wrapper">
            <div id="menu-wrapper-title">
              <h2>OpenPodium</h2>
            </div>
            {menuConfig.map((menuItem, index) => {
              if(path == menuItem.url) {
                return <Link key={index} to={menuItem.url} className="menu-item-active">{menuItem.name}</Link>
              }
              else {
                return <Link key={index} to={menuItem.url} className="menu-item">{menuItem.name}</Link>
              }
            })}
          </div>
          <div id="content-wrapper">
            Home - {auth.Username}
          </div>
        </div>
    )
}

export default Home
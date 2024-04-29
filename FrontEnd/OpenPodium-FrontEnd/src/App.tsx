/* eslint-disable @typescript-eslint/no-explicit-any */

import './App.css'

// Import Modules
import { useState, useEffect } from 'react';
import { NavigateFunction, useLocation, useNavigate } from 'react-router-dom';

// Models
import AuthObject from './Models/AuthObject';

// Redirect to the login page if there is no auth.
const checkForAuth = (state: any, navigate: NavigateFunction, setAuth: React.Dispatch<React.SetStateAction<AuthObject>>) => {
  if(state == null) {
    navigate("/login")
  }
  else {
    setAuth(state.auth)
  }
}

function App() {
  const navigate = useNavigate();
  const [auth, setAuth] = useState(new AuthObject());
  const state = useLocation().state;

  // Render the page, look for auth
  useEffect(() => {
    checkForAuth(state, navigate, setAuth);
  }, [state, navigate, setAuth])
  
  return (
    <>
      <p>Welcome back, {auth.Username}!</p>
    </>
  )
}

export default App

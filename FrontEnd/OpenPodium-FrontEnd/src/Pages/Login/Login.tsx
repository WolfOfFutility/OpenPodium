// Base Imports
import {useState} from 'react'
import "./Login.css"

// Dependencies
import {NavigateFunction, useNavigate} from 'react-router-dom'

// Functions
import PlatformVariables from "../../../PlatformVariables.config.ts"
import APIRequest from '../../Functions/APIRequest'

// Models
import AuthObject from '../../Models/AuthObject.ts'

// Send Login Request, taking the inputs from the login fields and passing it to the server
const sendLoginRequest = async (username: string, password: string, navigate: NavigateFunction) => {
    await APIRequest(`http://${PlatformVariables.Backend.hostAddress}:${PlatformVariables.Backend.hostPort}/login`, "POST", {username, password}, (res) => successNavigate(Object.assign(new AuthObject(), res), navigate), (err) => handleLoginFail(err))
}

// Handle the success of the Login Request, navigate and pass the client token
const successNavigate = (res: AuthObject, navigate: NavigateFunction) => {
    navigate("/home", {state: {auth: res}})
}

// Handle the failure of the login request
const handleLoginFail = (err: object) => {
    alert("Login Fail!")
    console.error(err)
}

// Login Page
const Login = () => {
    // Handle username and password fields being changed
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const navigate = useNavigate();

    // Render the page
    return (
        <div id="login-box">
            <div className="title-box">
                <h2>Login</h2>
            </div>
            <div className="input-group">
                <label id="username-field-label" htmlFor="username-field">Username:</label>
                <input type="text" name="username-field" id="username-field" onChange={(e) => setUsername(e.target.value)} />
                {/* <br /> <br /> */}
            </div>
            <div className="input-group">
                <label id="password-field-label" htmlFor="password-field">Password:</label>
                <input type="text" name="password-field" id="password-field" onChange={(e) => setPassword(e.target.value)} />
                {/* <br /> <br /> */}
            </div>
            <div id="login-box-button-group">
                <button id="login-button" onClick={() => sendLoginRequest(username, password, navigate)}>Login</button>
                <button>Create Account</button>
            </div>
        </div>
    )
}

export default Login
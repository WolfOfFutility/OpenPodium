// Base Imports
import {useState} from 'react'
import "./Login.css"

// Dependencies
import {Navigate} from 'react-router-dom'

// Functions
import PlatformVariables from "../../../PlatformVariables.config.ts"
import APIRequest from '../../Functions/APIRequest'

// Send Login Request, taking the inputs from the login fields and passing it to the server
const sendLoginRequest = async (username: string, password: string) => {
    await APIRequest(`http://${PlatformVariables.Backend.hostAddress}:${PlatformVariables.Backend.hostPort}/login`, "POST", {username, password}, (res) => successNavigate(res), (err) => handleLoginFail(err))
}

// Handle the success of the Login Request, navigate and pass the client token
const successNavigate = (res: object) => {
    // Navigate({to: "/", state: res})

    console.log(res)
}

// Handle the failure of the login request
const handleLoginFail = (err: object) => {
    console.log("Fail!")
    console.error(err)
}

// Login Page
const Login = () => {
    // Handle username and password fields being changed
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

    // Render the page
    return (
        <>
            <label id="username-field-label" htmlFor="username-field">Username:</label>
            <input type="text" name="username-field" id="username-field" onChange={(e) => setUsername(e.target.value)} /> <br /> <br />
            <label id="password-field-label" htmlFor="password-field">Password:</label>
            <input type="text" name="password-field" id="password-field" onChange={(e) => setPassword(e.target.value)} /> <br /> <br />
            <button id="login-button" onClick={() => sendLoginRequest(username, password)}>Login</button>
        </>
    )
}

export default Login
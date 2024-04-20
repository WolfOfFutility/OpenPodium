// Onload event for the page to test api on refresh
window.onload = (event) => {
    console.log("Loaded!");

    // getHello();

    podiumLogin("Admin", "admin")
    // addUserSecret();

    readUserSecrets();
}

// Constructor for generic API requests
const APIRequest = async (url, method = "GET", extraHeaders = {}, bodyData = {}, successCallback = (res) => {}, errorCallback = (err) => {}) => {
    // Set Default Headers
    let headers = new Headers();

    // Allow for extra headers to be set as well
    for(let extraHeader in extraHeaders) {
        headers.set(extraHeader, extraHeaders[extraHeader]);
    }

    let APISettings = {}

    // Set APISettings, GET cannot have a body
    if(method == "GET") {
        APISettings = {
            method,
            headers,
            "Access-Control-Allow-Origin": "*",
            "Access-Control-Allow-Credentials": "true",
            "Access-Control-Allow-Methods": "GET,HEAD,OPTIONS,POST,PUT",
            "Access-Control-Allow-Headers": "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers"
        }
    }
    else {
        APISettings = {
            method,
            headers,
            body: JSON.stringify(bodyData),
            "Access-Control-Allow-Origin": "*",
            "Access-Control-Allow-Credentials": "true",
            "Access-Control-Allow-Methods": "GET,HEAD,OPTIONS,POST,PUT",
            "Access-Control-Allow-Headers": "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers"
        }
    }
    
    // Call the api, allow for error handling and callbacks
    await fetch(url, APISettings)
    .then(response => response.text())
    .then(res => {
        try {
            successCallback(res);
        }
        catch(e) {
            console.error(e);
            return res;
        }
    })
    .catch(err => {
        try {
            errorCallback(err);
        }
        catch(e) {
            console.error(e);
            return err;
        }
    })
}

const handleAPISuccess = (res) => {
    console.log("Success Handler");
    console.log(res);
}

const handleAPIFailure = (err) => {
    console.log("Handle Failure Function")
    console.error(err);
}

const getHello = async () =>{
    await APIRequest("http://localhost:4000/hello", "GET", {}, {}, (res) => handleAPISuccess(res), (err) => handleAPIFailure(err));
}

const podiumLogin = async (username, password) => {
    await APIRequest("http://localhost:4000/login", "POST", {}, {username, password}, (res) => handleAPISuccess(res), (err) => handleAPIFailure(err))
}

const addUserSecret = async () => {
    await APIRequest("http://localhost:4000/secrets/new", "POST", {}, {}, (res) => handleAPISuccess(res), (err) => handleAPIFailure(err))
}

const readUserSecrets = async () => {
    await APIRequest("http://localhost:4000/secrets", "GET", {}, {}, (res) => handleAPISuccess(res), (err) => handleAPIFailure(err))
}
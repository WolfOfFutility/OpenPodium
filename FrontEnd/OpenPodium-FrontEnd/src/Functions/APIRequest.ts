// Constructor for generic API requests
const APIRequest = async (url: string, method : string = "GET", bodyData: object = {}, successCallback = (res: object) => {}, errorCallback = (err: object) => {}) => {
    // Set Default Headers
    const headers = new Headers();

    // Allow for extra headers to be set as well
    // for(const extraHeader in extraHeaders) {
    //     headers.set(extraHeader, extraHeaders[extraHeader]);
    // }

    let APISettings = {}

    // Set APISettings, GET cannot have a body
    if(method == "GET") {
        APISettings = {
            method,
            headers,
            "Access-Control-Allow-Origin": "*",
            "Access-Control-Allow-Credentials": "true",
            "Access-Control-Allow-Methods": "GET,HEAD,OPTIONS,POST,PUT,LIST",
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
            "Access-Control-Allow-Methods": "GET,HEAD,OPTIONS,POST,PUT,LIST",
            "Access-Control-Allow-Headers": "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers"
        }
    }
    
    // Call the api, allow for error handling and callbacks
    await fetch(url, APISettings)
    .then(response => response.json())
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

export default APIRequest
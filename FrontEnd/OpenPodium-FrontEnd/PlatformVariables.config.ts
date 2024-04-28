// Set some default variables to allow for un-complicated building
// Backend refers to the Go Server in the OpenPodium/Server folder
// Frontend refers to the Vite + React + TS application in OpenPodium/Frontend (here)

const PlatformVariables = {
    Backend : {
        hostAddress: "127.0.0.1",
        hostPort: 8080
    },
    Frontend: {
        hostAddress: "127.0.0.1",
        hostPort: 3000
    }
}

export default PlatformVariables
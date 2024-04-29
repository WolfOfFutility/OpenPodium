class AuthObject {
    Username: string = "";
    ClientToken: string = "";

    constructor(Username?: string, ClientToken?: string) {
        this.Username = Username ?? "";
        this.ClientToken = ClientToken ?? "";
    }
}

export default AuthObject
class Auth{
    token;
    user;
    constructor(){
        var token = sessionStorage.getItem("TROLLTOWER_TOKEN");
        var user = sessionStorage.getItem("TROLLTOWER_USER");
        this.token = token;
        this.user = JSON.parse(user);
    }

    getUser(){
        return this.user ? this.user : null;
    }
    getToken(){
        return this.token ? this.token : null;
    }
    setSessionStorage(response){
        sessionStorage.setItem("TROLLTOWER_USER", JSON.stringify(response.data.data.user))
        sessionStorage.setItem("TROLLTOWER_TOKEN", String(response.data.data.token))
    }
}
export default new Auth();
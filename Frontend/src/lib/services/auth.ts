import {user, isAuthenticated} from "$lib/stores/auth";

function sleep(ms: number) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function login(email: string, password: string) {
    await sleep(1000)
    const response = await fetch("http://localhost:8080/auth/login",{
        mode: 'no-cors',    
        method:"POST",
        body: JSON.stringify({email,password}),
        headers:{'Content-Type': 'application/json'}});
    console.log(response)
    user.set({"fullname":"Askhat Zhankulov", "email": email})
    isAuthenticated.set(true);
}

async function logout() {
    user.set({})
    isAuthenticated.set(false);
    return true
}

const auth = {
    login,
    logout
}

export default auth
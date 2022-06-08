import {user, isAuthenticated} from "$lib/stores/auth";

function sleep(ms: number) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function login(email: string, password: string) {
    await sleep(1000)
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
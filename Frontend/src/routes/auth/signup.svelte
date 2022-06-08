<div class="auth-form-container">
    <Logo />
    <div class="auth-form-main">
        <h2>Sign up</h2>
        <form on:submit|preventDefault={login}>
            <!-- <div class="input">
                <input type="text" >
            </div> -->
            <Label>Email</Label>
            <Textfield 
                variant="outlined" 
                bind:value={email} style="width: 100%; border-color: #A8DADC; margin-bottom: 24px;"
                helperLine$style="width: 100%;"
            >
                <Icon class="material-icons-outlined" slot="leadingIcon">account_circle</Icon>
            </Textfield>
            <Label>Password</Label>
            <Textfield 
            variant="outlined" 
            bind:value={password} 
            style="width: 100%; margin-bottom: 24px;"
            helperLine$style="width: 100%;"
            >
                <Icon class="material-icons-outlined" slot="leadingIcon">vpn_key</Icon>
            </Textfield>
            <Button style="width: 100%; margin-bottom: 24px;" variant="unelevated">

                <Label>Sign up</Label>
            </Button>

            <Button style="width: 100%;" variant="outlined">
                <Label>Google</Label>
            </Button>
        </form>
    </div>
</div>



<script lang="ts">
    import Textfield from '@smui/textfield';
    import Logo from "$lib/components/Logo.svelte";
    import Button, { Label } from '@smui/button';
    import Icon from '@smui/textfield/icon';
    import { writable,get } from 'svelte/store';
    import auth from "$lib/services/auth";
    import { goto } from '$app/navigation';

    let email = '';
    let password = '';

    async function login(){
        let logedIn = await auth.login(email, password);
        goto("/overview")
    }
</script>

<style>
    .auth-form-container{
        position: absolute;
        width: 500px;
        transform: translate(-50%,-50%);
        top: 50%;
        left: 50%;
    }
    .auth-form-main{
        padding: 24px;
        border-radius: 15px;
        background-color: #F9FAFD;
    }
    .auth-form-main h2{
        margin-top: 0;
        text-align: center;
    }
    .input{
        width: 100%;
        padding: 5px;
        border-radius: 8px;
        background-color: white;
        border: 1px solid #EDEDED;
        border-bottom: 5px solid #A8DADC;
    }
    .input input{
        outline: none;
        border: none;
        background-color: transparent;
    }
</style>
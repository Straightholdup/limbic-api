<script lang="ts">
    import Logo from "$lib/components/Logo.svelte";
    import { isAuthenticated } from "$lib/stores/auth";
    import { get } from "svelte/store";
    import { onMount } from 'svelte';

    
    let selected: EventTarget|null;
    let navLine={
        width:0,
        left:0,
    };
    onMount(()=>{
        selected=document.querySelector(".nav-left li")
    })
    $: if(selected!=null){
        navLine={
            width: selected.offsetWidth,
            left: selected.offsetLeft,
        }
    }
 
    function select(e: MouseEvent){
        selected=e.target
        let selectedSection=document.querySelector(selected.dataset.section)
        console.log(selectedSection)
        window.scrollTo(0,selectedSection.offsetTop)
    }
</script>

<nav>
    <div class="container">
        <Logo />
        <div>
            <div class="nav-line" style="width: {navLine.width}px; left:{navLine.left}px"></div>
            <ul class="nav-left">
                <li on:click={select} data-section="#slogan-section">Services</li>
                <li on:click={select} data-section="#examples-section">Examples</li>
                <li on:click={select} data-section="#pricing-section">Pricing</li>
            </ul>
        </div>
        <ul class="nav-right">
            {#if get(isAuthenticated)}
                <li><a on:click={()=>alert(1)}>Sign out</a></li>
            {:else}
                <li><a href="/auth/signup">Sign in</a></li>
            {/if}
        </ul>
    </div>
</nav>

<style>
    nav{
        background-color: #364143;
        color: white;
        position: sticky;
        right: 0;
        left: 0;
        top:0;
        z-index:100;
    }
    div{
        display: flex;
    }
    ul{
        display: flex;
        list-style: none;
        align-items: center;
    }
    .nav-line{
        height: 0.4em;
        background-color: white;
        display: block;
        position: absolute;
        top: 0;
        transition: all 0.75s;
    }
    li{
        margin: 0 1rem;
    }
    .nav-right{
        margin-left: auto;
    }
    .nav-left{
        padding: 0 3rem;
    }
</style>
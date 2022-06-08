<script lang="ts">
    import Logo from "$lib/components/Logo.svelte";
    import { isAuthenticated } from "$lib/stores/auth";
    import { get } from "svelte/store";
</script>

<nav>
    <div class="container">
        <Logo />
        <ul class="nav-left">
            <li>Services</li>
            <li>Examples</li>
            <li>Pricing</li>
        </ul>
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
        background-color: #151d1e;
        color: white;
    }
    div{
        display: flex;
    }
    ul{
        display: flex;
        list-style: none;
        align-items: center;
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
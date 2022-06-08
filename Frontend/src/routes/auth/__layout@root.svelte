<slot />

<script context="module" lang="ts">
    // import {get} from 'svelte/store';
    import {isAuthenticated} from "$lib/stores/auth";
    import type { LoadOutput, LoadInput } from '@sveltejs/kit/types.internal';
    import { writable,get } from 'svelte/store';

    export async function load(): Promise<LoadOutput> {
        if(get(isAuthenticated)) return { status: 302, redirect: '/overview' }
        else return {}
    }
</script>
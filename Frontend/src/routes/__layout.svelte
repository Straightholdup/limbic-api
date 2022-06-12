<script lang="ts">
    import Nav from '$lib/components/Nav.svelte';
    import Footer from '$lib/components/Footer.svelte';
</script>

<Nav />

<main>
    <slot />
</main>

<Footer />
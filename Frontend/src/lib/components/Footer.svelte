<script lang="ts">
    import LayoutGrid, { Cell } from '@smui/layout-grid';
    import Logo from "$lib/components/Logo.svelte";
    import { isAuthenticated } from "$lib/stores/auth";
    import { get } from "svelte/store";
    import Textfield from '@smui/textfield';
    import Button, { Label } from '@smui/button';

    let email="";
</script>

<footer>
    <div class="container">
        <LayoutGrid style="padding: 0;">
            <Cell span={3}>
                <h4>RESOURCES</h4>
                <ul>
                    <li>About Limbic</li>
                    <li>Documentation</li>
                    <li>Pricing</li>
                    <li>API</li>
                </ul>
            </Cell>
            <Cell span={3}>
                <h4>SECURITY</h4>
                <ul>
                    <li>Terms of Service</li>
                    <li>Privacy Policy</li>
                    <li>Cookies Settings</li>
                </ul>
            </Cell>
            <Cell  span={6}>
                <h4>Subscribe to our newsletter</h4>
                <LayoutGrid style="padding:0;">
                    <Cell span={8}>
                        <Textfield 
                        class="textfield-white"
                        variant="outlined" 
                        bind:value={email} 
                        style="width: 100%; height: 36px;"
                        helperLine$style="width: 100%;"
                        >
                    </Textfield>
                    </Cell>
                    <Cell span={4}>
                        <Button style="width: 100%; height: 100%;" variant="unelevated">
                            <Label>Sign up</Label>
                        </Button>
                    </Cell>
                </LayoutGrid>
            </Cell>
        </LayoutGrid>
    </div>
</footer>

<style>
    *
    :global(.textfield-white
      .mdc-notched-outline
      .mdc-notched-outline__leading) {
        width: 28px;
        background-color: white;
    }
    *
    :global(.textfield-white
        .mdc-notched-outline
        .mdc-notched-outline__trailing) {
        background-color: white;
    }
    *
    :global(.textfield-white
        .mdc-notched-outline){
        z-index: -100 !important;
    }
    footer{
        background-color: #364143;
        color: white;
        padding: 64px 0;
        margin-top: 64px;
    }
    h4{
        padding:0;
        margin-top: 0;
        margin-bottom: 8px;
    }
    ul{
        list-style: none;
        padding: 0;
    }
    ul li{
        margin-top: 4px;
    }
</style>


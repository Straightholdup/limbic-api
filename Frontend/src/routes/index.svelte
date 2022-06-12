<script>
     import Button from '@smui/button';
     import animation from '$lib/assets/animation.gif';
     import examples_background from '$lib/assets/examples-background.png';
     import LayoutGrid, { Cell }  from '@smui/layout-grid';
     import { Icon, Label } from '@smui/common';
</script>
<svelte:head>
  <title>Welcome</title>
</svelte:head>

<section id="slogan-section">
    <div class="container">
        <div class="wrapper">
            
            <div class="about-box">
                <h1 class="company-title">Limbic</h1>
                <h3>Words are not always important, but it is important with what emotions they were said.</h3>
            </div>
            <div class="signin-box">
                <Button href="/auth" variant="unelevated">
                    <Label>Create an account</Label>
                </Button>
                <a href="/auth">
                    Free trial
                </a>
            </div>
            <div class="animation-box">
                <img src="{animation}" alt="animation"/>
            </div>
        </div>
    </div>
</section>

<section id="examples-section" style="background-image: url({examples_background}); background-size:cover; background-position: center;">
    <div class="container">
        <h1 class="title-center">Examples</h1>
        <ul class="selector">
            <li>Emotion Recognition</li>
            <li>Realtime Emotion Recognition</li>
        </ul>
    </div>
</section>

<section id="pricing-section">
    <div class="container" style="display: flex; flex-direction:column; align-items:center;">
        <h1 class="title-center">Pricing</h1>
        <p style="max-width: 800px; text-align:center; font-size:24px;">To use our services, you must purchase one or more subscriptions. After the purchase, you will have a token. More information can be found in the documentation section</p>
        <LayoutGrid style="padding:0; width:100%;">
            <Cell >
                <div class="price-card">
                    <div class="line"></div>
                    <h1>
                        Standard
                    </h1>
                    <div class="price">
                        <h3>As low as</h3>
                        <h3>30$ per month</h3>
                        <Button  variant="unelevated">
                            <Label>Purchase</Label>
                        </Button>
                    </div>
                    <ul class="offers">
                        <li><Icon class="material-icons" style="color:#0077CC;" >check_circle</Icon>1000 delayed emotion recognitions</li>
                        <li><Icon class="material-icons" style="color:#0077CC;" >check_circle</Icon>500 realtime emotion recognitions</li>
                    </ul>
                </div>
            </Cell>
            <Cell >
                <div class="price-card">
                    <div class="line"></div>
                    <h1>
                        Gold
                    </h1>
                    <div class="price">
                        <h3>As low as</h3>
                        <h3>45$ per month</h3>
                        <Button  variant="unelevated">
                            <Label>Purchase</Label>
                        </Button>
                    </div>
                    <ul class="offers">
                        <li><Icon class="material-icons" style="color:#0077CC;" >check_circle</Icon>1000 delayed emotion recognitions</li>
                        <li><Icon class="material-icons" style="color:#0077CC;" >check_circle</Icon>500 realtime emotion recognitions</li>
                    </ul>
                </div>
            </Cell>
            <Cell >
                <div class="price-card">
                    <div class="line"></div>
                    <h1>
                        Premium
                    </h1>
                    <div class="price">
                        <h3>As low as</h3>
                        <h3>60$ per month</h3>
                        <Button  variant="unelevated">
                            <Label>Purchase</Label>
                        </Button>
                    </div>
                    <ul class="offers">
                        <li><Icon class="material-icons" style="color:#0077CC;" >check_circle</Icon>1000 delayed emotion recognitions</li>
                        <li><Icon class="material-icons" style="color:#0077CC;" >check_circle</Icon>500 realtime emotion recognitions</li>
                    </ul>
                </div>
            </Cell>
        </LayoutGrid>
    </div>
</section>

<style>
    .price-card .line{
        background-color: #0077CC;
        height: 12px;
    }
    .offers{
        list-style:none;
        padding: 0;
    }
    .offers li{
        display: flex;
        align-items: center;
    }

    .selector{
        color: white;
        list-style: none;
        display: flex;
        font-size: 24px;
        font-weight: bold;
    }
    .selector li:last-child{
        margin-left: 24px;
    }

    .title-center{
        text-align: center;
        font-size: 48px;
        letter-spacing: 10px;
        text-transform: uppercase;
    }
    #examples-section .title-center{
        color:white;
    }

    #examples-section{
        height: 100vh;
        padding:1vh 0;
        width: 100%;
    }
    .company-title{
        font-size: 96px;
        padding: 0;
        margin: 0;
        color: #1D3557;
    }
    .animation-box img{
        width: 100%;
        height: 100%;
        object-fit: cover;
    }
    .sad{
        height: 100vh;
    }
    
    #slogan-section{
        height: 100vh;
        display: flex;
    }
    p{
        width: 100%;
    }
    .wrapper{
        display: grid;
        grid-template-columns:  repeat(2, 1fr);
        grid-template-rows: repeat(3, 1fr);
        grid-template-areas: 
        "about animation"
        "about animation"
        "signin animation";
        height: 500px;
    }
    .about-box{
        grid-area: 1 / 1 / 3 / 2;
        align-self: center; 
    }
    .signin-box{
        grid-area: 3 / 1 / 4 / 2;
    }
    .animation-box{
        grid-area: 1 / 2 / 4 / 3;
    }
    #pricing-section{
        height: 100vh;
    }


</style>
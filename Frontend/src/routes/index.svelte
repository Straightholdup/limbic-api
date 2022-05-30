<script>
     import Button, { Label } from '@smui/button';
     import animation from '$lib/assets/animation.gif';
     import examples_background from '$lib/assets/examples-background.png';
</script>
<svelte:head>
  <title>Welcome</title>
</svelte:head>

<section class="slogan-section">
    <div class="container">
        <div class="wrapper">
            <div class="about-box">
                <h1 class="company-title">Limbic</h1>
                <h3>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur at volutpat sapien, at sollicitudin mi</h3>
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

<section class="exaples-section" style="background-image: url({examples_background}); background-size:cover; background-position: center;">
    <div class="container">
        <h1 class="title-center">Examples</h1>
        <ul class="selector">
            <li>Emotion Recognition</li>
            <li>Realtime Emotion Recognition</li>
        </ul>
    </div>
</section>

<section class="sad">
</section>

<style>
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
        color:white;
        text-transform: uppercase;
    }

    .exaples-section{
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
        height: 10000px;
    }
    .slogan-section{
        padding: 250px 0;
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
</style>
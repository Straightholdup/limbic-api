<div class="container"> 
    <h1>Overview</h1>

    <LayoutGrid style="padding: 0;">
        <GridCell span={12}>
            <div class="card">
                <h3>My Tokens</h3>
                <!-- <DataTable table$aria-label="People list" style="width: 100%;">
                    <Head>
                        <Row>
                            <Cell>Token name</Cell>
                            <Cell>Status</Cell>
                            <Cell numeric>Delayed SER</Cell>
                            <Cell numeric>Realtime SER</Cell>
                            <Cell></Cell>
                        </Row>
                    </Head>

                </DataTable> -->
                <table>
                    <thead>
                        <tr>
                            <th>Token name</th>
                            <th>Status</th>
                            <th>Delayed SER</th>
                            <th>Realtime SER</th>
                            <th class="icon"></th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>asdas</td>
                            <td>asdas</td>
                            <td>asdas</td>
                            <td>asdas</td>
                            <td>
                                <Icon class="material-icons-outlined" slot="leadingIcon">settings</Icon>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </GridCell>
        <GridCell span={5}>
            <div class="card">
                <h3>Workload</h3>
                <canvas bind:this={wrokload} width="400" height="400"></canvas>
            </div>
        </GridCell>
        <GridCell span={7}>
            <div class="card">
                <h3>Logs</h3>
                <table>
                    <thead>
                        <tr>
                            <th>Timestamp</th>
                            <th>Response time</th>
                            <th>SER type</th>
                            <th>Token</th>
                            <th>Emotion</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>2022-05-31 10:57:51.378 +06:00</td>
                            <td>0.257 s</td>
                            <td>Realtime</td>
                            <td>eyJhbGci...AQK_ZXw</td>
                            <td>Happy</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </GridCell>
    </LayoutGrid>
</div>

<script lang="ts">
    import LayoutGrid, { Cell as GridCell } from '@smui/layout-grid';
    import DataTable, { Head, Body, Row, Cell } from '@smui/data-table';
    import Icon from '@smui/textfield/icon';
    import { onMount } from 'svelte';
    import Chart from 'chart.js/auto';

    let wrokload: any;
    const today=new Date();
    const time=today.getHours()+":"+today.getMinutes()
    const labels = [time,time,time,time,time,time,time,];
    const data = {
    labels: labels,
    datasets: [{
    label: 'My First Dataset',
    data: [65, 59, 80, 81, 56, 55, 40],
    fill: false,
    borderColor: 'rgb(75, 192, 192)',
    tension: 0.1
    }]
    };
    const config = {
        type: 'line',
        data: data,
        options: {
            borderRadius: '30',
            responsive: true,
            cutout: '95%',
            spacing: 2,
            plugins: {
                legend: {
                    position: 'bottom',
                    display: true,
                    labels: {
                        usePointStyle: true,
                        padding: 20,
                        font: {
                            size: 14
                        }
                    }
                },
                title: {
                    display: true,
                    text: 'My Personal Portfolio'
                }
            }
        }
    };
      onMount(()=> {
        const ctx = wrokload.getContext('2d');
        // Initialize chart using default config set
        var myChart = new Chart(ctx, config);
      })
</script>

<style >
    h1{
        font-weight: lighter;
    }
    .card{
        background-color: #F9FAFD;
        padding: 16px;
        border-radius: 16px;
    }
    .card h3{
        margin-top: 0;
    }

    table{
        width: 100%;
        border-collapse: collapse;

        font-size: 14px;
    }
    th{
        text-align: left;
        border-bottom: 1px solid rgba(0,0,0,0.5);
    }
    th.icon{
        width: 48px;
    }
    td{
        padding: 0;
        border-bottom: 1px solid rgba(0,0,0,0.25);
    }
    tr:last-child td{
        border-bottom: none;
    }
    tr{
        white-space: nowrap;
    }

</style>
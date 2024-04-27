class UI {
    _refreshInterval
    _lights = 0;

    constructor() {
        this.setRefreshRate(3000)
        this.demoChart()
    }

    renderLights(num) {
        let lights = document.getElementsByTagName("start-lights")[0];
        lights.setAttribute("lights", num);
    }

    setRefreshRate(milliseconds) {
        setInterval(this._refreshInterval);
        setInterval(() => this.render(), milliseconds);
    }

    render() {
        console.log("render")
    }

    demoChart() {
        const ctx = document.getElementById('myChart');

        new Chart(ctx, {
            type: 'line',
            data: {
                datasets: [{
                    data: [12, 19, 3, 5, 2, 3],
                    borderWidth: 1
                }]
            },
            options: {
                scales: {
                    y: {
                        beginAtZero: true
                    }
                }
            }
        });
    }
}
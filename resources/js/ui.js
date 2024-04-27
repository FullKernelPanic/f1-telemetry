
class UI {
    _refreshInterval

    constructor() {
        this.setRefreshRate(3000)
        this.demoChart()
    }

    renderLights(num) {
        let lamps = document.querySelectorAll("#lamp-container circle");

        for (let i = 0; i < 5; i++) {
            if (i <= num - 1) {
                lamps[i].classList.remove("off");
                lamps[i].classList.add("on");
            } else {
                lamps[i].classList.remove("on");
                lamps[i].classList.add("off");
            }
        }
    }

    setRefreshRate(milliseconds){
        setInterval(this._refreshInterval);
        setInterval(() => this.render(), milliseconds);
    }

    render(){
        console.log("render")
    }

    demoChart() {
        const ctx = document.getElementById('myChart');

        new Chart(ctx, {
            type: 'bar',
            data: {
                labels: ['Red', 'Blue', 'Yellow', 'Green', 'Purple', 'Orange'],
                datasets: [{
                    label: '# of Votes',
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
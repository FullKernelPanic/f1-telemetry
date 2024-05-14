class UI {
    _refreshInterval
    /** @property {Chart} _throttleChart */
    _throttleChart

    /** @property {Number} _lastFrame */
    _lastFrame;

    constructor() {
        this.setRefreshRate(3000);
        this.demoChart();
    }

    renderLights(num) {
        let lights = document.getElementsByTagName("start-lights")[0];
        lights.setAttribute("lights", num);
    }

    updateSession(data) {
        document.getElementById("SummaryTrackName").innerHTML = data.TrackInfo.name;
        document.getElementById("SummaryFormula").innerHTML = data.formula;
        document.getElementById("SummarySessionType").innerHTML = data.sessionType;
        document.getElementById("SummaryTotalLaps").innerHTML = data.TrackInfo.totalLaps;
        document.getElementById("SummaryTemperature").innerHTML = data.environment.airTemperature + ' / ' + data.environment.trackTemperature;
        document.getElementById("SummaryWeather").innerHTML = data.environment.weather;
    }

    setRefreshRate(milliseconds) {
        setInterval(this._refreshInterval);
        setInterval(() => this.render(), milliseconds);
    }

    updateTelemetry(data) {
        this.addData(this._throttleChart, data.frameId, "throttle", data.telemetries[19].throttle)
        this.addData(this._throttleChart, data.frameId, "brake", data.telemetries[19].brake)
    }

    render() {
        console.log("render");
        this._throttleChart.update();
    }

    demoChart() {
        const ctx = document.getElementById('myChart');

        this._throttleChart = new Chart(ctx, {
            type: 'line',
            data: {
                labels: [],
                datasets: [
                    {
                        label: "throttle",
                        spanGaps: true,
                        data: [],
                        borderWidth: 1
                    },
                    {
                        label: "brake",
                        spanGaps: true,
                        data: [],
                        borderWidth: 1
                    }
                ]
            },
            options: {
                normalized: true,
                maxRotation: 0,
                minRotation: 0,
                sampleSize: 1,
                animation: false,
                bezierCurve: false,
                spanGaps: true,
                scales: {
                    y: {
                        type: 'linear',
                        min: 0,
                        max: 1
                    }
                },
                elements: {
                    point: {
                        radius: 0 // default to disabled in all datasets
                    }
                }
            }
        });
    }

    addData(chart, frame, datasetLabel, newData) {
        if (this._lastFrame !== frame) {
            this._lastFrame = frame;
            chart.data.labels.push(frame);
        }

        for (let i = 0; i < chart.data.datasets.length; i++) {
            if (chart.data.datasets[i].label === datasetLabel) {
                chart.data.datasets[i].data.push(newData);
            }
        }
    }
}
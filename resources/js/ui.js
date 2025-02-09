class UI {
    _refreshInterval
    /** @property {Chart} _throttleChart */
    _throttleChart

    /** @property {Number} _lastFrame */
    _lastFrame;

    _chart;

    /**
     * @property {State} _state
     * @private
     */
    _state;

    constructor(state) {
        this._state = state;
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
        // this.addData(this._throttleChart, data.frameId, "throttle", data.telemetries[19].throttle)
        // this.addData(this._throttleChart, data.frameId, "brake", data.telemetries[19].brake)
    }

    updateLapData(data) {
        //this._chart.series
    }

    render() {
        // this._throttleChart.update();

        this._chart.setOption({
            series: this.generateSeries(this._state.driverPositions())
        });

        this._chart.resize();
    }

    generateSeries() {
        const drivers = this._state.drivers().all();

        return this._state.driverPositions().all().map(function (positions, index) {
            if (drivers[index]) {
                let extendedPositions = positions.map((item, index) => [item, 'foo']);

                return {
                    name: drivers[index].name,
                    symbolSize: 28,
                    symbol: tyreSoftImg,
                    smooth: true,
                    type: 'line',
                    color: teamColors[drivers[index].teamId],
                    emphasis: {
                        focus: 'series'
                    },
                    endLabel: {
                        show: true,
                        formatter: '{a}',
                        distance: 20
                    },
                    lineStyle: {
                        width: 4,
                    },
                    data: positions
                };
            }
            return {};
        });
    }

    demoChart() {
        this._chart = echarts.init(document.getElementById('main'));

        // Specify the configuration items and data for the chart
        var option = {
            title: {
                text: 'Positions'
            },
            tooltip: {
                trigger: 'item'
            },
            grid: {
                left: 30,
                right: 110,
                bottom: 30,
                containLabel: true
            },
            xAxis: {
                type: 'category',
                splitLine: {
                    show: true
                },
                axisLabel: {
                    margin: 30,
                    fontSize: 16
                },
                boundaryGap: false,
                min: 0,
                max: 15,
            },
            yAxis: {
                type: 'value',
                axisLabel: {
                    margin: 30,
                    fontSize: 16,
                    formatter: '#{value}'
                },
                inverse: true,
                interval: 1,
                min: 1,
                max: 22
            }
        };

        // Display the chart using the configuration items and data just specified.
        this._chart.setOption(option);

    }

    addData(chart, frame, datasetLabel, newData) {

    }
}

function callback(){
    console.log('callback');
    console.log(arguments);
}

const teamColors = [
    '#6CD3BF', //Mercedes-AMG Petronas Formula One Team
    '#E8002D', //Scuderia Ferrari
    '#3671C6', //Red Bull Racing
    '#005AFF', //Williams Racing
    '#006F62', //Aston Martin Aramco Cognizant Formula One Team
    '#0090FF', //BWT Alpine F1 Team
    '#2B4562', //Visa Cash App RB Formula 1 Team
    '#B6BABD', //MoneyGram Haas F1 Team
    '#FF8000', //McLaren F1 Team
    '#900000', //Alfa Romeo
];

const tyreSoftImg = 'image://data:image/png;base64,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';
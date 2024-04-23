import { Chart } from "chart.js";

function NewSpeedChart(ctx: string): SpeedChart {
    return new SpeedChart(ctx)
}

class SpeedChart {
    private chart: Chart
    private speedByFrame: number[][]
    private lastFrameId: number

    constructor(ctx: string) {
        this.lastFrameId = 0
        this.speedByFrame = []
        this.chart = new Chart(ctx, {
            type: "line",
            data: {
                labels: [],
                datasets: [
                    //parsing: false,
                    //normalized: true
                ]
            },
            options: {
                plugins: {
                    tooltip: {
                        enabled: false
                    }
                },
                animation: false,
                elements: {
                    point: {
                        radius: 0
                    }
                },
                spanGaps: true,
                scales: {
                    x: {
                        type: "linear",
                        min: 0,
                        max: 60 * 180
                    },
                    y: {
                        type: "linear",
                        min: 0,
                        max: 400
                    }
                }
            }
        })
    }

    public clearParticipant(participantIndex: number) {
        if (this.chart.data.datasets[participantIndex]) {
            this.chart.data.datasets[participantIndex].data = []
        }
    }

    public pushData(frameId: number, speeds: number[]): void {
        this.speedByFrame.push(speeds)

        this.onRefresh(this.chart)
    }

    private onRefresh(chart: Chart) {
        for (var i: number = 0; i < this.speedByFrame.length; i++) {
            chart.data.labels?.push(i)
            for (var j: number = 0; j < this.speedByFrame[i].length; j++) {
                if (!chart.data.datasets[j]) {
                    chart.data.datasets[j] = {
                        type: "line",
                        label: 'Driver #' + j,
                        backgroundColor: 'rgb(255, 99, 132)',
                        borderColor: 'rgb(255, 159, 64)',
                        data: []
                    }
                }

                chart.data.datasets[j].data.push({
                    y: this.speedByFrame[i][j],
                    x: i
                })
            }
        }

        chart.update()
    }
}

export {
    NewSpeedChart,
    SpeedChart
}
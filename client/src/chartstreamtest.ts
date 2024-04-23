import { Chart, registerables as chartRegistables } from "chart.js";
import "chartjs-plugin-streaming";
import registerables from "chartjs-plugin-streaming";
import 'chartjs-adapter-moment';

var chart:Chart;

function createStreamChart(id:string): void{
    Chart.register(...registerables);
    Chart.register(...chartRegistables)

    chart = new Chart(id, {
        type: "line",
        data: {
            labels: ["foo", "bar"],
            datasets: [
                {
					label: "My First dataset",
					data: [
						1,2,
					],
				},
                {
					label: "My Second dataset",
					data: [
						1,2,
					],
				}
            ]
        },
        options: {
            scales: {
              x: {
                type: "realtime",
                realtime: {
                duration: 20000,
                refresh: 1000,
                delay: 2000,
                  onRefresh: function(chart) {
                    chart.data.datasets.forEach(function(dataset) {
                      dataset.data.push({
                        x: Date.now(),
                        y: Math.random()
                      });
                    });
                  }
                }
              }
            },
            interaction: {
              intersect: false
            }
          }
    })
}

export {
    createStreamChart
};

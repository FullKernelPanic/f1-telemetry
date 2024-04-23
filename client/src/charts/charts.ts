import { Chart, registerables as chartRegistables } from "chart.js";
import registerables from "chartjs-plugin-streaming";

export default function() {
    Chart.register(...chartRegistables)
    Chart.register(...registerables)
}
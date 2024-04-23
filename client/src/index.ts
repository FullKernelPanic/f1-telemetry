import { EventDispatcher, EventInterface } from 'EventDispatcher';
import { DataHandlerConnector, LapDataEvent, TelemetryEvent } from './events';
import charts from './charts/charts';
import 'chartjs-adapter-moment';
import { NewWebsocket } from './socket';
import { DataStorage, LapFinishedEvent, TelemetryInput } from './dataHandler';
import { NewSpeedChart, SpeedChart } from './charts/speedchart';

window.onload = (e) => {
    main()
}

var dataStorage: DataStorage

export function main(): void {
    eventDispatcher = new EventDispatcher()
    dataStorage = new DataStorage(eventDispatcher)

    var websocket: WebSocket = NewWebsocket(
        "ws://localhost:8080/socket",
        new DataHandlerConnector(eventDispatcher)
    )

    initCharts()

    eventDispatcher.on("telemetry", onTelemetry);
    eventDispatcher.on("lapData", onLapData);
    eventDispatcher.on("lapFinished", onLapFinished);
}

var eventDispatcher: EventDispatcher

var speedChart: SpeedChart

function initCharts() {
    charts()
    speedChart = NewSpeedChart("streamChart")
}

function onLapData(event: EventInterface): void {
    var lapDataEvent: LapDataEvent = event as LapDataEvent

    dataStorage.lapData.load(lapDataEvent.data)
}

function onTelemetry(event: EventInterface): void {
    var telemetryEvent: TelemetryEvent = event as TelemetryEvent

    var speeds: number[] = telemetryEvent.data.telemetries.map((t: TelemetryInput): number => {
        return t.speed
    })

    speedChart.pushData(
        telemetryEvent.data.frameId,
        speeds
    )
}

function onLapFinished(event: EventInterface): void {
    var lapFinishedEvent:LapFinishedEvent = event as LapFinishedEvent
    
    speedChart.clearParticipant(lapFinishedEvent.participantIndex)
}
import { EventDispatcher, EventInterface } from "EventDispatcher";
import { CarTelemetries, LapDatas } from "./dataHandler";
import { Listener } from "./socket";

class DataHandlerConnector implements Listener {
    private dispatcher: EventDispatcher

    constructor(dispatcher: EventDispatcher) {
        this.dispatcher = dispatcher;
    }

    public onData(data: { type: string, data: Object }): void {
        switch (data.type){
        case "telemetry":
            this.dispatcher.trigger(new TelemetryEvent(data.data as CarTelemetries))
            break;
        case "lapData":
            this.dispatcher.trigger(new LapDataEvent(data.data as LapDatas))
            break;
        }
    }
}

class TelemetryEvent implements EventInterface {
    public type: string = "telemetry"
    public data: CarTelemetries

    constructor(data: CarTelemetries) {
        this.data = data
    }
}

class LapDataEvent implements EventInterface {
    public type: string = "lapData"
    public data: LapDatas

    constructor(data: LapDatas) {
        this.data = data
    }
}

export {
    DataHandlerConnector,
    TelemetryEvent,
    LapDataEvent
}
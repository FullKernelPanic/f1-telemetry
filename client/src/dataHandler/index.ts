import { EventDispatcher, EventInterface } from "EventDispatcher"

class DataStorage {
    public session: SessionHandler
    public telemetry: TelemetryHandler
    public lapData: LapDataHandler
    public dispatcher: EventDispatcher

    constructor(dispatcher: EventDispatcher) {
        this.dispatcher = dispatcher
        this.session = new SessionHandler()
        this.telemetry = new TelemetryHandler()
        this.lapData = new LapDataHandler(this.dispatcher)
    }

    public OnData(data: { type: string, data: Object }): void {
        switch (data.type) {
            case "session":
                this.session.load(data.data as SessionInput)
                break;
            case "telemetry":
                this.telemetry.load(data.data as CarTelemetries)
        }
    }
}

class LapDataHandler {
    private currentLaps: number[]
    private dispatcher: EventDispatcher

    constructor(dispatcher: EventDispatcher) {
        this.currentLaps = []
        this.dispatcher = dispatcher
    }

    public load(lapDatas: LapDatas) {
        for (var i: number = 0; i < lapDatas.lapDatas.length; i++) {
            if (this.currentLaps[i] != lapDatas.lapDatas[i].currentLapNum) {
                this.currentLaps[i] = lapDatas.lapDatas[i].currentLapNum

                this.dispatcher.trigger(new LapFinishedEvent(i))
            }
        }
    }
}

class LapFinishedEvent implements EventInterface {
    public type: string = "lapFinished"
    public participantIndex: number

    constructor(participantIndex: number) {
        this.participantIndex = participantIndex
    }
}

class TelemetryHandler {
    private speed: number[][]

    constructor() {
        this.speed = []
        this.speed = this.speed.fill([], 0, 21)
    }

    public load(input: CarTelemetries) {
        for (var i = 0; i < input.telemetries.length; i++) {
            !this.speed[i] ? this.speed[i] = [] : null
            this.speed[i].push(input.telemetries[i].speed)
        }
    }

    public speedByCars(): number[][] {
        return this.speed
    }
}

class SessionHandler {
    public AIDifficulty: number = 0
    public formula: string = "unknown"
    public isGamePaused: boolean = true
    public sessionDuration: number = 0
    public sessionTimeLeft: number = 0
    public sessionType: string = "unkown"

    public load(input: SessionInput) {
        this.AIDifficulty = input.AIDifficulty
        this.formula = input.formula
        this.isGamePaused = input.isGamePaused
        this.sessionDuration = input.sessionDuration
        this.sessionTimeLeft = input.sessionTimeLeft
        this.sessionType = input.sessionType
    }
}

interface SessionInput {
    AIDifficulty: number
    assist: AssistInput
    environment: Object
    formula: string
    isGamePaused: boolean
    sessionDuration: number
    sessionTimeLeft: number
    sessionType: string
};

interface AssistInput {
    anti_lock_brakes: boolean
    braking_assist: string
    drs_assist: boolean
    dynamic_racing_line: number
    dynamic_racing_line_type: number
    ers_assist: boolean
    gearbox_assist: number
    pit_assist: boolean
    pit_release_assist: boolean
    steering_assist: boolean
    traction_control: number
}

interface TelemetryInput {
    speed: number
    throttle: number
    steer: number
    brake: number
    brakesTemperature: number[]
    clutch: number
    drs: boolean
    engineRPM: number
    engineTemperature: number
    gear: number
    revLightsBitValue: number
    revLightsPercent: number
}

interface CarTelemetries {
    frameId: number
    telemetries: TelemetryInput[]
}

interface LapDatas {
    frameId: number
    lapDatas: LapData[]
}

interface LapData {
    carPosition: number
    currentLapInvalid: boolean
    currentLapNum: number
    currentLapTimeInMS: number
    driverStatus: number
    gridPosition: number
    lapDistance: number
    lastLapTimeInMS: number
    numPitStops: number
    numUnservedDriveThroughPens: number
    numUnservedStopGoPens: number
    penalties: number
    pitLaneTimeInLaneInMS: number
    pitLaneTimerActive: number
    pitStatus: number
    pitStopShouldServePen: number
    pitStopTimerInMS: number
    resultStatus: number
    safetyCarDelta: number
    sector: number
    sector1TimeInMS: number
    sector2TimeInMS: number
    totalDistance: number
    warnings: number
}

export {
    DataStorage,
    CarTelemetries,
    TelemetryInput,
    LapDatas,
    LapFinishedEvent
}
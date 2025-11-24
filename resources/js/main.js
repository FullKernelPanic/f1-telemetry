function main() {
    new App(document.querySelector('meta[name="websocket"]').content)
}

class App {
    /** @property {WebSocketHandler} _websocket */
    _websocket;
    /** @property {UI} _ui */
    _ui;
    /** @property {State} _state */
    _state;

    /**
     * @param {string} webSocketUrl
     */
    constructor(webSocketUrl) {
        this._state = new State();

        this._ui = new UI(this._state);

        this._websocket = new WebSocketHandler(webSocketUrl, this._ui);
        this._websocket.addEventListener("button", this.onButton.bind(this));
        this._websocket.addEventListener("startLight", this.onStartLight.bind(this));
        this._websocket.addEventListener("session", this.onSession.bind(this));
        this._websocket.addEventListener("carDamage", this.onCarDamage.bind(this));
        this._websocket.addEventListener("telemetry", this.onTelemetry.bind(this));
        this._websocket.addEventListener("participants", this.onParticipants.bind(this));
        this._websocket.addEventListener("lapData", this.onLapData.bind(this));
        this._websocket.addEventListener("sessionHistory", this.onSessionHistory.bind(this));

        window.app = this;
    }

    /**
     * @param {TelemetryEvent} event
     */
    onButton(event) {
    }

    /**
     * @param {TelemetryEvent} event
     */
    onLapData(event) {
        this._state.driverPositions().update(event.data);
        //this._ui.updateLapData(event.data);
    }

    /**
     * @param {TelemetryEvent} event
     */
    onParticipants(event) {
        this._state.drivers().update(event.data.drivers);
    }

    /**
     * @param {TelemetryEvent} event
     */
    onTelemetry(event) {

        this._ui.updateTelemetry(event.data);
    }

    /**
     * @param {TelemetryEvent} event
     */
    onCarDamage(event) {
    }

    /**
     * @param {TelemetryEvent} event
     */
    onStartLight(event) {
        this._ui.renderLights(event.data.NumberOfLights);
    }

    /**
     * @param {TelemetryEvent} event
     */
    onSession(event) {
        this._ui.updateSession(event.data);
    }

    /**
     * @param {TelemetryEvent} event
     */
    onSessionHistory(event) {
        //console.log(event.data);
    }
}

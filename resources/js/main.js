function main() {
    new App(document.querySelector('meta[name="websocket"]').content)
}

class App {
    /** @property {WebSocketHandler} _websocket */
    _websocket;
    /** @property {UI} _ui */
    _ui;

    /**
     * @param {string} webSocketUrl
     */
    constructor(webSocketUrl) {
        this._ui = new UI();
        this._websocket = new WebSocketHandler(webSocketUrl, this._ui);
        this._websocket.addEventListener("button", this.onButton.bind(this));
        this._websocket.addEventListener("startLight", this.onStartLight.bind(this));
        this._websocket.addEventListener("session", this.onSession.bind(this));
        this._websocket.addEventListener("carDamage", this.onCarDamage.bind(this));
    }

    /**
     * @param {TelemetryEvent} event
     */
    onButton(event) {
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

    onSession(event) {
        this._ui.updateSession(event.data);
    }
}

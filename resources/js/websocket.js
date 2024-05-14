class TelemetryEvent extends Event {
    /** @property {Object} data */
    data;

    constructor(type, data) {
        super(type);
        this.data = data;
    }
}

class WebSocketHandler extends EventTarget {
    static buttonEvent = new TelemetryEvent('button');
    static sessionEvent = new TelemetryEvent('session');
    static carDamageEvent = new TelemetryEvent('carDamage');
    static lapDataEvent = new TelemetryEvent('lapData');
    static sessionHistoryEvent = new TelemetryEvent('sessionHistory');
    static motionEvent = new TelemetryEvent('motion');
    static telemetryEvent = new TelemetryEvent('telemetry');
    static participantsEvent = new TelemetryEvent('participants');
    static carSetupEvent = new TelemetryEvent('carSetup');
    static startLightEvent = new TelemetryEvent('startLight');

    /**
     * @param {string} url
     * @param {UI} ui
     */
    constructor(url, ui) {
        super();

        this.ui = ui
        this.ws = new WebSocket(url)
        this.ws.addEventListener("open", this.openHandler.bind(this));
        this.ws.addEventListener("close", this.closeHandler.bind(this));
        this.ws.addEventListener("error", this.errorHandler.bind(this));
        this.ws.addEventListener("message", this.messageHandler.bind(this));
    }

    openHandler() {
        console.log("open!");
    }

    /**
     * @param {CloseEvent} data
     */
    closeHandler(data) {
        console.log("close!");
    }

    errorHandler() {
        console.log("error!");
    }

    /**
     * @param {MessageEvent} messageEvent
     */
    messageHandler(messageEvent) {
        const data = JSON.parse(messageEvent.data);

        if (!data.type) {
            return;
        }

        this.dispatchEvent(new TelemetryEvent(data.type, data.data));
    }
}

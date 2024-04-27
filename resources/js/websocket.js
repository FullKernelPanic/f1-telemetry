class WebSocketHandler {
    /**
     * @param {string} url
     * @param {UI} ui
     */
    constructor(url, ui) {
        this.ui = ui
        this.ws = new WebSocket(url)
        this.ws.addEventListener("open", this.openHandler.bind(this))
        this.ws.addEventListener("close", this.closeHandler.bind(this))
        this.ws.addEventListener("error", this.errorHandler.bind(this))
        this.ws.addEventListener("message", this.messageHandler.bind(this))
    }

    /**
     * @param {Event} data
     */
    openHandler(data) {
        console.log("open!");
        console.log(data);
    }

    /**
     * @param {CloseEvent} data
     */
    closeHandler(data) {
        console.log("close!");
        console.log(data);
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

        switch (data.type) {
            case "button":
                break;
            case "carDamage":
                break;
            case "sessionHistory":
                break;
            case "motion":
                break;
            case "telemetry":
                break;
            case "participants":
                break;
            case "lapData":
                this.handleLapData(data);
                break;
            case "session":
                this.handleSession(data)
                break;
            case "carSetup":
                this.handleCarSetup(data)
                break;
            case "startLight":
                this.handleStartLight(data)
                break;
            default:
                console.log(data.type);
        }
    }

    handleSession(data) {
    }

    handleCarSetup(data) {
    }

    handleLapData(data) {
    }

    handleStartLight(data) {
        this.ui.renderLights(data.data.NumberOfLights);
    }
}


function main() {
    new App(document.querySelector('meta[name="websocket"]').content)
}

class App {
    _websocket
    _ui

    /**
     * @param {string} webSocketUrl
     */
    constructor(webSocketUrl) {
        this._ui = new UI();
        this._websocket = new WebSocketHandler(webSocketUrl, this._ui);
    }
}

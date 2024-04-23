
interface Listener {
    onData(data:{type:string,data:Object}): void
}

function NewWebsocket(socketUrl: string, listener:Listener): WebSocket {
    const socket = new WebSocket(socketUrl)

    var callback = (e: Event) => console.log("Event: " + e.type, e);
    
    socket.onclose = callback;
    socket.onerror = callback;
    socket.onopen = callback;

    socket.onmessage = (e: MessageEvent) => {
        let data = JSON.parse(e.data)
        if (!data) {
            console.error("Unproccessable data")
            return;
        } 

        listener.onData(data)
    };

    return socket
}

export {
    NewWebsocket,
    Listener
}
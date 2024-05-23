var socket = new WebSocket("ws://localhost:8080/ws");

let connect = cb => {
    console.log("connecting");

    socket.onopen = () => {
        console.log("Successfully Connected");
    };

    socket.onmessage = msg => {
        console.log(msg);
        cb(msg);
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };
}

let sendMsg = msg => {
    if (socket.readyState === WebSocket.OPEN){
        console.log(msg)
        socket.send(msg);
    } else {
        console.log('The socket is not open.');
    }
}

export { connect, sendMsg };
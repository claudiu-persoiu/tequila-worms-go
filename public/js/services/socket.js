wormApp.factory('socketService', ["$window", function (window) {

    const url = "ws://127.0.0.1:3000/ws"

    let socket, socketOpened = false;

    const connect = function () {
        socket = new window.WebSocket(url);

        socket.onopen = function () {
            socketOpened = true;
            // errorBlock.style.display = 'none';
        };

        socket.onmessage = function (event) {

            const message = JSON.parse(event.data);
            if (listeners[message.action]) {
                listeners[message.action].forEach(l => l(message.value))
            //     for (let l of listeners[message.action]) {
            //         l(message.value)
            //     }
            }
        };

        socket.onclose = function () {
            socketOpened = false;
            // errorBlock.style.display = 'block';
        };
    };

    connect();

    const listeners = {};

    return {
        on: function (action, callback) {
            if (!listeners[action]) {
                listeners[action] = [];
            }
            listeners[action].push(callback)
            console.log("on", action, callback)
        },
        emit: function (action, value) {
            if (socketOpened) {
                return socket.send(JSON.stringify({
                    "action": action,
                    "value": value,
                }));
            }
            console.log("emit", action, value)
        }
    }
}]);

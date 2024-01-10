wormApp.factory('socketService', ["socketFactory", "$window", function (socketFactory, window) {
    const socket = function () {
        const url = `${location.protocol === 'https:' ? 'wss' : 'ws'}://${location.host + location.pathname}ws`

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
                }

                if (listenersOnce[message.action]) {
                    listenersOnce[message.action].forEach(l => l(message.value))
                    listenersOnce[message.action] = []
                }
            };

            socket.onclose = function () {
                socketOpened = false;
                // errorBlock.style.display = 'block';
            };
        };

        connect();

        let listeners = {};
        let listenersOnce = {};

        return {
            on: function (action, callback) {
                if (!listeners[action]) {
                    listeners[action] = [];
                }
                listeners[action].push(callback)
            },
            emit: function (action, value) {
                if (socketOpened) {
                    return socket.send(JSON.stringify({
                        "action": action,
                        "value": value,
                    }));
                }
                console.log("emit", action, value)
            },
            once: function (action, callback) {
                if (!listenersOnce[action]) {
                    listenersOnce[action] = [];
                }
                listenersOnce[action].push(callback)
            },
            removeListener: function () {},
            removeAllListeners: function () {
                listeners = {}
                listenersOnce = {}
            },
            connect: function () {},
            disconnect: function () {},
        }
    }
    const options = {
        ioSocket: socket()
    }

    return socketFactory(options)
}]);

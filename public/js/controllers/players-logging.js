wormApp.controller('PlayersLoggingController', ['$scope', 'socketService', function ($scope, socketService) {

    $scope.messages = [];
    const max = 3;

    const addMessage = function (message) {
        $scope.messages.unshift(message);

        if ($scope.messages.length >= max) {
            $scope.messages = $scope.messages.slice(0, max);
        }
    };

    socketService.on('dead worm', function (message) {
        addMessage(message);
    });

    setInterval(function () {
        $scope.messages.pop();
    }, 10000);

}]);

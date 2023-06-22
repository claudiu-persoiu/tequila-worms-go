wormApp.controller('CanvasController', ['$scope', 'socketService', function ($scope, socketService) {

    socketService.on('matrix size', function (value) {
        $scope.matrixSize = value;
    });

    socketService.on('worm data', function (value) {
        $scope.wormData = value;
    });

    $scope.swipeUp = function () {
        socketService.emit('new direction', 'down');
    };

    $scope.swipeDown = function () {
        socketService.emit('new direction', 'up');
    };

    $scope.swipeLeft = function () {
        socketService.emit('new direction', 'left');
    };

    $scope.swipeRight = function () {
        socketService.emit('new direction', 'right');
    };

}]);
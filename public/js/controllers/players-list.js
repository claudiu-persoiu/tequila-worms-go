wormApp.controller('PlayersListController', ['$scope', 'socketService', function ($scope, socketService) {

    socketService.on('player list', function (worms) {
        $scope.worms = worms;
    });

}]);
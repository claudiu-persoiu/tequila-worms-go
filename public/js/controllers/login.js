
wormApp.controller('LoginController', ['$scope', 'socketService', '$uibModal', function ($scope, socketService, $uibModal) {

    var setName = function (name) {
        if (name && name.length) {
            $scope.name = name;
            socketService.emit('start game', $scope.name);
            document.addEventListener('keydown', keyDirectionUpdate, false);
        }
        $scope.modalOpened = false;
    };

    $scope.modalOpen = false;

    var openModal = function () {

        if ($scope.modalOpened) {
            return;
        }

        $scope.modalOpened = true;

        var modalInstance = $uibModal.open({
            animation: true,
            templateUrl: 'templates/login-modal.html',
            controller: 'LoginModalController',
            backdrop: 'static',
            keyboard: false,
            resolve: {
                user: {
                    name: $scope.name
                }
            }
        });

        modalInstance.result.then(setName);
    };

    if (!$scope.name) {
        openModal();
    }

    var stopGame = function () {
        document.removeEventListener('keydown', keyDirectionUpdate);
        openModal();
    };

    socketService.on('disconnect', stopGame);
    socketService.on('you dead', stopGame);

    var directionMap = {
        '40': 'up',
        '38': 'down',
        '37': 'left',
        '39': 'right'
    };

    var keyDirectionUpdate = function (e) {
        var direction = directionMap[e.keyCode];
        if (direction) {
            socketService.emit('new direction', directionMap[e.keyCode]);
            e.preventDefault();
        }
    };

}]);
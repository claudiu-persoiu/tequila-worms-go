wormApp.controller('LoginModalController', ['$scope', '$uibModalInstance', 'user', function ($scope, $uibModalInstance, user) {

    $scope.name = user.name;

    $scope.ok = function () {
        if ($scope.name.length) {
            $uibModalInstance.close($scope.name);
        }

    };

}]);
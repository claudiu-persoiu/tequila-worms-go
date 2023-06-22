
var wormApp = angular.module('worms', ['ng', 'ui.bootstrap', 'swipe']);

wormApp.directive('loginUser', [function () {
    return {
        controller: 'LoginController',
        templateUrl: 'templates/login.html'
    };
}]).directive('canvasContainer', ['drawingService', '$window', function (drawingService, $window) {

    var link = function (scope, element, attrs) {

        var headColor = '#CCCCCC';

        drawingService.setCanvas(element.find('canvas')[0])
            .calculateProportions();

        var drawWorm = function (worm) {
            drawingService.setPieceColor(headColor);
            worm.pieces.forEach(function (piece) {
                drawingService.drawPiece(piece);
                drawingService.setPieceColor(worm.color);
            });
        };

        var drawMatrix = function (worms) {
            drawingService.clearCanvas();
            worms.forEach(drawWorm);
        };

        angular.element($window).bind('resize', function() {
            drawingService.calculateProportions()
                .calculateElementSize();
        });

        scope.$watch('matrixSize', function(matrixSize) {
            if (matrixSize) {
                drawingService.setMatrixSize(matrixSize)
                    .calculateElementSize();
            }
        });

        scope.$watch('wormData', function (worms) {
            if (worms) {
                drawMatrix(worms);
            }
        });
    };

    return {
        controller: 'CanvasController',
        templateUrl: 'templates/canvas.html',
        link: link
    };
}]).directive('playersList', [function () {
    return {
        controller: 'PlayersListController',
        templateUrl: 'templates/players-list.html'
    };
}]).directive('playersLogging', [function () {
    return {
        controller: 'PlayersLoggingController',
        templateUrl: 'templates/players-logging.html'
    };
}]);

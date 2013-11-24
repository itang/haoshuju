/**
 * app.js
 */

var api = {};
api.getAppInfo = function(http, call){
  http.get("/appinfo").success(function(app){
    call(app);
    console.log(app);
  });
}

function AppInfoCtrl($scope,$http) {
  api.getAppInfo($http, function(app){
    $scope.app = app;
    $scope.restApis = app.restApis;
  });
}

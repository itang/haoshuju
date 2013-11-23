/**
 * app.js
 */

var api = {};
api.getAppInfo = function(http, call){
  http.get("/appinfo").success(function(appInfo){
    call(appInfo);
    console.log(appInfo);
  });
}

function AppInfoCtrl($scope,$http) {
  api.getAppInfo($http, function(appInfo){
    $scope.app = appInfo.app;
    $scope.restApis = appInfo.restApis;
  });
}

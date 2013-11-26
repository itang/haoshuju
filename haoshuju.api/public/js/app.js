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

api.getClientApps = function(http, call){
  http.get("/clientapps").success(function(ret){
    call(ret);
    console.log(ret);
  });
}

function AppInfoCtrl($scope,$http) {
  api.getAppInfo($http, function(app){
    $scope.app = app;
    $scope.restApis = app.restApis;
  });

  api.getClientApps($http, function(apps){
    $scope.clientApps = apps;
  });
}

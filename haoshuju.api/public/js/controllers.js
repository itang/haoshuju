'use strict';

/* Controllers */

var haoshujuControllers = angular.module('haoshujuControllers', []);

haoshujuControllers.controller('AppInfoCtrl', ['$scope', '$http',
  function AppInfoCtrl($scope, $http) {
    api.getAppInfo($http, function(app){
      $scope.app = app;
      $scope.restApis = app.restApis;
    });

    api.getClientApps($http, function(apps){
      $scope.clientApps = apps;
    });
  }]);

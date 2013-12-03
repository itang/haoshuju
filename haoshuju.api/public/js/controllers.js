'use strict';

/* Controllers */

var haoshujuControllers = angular.module('haoshujuControllers', []);

haoshujuControllers.controller('AppInfoCtrl', ['$scope', '$http',
  function AppInfoCtrl($scope, $http) {
    api.getAppInfo($http, function(app){
      $scope.app = app;
      $scope.modules = app.modules;
    });

    api.getClientApps($http, function(apps){
      $scope.clientApps = apps;

      for(var i =0; i<apps.length; i++){
        api.checkHostAlive(apps[i].hostname + ":" + apps[i].httpport, i, $http, function(i, ret){
          $scope.clientApps[i].alive = ret.alive;
        });
      }
    });
  }]);

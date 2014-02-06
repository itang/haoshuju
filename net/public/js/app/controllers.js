'use strict';

/* Controllers */

var haoshujuControllers = angular.module('haoshujuControllers', []);

haoshujuControllers.controller('VersionController', ['$scope', '$http', '$timeout',
  function VersionController($scope, $http, $timeout) {
    $http.get('/app/version').success(function(reps) {
      $scope.version = reps.data;
    });
    $scope.$watch("serverTime", function(newValue, oldValue){
      $timeout(function() {
        $http.get('/app/servertime').success(function(reps) {
          if(reps.code == 0){
            $scope.serverTime = reps.data.str;
          }else {
            $scope.serverTime = reps.data;
          }
        });
      }, 1000);
    });
  }]);

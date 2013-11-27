'use strict';

/* Controllers */

var haoshujuControllers = angular.module('haoshujuControllers', []);

haoshujuControllers.controller('VersionController', ['$scope', '$http', '$timeout',
  function VersionController($scope, $http, $timeout) {
    $http.get('/app/version').success(function(data) {
      $scope.version = data.Data;
    });
    $scope.$watch("serverTime", function(newValue, oldValue){
      $timeout(function() {
        $http.get('/app/servertime').success(function(data) {
          $scope.serverTime = data.Data;
        });
      }, 1000);
    });
  }]);

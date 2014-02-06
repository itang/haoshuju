'use strict';

/* App Module */

var haoshujuApp = angular.module("haoshujuApp", [
  'haoshujuControllers',
  'haoshujuFilters',
  'haoshujuServices']);

haoshujuApp.config(["$interpolateProvider", function($interpolateProvider){
  $interpolateProvider.startSymbol("[[");
  $interpolateProvider.endSymbol("]]");
  }]);

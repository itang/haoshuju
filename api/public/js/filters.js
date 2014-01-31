'use strict';

/* Filters */

angular.module('haoshujuFilters', []).filter('statusName', function() {
  return function(input) {
    return input == 0 ? 'Valid' : 'Invalid';
  };
});

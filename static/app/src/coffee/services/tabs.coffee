angular.module('horodata').factory("tabsService", [
  "$rootScope"
  ($rootScope) ->

    current = null

    $rootScope.$on("$routeChangeStart", -> current = null)

    return {
      get: -> current
      set: (tab) -> current = tab
    }
])

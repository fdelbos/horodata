angular.module("horodata").directive("appMenuSidenav", [
  "$mdSidenav"
  "$http"
  "$location"
  "apiService"
  "$routeParams"
  "groupNewService"
  (
    $mdSidenav
    $http
    $location
    apiService
    $routeParams
    groupNewService
  )->

    l = (scope, elem) ->
      scope.toggleSidenav = -> $mdSidenav("sidenav").toggle()

      groupNewService.fetch()
      scope.groups = -> groupNewService.listing()

      scope.changeGroup = (url) ->
        scope.toggleSidenav()
        $location.path(url)

      scope.$on("$routeChangeSuccess", ->
        scope.currentGroupUrl = $routeParams.group
      )


    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/menu/sidenav.html"
    }
])

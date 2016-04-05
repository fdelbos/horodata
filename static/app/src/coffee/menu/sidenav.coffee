angular.module("horodata").directive("appMenuSidenav", [
  "$mdSidenav"
  "$http"
  "$location"
  "apiService"
  "$routeParams"
  "groupService"
  (
    $mdSidenav
    $http
    $location
    apiService
    $routeParams
    groupService
  )->

    l = (scope, elem) ->
      scope.toggleSidenav = -> $mdSidenav("sidenav").toggle()
      scope.closeSidenav = -> $mdSidenav("sidenav").close()

      groupService.fetch()
      scope.groups = -> groupService.listing()

      scope.changeGroup = (url) ->
        scope.closeSidenav()
        $location.path(url)

      scope.$on("$routeChangeSuccess", ->
        scope.currentGroupUrl = $routeParams.group
        scope.closeSidenav()
      )

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/menu/sidenav.html"
    }
])

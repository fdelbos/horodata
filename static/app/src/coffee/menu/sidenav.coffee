angular.module("horodata").directive("appMenuSidenav", [
  "$mdSidenav"
  "$http"
  "$location"
  "apiService"
  "$routeParams"
  (
    $mdSidenav
    $http
    $location
    apiService
    $routeParams
  )->

    l = (scope, elem) ->
      scope.toggleSidenav = -> $mdSidenav("sidenav").toggle()

      $http.get("#{apiService.get()}/groups").then(
        (resp) ->
          scope.groups = resp.data.data.results
      )

      #$routeParams.group
      scope.changeGroup = (url) -> $location.path(url)
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

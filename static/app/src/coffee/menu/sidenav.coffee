angular.module("horodata").directive("appMenuSidenav", [
  "$mdSidenav"
  "$http"
  "apiService"
  (
    $mdSidenav
    $http
    apiService
  )->

    l = (scope, elem) ->
      scope.toggleSidenav = -> $mdSidenav("sidenav").toggle()

      $http.get("#{apiService.get()}/groups").then(
        (resp) ->
          scope.groups = resp.data.data.results
      )

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/menu/sidenav.html"
    }
])

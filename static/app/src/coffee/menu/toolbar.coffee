angular.module("horodata").directive("appMenuToolbar", [
  "titleService"
  "userService"
  "homeService"
  (titleService, userService, homeService) ->

    l = (scope, elem) ->
      scope.MainTitle = titleService.get
      userService.get (u) -> scope.user = u
      scope.home = homeService.get()

      scope.openMenu = ($mdOpenMenu, $event) -> $mdOpenMenu($event)

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/menu/toolbar.html"
    }
])

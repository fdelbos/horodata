angular.module("horodata").directive("appMenuToolbar", [
  "titleService"
  "userService"
  "homeService"
  "$mdMedia"
  "$mdBottomSheet"
  (titleService, userService, homeService, $mdMedia, $mdBottomSheet) ->

    l = (scope, elem) ->
      scope.MainTitle = titleService.get
      userService.get (u) -> scope.user = u
      scope.home = homeService.get()

      scope.openMenu = ($mdOpenMenu, $event) ->
        if $mdMedia('xs') || $mdMedia('sm')
          $mdBottomSheet.show
            templateUrl: "horodata/menu/bottom_sheet.html"
        else $mdOpenMenu($event)

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/menu/toolbar.html"
    }
])

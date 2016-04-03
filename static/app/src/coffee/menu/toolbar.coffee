angular.module("horodata").directive("appMenuToolbar", [
  "titleService"
  "userService"
  "homeService"
  "popupService"
  "$mdMedia"
  "$mdBottomSheet"
  (titleService, userService, homeService, popupService, $mdMedia, $mdBottomSheet) ->

    l = (scope, elem) ->
      scope.MainTitle = titleService.get
      userService.get (u) -> scope.user = u
      scope.home = homeService.get()

      scope.openMenu = ($mdOpenMenu, $event) ->
        if $mdMedia('xs') || $mdMedia('sm')
          $mdBottomSheet.show
            templateUrl: "horodata/menu/bottom_sheet.html"
        else $mdOpenMenu($event)

      scope.showProfile = (ev) ->
        popupService("horodata/views/profile.html", "Profile", scope, ev)

      scope.showQuotas = (ev) ->
        popupService("horodata/views/quotas.html", "Quotas", scope, ev)

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/menu/toolbar.html"
    }
])

angular.module("horodata").directive("appMenuToolbar", [
  "titleService"
  "userService"
  "homeService"
  "popupService"
  "$mdMedia"
  "$mdBottomSheet"
  "$location"
  (titleService, userService, homeService, popupService, $mdMedia, $mdBottomSheet, $location) ->

    l = (scope, elem) ->
      scope.MainTitle = titleService.get
      userService.get (u) -> scope.user = u
      scope.home = homeService.get()

      scope.openMenu = ($mdOpenMenu, $event) ->
        if $mdMedia('xs') || $mdMedia('sm')
          $mdBottomSheet.show
            templateUrl: "horodata/menu/bottom_sheet.html"
        else $mdOpenMenu($event)

      scope.goToBilling = (ev) ->
        $mdBottomSheet.hide()
        $location.path("billing")

      scope.showProfile = (ev) ->
        $mdBottomSheet.hide()
        popupService("horodata/views/profile.html", "Profile", scope, ev)

      scope.showQuotas = (ev) ->
        $mdBottomSheet.hide()
        popupService("horodata/views/quotas.html", "Quotas", scope, ev)

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/menu/toolbar.html"
    }
])

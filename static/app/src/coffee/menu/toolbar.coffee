angular.module("horodata").directive("appMenuToolbar", [
  "titleService"
  (titleService) ->

    l = (scope, elem) ->
      scope.MainTitle = titleService.get

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/menu/toolbar.html"
    }
])

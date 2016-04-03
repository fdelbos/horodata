angular.module("horodata").directive("appWidgetsQuota", [
  ->

    l = (scope) ->

      genPercent = ->
        scope.percent = Math.floor(scope.current / scope.max * 100)

      scope.$watch("current", -> genPercent())
      scope.$watch("max", -> genPercent())

    return {
      link: l
      scope:
        label: "@"
        current: "="
        max: "="
      restrict: "E"
      templateUrl: "horodata/widgets/quota.html"
    }
])

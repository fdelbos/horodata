angular.module("horodata").directive("appWidgetsStatsContainer", [
  ->

    return {
      transclude: true
      scope:
        caption: "@"
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/stats/container.html"
    }
])

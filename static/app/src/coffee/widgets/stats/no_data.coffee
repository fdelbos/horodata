angular.module("horodata").directive("appWidgetsStatsNoData", [
  ->
    return {
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/stats/no_data.html"
    }
])

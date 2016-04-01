angular.module("horodata").directive("appWidgetsStatsTaskTime", [
  "statsService"
  (statsService)->

    l = (scope) ->
      scope.stats = statsService

      statsService.fetch(scope.group.url, "task_time", scope.search, (data) =>
        scope.data = data
      )

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/stats/task_time.html"
    }
])

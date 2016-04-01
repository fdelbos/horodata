angular.module("horodata").directive("appWidgetsStatsGuestTime", [
  "statsService"
  (statsService)->

    l = (scope) ->
      scope.stats = statsService

      statsService.fetch(scope.group.url, "guest_time", scope.search, (data) =>
        scope.data = data
      )

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/stats/guest_time.html"
    }
])

angular.module("horodata").directive("appWidgetsStatsCustomerTime", [
  "statsService"
  (statsService)->

    l = (scope) ->
      scope.stats = statsService

      statsService.fetch(scope.group.url, "customer_time", scope.search, (data) =>
        scope.data = data
      )

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/stats/customer_time.html"
    }
])

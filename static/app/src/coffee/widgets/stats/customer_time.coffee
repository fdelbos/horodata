angular.module("horodata").directive("appWidgetsStatsCustomerTime", [
  "statsService"
  "statsFilterService"
  (statsService, statsFilterService)->

    l = (scope) ->
      scope.stats = statsService
      scope.filter = statsFilterService

      update = ->
        statsService.fetch(scope.group.url, "customer_time", (data) =>
          scope.data = data)
      update()

      scope.$watch("filter", (v, o) ->
        if v.begin == o.begin and v.end == o.end then return
        update()
      , true)

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/stats/customer_time.html"
    }
])

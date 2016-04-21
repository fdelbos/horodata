angular.module("horodata").directive("appWidgetsStatsCustomer", [
  "statsService"
  "statsFilterService"
  (statsService, statsFilterService)->

    l = (scope) ->
      scope.stats = statsService
      scope.filter = statsFilterService

      updateTime = ->
        scope.time = null
        statsService.fetch(scope.group.url, "customer_time", (data) -> scope.time = data)
      updateTime()

      updateCost = ->
        scope.cost = null
        statsService.fetch(scope.group.url, "customer_cost", (data) -> scope.cost = data)
      updateCost()

      scope.formatTime = (value, ration, id) ->
        d = moment.duration(value, 'seconds')
        minutes = d.minutes()
        if minutes < 10 then minutes = "0#{minutes}"
        "#{moment.duration(value, 'seconds').hours()}h#{minutes}"

      scope.formatCost = (value, ration, id) -> value

      scope.$watch("filter", (v, o) ->
        if v.begin == o.begin and v.end == o.end then return
        scope.listing = null
        updateTime()
        updateCost()
      , true)

      scope.$watchGroup(["cost", "time"], (v) ->
        if !v[0]? || !v[1]? then return
        scope.listing = _.merge(scope.cost, scope.time)
      )


    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/stats/customer.html"
    }
])

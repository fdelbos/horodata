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
        "#{parseInt(moment.duration(value, 'seconds').asHours())}h#{minutes}"

      scope.formatCost = (value, ration, id) -> value

      scope.$watch("filter", (v, o) ->
        if v.begin == o.begin and v.end == o.end then return
        scope.listing = null
        updateTime()
        updateCost()
      , true)

      scope.$watchGroup(["cost", "time"], (v) ->
        if !v[0]? || !v[1]? then return

        m = {}
        m[i.customer_id] = {cost: i.cost} for i in scope.cost
        m[i.customer_id].duration = i.duration for i in scope.time

        listing = []
        for k, v of m
          listing.push {
            cost: v.cost
            duration: scope.formatTime(v.duration)
            name: scope.customers[k].name
            id: k
          }
        scope.listing = _.sortBy(listing, (i) -> i.name.toLowerCase())
      )


    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/stats/customer.html"
    }
])

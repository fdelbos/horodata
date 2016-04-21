angular.module("horodata").directive("appWidgetsStatsGuestTime", [
  "statsService"
  "statsFilterService"
  (statsService, statsFilterService)->

    l = (scope) ->
      scope.stats = statsService
      scope.filter = statsFilterService

      scope.tableData = []

      generateCosts = (data) ->
        guestRates = {}
        for g in scope.group.guests
          if g.rate? and parseFloat(g.rate) != NaN
            guestRates[g.id] = parseFloat g.rate
        costs = []

        for i in data
          if guestRates[i.guest_id]?
            c = (i.duration / 3600) * guestRates[i.guest_id]
            costs.push {
              guest_id: i.guest_id
              cost: c
            }
            scope.tableData.push {
              id: i.guest_id
              cost: scope.formatCost(c)
              duration: scope.formatDuration(i.duration)
            }
        return costs

      update = ->
        statsService.fetch(scope.group.url, "guest_time", (data) ->
          scope.data = data
          scope.costs = generateCosts(data)
        )
      update()

      scope.formatCost = (value, ratio, id) ->
        d3.format('.2f')(value)

      scope.formatDuration = (value, ration, id) ->
        d = moment.duration(value, 'seconds')
        minutes = d.minutes()
        if minutes < 10 then minutes = "0#{minutes}"
        "#{moment.duration(value, 'seconds').hours()}h#{minutes}"


      scope.$watch("filter", (v, o) ->
        if v.begin == o.begin and v.end == o.end then return
        update()
      , true)

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/stats/guest_time.html"
    }
])

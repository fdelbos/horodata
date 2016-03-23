angular.module("horodata").directive("appWidgetsStats", [
  ->

    l = (scope) ->
      scope.availableStats = [
        {
          id: "customer_time"
          label: "Repartition du temps par dossier."
        }
      ]

      scope.selected = null

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/stats.html"
    }
])

angular.module("horodata").directive("appWidgetsStats", [
  ->

    l = (scope) ->
      scope.availableStats = [
        {
          id: "customer_time"
          label: "Repartition du temps par dossier."
        }
        {
          id: "task_time"
          label: "Repartition du temps par tâche."
        }
        {
          id: "guest_time"
          label: "Repartition du temps par utilisateur."
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

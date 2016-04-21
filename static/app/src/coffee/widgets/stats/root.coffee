angular.module("horodata").directive("appWidgetsStats", [
  "statsFilterService"
  (statsFilterService)->

    l = (scope) ->

      scope.filter = statsFilterService
      scope.today = new Date()

      scope.availableStats = [
        {
          id: "customer"
          label: "Dossiers"
        }
        {
          id: "task_time"
          label: "Tâches"
        }
        {
          id: "guest_time"
          label: "Utilisateurs"
        }
      ]

      scope.selected = null

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/stats/root.html"
    }
])

angular.module("horodata").directive("appWidgetsNewTask", [
  "groupNewService"
  (groupNewService)->

    l = (scope) ->
      scope.showNewTaskDialog = (ev) -> groupNewService.open(ev)

    return {
      link: l
      restrict: "E"
      templateUrl: "horodata/widgets/new_task.html"
    }
])

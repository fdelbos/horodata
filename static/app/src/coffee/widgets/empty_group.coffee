angular.module("horodata").directive("appWidgetsEmptyGroup", [
  -> return {
      restrict: "E"
      templateUrl: "horodata/widgets/empty_group.html"
    }
])

angular.module("horodata").directive("appWidgetsEmptyGroupBoxed", [
  -> return {
      restrict: "E"
      templateUrl: "horodata/widgets/empty_group_boxed.html"
    }
])

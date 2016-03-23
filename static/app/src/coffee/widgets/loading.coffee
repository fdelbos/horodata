angular.module("horodata").directive("appWidgetsLoading", [
  ->
    return {
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/loading.html"
    }
])

angular.module("horodata").directive("appWidgetsConfiguration", [
  ->
    return {
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/configuration/root.html"
    }
])

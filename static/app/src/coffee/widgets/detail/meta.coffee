angular.module("horodata").directive("appWidgetsDetailMeta", [
  ->
    return {
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/detail/meta.html"
    }
])

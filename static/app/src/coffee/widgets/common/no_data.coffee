angular.module("horodata").directive("appWidgetsCommonNoData", [
  ->
    return {
      scope:
        begin: "="
        end: "="
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/common/no_data.html"
    }
])

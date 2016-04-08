angular.module("horodata").directive("appWidgetsCommonLoading", [
  ->
    return {
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/common/loading.html"
    }
])

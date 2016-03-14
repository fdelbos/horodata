angular.module("horodata").directive("appWidgetsSearchBar", [
  ->
    return {
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/search_bar.html"
    }
])

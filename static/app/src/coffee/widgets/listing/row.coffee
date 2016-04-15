angular.module("horodata").directive("appWidgetsListingRow", [
  ->
    return {
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/listing/row.html"
    }
])

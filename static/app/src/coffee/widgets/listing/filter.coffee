angular.module("horodata").directive("appWidgetsListingFilter", [
  ->

    l = (scope) ->
      scope.today = new Date()

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/listing/filter.html"
    }
])

angular.module("horodata").directive("appWidgetsSearchBar", [
  ->

    l = (scope) ->
      scope.today = new Date()

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/search_bar.html"
    }
])

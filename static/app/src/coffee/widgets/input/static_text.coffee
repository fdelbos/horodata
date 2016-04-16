angular.module("horodata").directive("appWidgetsInputStaticText", [
  ->
    return {
      scope:
        text: "@"
        caption: "@"
        help: "@"
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/input/static_text.html"
    }
])

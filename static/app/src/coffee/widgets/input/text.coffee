angular.module("horodata").directive("appWidgetsInputText", [
  ->
    return {
      scope:
        value: "="
        error: "="
        caption: "@"
        subCaption: "@"
        help: "@"
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/input/text.html"
    }
])

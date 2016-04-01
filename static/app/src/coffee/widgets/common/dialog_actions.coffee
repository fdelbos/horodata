angular.module("horodata").directive("appWidgetsCommonDialogActions", [
  "$mdDialog"
  ($mdDialog)->

    l = (scope) ->
      scope.hide = -> $mdDialog.hide()

    return {
      link: l
      transclude: true
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/common/dialog_actions.html"
    }
])

angular.module("horodata").directive("appWidgetsCommonDialogToolbar", [
  "$mdDialog"
  ($mdDialog)->

    l = (scope) ->

      scope.hide = -> $mdDialog.hide()

    return {
      link: l
      scope:
        warn: "="
      transclude: true
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/common/dialog_toolbar.html"
    }
])

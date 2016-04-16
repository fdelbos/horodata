angular.module("horodata").directive("appWidgetsCommonDialogActions", [
  "$mdDialog"
  ($mdDialog)->

    l = (scope, el, attr) ->
      scope.hide = -> $mdDialog.hide()
      
      if attr.close && !_.isEmpty(attr.close)
        scope.close = attr.close

    return {
      link: l
      transclude: true
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/common/dialog_actions.html"
    }
])

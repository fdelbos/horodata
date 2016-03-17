angular.module('horodata').factory("popupService", [
  "$mdMedia"
  "$mdDialog"
  ($mdMedia, $mdDialog)->
    
    return (tmpl, ctrl, scope, ev) ->
      fullscreen = $mdMedia('xs') || $mdMedia('sm')
      $mdDialog.show
        controller: ctrl
        templateUrl: tmpl
        parent: angular.element(document.body),
        targetEvent: ev,
        preserveScope: true
        scope: scope
        clickOutsideToClose:true
        escapeToClose: true
        fullscreen: fullscreen
])

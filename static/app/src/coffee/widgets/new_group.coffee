angular.module("horodata").directive("appWidgetsNewGroup", [
  "popupService"
  (popupService)->

    l = (scope) ->

      scope.showNewGroupDialog = (ev) ->
        popupService(
          "horodata/widgets/new_group_form.html"
          "newGroupDialog"
          scope, ev)

    return {
      link: l
      restrict: "E"
      templateUrl: "horodata/widgets/new_group.html"
    }
])

angular.module("horodata").controller("newGroupDialog", [
  "$scope",
  "$mdDialog",
  "$mdToast",
  "$http",
  "$location",
  "apiService"
  "groupNewService"
  ($scope, $mdDialog, $mdToast, $http, $location, apiService, groupNewService)->
    $scope.name = ""
    $scope.errors = null

    $scope.close = -> $mdDialog.hide()

    $scope.send = ->
      $http.post("#{apiService.get()}/groups", {name: $scope.name}).then(
        (resp) ->
          group = resp.data.data
          $mdDialog.hide()
          $mdToast.showSimple("Le nouveau groupe '#{group.name}' a été créé.")
          $location.path("/#{group.url}")
          groupNewService.fetch()
        (resp) -> $scope.errors = resp.data.errors
      )

])

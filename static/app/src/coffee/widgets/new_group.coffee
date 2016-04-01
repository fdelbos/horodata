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
    $scope.loading = false

    $scope.send = ->
      $scope.loading = true
      $http.post("#{apiService.get()}/groups", {name: $scope.name}).then(
        (resp) ->
          group = resp.data.data
          $mdDialog.hide()
          $mdToast.showSimple("Nouveau groupe '#{group.name}' créé")
          $location.path("/#{group.url}")
          groupNewService.fetch()
        (resp) ->
          $scope.loading = false
          $scope.errors = resp.data.errors
      )

])

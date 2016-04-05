angular.module("horodata").directive("appWidgetsConfigurationDelete", [
  "popupService"
  (popupService)->

    l = (scope, elem, attr) ->

      scope.deleteGroup = (ev)->
        popupService(
          "horodata/widgets/configuration/delete_confirm.html",
          "appWidgetsConfigurationDeleteConfirm"
          scope, ev)

    return {
      link: l
      restrict: "E"
      templateUrl: "horodata/widgets/configuration/delete.html"
    }
])

angular.module("horodata").controller("appWidgetsConfigurationDeleteConfirm", [
  "$scope",
  "$mdDialog",
  "$mdToast",
  "$http",
  "$location",
  "apiService"
  "groupService"
  ($scope, $mdDialog, $mdToast, $http, $location, apiService, groupService)->
    $scope.loading = false

    $scope.delete = ->
      $scope.loading = true
      $http.delete("#{apiService.get()}/groups/#{$scope.group.url}").then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Groupe '#{$scope.group.name}' supprimé")
          $location.path("/")
          groupService.fetch()
        (resp) ->
          $scope.errors = resp.data.errors
          $scope.loading = false
      )
])

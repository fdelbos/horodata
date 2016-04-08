angular.module("horodata").directive("appWidgetsConfigurationLeave", [
  "popupService"
  (popupService)->

    l = (scope, elem, attr) ->

      scope.leaveGroup = (ev)->
        popupService(
          "horodata/widgets/configuration/leave_confirm.html",
          "appWidgetsConfigurationLeaveConfirm"
          scope, ev)

    return {
      link: l
      restrict: "E"
      templateUrl: "horodata/widgets/configuration/leave.html"
    }
])

angular.module("horodata").controller("appWidgetsConfigurationLeaveConfirm", [
  "$scope",
  "$mdDialog",
  "$mdToast",
  "$http",
  "$location",
  "apiService"
  "groupService"
  ($scope, $mdDialog, $mdToast, $http, $location, apiService, groupService)->
    $scope.loading = false

    $scope.leave = ->
      $scope.loading = true
      $http.post("#{apiService.get()}/groups/#{$scope.group.url}/leave").then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Vous avez quitté le groupe '#{$scope.group.name}'")
          $location.path("/")
          groupService.fetch()
        (resp) ->
          $scope.errors = resp.data.errors
          $scope.loading = false
      )
])

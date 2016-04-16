angular.module("horodata").controller("Profile", [
  "$scope"
  "$mdDialog"
  "$mdToast"
  "$http"
  "apiService"
  "userService"
  ($scope, $mdDialog, $mdToast, $http, apiService, userService)->
    $scope.errors = null
    $scope.loading = false

    $scope.name = $scope.user.name

    $scope.update = ->
      $scope.loading = true
      $http.put("#{apiService.get()}/users/me", {name: $scope.name}).then(
        (resp) ->
          $scope.user.name = $scope.name
          $mdDialog.hide()
          $mdToast.showSimple("Votre profile a bien été modifié")
        (resp) ->
          $scope.errors = resp.data.errors
          $scope.loading = false
      )

])

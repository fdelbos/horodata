angular.module("horodata").controller("Contact", [
  "$scope"
  "$mdDialog"
  "$http"
  "apiService"
  "$mdToast"
  ($scope, $mdDialog, $http, apiService, $mdToast)->

    $scope.message = null

    $scope.send = ->
      $scope.loading = true
      $http.post("#{apiService.get()}/users/contact_message", $scope.message).then(
        (resp) ->
          $scope.loading = false
          $mdDialog.hide()
          $mdToast.showSimple("Votre message a bien ete envoye")
        (resp) ->
          $scope.loading = false
          $scope.errors = resp.data.errors
      )

])

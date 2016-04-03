angular.module("horodata").controller("Quotas", [
  "$scope"
  "$mdDialog"
  "$http"
  "apiService"
  ($scope, $mdDialog, $http, apiService)->

    $scope.loading = true
    $http.get("#{apiService.get()}/users/me/quotas").then(
      (resp) ->
        $scope.loading = false
        $scope.quotas = resp.data.data
      (resp) ->
        $scope.loading = false
    )

    $scope.send = ->
      $scope.loading = true

])

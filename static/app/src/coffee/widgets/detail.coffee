angular.module("horodata").controller("detailDialog", [
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

    $scope.close = -> $mdDialog.hide()

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

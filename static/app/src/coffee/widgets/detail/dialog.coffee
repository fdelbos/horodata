angular.module("horodata").controller("detailDialog", [
  "$scope",
  "$mdDialog",
  "$mdToast",
  "$http",
  "$location",
  "apiService"
  "groupService"
  ($scope, $mdDialog, $mdToast, $http, $location, apiService, groupService)->
    $scope.name = ""
    $scope.errors = null
    $scope.loading = false

    $scope.canEdit = false
    if $scope.isAdmin then $scope.canEdit = true
    else if moment($scope.detailJob.created).isSame(new Date(), "day")
      $scope.canEdit = true

    $scope.hours = [0..12]
    $scope.minutes = (x for x in [0..55] by 5)

    $scope.detailJob.hours = Math.floor($scope.detailJob.duration / 3600)
    $scope.detailJob.minutes = Math.floor(($scope.detailJob.duration % 3600) / 60)

    $scope.close = -> $mdDialog.hide()

    $scope.send = ->
      $scope.loading = true
      $http.post("#{apiService.get()}/groups", {name: $scope.name}).then(
        (resp) ->
          group = resp.data.data
          $mdDialog.hide()
          $mdToast.showSimple("Nouveau groupe '#{group.name}' créé")
          $location.path("/#{group.url}")
          groupService.fetch()
        (resp) ->
          $scope.loading = false
          $scope.errors = resp.data.errors
      )
])

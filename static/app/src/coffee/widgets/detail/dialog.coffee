angular.module("horodata").controller("detailDialog", [
  "$scope",
  "$mdDialog",
  "$mdToast",
  "$http",
  "apiService"
  "listingService"
  ($scope, $mdDialog, $mdToast, $http, apiService, listingService)->
    $scope.name = ""
    $scope.errors = null
    $scope.loading = false

    $scope.deleteMode = false
    $scope.toggleDeleteMode = -> $scope.deleteMode = !$scope.deleteMode

    $scope.canEdit = false
    if $scope.isAdmin then $scope.canEdit = true
    else if moment($scope.detailJob.created).isSame(new Date(), "day")
      $scope.canEdit = true

    $scope.hours = [0..12]
    $scope.minutes = (x for x in [0..55] by 5)

    $scope.detailJob.hours = Math.floor($scope.detailJob.duration / 3600)
    $scope.detailJob.minutes = Math.floor(($scope.detailJob.duration % 3600) / 60)

    $scope.close = -> $mdDialog.hide()

    $scope.update = ->
      $scope.loading = true

      job =
        duration:$scope.detailJob.hours * 3600 + $scope.detailJob.minutes * 60
        task: parseInt $scope.detailJob.task_id
        customer:  parseInt $scope.detailJob.customer_id
        comment:  $scope.detailJob.comment

      $http.put("#{apiService.get()}/groups/#{$scope.group.url}/jobs/#{$scope.detailJob.id}", job).then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Saisie mise à jour")
          listingService.get().reload()
        (resp) ->
          $scope.loading = false
          $scope.errors = resp.data.errors
      )

    $scope.delete = ->
      $scope.loading = true
      $http.delete("#{apiService.get()}/groups/#{$scope.group.url}/jobs/#{$scope.detailJob.id}").then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Saisie supprimee")
          listingService.get().reload()
        (resp) ->
          $scope.loading = false
          $scope.errors = resp.data.errors
      )

])

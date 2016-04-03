angular.module("horodata").directive("appWidgetsNewTask", [
  "groupNewService"
  (groupNewService)->

    l = (scope) ->
      scope.showNewTaskDialog = (ev) -> groupNewService.open(ev)

    return {
      link: l
      restrict: "E"
      templateUrl: "horodata/widgets/new_task.html"
    }
])


angular.module("horodata").controller("newTaskDialog", [
  "$scope",
  "$mdDialog",
  "$mdToast",
  "$http",
  "$location",
  "apiService"
  "listingService"
  ($scope, $mdDialog, $mdToast, $http, $location, apiService, listingService)->
    $scope.task =
      minutes: 0
      hours: 0
    $scope.errors = null
    $scope.loading = false
    $scope.quotaError = null

    $scope.hours = [0..12]
    $scope.minutes = (x for x in [0..55] by 5)

    $scope.send = ->
      task =
        duration: $scope.task.hours * 3600 + $scope.task.minutes * 60
        task: parseInt $scope.task.task
        customer:  parseInt $scope.task.customer
        comment:  $scope.task.comment
      $scope.loading = true
      $http.post("#{apiService.get()}/groups/#{$scope.group.url}/jobs",task).then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Nouvelle tâche saisie")
          listingService.listing().fetch()
        (resp) ->
          $scope.loading = false
          if resp.status == 429 && _.get(resp, "data.errors.type") == "quota"
            $scope.quotaError = resp.data.errors
          else $scope.errors = resp.data.errors
      )

])

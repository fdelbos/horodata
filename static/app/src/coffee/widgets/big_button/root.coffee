angular.module("horodata").directive("appWidgetsBigButton", [
  "tabsService"
  "popupService"
  (tabsService, popupService)->

    l = (scope) ->
      scope.currentTab = ->
        switch tabsService.get()
          when 0 then "jobs"
          when 1 then "export"
          else null

      scope.newDialog = (ev) ->
        if scope.currentTab() == "jobs"
          popupService(
            "horodata/widgets/big_button/new_task.html"
            "newTaskDialog"
            scope, ev)
        if scope.currentTab() == "export"
          popupService(
            "horodata/widgets/big_button/export.html"
            "exportDialog"
            scope, ev)

    return {
      link: l
      restrict: "E"
      templateUrl: "horodata/widgets/big_button/root.html"
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
  "groupService"
  ($scope, $mdDialog, $mdToast, $http, $location, apiService, listingService, groupService)->

    $scope.group = groupService.get()
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


angular.module("horodata").controller("exportDialog", [
  "$scope"
  "$mdDialog"
  "apiService"
  "groupService"
  "statsFilterService"
  ($scope, $mdDialog, apiService, groupService, statsFilterService)->

    $scope.filter = statsFilterService
    $scope.group = groupService.get()
    $scope.export =
      fileType: "xlsx"

    $scope.url = "#{apiService.get()}/groups/#{$scope.group.url}/export"
])

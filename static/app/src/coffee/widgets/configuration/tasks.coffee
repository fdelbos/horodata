angular.module("horodata").directive("appWidgetsConfigurationTasks", [
  "popupService"
  (popupService)->

    l = (scope, elem, attr) ->

      scope.tasks =
        current: null
        selected: null
        edit: (ev) ->
          @current = _.cloneDeep(_.find(scope.group.tasks, {id: parseInt @selected}))
          popupService(
            "horodata/widgets/configuration/tasks_form.html"
            "appWidgetsConfigurationTasksDialog"
            scope, ev)
        create: (ev) ->
          @current =
            name: ""
            comment_mandatory: false
          popupService(
            "horodata/widgets/configuration/tasks_form.html"
            "appWidgetsConfigurationTasksDialog"
            scope, ev)

    return {
      link: l
      restrict: "E"
      templateUrl: "horodata/widgets/configuration/tasks.html"
    }
])

angular.module("horodata").controller("appWidgetsConfigurationTasksDialog", [
  "$scope"
  "$mdDialog"
  "$mdToast"
  "$http"
  "$location"
  "apiService"
  ($scope, $mdDialog, $mdToast, $http, $location, apiService)->
    $scope.errors = null
    $scope.loading = false

    update = (t) ->
      idx = _.findIndex($scope.group.tasks, {id: t.id})
      $scope.group.tasks[idx] = $scope.tasks.current
      $scope.group.tasks = _.sortBy($scope.group.tasks, ["name"])

    $scope.create = ->
      $scope.loading = true
      $http.post("#{apiService.get()}/groups/#{$scope.group.url}/tasks", $scope.tasks.current).then(
        (resp) ->
          $scope.$emit("group.reload")
          $mdDialog.hide()
          $mdToast.showSimple("Nouvelle tâche '#{$scope.tasks.current.name}' ajoutée")
        (resp) ->
          $scope.loading = false
          $scope.errors = resp.data.errors
      )

    $scope.edit = ->
      $scope.loading = true
      $http.put("#{apiService.get()}/groups/#{$scope.group.url}/tasks/#{ $scope.tasks.selected }", $scope.tasks.current).then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Tâche '#{$scope.tasks.current.name}' modifiée")
          update($scope.tasks.current)
          $scope.tasks.selected = null
        (resp) ->
          $scope.loading = false
          $scope.errors = resp.data.errors
      )

    $scope.delete = ->
      $scope.loading = true
      $http.delete("#{apiService.get()}/groups/#{$scope.group.url}/tasks/#{ $scope.tasks.selected }", $scope.tasks.current).then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Tâche '#{$scope.tasks.current.name}' supprimée")
          $scope.group.tasks.splice(_.findIndex($scope.group.tasks, {id: parseInt $scope.tasks.selected}), 1)
          $scope.tasks.selected = null
        (resp) ->
          $scope.loading = false
          $scope.errors = resp.data.errors
      )



])

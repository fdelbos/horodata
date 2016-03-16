angular.module("horodata").directive("appWidgetsConfigurationTasks", [
  "$mdDialog"
  "$mdMedia"
  ($mdDialog, $mdMedia)->

    l = (scope, elem, attr) ->

      openForm = (ev) ->
        fullscreen = $mdMedia('xs') || $mdMedia('sm')

        $mdDialog.show({
          controller: "appWidgetsConfigurationTasksDialog",
          templateUrl: "horodata/widgets/configuration/tasks_form.html",
          parent: angular.element(document.body),
          targetEvent: ev,
          preserveScope: true,
          scope: scope,
          clickOutsideToClose:true,
          escapeToClose: true,
          fullscreen: fullscreen
        })

      scope.tasks =
        current: null
        selected: null
        edit: ->
          @current = _.cloneDeep(_.find(scope.group.tasks, {id: parseInt @selected}))
          openForm()
        create: ->
          @current =
            name: ""
            comment_mandatory: false
          openForm()

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

    $scope.close = -> $mdDialog.hide()

    update = (t) ->
      idx = _.findIndex($scope.group.tasks, {id: t.id})
      $scope.group.tasks[idx] = $scope.tasks.current
      $scope.group.tasks = _.sortBy($scope.group.tasks, ["name"])

    $scope.create = ->
      $http.post("#{apiService.get()}/groups/#{$scope.group.url}/tasks", $scope.tasks.current).then(
        (resp) ->
          $scope.$emit("group.reload")
          $mdDialog.hide()
          $mdToast.showSimple("Nouveau type de tâche: '#{$scope.tasks.current.name}' ajouté.")
        (resp) -> $scope.errors = resp.tasks.errors
      )

    $scope.edit = ->
      $http.put("#{apiService.get()}/groups/#{$scope.group.url}/tasks/#{ $scope.tasks.selected }", $scope.tasks.current).then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Type de tâche: '#{$scope.tasks.current.name}' mis a jour.")
          update($scope.tasks.current)
          $scope.tasks.selected = null
        (resp) -> $scope.errors = resp.data.errors
      )

    $scope.delete = ->
      $http.delete("#{apiService.get()}/groups/#{$scope.group.url}/tasks/#{ $scope.tasks.selected }", $scope.tasks.current).then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Type de tâche: '#{$scope.tasks.current.name}' supprimé.")
          $scope.group.tasks.splice(_.findIndex($scope.group.tasks, {id: parseInt $scope.tasks.selected}), 1)
          $scope.task.selected = null
        (resp) -> $scope.errors = resp.data.errors
      )



])

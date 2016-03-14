angular.module("horodata").directive("appWidgetsConfigurationTasks", [
  "$mdDialog",
  "$mdMedia",
  ($mdDialog, $mdMedia)->

    l = (scope, elem, attr) ->

      scope.showConfigurationTasksDialog = (ev) ->
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

    return {
      link: l
      restrict: "E"
      templateUrl: "horodata/widgets/configuration/tasks.html"
    }
])

angular.module("horodata").controller("appWidgetsConfigurationTasksDialog", [
  "$scope",
  "$mdDialog",
  "$mdToast",
  "$http",
  "$location",
  "apiService"
  ($scope, $mdDialog, $mdToast, $http, $location, apiService)->
    $scope.name = ""
    $scope.errors = null

    $scope.close = -> $mdDialog.hide()

    # $scope.selectedTask = null

    $scope.send = ->
      $http.post("#{apiService.get()}/groups", {name: $scope.name}).then(
        (resp) ->
          group = resp.data.data
          $mdDialog.hide()
          $mdToast.showSimple("Nouveau groupe '#{group.name}' sauvegarde.")
          $location.path("/group/#{group.url}")
        (resp) -> $scope.errors = resp.data.errors
      )

])

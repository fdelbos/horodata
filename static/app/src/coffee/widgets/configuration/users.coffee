angular.module("horodata").directive("appWidgetsConfigurationUsers", [
  "$mdDialog",
  "$mdMedia",
  ($mdDialog, $mdMedia)->

    l = (scope, elem, attr) ->

      scope.showConfigurationUsersDialog = (ev) ->
        fullscreen = $mdMedia('xs') || $mdMedia('sm')

        $mdDialog.show({
          controller: "appWidgetsConfigurationUsersDialog",
          templateUrl: "horodata/widgets/configuration/user_form.html",
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
      templateUrl: "horodata/widgets/configuration/users.html"
    }
])

angular.module("horodata").controller("appWidgetsConfigurationUsersDialog", [
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

angular.module("horodata").directive("appWidgetsConfigurationCustomers", [
  "$mdDialog",
  "$mdMedia",
  ($mdDialog, $mdMedia)->

    l = (scope, elem, attr) ->

      scope.showConfigurationCustomersDialog = (ev) ->
        
        fullscreen = $mdMedia('xs') || $mdMedia('sm')

        $mdDialog.show({
          controller: "appWidgetsConfigurationCustomersDialog",
          templateUrl: "horodata/widgets/configuration/customers_form.html",
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
      templateUrl: "horodata/widgets/configuration/customers.html"
    }
])

angular.module("horodata").controller("appWidgetsConfigurationCustomersDialog", [
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

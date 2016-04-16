angular.module("horodata").directive("appWidgetsConfigurationCustomers", [
  "popupService"
  (popupService)->

    l = (scope, elem, attr) ->

      scope.data =
        current: null
        selected: null
        edit: (ev)->
          @current = _.cloneDeep(_.find(scope.group.customers, {id: parseInt @selected}))
          popupService(
            "horodata/widgets/configuration/customers_edit_form.html"
            "appWidgetsConfigurationCustomersDialog"
            scope, ev)

        create: (ev)->
          @current = {customers: ""}
          popupService(
            "horodata/widgets/configuration/customers_create_form.html"
            "appWidgetsConfigurationCustomersDialog"
            scope, ev)

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
  "apiService",
  ($scope, $mdDialog, $mdToast, $http, $location, apiService)->
    $scope.errors = null
    $scope.loading = false

    $scope.create = ->
      $scope.loading = true
      $http.post("#{apiService.get()}/groups/#{$scope.group.url}/customers", $scope.data.current).then(
        (resp) ->
          total = resp.data.data.total
          $scope.$emit("group.reload")
          $mdDialog.hide()
          if total == 1
            $mdToast.showSimple("1 nouveau dossier ajouté")
          else
            $mdToast.showSimple("#{total} nouveaux dossiers ajoutés")
        (resp) ->
          $scope.errors = resp.data.errors
          $scope.loading = false
      )

    $scope.edit = ->
      $scope.loading = true
      $http.put("#{apiService.get()}/groups/#{$scope.group.url}/customers/#{ $scope.data.selected }", $scope.data.current).then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Dossier '#{$scope.data.current.name}' modifié")
          $scope.$emit("group.reload")
          $scope.data.selected = null
        (resp) ->
          $scope.errors = resp.data.errors
          $scope.loading = false
      )

    $scope.delete = ->
      $scope.loading = true
      $http.delete("#{apiService.get()}/groups/#{$scope.group.url}/customers/#{ $scope.data.selected }").then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Dossier '#{$scope.data.current.name}' supprimé")
          $scope.$emit("group.reload")
          $scope.data.selected = null
        (resp) ->
          $scope.errors = resp.data.errors
          $scope.loading = false
      )

])

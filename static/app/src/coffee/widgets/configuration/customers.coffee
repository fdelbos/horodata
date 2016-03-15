angular.module("horodata").directive("appWidgetsConfigurationCustomers", [
  "$mdDialog",
  "$mdMedia",
  ($mdDialog, $mdMedia)->

    l = (scope, elem, attr) ->

      openCreateForm = (ev) ->
        fullscreen = $mdMedia('xs') || $mdMedia('sm')
        $mdDialog.show
          controller: "appWidgetsConfigurationCustomersDialog"
          templateUrl: "horodata/widgets/configuration/customers_create_form.html"
          parent: angular.element(document.body)
          targetEvent: ev
          preserveScope: true
          scope: scope
          clickOutsideToClose:true
          escapeToClose: true
          fullscreen: fullscreen

      openEditForm = (ev) ->
        fullscreen = $mdMedia('xs') || $mdMedia('sm')
        $mdDialog.show
          controller: "appWidgetsConfigurationCustomersDialog"
          templateUrl: "horodata/widgets/configuration/customers_edit_form.html"
          parent: angular.element(document.body)
          targetEvent: ev
          preserveScope: true
          scope: scope
          clickOutsideToClose:true
          escapeToClose: true
          fullscreen: fullscreen

      scope.customers =
        current: null
        selected: null
        edit: ->
          @current = _.cloneDeep(_.find(scope.group.customers, {id: parseInt @selected}))
          openEditForm()
        create: ->
          @current = {customers: ""}
          openCreateForm()

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
    $scope.errors = null

    $scope.close = -> $mdDialog.hide()

    update = (t) ->
      idx = _.findIndex($scope.group.customers, {id: t.id})
      $scope.group.customers[idx] = $scope.customers.current
      $scope.group.customers = _.sortBy($scope.group.customers, ["name"])

    $scope.create = ->
      $http.post("#{apiService.get()}/groups/#{$scope.group.url}/customers", $scope.customers.current).then(
        (resp) ->
          total = resp.data.data.total
          $scope.$emit("group.reload")
          $mdDialog.hide()
          if total == 1
            $mdToast.showSimple("1 nouveau dossier a été ajouté.")
          else
            $mdToast.showSimple("#{total} nouveaux dossiers ont été ajoutés.")
        (resp) -> $scope.errors = resp.data.errors
      )

    $scope.edit = ->
      $http.put("#{apiService.get()}/groups/#{$scope.group.url}/customers/#{ $scope.customers.selected }", $scope.customers.current).then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Dossier: '#{$scope.customers.current.name}' mis a jour.")
          update($scope.customers.current)
          $scope.customers.selected = null
        (resp) -> $scope.errors = resp.data.errors
      )

    $scope.delete = ->
      $http.delete("#{apiService.get()}/groups/#{$scope.group.url}/customers/#{ $scope.customers.selected }", $scope.customers.current).then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Dossier: '#{$scope.customers.current.name}' supprimé.")
          $scope.group.customers.splice(_.findIndex($scope.group.customers, {id: parseInt $scope.customers.selected}), 1)
          $scope.customers.selected = null
        (resp) -> $scope.errors = resp.data.errors
      )

])

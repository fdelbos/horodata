angular.module("horodata").directive("appWidgetsConfigurationGuests", [
  "popupService"
  (popupService)->

    l = (scope, elem, attr) ->
      scope.guests =
        current: null
        selected: null
        edit: (ev) ->
          @current = _.cloneDeep(_.find(scope.group.guests, {id: parseInt @selected}))
          popupService(
            "horodata/widgets/configuration/guests_edit_form.html",
            "appWidgetsConfigurationGuestsDialog"
            scope, ev)
        create: (ev) ->
          @current =
            email: ""
            admin: false
            rate: 0
          popupService(
            "horodata/widgets/configuration/guests_create_form.html",
            "appWidgetsConfigurationGuestsDialog"
            scope, ev)

    return {
      link: l
      restrict: "E"
      templateUrl: "horodata/widgets/configuration/guests.html"
    }
])


angular.module("horodata").controller("appWidgetsConfigurationGuestsDialog", [
  "$scope"
  "$mdDialog"
  "$mdToast"
  "$http"
  "$location"
  "apiService"
  ($scope, $mdDialog, $mdToast, $http, $location, apiService)->
    $scope.errors = null
    $scope.loading = false

    $scope.close = -> $mdDialog.hide()

    update = (t) ->
      idx = _.findIndex($scope.group.guests, {id: t.id})
      $scope.group.guests[idx] = $scope.guests.current
      $scope.group.guests = _.sortBy($scope.group.guests, ["name"])

    $scope.create = ->
      $scope.loading = true
      $http.post("#{apiService.get()}/groups/#{$scope.group.url}/guests", $scope.guests.current).then(
        (resp) ->
          $scope.$emit("group.reload")
          $mdDialog.hide()
          $mdToast.showSimple("Nouvel utilisateur '#{$scope.guests.current.email}' ajouté")
        (resp) ->
          $scope.loading = false
          $scope.errors = resp.data.errors
      )

    $scope.edit = ->
      $scope.loading = true
      $http.put("#{apiService.get()}/groups/#{$scope.group.url}/guests/#{ $scope.guests.selected }", $scope.guests.current).then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Utilisateur '#{$scope.guests.current.email}' modifié")
          update($scope.guests.current)
          $scope.guests.selected = null
        (resp) ->
          $scope.loading = false
          $scope.errors = resp.data.errors
      )

    $scope.delete = ->
      $scope.loading = true
      $http.delete("#{apiService.get()}/groups/#{$scope.group.url}/guests/#{ $scope.guests.selected }", $scope.guests.current).then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Utilisateur '#{$scope.guests.current.email}' supprimé")
          $scope.group.guests.splice(_.findIndex($scope.group.guests, {id: parseInt $scope.guests.selected}), 1)
          $scope.task.selected = null
        (resp) ->
          $scope.loading = false
          $scope.errors = resp.data.errors
      )



])

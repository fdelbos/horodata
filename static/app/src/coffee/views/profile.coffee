angular.module("horodata").controller("Profile", [
  "$scope"
  "$mdDialog"
  "$mdToast"
  "$http"
  "apiService"
  "userService"
  ($scope, $mdDialog, $mdToast, $http, apiService, userService)->
    $scope.errors = null
    $scope.loading = false

    $scope.name = $scope.user.name

    $scope.send = ->
      $scope.loading = true


    # userService.get (u) ->
    #   $scope.name = u.name

    # $scope.edit = ->
    #   $scope.loading = true
    #   $http.put("#{apiService.get()}/groups/#{$scope.group.url}/customers/#{ $scope.customers.selected }", $scope.customers.current).then(
    #     (resp) ->
    #       $mdDialog.hide()
    #       $mdToast.showSimple("Dossier '#{$scope.customers.current.name}' modifié")
    #       update($scope.customers.current)
    #       $scope.customers.selected = null
    #     (resp) ->
    #       $scope.errors = resp.data.errors
    #       $scope.loading = false
    #   )
    #
    # $scope.delete = ->
    #   $scope.loading = true
    #   $http.delete("#{apiService.get()}/groups/#{$scope.group.url}/customers/#{ $scope.customers.selected }").then(
    #     (resp) ->
    #       $mdDialog.hide()
    #       $mdToast.showSimple("Dossier '#{$scope.customers.current.name}' supprimé")
    #       $scope.group.customers.splice(_.findIndex($scope.group.customers, {id: parseInt $scope.customers.selected}), 1)
    #       $scope.customers.selected = null
    #     (resp) ->
    #       $scope.errors = resp.data.errors
    #       $scope.loading = false
    #   )


])

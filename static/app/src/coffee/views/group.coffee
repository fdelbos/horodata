angular.module("horodata").controller("Group", [
  "$http"
  "$routeParams"
  "$scope"
  "titleService"
  "userService"
  "apiService"
  ($http, $routeParams, $scope, titleService, userService, apiService)->

    $scope.maxDate = new Date()
    $scope.endDate = $scope.maxDate
    $scope.beginDate = moment().subtract(1, 'months').toDate()

    fetchUsers = ->
      $http.get("#{apiService.get()}/groups/#{$routeParams.group}/users").then(
        (resp) -> $scope.users = resp.data.data)

    $scope.isOwner = false

    userService.get((u) ->
      $http.get("#{apiService.get()}/groups/#{$routeParams.group}").then(
        (resp) ->
          $scope.group = resp.data.data
          if $scope.group.owner.login == u.login
            $scope.isOwner = true
          titleService.set($scope.group.name)
          fetchUsers()))
])

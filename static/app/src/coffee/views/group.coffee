angular.module("horodata").controller("Group", [
  "$http"
  "$routeParams"
  "$scope"
  "titleService"
  "userService"
  "apiService"
  ($http, $routeParams, $scope, titleService, userService, apiService)->

    $scope.isGroupView = true

    $scope.maxDate = new Date()
    $scope.endDate = $scope.maxDate
    $scope.beginDate = moment().subtract(1, 'months').toDate()

    fetchUsers = ->
      $http.get("#{apiService.get()}/groups/#{$routeParams.group}/users").then(
        (resp) -> $scope.users = resp.data.data)

    $scope.isOwner = false

    getGroup = ->
      $http.get("#{apiService.get()}/groups/#{$routeParams.group}").then(
        (resp) ->
          $scope.group = resp.data.data
          if $scope.group.owner.login == $scope.user.login
            $scope.isOwner = true
          titleService.set($scope.group.name, true)
      )

    userService.get((u) ->
      $scope.user = u
      getGroup()
      if $scope.isOwner then fetchUsers())

    $scope.$on("group.reload", (e) ->
      e.stopPropagation()
      getGroup())
])

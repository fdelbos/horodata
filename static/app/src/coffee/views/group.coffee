angular.module("horodata").controller("Group", [
  "$http"
  "$routeParams"
  "$scope"
  "titleService"
  "userService"
  "apiService"
  "groupService"
  "popupService"
  "listingService"
  "tabsService"
  ($http, $routeParams, $scope, titleService, userService, apiService, groupService, popupService, listingService, tabsService)->

    $scope.isGroupView = true
    $scope.isAdmin = false
    $scope.isLoading = false
    $scope.group = null
    $scope.groupError = null
    $scope.selectedTab =
      id: 0

    $scope.selectTab = (i)->
      $scope.selectedTab.id = i
      if i == 0 then tabsService.set "jobs"
      else if i == 1 and $scope.isAdmin == true then tabsService.set "export"
      else tabsService.set null

    $scope.goLeft = -> $scope.selectTab($scope.selectedTab.id - 1)
    $scope.goRight = -> $scope.selectTab($scope.selectedTab.id + 1)

    $scope.$watch("selectedTab.id", (v, o) -> if v != o then $scope.selectTab(v))

    getGroup = ->
      $scope.isLoading = true
      $http.get("#{apiService.get()}/groups/#{$routeParams.group}").then(
        (resp) ->
          if !$scope.group? then $scope.selectTab(0)

          $scope.group = resp.data.data
          groupService.set($scope.group)

          $scope.isAdmin = $scope.group.guests?
          $scope.isOwner = $scope.user.id == $scope.group.owner

          $scope.tasks = _.keyBy($scope.group.tasks, 'id')
          $scope.customers = _.keyBy($scope.group.customers, 'id')
          $scope.guests = _.keyBy($scope.group.guests, 'id')
          titleService.set($scope.group.name, true)
          $scope.isLoading = false
        (resp) ->
          $scope.isLoading = false
          $scope.groupError = switch resp.status
            when 403 then "Forbidden"
            when 404 then "NotFound"
            else "unknow"
      )

    userService.get((u) ->
      $scope.user = u
      getGroup())

    $scope.$on("group.reload", (e) ->
      e.stopPropagation()
      getGroup())

])

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

    getGroup = ->
      $scope.isLoading = true
      $http.get("#{apiService.get()}/groups/#{$routeParams.group}").then(
        (resp) ->
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

    $scope.selectTab = (i)->
      $scope.selectedTab = i
      tabsService.set(i)

    $scope.$watch("selectedTab", (v, o) -> if v != o then $scope.selectTab(v))
    $scope.selectTab(0)

])

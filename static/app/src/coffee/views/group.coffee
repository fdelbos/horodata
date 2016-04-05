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

    $scope.search =
      begin: moment().subtract(1, 'months').toDate()
      end: new Date()
      customer: null
      guest: null


    $scope.$watch("search", (v) ->
      if !v? then return
      listingService.search($routeParams.group, v)
      listingService.listing().fetch(1)
    , true)
    $scope.isAdmin = false

    getGroup = ->
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
      )

    userService.get((u) ->
      $scope.user = u
      getGroup())

    $scope.$on("group.reload", (e) ->
      e.stopPropagation()
      getGroup())

    # groupService.set (ev)->
    #   popupService(
    #     "horodata/widgets/new_task_form.html"
    #     "newTaskDialog"
    #     $scope, ev)

    $scope.selectTab = (i)->
      $scope.selectedTab = i
      tabsService.set(i)

    $scope.$watch("selectedTab", (v, o) -> if v != o then $scope.selectTab(v))
    $scope.selectTab(0)

])

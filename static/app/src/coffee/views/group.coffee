angular.module("horodata").controller("Group", [
  "$http"
  "$routeParams"
  "$scope"
  "titleService"
  "userService"
  "apiService"
  "groupNewService"
  "popupService"
  "listingService"
  ($http, $routeParams, $scope, titleService, userService, apiService, groupNewService, popupService, listingService)->

    $scope.isGroupView = true
    $scope.selectedTab = 0

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
          $scope.isAdmin = $scope.group.guests?
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

    groupNewService.set (ev)->
      popupService(
        "horodata/widgets/new_task_form.html"
        "newTaskDialog"
        $scope, ev)

    $scope.selectTab = (i)-> $scope.selectedTab = i

])

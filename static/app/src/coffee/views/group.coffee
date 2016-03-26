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
        "horodata/views/new_task_form.html"
        "groupNewTaskDialog"
        $scope, ev)

])

angular.module("horodata").controller("groupNewTaskDialog", [
  "$scope",
  "$mdDialog",
  "$mdToast",
  "$http",
  "$location",
  "apiService"
  "listingService"
  ($scope, $mdDialog, $mdToast, $http, $location, apiService, listingService)->
    $scope.task =
      minutes: 0
      hours: 0
    $scope.errors = null
    $scope.loading = false

    $scope.hours = [0..12]
    $scope.minutes = (x for x in [0..55] by 5)

    $scope.close = -> $mdDialog.hide()

    $scope.send = ->
      task =
        duration: $scope.task.hours * 3600 + $scope.task.minutes * 60
        task: parseInt $scope.task.task
        customer:  parseInt $scope.task.customer
        comment:  $scope.task.comment
      $scope.loading = true
      $http.post("#{apiService.get()}/groups/#{$scope.group.url}/jobs",task).then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Nouvelle tâche saisie")
          listingService.listing().fetch()
        (resp) ->
          $scope.loading = false
          $scope.errors = resp.data.errors
      )

])

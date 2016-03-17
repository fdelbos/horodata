angular.module("horodata").controller("Group", [
  "$http"
  "$routeParams"
  "$scope"
  "titleService"
  "userService"
  "apiService"
  "groupNewService"
  "$mdMedia"
  "$mdDialog"
  ($http, $routeParams, $scope, titleService, userService, apiService, groupNewService, $mdMedia, $mdDialog)->

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

    groupNewService.set (ev)->
      fullscreen = $mdMedia('xs') || $mdMedia('sm')
      $mdDialog.show
        controller: "groupNewTaskDialog"
        templateUrl: "horodata/views/new_task_form.html"
        parent: angular.element(document.body)
        targetEvent: ev
        preserveScope: true
        scope: $scope
        clickOutsideToClose:true
        escapeToClose: true
        fullscreen: fullscreen
])

angular.module("horodata").controller("groupNewTaskDialog", [
  "$scope",
  "$mdDialog",
  "$mdToast",
  "$http",
  "$location",
  "apiService"
  ($scope, $mdDialog, $mdToast, $http, $location, apiService)->
    $scope.task = {}
    $scope.errors = null

    $scope.close = -> $mdDialog.hide()

    $scope.send = ->
      $http.post("#{apiService.get()}/groups", {name: $scope.name}).then(
        (resp) ->
          group = resp.data.data
          $mdDialog.hide()
          $mdToast.showSimple("Nouveau groupe '#{group.name}' sauvegarde.")
          $location.path("/group/#{group.url}")
        (resp) -> $scope.errors = resp.data.errors
      )

])

angular.module("horodata").controller("Profile", [
  "$scope"
  "$mdDialog"
  "$mdToast"
  "$http"
  "apiService"
  "userService"
  "homeService"
  ($scope, $mdDialog, $mdToast, $http, apiService, userService, homeService)->
    $scope.errors = null
    $scope.loading = false

    $scope.name = $scope.user.name


    $scope.update = ->
      $scope.loading = true
      hasImage = false

      errorHandler = (resp) ->
        $scope.errors = resp.data.errors
        $scope.loading = false

      okHandler = (resp) ->
        $scope.user.name = $scope.name
        $mdDialog.hide()
        if hasImage
          t = $mdToast.simple()
          t.textContent("Votre profil a bien été modifié (rechargez la page pour voir votre nouvelle photo).")
          t.hideDelay = 8000
          $mdToast.show(t)
        else
          $mdToast.showSimple("Votre profil a bien été modifié")

      formData = new FormData()
      formData.append("name", $scope.name)
      if $scope.files.length > 0
        hasImage = true
        angular.forEach($scope.files, (f) -> formData.append('file', f.lfFile))

      $http.put("#{apiService.get()}/users/me", formData, {
        transformRequest: angular.identity
        headers: {'Content-Type': undefined}
      }).then(okHandler, errorHandler)

])

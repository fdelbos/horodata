angular.module("horodata").controller("Index", [
  "$http"
  "$scope"
  "userService"
  "titleService"
  ($http, $scope, userService, titleService)->

    titleService.set("Choisissez un Groupe")

])

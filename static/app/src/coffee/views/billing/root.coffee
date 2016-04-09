angular.module("horodata").controller("Billing", [
  "$http"
  "$scope"
  "userService"
  "titleService"
  ($http, $scope, userService, titleService)->

    titleService.set("Abonnement")

    $scope.plans =
      current: "free"
      free:
        code: "free"
        name: "Gratuit"
        price: 0
        groups: 1
        guests: 2
        jobs: 15
      small:
        code: "small"
        name: "10 utilisateurs"
        price: 10
        groups: 2
        guests: 10
        jobs: 500
      medium:
        code: "medium"
        name: "30 utilisateurs"
        price: 20
        groups: 5
        guests: 30
        jobs: 1500
      large:
        code: "large"
        name: "100 utilisateurs"
        price: 50
        groups: 15
        guests: 100
        jobs: 5000

])

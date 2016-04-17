angular.module("horodata").controller("Billing", [
  "$scope"
  "$http"
  "apiService"
  "titleService"
  ($scope, $http, apiService, titleService)->

    titleService.set("Abonnement")

    $scope.plans =
      current: null
      free:
        code: "free"
        name: "Gratuit"
        price: 0
        priceTTC: 0
        groups: 1
        guests: 2
        jobs: 15
      small:
        code: "small"
        name: "10 utilisateurs"
        price: 10
        priceTTC: 12
        groups: 2
        guests: 10
        jobs: 500
      medium:
        code: "medium"
        name: "30 utilisateurs"
        price: 20
        priceTTC: 24
        groups: 5
        guests: 30
        jobs: 1500
      large:
        code: "large"
        name: "100 utilisateurs"
        price: 50
        priceTTC: 60
        groups: 15
        guests: 100
        jobs: 5000

    reload = ->
      $scope.loading = true
      $http.get("#{apiService.get()}/billing/plan").then(
        (resp) ->
          $scope.loading = false
          $scope.plans.current = resp.data.data.plan
          if resp.data.data.end?
            $scope.plans.end = resp.data.data.end
          else $scope.plans.end = null
      )
    reload()

    $scope.$on("plan.reload", (e)->
      e.stopPropagation()
      reload()
    )
])

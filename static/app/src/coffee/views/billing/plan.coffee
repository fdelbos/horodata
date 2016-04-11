angular.module("horodata").directive("billingPlan", [
  "popupService"
  (popupService) ->

    l = (scope, elem, attr) ->

      scope.select = (ev)->
        popupService(
          "horodata/views/billing/plan_change.html",
          "BillingPlanChange"
          scope, ev)

    return {
      link: l
      scope:
        plan: "="
        current: "="
      restrict: "E"
      templateUrl: "horodata/views/billing/plan.html"
    }
])

angular.module("horodata").controller("BillingPlanChange", [
  "$scope"
  "$mdDialog"
  "$mdToast"
  "$http"
  "$location"
  "apiService"
  "groupService"
  ($scope, $mdDialog, $mdToast, $http, $location, apiService, groupService)->
    $scope.loading = false
    $scope.preLoading = true

    $scope.addr =
      current: null
      loaded: false

    $scope.card =
      current: null
      loaded: false

    getAddr = ->
      $http.get("#{apiService.get()}/billing/address").then(
        (resp) ->
          $scope.addr.loaded = true
          $scope.addr.current = resp.data.data
        (resp) -> $scope.card.loaded = true
      )
    getAddr()

    getCard = ->
      $http.get("#{apiService.get()}/billing/card").then(
        (resp) ->
          $scope.card.loaded = true
          $scope.card.current = resp.data.data
        (resp) -> $scope.card.loaded = true
      )
    getCard()

    w = $scope.$watchGroup(["addr.loaded", "card.loaded"], (v) ->
      console.log v
      if !v[0] || !v[1] then return
      $scope.preLoading = false
      w()
    )
])

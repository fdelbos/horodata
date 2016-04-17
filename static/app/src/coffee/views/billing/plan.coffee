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
        recap: "="
        end: "="
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
    $scope.ready = false
    $scope.errors = null
    $scope.done = false

    $scope.addr =
      current: null
      loaded: false

    $scope.card =
      current: null
      loaded: false

    $scope.period =
      end: null
      loaded: false

    getAddr = ->
      $http.get("#{apiService.get()}/billing/address").then(
        (resp) ->
          $scope.addr.loaded = true
          $scope.addr.current = resp.data.data
        (resp) -> $scope.addr.loaded = true
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

    if $scope.plan.code == "free"
      getEndPeriod = ->
        $http.get("#{apiService.get()}/billing/end_period").then(
          (resp) ->
            $scope.period.loaded = true
            $scope.period.end = resp.data.data.end
          (resp) -> $scope.period.loaded = true
        )
      getEndPeriod()

      w = $scope.$watchGroup(["addr.loaded", "card.loaded", "period.loaded"], (v) ->
        if !v[0] || !v[1] || !v[2] then return
        $scope.preLoading = false
        $scope.ready = $scope.addr.current? && $scope.card.current?
        w()
      )

    else
      w = $scope.$watchGroup(["addr.loaded", "card.loaded"], (v) ->
        if !v[0] || !v[1] then return
        $scope.preLoading = false
        $scope.ready = $scope.addr.current? && $scope.card.current?
        w()
      )

    $scope.form =
      accept: false

    $scope.validate = ->
      if $scope.plan.code != "free" and $scope.form.accept == false
        $scope.errors =
          accept: "Vous devez cocher cette case pour continuer."
        return
      $scope.loading = true
      $http.post("#{apiService.get()}/billing/change_plan", {plan: $scope.plan.code}).then(
        (resp) ->
          $scope.loading = false
          $scope.done = true
          $mdToast.showSimple("Abonnement mis a jour")
          # if $scope.plan.code != "free" then $scope.current = $scope.plan.code
          $scope.$emit("plan.reload")
        (resp) ->
          $scope.loading = false
      )

])

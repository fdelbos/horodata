angular.module("horodata").directive("billingCard", [
  "popupService"
  "$http"
  "apiService"
  (popupService, $http, apiService)->

    l = (scope, elem, attr) ->

      scope.editCard = (ev)->
        popupService(
          "horodata/views/billing/card_edit.html",
          "BillingCardEdit"
          scope, ev)

      scope.card =
        current: null

      get = ->
        scope.loading = true
        $http.get("#{apiService.get()}/billing/card").then(
          (resp) ->
            scope.loading = false
            scope.card.current = resp.data.data
          (resp) ->
            scope.loading = false
            scope.card.current = null
        )

      get()

    return {
      link: l
      restrict: "E"
      templateUrl: "horodata/views/billing/card.html"
    }
])

angular.module("horodata").controller("BillingCardEdit", [
  "$scope"
  "$mdDialog"
  "$mdToast"
  "$http"
  "apiService"
  "stripeService"
  ($scope, $mdDialog, $mdToast, $http, apiService, stripeService)->
    $scope.loading = false
    $scope.errors = null


    stripeService.get((key) -> Stripe.setPublishableKey(key))

    $scope.months = []
    current = 0
    for i in moment.months()
      current += 1
      $scope.months.push({id: current, label: i})

    $scope.years = [moment().year()..(moment().year() + 30)]

    $scope.card.new =
      number: null
      cvc: null
      exp_month: null
      exp_year: null

    $scope.update = ->
      $scope.loading = true

      $scope.errors = {}
      if !$scope.card.new.number? || $scope.card.new.number == ""
        $scope.errors.number = "Ce champ est obligatoire."
      if !$scope.card.new.cvc? || $scope.card.new.cvc == ""
        $scope.errors.cvc = "Ce champ est obligatoire."
      if !$scope.card.new.exp_month?
        $scope.errors.exp_month = "Ce champ est obligatoire."
      if !$scope.card.new.exp_year?
        $scope.errors.exp_year = "Ce champ est obligatoire."

      if !_.isEmpty($scope.errors)
        $scope.loading = false
        return

      Stripe.card.createToken({
        number: $scope.card.new.number,
        cvc: $scope.card.new.cvc,
        exp_month: $scope.card.new.exp_month,
        exp_year: $scope.card.new.exp_year}, (status, resp) =>
          if resp.error?
            $scope.loading = false
            $scope.errors[resp.error.param] = resp.error.message
            $scope.$apply()
            return

          $http.post("#{apiService.get()}/billing/card", {token: resp.id}).then(
            (resp) ->
              $mdDialog.hide()
              $mdToast.showSimple("Nouvelle carte de credit enregistree.")
              $scope.card.current = resp.data.data
              $scope.loading = false
            (resp) ->
              $scope.loading = false
          )
      )
])

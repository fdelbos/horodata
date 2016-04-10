angular.module("horodata").directive("billingAddr", [
  "popupService"
  "$http"
  "apiService"
  (popupService, $http, apiService)->

    l = (scope, elem, attr) ->

      scope.editAddr = (ev)->
        popupService(
          "horodata/views/billing/addr_edit.html",
          "BillingAddrEdit"
          scope, ev)

      scope.addr =
        current: null

      get = ->
        scope.loading = true
        $http.get("#{apiService.get()}/billing/address").then(
          (resp) ->
            scope.loading = false
            scope.addr.current = resp.data.data
          (resp) ->
            scope.loading = false
            scope.addr.current = null
        )

      get()

    return {
      link: l
      restrict: "E"
      templateUrl: "horodata/views/billing/addr.html"
    }
])

angular.module("horodata").controller("BillingAddrEdit", [
  "$scope"
  "$mdDialog"
  "$mdToast"
  "$http"
  "$location"
  "apiService"
  "groupService"
  ($scope, $mdDialog, $mdToast, $http, $location, apiService, groupService)->
    $scope.loading = false

    $scope.update = ->
      $scope.loading = true
      $http.post("#{apiService.get()}/billing/address", $scope.addr.current).then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Nouvelle adresse enregistree.")
          $scope.addr.current = resp.data.data
          $scope.loading = false
        (resp) ->
          $scope.errors = resp.data.errors
          $scope.loading = false
      )
])

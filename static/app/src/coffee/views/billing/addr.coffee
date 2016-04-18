angular.module("horodata").directive("billingAddr", [
  "popupService"
  "$http"
  "apiService"
  (popupService, $http, apiService)->

    l = (scope, elem, attr) ->

      scope.addr =
        current: null
        edit: null

      scope.editAddr = (ev)->
        if !scope.addr.current?
           scope.addr.edit = {}
        else
          scope.addr.edit = _.cloneDeep(scope.addr.current)
        popupService(
          "horodata/views/billing/addr_edit.html",
          "BillingAddrEdit"
          scope, ev)

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
    $scope.errors = null

    $scope.update = ->
      $scope.loading = true
      $http.post("#{apiService.get()}/billing/address", $scope.addr.edit).then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Nouvelle adresse enregistrée.")
          $scope.addr.current = resp.data.data
          $scope.loading = false
        (resp) ->
          $scope.errors = resp.data.errors
          $scope.loading = false
      )
])

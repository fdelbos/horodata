angular.module("horodata").directive("billingInvoices", [
  "$http"
  "apiService"
  ($http, apiService)->

    l = (scope, elem, attr) ->

      scope.url = "#{apiService.get()}/billing/invoices"

      scope.loading = true
      $http.get("#{apiService.get()}/billing/invoices").then(
        (resp) ->
          scope.invoices = resp.data.data
          scope.loading = false
        (resp) ->
          scope.loading = false
      )


    return {
      link: l
      restrict: "E"
      templateUrl: "horodata/views/billing/invoices.html"
    }
])

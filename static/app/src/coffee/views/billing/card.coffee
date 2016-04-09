angular.module("horodata").directive("billingCard", [
  "popupService"
  (popupService)->

    l = (scope, elem, attr) ->

    return {
      link: l
      restrict: "E"
      templateUrl: "horodata/views/billing/card.html"
    }
])

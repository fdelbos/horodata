angular.module("horodata").directive("billingPlan", [
  ->

    l = (scope, elem, attr) ->

    return {
      link: l
      scope:
        plan: "="
        current: "="
      restrict: "E"
      templateUrl: "horodata/views/billing/plan.html"
    }
])

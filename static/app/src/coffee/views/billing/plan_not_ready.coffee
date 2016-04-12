angular.module("horodata").directive("billingPlanNotReady", [
  ->
    return {
      restrict: "E"
      templateUrl: "horodata/views/billing/plan_not_ready.html"
    }
])

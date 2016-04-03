angular.module("horodata").directive("appWidgetsCommonQuotaError", [
  ->

    l = (scope) ->

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/common/quota_error.html"
    }
])

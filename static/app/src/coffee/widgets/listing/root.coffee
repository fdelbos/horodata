angular.module("horodata").directive("appWidgetsListing", [
  "listingService"
  "$timeout"
  "$location"
  "popupService"
  (listingService, $timeout, $location, popupService) ->

    l = (scope) ->

      scope.listing = listingService

      scope.goTo = (page) ->
        $location.search("page", page)
        listingService.listing().fetch(page)

      scope.showDetail = (ev, job) ->
        scope.detailJob = _.cloneDeep job
        popupService(
          "horodata/widgets/detail/dialog.html"
          "detailDialog"
          scope, ev)

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/listing/root.html"
    }
])

angular.module("horodata").directive("appWidgetsListing", [
  "listingService"
  "$timeout"
  "$location"
  (listingService, $timeout, $location) ->

    l = (scope) ->

      scope.listing = listingService

      scope.goTo = (page) ->
        $location.search("page", page)
        listingService.listing().fetch(page)

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/listing.html"
    }
])

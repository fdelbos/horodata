angular.module("horodata").directive("appWidgetsListing", [
  "listingService"
  "$timeout"
  (listingService, $timeout) ->

    l = (scope) ->

      scope.tasks = {}
      for i in scope.group.tasks
        scope.tasks[i.id] = i.name

      scope.customers = {}
      for i in scope.group.customers
        scope.customers[i.id] = i.name

      scope.listing = listingService.listing()
      scope.listing.fetch(0)


    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/listing.html"
    }
])

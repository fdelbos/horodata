angular.module("horodata").directive("appWidgetsListing", [
  "listingService"
  "$timeout"
  "$location"
  "popupService"
  (listingService, $timeout, $location, popupService) ->

    l = (scope) ->

      scope.tasks = {}
      for i in scope.group.tasks
        scope.tasks[i.id] = i.name

      scope.customers = {}
      for i in scope.group.customers
        scope.customers[i.id] = i.name

      scope.listing = listingService
      #scope.listing.fetch(0)

      scope.goTo = (page) ->
        $location.search("page", page)
        listingService.listing().fetch(page)

      scope.showDetail = (ev, job) ->
        scope.detailJob = _.cloneDeep job
        console.log job
        popupService(
          "horodata/widgets/detail.html"
          "detailDialog"
          scope, ev)

    return {
      link: l
      replace: true
      restrict: "E"
      templateUrl: "horodata/widgets/listing.html"
    }
])

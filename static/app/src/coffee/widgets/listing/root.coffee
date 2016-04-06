angular.module("horodata").directive("appWidgetsListing", [
  "listingService"
  "$timeout"
  "$location"
  "popupService"
  (listingService, $timeout, $location, popupService) ->

    l = (scope) ->

      scope.search =
        begin: moment().subtract(1, 'months').toDate()
        end: new Date()
        customer: null
        guest: null

      scope.$watch("search", (v) ->
        if !v? then return
        listingService.search(scope.group.url, v)
        scope.listing = listingService.get()
        scope.listing.reload()
      , true)

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

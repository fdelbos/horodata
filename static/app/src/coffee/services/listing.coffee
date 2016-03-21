angular.module('horodata').factory("listingService", [
  "$http"
  "apiService"
  ($http, apiService)->

    class Listing
      constructor: (@groupUrl, begin, end, customer, guest)->
        @size = 50
        @list = null
        @loading = false
        @total = -1
        @params =
          begin: moment(begin).format('YYYY-MM-DD')
          end: moment(end).format('YYYY-MM-DD')
          customer: customer
          guest: guest

      pages: (page) ->
        if @total == -1 then return []

      fetch: (page) =>
        if @loading then return
        @loading = true
        params = _.cloneDeep(@params)
        params.offset = page * @size
        params.size = @size
        $http.get("#{apiService.get()}/groups/#{@groupUrl}/jobs", {params: params}).then(
          (resp) =>
            @list = resp.data.data.results
            @loading = false
            @total = resp.data.data.total
          (resp) =>
            console.log resp.error
            @loading = false
        )

    listing = {}

    return {
        data: ->
          if !listing.list? || listing.loading then null
          listing.list
        listing: -> listing
        search: (groupUrl, params) ->
          listing = new Listing(groupUrl, params.begin, params.end, params.customer, params.guest)
    }
])

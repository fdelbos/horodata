angular.module('horodata').factory("listingService", [
  "$http"
  "apiService"
  ($http, apiService)->

    class Listing
      constructor: (@groupUrl, begin, end, customer, guest)->
        @size = 100
        @page = 1
        @list = []
        @loading = false
        @total = -1
        @params =
          begin: moment(begin).format('YYYY-MM-DD')
          end: moment(end).format('YYYY-MM-DD')
          customer: customer
          guest: guest

      hasMore: -> (@page * @size) < @total

      next: -> if @hasMore() then @_fetch(@page + 1)

      reload: ->
        @page = 1
        @list = []
        @_fetch(1)

      _fetch: (page) =>
        if @loading then return
        @loading = true
        params = _.cloneDeep(@params)
        if !page? then page = @page
        @page = page
        params.offset = (page - 1) * @size
        params.size = @size
        $http.get("#{apiService.get()}/groups/#{@groupUrl}/jobs", {params: params}).then(
          (resp) =>
            @list.push(i)  for i in resp.data.data.results
            @loading = false
            @total = resp.data.data.total
          (resp) =>
            console.log resp.error
            @loading = false
        )

    listingInstance = {}

    return {
        data: ->
          if !listing.list? || listing.loading then null
          listing.list
        get: -> listingInstance
        search: (groupUrl, params) ->
          listingInstance = new Listing(groupUrl, params.begin, params.end, params.customer, params.guest)
    }
])

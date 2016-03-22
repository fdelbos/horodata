angular.module('horodata').factory("listingService", [
  "$http"
  "apiService"
  ($http, apiService)->

    class Listing
      constructor: (@groupUrl, begin, end, customer, guest)->
        @size = 100
        @page = 1
        @list = null
        @loading = false
        @total = -1
        @params =
          begin: moment(begin).format('YYYY-MM-DD')
          end: moment(end).format('YYYY-MM-DD')
          customer: customer
          guest: guest

      pages: ->
        if @total == -1 then return null
        res =
          page: @page
          prev: if @page == 1 then null else @page - 1
          next: if @page * @size > @total then null else @page + 1
        return res


      fetch: (page) =>
        if @loading then return
        @loading = true
        params = _.cloneDeep(@params)
        if !page? then page = @page
        @page = page
        params.offset = (page - 1) * @size
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
        pages: -> listing.pages()
    }
])

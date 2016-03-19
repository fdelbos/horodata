angular.module('horodata').factory("listingService", [
  "$http"
  "apiService"
  ($http, apiService)->

    # class Listing
    #   constructor: (@groupUrl, begin, end, customer, guest)->
    #     @total = -1
    #     @list = []
    #     @loading = false
    #     @params =
    #       begin: moment(begin).format('YYYY-MM-DD')
    #       end: moment(end).format('YYYY-MM-DD')
    #       customer: customer
    #       guest: guest
    #
    #   fetch: ->
    #     if @loading then return
    #     @loading = true
    #     $http.get("#{apiService.get()}/groups/#{@groupUrl}/jobs", {params: @params}).then(
    #       (resp) =>
    #         for i in resp.data.data.results
    #           @list.push i
    #         @total = resp.data.data.total
    #         @loading = false
    #       (resp) =>
    #         console.log resp.error
    #         @loading = false
    #     )
    #
    #   getLength: ->
    #     if @total < 0 then return 100
    #     else return @total
    #
    #   getItemAtIndex: (idx) ->
    #     if idx > @total && @total > 0
    #       console.log "getItemAtIndex #{idx}, too large"
    #       return null
    #     if idx > @list.length
    #       console.log "getItemAtIndex #{idx}, need to fetch"
    #       if @loading
    #         console.log "already loading"
    #         return null
    #       else @fetch()
    #     else
    #       console.log "getItemAtIndex #{idx}, found!"
    #       console.log @list[idx]
    #       return @list[idx]
    #
    #   nextPage: =>
    #     console.log "next page"
    #     console.log @list
    #     @fetch()

    # class Listing
    #   constructor: (@groupUrl, begin, end, customer, guest)->
    #     @list = []
    #     @loading = false
    #     @params =
    #       begin: moment(begin).format('YYYY-MM-DD')
    #       end: moment(end).format('YYYY-MM-DD')
    #       customer: customer
    #       guest: guest
    #
    #   fetch: =>
    #     if @loading then return
    #     @loading = true
    #     params = _.cloneDeep(@params)
    #     params.offset = @list.length
    #     $http.get("#{apiService.get()}/groups/#{@groupUrl}/jobs", {params: params}).then(
    #       (resp) =>
    #         for i in resp.data.data.results
    #           @list.push i
    #         @loading = false
    #         console.log @list
    #       (resp) =>
    #         console.log resp.error
    #         @loading = false
    #     )
    #
    #   nextPage: =>
    #     console.log "next page"
    #     @fetch()


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
        listing: -> listing
        search: (groupUrl, params) ->
          listing = new Listing(groupUrl, params.begin, params.end, params.customer, params.guest)
    }
])

angular.module('horodata').factory("groupNewService", [
  "apiService"
  "$http"
   (apiService, $http)->

    callback = null

    groups = []

    fetchListing = ->
      $http.get("#{apiService.get()}/groups").then(
        (resp) ->
          groups = resp.data.data.results
      )

    return {
      set: (fn) -> callback = fn
      open: (ev) -> callback(ev)
      listing: ->
        groups
      fetch: ->
        fetchListing()
    }
])

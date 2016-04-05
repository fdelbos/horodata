angular.module('horodata').factory("groupService", [
  "apiService"
  "$http"
   (apiService, $http)->

    current = null
    groups = []

    fetchListing = ->
      $http.get("#{apiService.get()}/groups").then(
        (resp) ->
          groups = resp.data.data.results
      )

    return {
      set: (group) -> current = group
      get: -> current
      listing: ->
        groups
      fetch: ->
        fetchListing()
    }
])

angular.module('horodata').factory("stripeService", [
  "$http",
  "apiService",
  ($http, apiService) ->

    key = null
    promise = $http.get("#{apiService.get()}/billing/stripe_key")

    get = (cb) ->
      if !key?
        promise.then (payload) ->
          key = payload.data.data
          cb(key)
      else cb(key)

    return {
      get: get
    }
])

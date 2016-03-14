angular.module('horodata').factory("userService", [
  "$http",
  "apiService",
  ($http, apiService) ->

    user = null
    promise = $http.get("#{apiService.get()}/users/me")

    get = (cb) ->
      if !user?
        promise.then (payload) ->
          user = payload.data.data
          cb(user)
      else cb(user)

    update = (u) -> user = u

    return {
      get: get
      update: update
    }
])

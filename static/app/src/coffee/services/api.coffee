angular.module('horodata').factory("apiService", [
   ->
    root = $("api").attr("href")

    return {
      get: -> root
    }
])

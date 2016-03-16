angular.module('horodata').factory("apiService", [
   ->
    root = document.getElementsByTagName("api")[0].getAttribute("href")

    return {
      get: -> root
    }
])

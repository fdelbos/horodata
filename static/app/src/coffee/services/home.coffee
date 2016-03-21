angular.module('horodata').factory("homeService", [
   ->
    root = document.getElementsByTagName("home")[0].getAttribute("href")

    return {
      get: -> root
    }
])

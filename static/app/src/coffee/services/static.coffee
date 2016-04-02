angular.module('horodata').factory("staticService", [
   ->
    root = document.getElementsByTagName("static")[0].getAttribute("href")

    return {
      get: -> root
    }
])

angular.module('horodata').factory("titleService", [
   ->
    title = {
      title: ""
    }

    return {
      get: -> title
      set: (t) -> title.title = t
    }
])

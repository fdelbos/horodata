angular.module('horodata').factory("titleService", [
   ->
    title = {
      title: ""
      canEdit: false
    }

    return {
      get: -> title
      set: (t, canEdit = false) ->
        title.title = t
        title.canEdit = canEdit
    }
])

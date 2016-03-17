angular.module('horodata').factory("groupNewService", [
   ->
    callback = null

    return {
      set: (fn) -> callback = fn
      open: (ev) -> callback(ev)
    }
])

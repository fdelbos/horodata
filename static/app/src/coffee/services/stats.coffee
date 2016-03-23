angular.module('horodata').factory("statsService", [
  "apiService"
  "$http"
   (apiService, $http)->

    loading = false

    fetch = (group, stat, params, cb)->
      p =
        begin: moment(params.begin).format('YYYY-MM-DD')
        end: moment(params.end).format('YYYY-MM-DD')
        guest: params.guest


      loading = true
      $http.get("#{apiService.get()}/groups/#{group}/stats/#{stat}", {params: p}).then(
        (resp) ->
          loading = false
          cb(resp.data.data)
        (resp) ->
          cd(null)
          loading = false
      )

    return {
      fetch: (group, stat, params, cb)-> fetch(group, stat, params, cb)
      data: -> data
      loading: -> loading
    }
])

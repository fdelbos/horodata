angular.module('horodata').factory("statsService", [
  "apiService"
  "$http"
  "statsFilterService"
  (apiService, $http, statsFilterService)->

    loading = false

    fetch = (group, stat, cb)->

      begin = moment(statsFilterService.begin).format('YYYY-MM-DD')
      end = moment(statsFilterService.end).format('YYYY-MM-DD')

      loading = true
      $http.get("#{apiService.get()}/groups/#{group}/stats/#{stat}", {params: {begin: begin, end: end}}).then(
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

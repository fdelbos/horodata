angular.module('horodata').factory("statsFilterService", [
  ->

    begin = moment().subtract(1, 'months').toDate()
    end = new Date()

    urlParams = ->
      p = params()
      "?begin=#{p.begin}&end=#{p.end}"

    return {
      begin: begin
      end: end
      urlParams: urlParams
    }
])

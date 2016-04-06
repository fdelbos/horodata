angular.module('horodata').factory("statsFilterService", [
  ->

    begin = moment().subtract(1, 'months').toDate()
    end = new Date()

    urlParams = ->

      "?begin=#{moment(begin).format('YYYY-MM-DD')}&end=#{moment(end).format('YYYY-MM-DD')}"

    return {
      begin: begin
      end: end
      urlParams: urlParams
    }
])

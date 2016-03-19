angular.module("horodata").filter("Day", [
  -> return (input) -> moment(input).format('LL')
])

angular.module("horodata").filter("Ago", [
  -> return (input) -> moment(input).fromNow()
])

angular.module("horodata").filter("Duration", [
  -> return (input) ->
    d = moment.duration(input, 'seconds')
    minutes = d.minutes()
    if minutes == 0 then minutes = "00"
    hours = d.hours()
    return "#{hours}h#{minutes}"
])

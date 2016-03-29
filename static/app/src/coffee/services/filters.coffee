angular.module("horodata").filter("Day", [
  -> return (input) -> moment(input).format('LL')
])

angular.module("horodata").filter("Ago", [
  -> return (input) -> moment(input).fromNow()
])

angular.module("horodata").filter("Date", [
  -> return (input) -> moment(input).format('LLLL')
])

angular.module("horodata").filter("Duration", [
  -> return (input) ->
    d = moment.duration(input, 'seconds')
    minutes = d.minutes()
    if minutes < 10 then minutes = "0#{minutes}"
    hours = d.hours()
    return "#{hours}h#{minutes}"
])

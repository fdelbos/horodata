angular.module("horodata", ["ngMaterial", "ngRoute", "ngMessages"])

.config([
  "$mdDateLocaleProvider"
  "$mdThemingProvider"
  "$locationProvider"
  "$routeProvider"
  (
    $mdDateLocaleProvider
    $mdThemingProvider
    $locationProvider
    $routeProvider
  ) ->

    $mdThemingProvider.theme('default')
      .primaryPalette('blue')
      .accentPalette('pink')

    $locationProvider.html5Mode(true)

    $routeProvider
      .when("/",
        templateUrl: "horodata/views/index.html"
        controller: "Index")
      .when("/group/:group",
        templateUrl: "horodata/views/group.html"
        controller: "Group")

    months = ["Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre"]
    $mdDateLocaleProvider.month = months
    $mdDateLocaleProvider.days = ['dimanche', 'lundi', 'mardi', 'mercredi', 'jeudi', 'venredi', 'samedi']
    $mdDateLocaleProvider.shortDays = ['Di', 'Lu', 'Ma', 'Me', 'Je', 'Ve', 'Sa'];
    $mdDateLocaleProvider.firstDayOfWeek = 1;
    $mdDateLocaleProvider.msgCalendar = 'Calendrier';
    $mdDateLocaleProvider.msgOpenCalendar = 'Ouvrir le calendrier';
    $mdDateLocaleProvider.monthHeaderFormatter = (date) ->
      months[date.getMonth()] + ' ' + date.getFullYear()
])

.run([
  "$http"
  ($http)->
    $http.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest'

])

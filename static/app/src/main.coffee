angular.module("horodata", ["ngMaterial", "ngRoute", "ngMessages", "gridshore.c3js.chart", "lfNgMdFileInput"])

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
    $mdThemingProvider.setDefaultTheme('default')


    $locationProvider.html5Mode(true)

    $routeProvider
      .when("/",
        templateUrl: "horodata/views/index.html"
        controller: "Index")
      .when("/billing",
        templateUrl: "horodata/views/billing/root.html"
        controller: "Billing")
      .when("/:group",
        templateUrl: "horodata/views/group.html"
        controller: "Group")

    # Dates and calendar

    moment.locale('fr')

    months = ["Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre"]
    $mdDateLocaleProvider.month = months
    $mdDateLocaleProvider.days = ['dimanche', 'lundi', 'mardi', 'mercredi', 'jeudi', 'venredi', 'samedi']
    $mdDateLocaleProvider.shortDays = ['Di', 'Lu', 'Ma', 'Me', 'Je', 'Ve', 'Sa'];
    $mdDateLocaleProvider.firstDayOfWeek = 1;
    $mdDateLocaleProvider.msgCalendar = 'Calendrier';
    $mdDateLocaleProvider.msgOpenCalendar = 'Ouvrir le calendrier';
    $mdDateLocaleProvider.monthHeaderFormatter = (date) ->
      months[date.getMonth()] + ' ' + date.getFullYear()
    $mdDateLocaleProvider.parseDate = (dateString) ->
      if moment(dateString, 'L', true).isValid()
        return m.toDate()
      else return new Date(NaN)
    $mdDateLocaleProvider.formatDate = (date) -> moment(date).format('L')

])

.run([
  "$http"
  ($http)->
    $http.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest'

])

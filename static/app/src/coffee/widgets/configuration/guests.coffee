angular.module("horodata").directive("appWidgetsConfigurationGuests", [
  "popupService"
  (popupService)->

    l = (scope, elem, attr) ->
      scope.guests =
        current: null
        selected: null
        edit: (ev) ->
          @current = _.cloneDeep(_.find(scope.group.guests, {id: parseInt @selected}))
          popupService(
            "horodata/widgets/configuration/guests_edit_form.html",
            "appWidgetsConfigurationGuestsDialog"
            scope, ev)
        create: (ev) ->
          @current =
            email: ""
            admin: false
            rate: "0"
          popupService(
            "horodata/widgets/configuration/guests_create_form.html",
            "appWidgetsConfigurationGuestsDialog"
            scope, ev)

    return {
      link: l
      restrict: "E"
      templateUrl: "horodata/widgets/configuration/guests.html"
    }
])


angular.module("horodata").controller("appWidgetsConfigurationGuestsDialog", [
  "$scope"
  "$mdDialog"
  "$mdToast"
  "$http"
  "$location"
  "apiService"
  ($scope, $mdDialog, $mdToast, $http, $location, apiService)->
    $scope.errors = null
    $scope.loading = false
    $scope.quotaError = null

    update = (t) ->
      idx = _.findIndex($scope.group.guests, {id: t.id})
      $scope.group.guests[idx] = $scope.guests.current
      $scope.group.guests = _.sortBy($scope.group.guests, ["name"])

    $scope.create = ->
      $scope.loading = true
      $http.post("#{apiService.get()}/groups/#{$scope.group.url}/guests", $scope.guests.current).then(
        (resp) ->
          $scope.$emit("group.reload")
          $mdDialog.hide()
          $mdToast.showSimple("Nouvel utilisateur '#{$scope.guests.current.email}' ajouté")
        (resp) ->
          $scope.loading = false
          if resp.status == 429 && _.get(resp, "data.errors.type") == "quota"
            $scope.quotaError = resp.data.errors
          else $scope.errors = resp.data.errors
      )

    $scope.edit = ->
      $scope.loading = true
      # data =
      #   email: $scope.guests.current.email
      #   admin: $scope.guests.current.admin
      #   rate: $scope.guests.current.rate * 100
      $http.put("#{apiService.get()}/groups/#{$scope.group.url}/guests/#{ $scope.guests.selected }", $scope.guests.current).then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Utilisateur '#{$scope.guests.current.email}' modifié")
          update($scope.guests.current)
          $scope.guests.selected = null
        (resp) ->
          $scope.loading = false
          $scope.errors = resp.data.errors
      )

    $scope.delete = ->
      $scope.loading = true
      $http.delete("#{apiService.get()}/groups/#{$scope.group.url}/guests/#{ $scope.guests.selected }", $scope.guests.current).then(
        (resp) ->
          $mdDialog.hide()
          $mdToast.showSimple("Utilisateur '#{$scope.guests.current.email}' supprimé")
          $scope.group.guests.splice(_.findIndex($scope.group.guests, {id: parseInt $scope.guests.selected}), 1)
          $scope.guests.selected = null
        (resp) ->
          $scope.loading = false
          $scope.errors = resp.data.errors
      )
])

angular.module("horodata").directive("validDecimal", [
  ->

    l = (scope, element, attrs, ngModelCtrl) ->
      if !ngModelCtrl? then return

      ngModelCtrl.$parsers.push (v) ->
        if !v? then v = ''

        clean = v.replace(",", ".")
        clean = clean.replace(/[^0-9\.]/g, '')
        decimalCheck = clean.split('.')

        if decimalCheck[1]?
          decimalCheck[1] = decimalCheck[1].slice(0,2)
          clean = "#{decimalCheck[0]}.#{decimalCheck[1]}"

        if v != clean
          ngModelCtrl.$setViewValue(clean)
          ngModelCtrl.$render()

        return clean

      element.bind('keypress', (ev) ->
        if ev.keyCode == 32 then event.preventDefault()
      )

    return {
      link: l
      require: '?ngModel'
    }
])

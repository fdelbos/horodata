/**
 * HoroData Javascript Interface
 * Version: 0.0.1
 * Copyright © 2016 Hyperboloide. All rights reserved.
*/
angular.module("horodata", ["ngMaterial", "ngRoute", "ngMessages"]).config([
  "$mdDateLocaleProvider", "$mdThemingProvider", "$locationProvider", "$routeProvider", function($mdDateLocaleProvider, $mdThemingProvider, $locationProvider, $routeProvider) {
    var months;
    $mdThemingProvider.theme('default').primaryPalette('blue').accentPalette('pink');
    $mdThemingProvider.setDefaultTheme('default');
    $locationProvider.html5Mode(true);
    $routeProvider.when("/", {
      templateUrl: "horodata/views/index.html",
      controller: "Index"
    }).when("/group/:group", {
      templateUrl: "horodata/views/group.html",
      controller: "Group"
    });
    months = ["Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre"];
    $mdDateLocaleProvider.month = months;
    $mdDateLocaleProvider.days = ['dimanche', 'lundi', 'mardi', 'mercredi', 'jeudi', 'venredi', 'samedi'];
    $mdDateLocaleProvider.shortDays = ['Di', 'Lu', 'Ma', 'Me', 'Je', 'Ve', 'Sa'];
    $mdDateLocaleProvider.firstDayOfWeek = 1;
    $mdDateLocaleProvider.msgCalendar = 'Calendrier';
    $mdDateLocaleProvider.msgOpenCalendar = 'Ouvrir le calendrier';
    return $mdDateLocaleProvider.monthHeaderFormatter = function(date) {
      return months[date.getMonth()] + ' ' + date.getFullYear();
    };
  }
]).run([
  "$http", function($http) {
    return $http.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
  }
]);

angular.module("horodata").run(["$templateCache", function($templateCache) {$templateCache.put("horodata/menu/sidenav.html","<md-sidenav class=\"md-sidenav-left md-whiteframe-z2\" md-component-id=sidenav md-is-locked-open=\"$mdMedia(\'gt-sm\')\" layout=column><md-toolbar hide-gt-sm><div class=md-toolbar-tools><h1><md-icon class=md-24>access_time</md-icon>&nbsp; Horo Data</h1><span flex></span><md-button ng-click=toggleSidenav() class=md-icon-button><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-content layout-padding hide show-gt-sm><div layout=column layout-align=\"center center\"><div class=md-display-3><i class=material-icons style=\"font-size: 64px;\">access_time</i></div><div class=md-headline>Horo Data</div></div></md-content><md-content flex><section><md-subheader class=md-accent><div layout=row layout-align=\"space-between center\"><span class=md-headline>Groupes</span><app-widgets-new-group layout-align=end></app-widgets-new-group></div></md-subheader><md-list flex layout=column class=md-body-1><md-list-item ng-repeat=\"group in groups\" ng-href=\"./group/{{ group.url}}\">{{group.name}}</md-list-item></md-list><section></section></section></md-content></md-sidenav>");
$templateCache.put("horodata/menu/toolbar.html","<div><app-widgets-new-task></app-widgets-new-task><md-toolbar><div class=md-toolbar-tools><md-button hide-gt-sm class=md-icon-button aria-label=Settings ng-click=toggleSidenav()><md-icon class=md-24>menu</md-icon></md-button><h2><span>{{ MainTitle().title }}</span></h2><span flex></span><md-button class=\"md-fab md-mini\" aria-label=Favorite><md-icon class=md-24>account_circle</md-icon></md-button></div></md-toolbar></div>");
$templateCache.put("horodata/views/group.html","<div><app-widgets-search-bar></app-widgets-search-bar><md-content><md-tabs md-dynamic-height md-border-bottom><md-tab label=Heures><app-widgets-empty-group ng-if=\"group.tasks.length == 0 || group.customers.length == 0\" group=group><app-widgets-empty-group></app-widgets-empty-group></app-widgets-empty-group></md-tab><md-tab label=Statistiques><app-widgets-empty-group ng-if=\"group.tasks.length == 0 || group.customers.length == 0\" group=group><app-widgets-empty-group></app-widgets-empty-group></app-widgets-empty-group></md-tab><md-tab label=Configuration layout=column layout-margin><app-widgets-configuration-guests></app-widgets-configuration-guests><md-divider></md-divider><app-widgets-configuration-customers></app-widgets-configuration-customers><md-divider></md-divider><app-widgets-configuration-tasks></app-widgets-configuration-tasks></md-tab></md-tabs></md-content></div>");
$templateCache.put("horodata/views/index.html","<div><app-new-group></app-new-group></div>");
$templateCache.put("horodata/views/new_task_form.html","<md-dialog aria-label=\"Nouvelle Tâche\" flex=40><div ng-if=\"group.tasks.length == 0 || group.customers.length == 0\"><md-dialog-content><div class=md-dialog-content><div ng-if=\"group.tasks.length == 0 && group.customers.length == 0\"><div class=\"md-headline input-error\">Aucune tâches ni aucun dossier enregistre</div><p>Vous devez definir des tâches et des dossiers pour la saisie. Cliquez sur l\'onglet configuration et ajoutez les.</p></div><div ng-if=\"group.tasks.length > 0 && group.customers.length == 0\"><div class=\"md-headline input-error\">Aucun dossier enregistre</div><p>Vous devez definir des dossiers pour la saisie. Cliquez sur l\'onglet configuration et ajoutez les.</p></div><div ng-if=\"group.tasks.length == 0 && group.customers.length > 0\"><div class=\"md-headline input-error\">Aucune tâches definie</div><p>Vous devez definir des tâches pour la saisie. Cliquez sur l\'onglet configuration et ajoutez les.</p></div></div></md-dialog-content><md-dialog-actions><md-button ng-click=close() class=md-raised>Fermer</md-button></md-dialog-actions></div><form ng-if=\"group.tasks.length > 0 && group.customers.length > 0\" name=newTaskForm><md-toolbar><div class=md-toolbar-tools><h2>Nouvelle Tâche</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><div layout=column><md-input-container flex><label>Selectionnez un dossier</label><md-select ng-model=task.customer><md-option ng-repeat=\"c in group.customers\" value=\"{{ c.id }}\">{{ c.name }}</md-option></md-select></md-input-container><md-input-container flex><label>Selectionnez une tâche</label><md-select ng-model=task.task><md-option ng-repeat=\"t in group.tasks\" value=\"{{ t.id }}\">{{ t.name }}</md-option></md-select></md-input-container><div layout=row><md-input-container flex><label>Heures</label><md-select ng-model=task.hours><md-option ng-repeat=\"h in hours\" value=\"{{ h }}\">{{ h }} <span ng-if=\"h > 1\">Heures</span> <span ng-if=\"h <= 1\">Heure</span></md-option></md-select></md-input-container><md-input-container flex><label>Minutes</label><md-select ng-model=task.minutes><md-option ng-repeat=\"t in group.tasks\" value=\"{{ t.id }}\">{{ t.name }}</md-option></md-select></md-input-container></div><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.comment}\"><label>Commentaire</label> <textarea md-no-autogrow ng-model=task.comment rows=5 md-select-on-focus></textarea> <small ng-if=errors.comment class=input-error>{{ errors.comment }}</small></md-input-container></div></div></md-dialog-content><md-dialog-actions><md-button ng-click=send() class=\"md-primary md-raised\">Creer</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/empty_group.html","<div layout-align=\"center center\" layout=row><md-whiteframe class=\"md-whiteframe-1dp md-accent\" flex=100 flex-sm=70 flex-md=60 flex-gt-md=50 layout-padding layout-margin layout=column layout-align=\"center center\"><div ng-if=\"group.tasks.length == 0 && group.customers.length == 0\"><div class=\"md-headline input-error\">Aucune tâches ni aucun dossier enregistre</div><p>Vous devez definir des tâches et des dossiers pour la saisie. Cliquez sur l\'onglet configuration et ajoutez les.</p></div><div ng-if=\"group.tasks.length > 0 && group.customers.length == 0\"><div class=\"md-headline input-error\">Aucun dossier enregistre</div><p>Vous devez definir des dossiers pour la saisie. Cliquez sur l\'onglet configuration et ajoutez les.</p></div><div ng-if=\"group.tasks.length == 0 && group.customers.length > 0\"><div class=\"md-headline input-error\">Aucune tâches definie</div><p>Vous devez definir des tâches pour la saisie. Cliquez sur l\'onglet configuration et ajoutez les.</p></div></md-whiteframe></div>");
$templateCache.put("horodata/widgets/new_group.html","<md-button class=\"md-fab md-primary md-mini\" ng-click=showNewGroupDialog($event) aria-label=\"Nouveau Groupe\"><md-icon class=md-24>add</md-icon><md-tooltip md-direction=top>Ajouter un groupe</md-tooltip></md-button>");
$templateCache.put("horodata/widgets/new_group_form.html","<md-dialog aria-label=\"Nouveau Groupe\" flex=40><form name=newGroupForm><md-toolbar><div class=md-toolbar-tools><h2>Nouveau Groupe</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.name}\"><label>Nom du Groupe</label> <input type=text ng-model=name> <small ng-if=errors.name class=input-error>{{ errors.name }}</small></md-input-container></div></md-dialog-content><md-dialog-actions><md-button ng-click=send() class=\"md-primary md-raised\">Creer</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/new_task.html","<md-button ng-if=MainTitle().canEdit class=\"md-fab md-fab-bottom-right\" aria-label=Add ng-click=showNewTaskDialog($event)><md-tooltip md-direction=left>Ajouter un tâche</md-tooltip><md-icon class=md-48 style=\"margin: -2px 0px 0px -1px\">access_time</md-icon></md-button>");
$templateCache.put("horodata/widgets/search_bar.html","<div><div layout=row layout-align=\"space-around center\" hide show-gt-md layout-padding><md-datepicker flex=25 ng-model=beginDate md-placeholder=\"Date debut\" md-max-date=endDate></md-datepicker><md-datepicker flex=25 ng-model=endDate md-placeholder=\"Date fin\" md-max-date=maxDate></md-datepicker><md-input-container flex=25><label>Dossier</label><md-select ng-model=customer><md-option ng-repeat=\"c in group.customers\" value=\"{{ c.id }}\">{{ c.name }}</md-option></md-select></md-input-container><md-input-container flex=25><label>Utilisateur</label><md-select ng-model=user><md-option ng-repeat=\"u in users\" value=\"{{ u.email }}\">{{ u.name }}</md-option></md-select></md-input-container></div><md-divider></md-divider></div>");
$templateCache.put("horodata/widgets/configuration/customers.html","<div layout-margin layout-padding><div layout=column flex-sm=70 flex-md=80 flex-gt-md=60><div layout=row layout-align=\"space-between center\"><div class=md-display-1>Dossiers</div><md-button ng-click=customers.create($event) class=\"md-raised md-primary\" hide show-gt-sm><md-icon class=md-24>add</md-icon>Ajouter</md-button><md-button ng-click=customers.create($event) class=\"md-fab md-primary\" hide-gt-sm aria-label=\"Ajouter un dossier\"><md-tooltip md-direction=top>Ajouter un Dossier</md-tooltip><md-icon class=md-36>add</md-icon></md-button></div><p class=md-body-1>Un dossier peut representer un client ou un projet. Chaque nouvelle tache est attachee a un dossier.</p><md-whiteframe ng-if=\"group.customers.length == 0\" class=\"md-whiteframe-1dp md-accent\" layout-padding layout=column layout-align=\"center center\"><div class=\"md-headline input-error\">Aucune dossier enregistre</div><div>Vous devez definir des dossiers pour la saisie.</div></md-whiteframe><div layout=row ng-if=\"group.customers.length > 0\"><md-input-container flex><label>Selectionnez un dossiers pour l\'editer</label><md-select ng-model=customers.selected><md-option ng-repeat=\"c in group.customers\" value=\"{{ c.id }}\">{{ c.name }}</md-option></md-select></md-input-container><md-button ng-if=customers.selected ng-click=customers.edit($event) class=\"md-fab md-accent\" aria-label=\"Editer un dossier\"><md-tooltip md-direction=top>Modifier le dossier</md-tooltip><md-icon class=md-36>edit</md-icon></md-button></div></div></div>");
$templateCache.put("horodata/widgets/configuration/customers_create_form.html","<md-dialog aria-label=\"Nouveau Dossier\" flex=40><form name=newCustomerForm><md-toolbar><div class=md-toolbar-tools><h2>Nouveaux Dossiers</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.customers}\"><label>Noms des Dossiers a ajouter (saisissez un dossier par ligne).</label> <textarea md-no-autogrow ng-model=customers.current.customers rows=5 md-select-on-focus></textarea> <small ng-if=errors.customers class=input-error>{{ errors.customers }}</small></md-input-container></div></md-dialog-content><md-dialog-actions><md-button ng-click=create() class=\"md-primary md-raised\">Ajouter</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/configuration/customers_edit_form.html","<md-dialog aria-label=\"Modifier un dossier\" flex=40><form name=editCustomerForm><md-toolbar><div class=md-toolbar-tools><h2>Modifier un dossier</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.name}\"><label>Nom du dossier</label> <input md-maxlength=40 type=text ng-model=customers.current.name> <small ng-if=errors.name class=input-error>{{ errors.name }}</small></md-input-container></div></md-dialog-content><md-dialog-actions><md-button ng-click=edit() class=\"md-primary md-raised\">Editer</md-button><md-button ng-click=delete() class=\"md-warn md-raised\">Supprimer</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/configuration/guests.html","<div layout-margin layout-padding><div layout=column flex-sm=70 flex-md=80 flex-gt-md=60><div layout=row layout-align=\"space-between center\"><div class=md-display-1>Utilisateurs</div><md-button ng-click=guests.create($event) class=\"md-raised md-primary\" hide show-gt-sm><md-icon class=md-24>add</md-icon>Ajouter</md-button><md-button ng-click=guests.create($event) class=\"md-fab md-primary\" hide-gt-sm aria-label=\"Ajouter une tâche\"><md-tooltip md-direction=top>Ajouter un utilisateur</md-tooltip><md-icon class=md-36>add</md-icon></md-button></div><p class=md-body-1>Pour que vos collaborateurs puissent enregister leur travail, vous devez les inviter.</p><div layout=row ng-if=\"group.guests.length > 0\"><md-input-container flex><label>Selectionnez un utilisateur pour l\'editer</label><md-select ng-model=guests.selected><md-option ng-repeat=\"g in group.guests\" value=\"{{ g.id }}\"><span ng-if=!g.name>{{ g.email }}</span> <span ng-if=g.name>{{ g.name }} <small>{{ g.email }}</small></span></md-option></md-select></md-input-container><md-button ng-if=guests.selected ng-click=guests.edit($event) class=\"md-fab md-accent\" aria-label=\"Editer une tâche\"><md-tooltip md-direction=top>Editer l\'utilisateur</md-tooltip><md-icon class=md-36>edit</md-icon></md-button></div></div></div>");
$templateCache.put("horodata/widgets/configuration/guests_create_form.html","<md-dialog aria-label=\"Nouvel Utilisateur\" flex=40><form name=newUserForm><md-toolbar><div class=md-toolbar-tools><h2>Nouvel Utilisateur</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.email}\"><label>Adresse Email</label> <input type=text md-maxlength=100 ng-model=guests.current.email> <small ng-if=errors.email class=input-error>{{ errors.email }}</small></md-input-container><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.rate}\"><label>Taux Horraire</label> <input type=number ng-model=guests.current.rate ng-value=0> <small ng-if=errors.rate class=input-error>{{ errors.rate }}</small></md-input-container><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.admin}\"><md-switch ng-model=guests.current.admin aria-label=Administrateur>Administrateur</md-switch><div class=input-hint>Cochez cette case si vous souhaitez deleguer la gestion de ce group a cet utilisateur.</div><small ng-if=errors.admin class=input-error>{{ errors.admin }}</small></md-input-container></div></md-dialog-content><md-dialog-actions><md-button ng-click=create() class=\"md-primary md-raised\">Ajouter</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/configuration/guests_edit_form.html","<md-dialog aria-label=\"Nouvel Utilisateur\" flex=40><form name=newUserForm><md-toolbar><div class=md-toolbar-tools><h2>Mdifier l\' utilisateur {{ guests.current.email }}</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.rate}\"><label>Taux Horraire</label> <input type=number ng-model=guests.current.rate ng-value=0> <small ng-if=errors.rate class=input-error>{{ errors.rate }}</small></md-input-container><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.admin}\"><md-switch ng-model=guests.current.admin aria-label=Administrateur>Administrateur</md-switch><div class=input-hint>Cochez cette case si vous souhaitez deleguer la gestion de ce group a cet utilisateur.</div><small ng-if=errors.admin class=input-error>{{ errors.admin }}</small></md-input-container></div></md-dialog-content><md-dialog-actions><md-button ng-click=edit() class=\"md-primary md-raised\">Editer</md-button><md-button ng-click=delete() class=\"md-warn md-raised\">Supprimer</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/configuration/tasks.html","<div layout-margin layout-padding><div layout=column flex-sm=70 flex-md=80 flex-gt-md=60><div layout=row layout-align=\"space-between center\"><div class=md-display-1>Types de Tâches</div><md-button ng-click=tasks.create($event) class=\"md-raised md-primary\" hide show-gt-sm><md-icon class=md-24>add</md-icon>Ajouter</md-button><md-button ng-click=tasks.create($event) class=\"md-fab md-primary\" hide-gt-sm aria-label=\"Ajouter une tâche\"><md-tooltip md-direction=top>Ajouter une tâche</md-tooltip><md-icon class=md-36>add</md-icon></md-button></div><p class=md-body-1>Pour categoriser chaque tâche, les utilisateurs doivent choisir un \"type de tâche\". Chaque type de tâche peut etre associe a un commentaire pour plus de precisions.</p><md-whiteframe ng-if=\"group.tasks.length == 0\" class=\"md-whiteframe-1dp md-accent\" layout-padding layout=column layout-align=\"center center\"><div class=\"md-headline input-error\">Aucune tâches definie</div><div>Vous devez definir des tâches pour la saisie.</div></md-whiteframe><div layout=row ng-if=\"group.tasks.length > 0\"><md-input-container flex><label>Selectionnez un type de tâche pour l\'editer</label><md-select ng-model=tasks.selected><md-option ng-repeat=\"t in group.tasks\" value=\"{{ t.id }}\">{{ t.name }} <small ng-if=t.comment_mandatory>(Commentaire obligatoire)</small></md-option></md-select></md-input-container><md-button ng-if=tasks.selected ng-click=tasks.edit($event) class=\"md-fab md-accent\" aria-label=\"Editer une tâche\"><md-tooltip md-direction=top>Editer la tâche</md-tooltip><md-icon class=md-36>edit</md-icon></md-button></div></div></div>");
$templateCache.put("horodata/widgets/configuration/tasks_form.html","<md-dialog aria-label=\"Nouveau type de tâche\" flex=40><form name=newTaskForm><md-toolbar><div class=md-toolbar-tools><h2>Nouveau type de tâche</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.name}\"><label>Nom de la tâche</label> <input md-maxlength=30 type=text ng-model=tasks.current.name> <small ng-if=errors.name class=input-error>{{ errors.name }}</small></md-input-container><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.admin}\"><md-switch ng-model=tasks.current.comment_mandatory aria-label=Administrateur>Commentaire obligatoire</md-switch><div class=input-hint>Force l\'utilisateur a associer un commentaire a la chaque tâche de ce type.</div></md-input-container></div></md-dialog-content><md-dialog-actions><md-button ng-if=!tasks.current.id ng-click=create() class=\"md-primary md-raised\">Ajouter</md-button><md-button ng-if=tasks.current.id ng-click=edit() class=\"md-primary md-raised\">Editer</md-button><md-button ng-if=tasks.current.id ng-click=delete() class=\"md-warn md-raised\">Supprimer</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");}]);
angular.module("horodata").directive("appMenuSidenav", [
  "$mdSidenav", "$http", "apiService", function($mdSidenav, $http, apiService) {
    var l;
    l = function(scope, elem) {
      scope.toggleSidenav = function() {
        return $mdSidenav("sidenav").toggle();
      };
      return $http.get((apiService.get()) + "/groups").then(function(resp) {
        return scope.groups = resp.data.data.results;
      });
    };
    return {
      link: l,
      replace: true,
      restrict: "E",
      templateUrl: "horodata/menu/sidenav.html"
    };
  }
]);

angular.module("horodata").directive("appMenuToolbar", [
  "titleService", function(titleService) {
    var l;
    l = function(scope, elem) {
      return scope.MainTitle = titleService.get;
    };
    return {
      link: l,
      replace: true,
      restrict: "E",
      templateUrl: "horodata/menu/toolbar.html"
    };
  }
]);

angular.module('horodata').factory("apiService", [
  function() {
    var root;
    root = document.getElementsByTagName("api")[0].getAttribute("href");
    return {
      get: function() {
        return root;
      }
    };
  }
]);

angular.module('horodata').factory("groupNewService", [
  function() {
    var callback;
    callback = null;
    return {
      set: function(fn) {
        return callback = fn;
      },
      open: function(ev) {
        return callback(ev);
      }
    };
  }
]);

angular.module('horodata').factory("popupService", [
  "$mdMedia", "$mdDialog", function($mdMedia, $mdDialog) {
    return function(tmpl, ctrl, scope, ev) {
      var fullscreen;
      fullscreen = $mdMedia('xs') || $mdMedia('sm');
      return $mdDialog.show({
        controller: ctrl,
        templateUrl: tmpl,
        parent: angular.element(document.body),
        targetEvent: ev,
        preserveScope: true,
        scope: scope,
        clickOutsideToClose: true,
        escapeToClose: true,
        fullscreen: fullscreen
      });
    };
  }
]);

angular.module('horodata').factory("searchUsersService", [
  function() {
    var Search;
    Search = (function() {
      function Search(items1) {
        var i, j, obj, ref;
        this.items = items1;
        this.selected = null;
        this.searchText = "";
        this.index = lunr(function() {
          this.field('name');
          this.field('email');
          return this.ref('index');
        });
        for (i = j = 0, ref = this.items.length - 1; 0 <= ref ? j <= ref : j >= ref; i = 0 <= ref ? ++j : --j) {
          obj = {
            name: this.items[i].name,
            email: this.items[i].email,
            index: i
          };
          this.items[i] = obj;
          this.index.add(obj);
        }
      }

      Search.prototype.search = function(query) {
        var data, i, j, len, res;
        if (query === "") {
          return this.items;
        }
        res = this.index.search(query);
        data = [];
        for (j = 0, len = res.length; j < len; j++) {
          i = res[j];
          data.push(this.items[i.ref]);
        }
        return data;
      };

      Search.prototype.select = function(item) {
        return this.selected = item;
      };

      return Search;

    })();
    return {
      get: function(items) {
        return new Search(items);
      }
    };
  }
]);

angular.module('horodata').factory("titleService", [
  function() {
    var title;
    title = {
      title: "",
      canEdit: false
    };
    return {
      get: function() {
        return title;
      },
      set: function(t, canEdit) {
        if (canEdit == null) {
          canEdit = false;
        }
        title.title = t;
        return title.canEdit = canEdit;
      }
    };
  }
]);

angular.module('horodata').factory("userService", [
  "$http", "apiService", function($http, apiService) {
    var get, promise, update, user;
    user = null;
    promise = $http.get((apiService.get()) + "/users/me");
    get = function(cb) {
      if (user == null) {
        return promise.then(function(payload) {
          user = payload.data.data;
          return cb(user);
        });
      } else {
        return cb(user);
      }
    };
    update = function(u) {
      return user = u;
    };
    return {
      get: get,
      update: update
    };
  }
]);

angular.module("horodata").controller("Group", [
  "$http", "$routeParams", "$scope", "titleService", "userService", "apiService", "groupNewService", "popupService", function($http, $routeParams, $scope, titleService, userService, apiService, groupNewService, popupService) {
    var fetchUsers, getGroup;
    $scope.isGroupView = true;
    $scope.maxDate = new Date();
    $scope.endDate = $scope.maxDate;
    $scope.beginDate = moment().subtract(1, 'months').toDate();
    fetchUsers = function() {
      return $http.get((apiService.get()) + "/groups/" + $routeParams.group + "/users").then(function(resp) {
        return $scope.users = resp.data.data;
      });
    };
    $scope.isOwner = false;
    getGroup = function() {
      return $http.get((apiService.get()) + "/groups/" + $routeParams.group).then(function(resp) {
        $scope.group = resp.data.data;
        if ($scope.group.owner.login === $scope.user.login) {
          $scope.isOwner = true;
        }
        return titleService.set($scope.group.name, true);
      });
    };
    userService.get(function(u) {
      $scope.user = u;
      getGroup();
      if ($scope.isOwner) {
        return fetchUsers();
      }
    });
    $scope.$on("group.reload", function(e) {
      e.stopPropagation();
      return getGroup();
    });
    return groupNewService.set(function(ev) {
      return popupService("horodata/views/new_task_form.html", "groupNewTaskDialog", $scope, ev);
    });
  }
]);

angular.module("horodata").controller("groupNewTaskDialog", [
  "$scope", "$mdDialog", "$mdToast", "$http", "$location", "apiService", function($scope, $mdDialog, $mdToast, $http, $location, apiService) {
    $scope.task = {};
    $scope.errors = null;
    $scope.close = function() {
      return $mdDialog.hide();
    };
    return $scope.send = function() {
      return $http.post((apiService.get()) + "/groups", {
        name: $scope.name
      }).then(function(resp) {
        var group;
        group = resp.data.data;
        $mdDialog.hide();
        $mdToast.showSimple("Nouveau groupe '" + group.name + "' sauvegarde.");
        return $location.path("/group/" + group.url);
      }, function(resp) {
        return $scope.errors = resp.data.errors;
      });
    };
  }
]);

angular.module("horodata").controller("Index", [
  "$http", "$scope", "userService", "titleService", function($http, $scope, userService, titleService) {
    return titleService.set("Choisissez un Groupe");
  }
]);

angular.module("horodata").directive("appWidgetsEmptyGroup", [
  function() {
    return {
      restrict: "E",
      templateUrl: "horodata/widgets/empty_group.html"
    };
  }
]);

angular.module("horodata").directive("appWidgetsNewGroup", [
  "popupService", function(popupService) {
    var l;
    l = function(scope) {
      return scope.showNewGroupDialog = function(ev) {
        return popupService("horodata/widgets/new_group_form.html", "newGroupDialog", scope, ev);
      };
    };
    return {
      link: l,
      restrict: "E",
      templateUrl: "horodata/widgets/new_group.html"
    };
  }
]);

angular.module("horodata").controller("newGroupDialog", [
  "$scope", "$mdDialog", "$mdToast", "$http", "$location", "apiService", function($scope, $mdDialog, $mdToast, $http, $location, apiService) {
    $scope.name = "";
    $scope.errors = null;
    $scope.close = function() {
      return $mdDialog.hide();
    };
    return $scope.send = function() {
      return $http.post((apiService.get()) + "/groups", {
        name: $scope.name
      }).then(function(resp) {
        var group;
        group = resp.data.data;
        $mdDialog.hide();
        $mdToast.showSimple("Nouveau groupe '" + group.name + "' sauvegarde.");
        return $location.path("/group/" + group.url);
      }, function(resp) {
        return $scope.errors = resp.data.errors;
      });
    };
  }
]);

angular.module("horodata").directive("appWidgetsNewTask", [
  "groupNewService", function(groupNewService) {
    var l;
    l = function(scope) {
      return scope.showNewTaskDialog = function(ev) {
        return groupNewService.open(ev);
      };
    };
    return {
      link: l,
      restrict: "E",
      templateUrl: "horodata/widgets/new_task.html"
    };
  }
]);

angular.module("horodata").directive("appWidgetsSearchBar", [
  function() {
    return {
      replace: true,
      restrict: "E",
      templateUrl: "horodata/widgets/search_bar.html"
    };
  }
]);

angular.module("horodata").directive("appWidgetsConfigurationCustomers", [
  "popupService", function(popupService) {
    var l;
    l = function(scope, elem, attr) {
      return scope.customers = {
        current: null,
        selected: null,
        edit: function(ev) {
          this.current = _.cloneDeep(_.find(scope.group.customers, {
            id: parseInt(this.selected)
          }));
          return popupService("horodata/widgets/configuration/customers_edit_form.html", "appWidgetsConfigurationCustomersDialog", scope, ev);
        },
        create: function(ev) {
          this.current = {
            customers: ""
          };
          return popupService("horodata/widgets/configuration/customers_create_form.html", "appWidgetsConfigurationCustomersDialog", scope, ev);
        }
      };
    };
    return {
      link: l,
      restrict: "E",
      templateUrl: "horodata/widgets/configuration/customers.html"
    };
  }
]);

angular.module("horodata").controller("appWidgetsConfigurationCustomersDialog", [
  "$scope", "$mdDialog", "$mdToast", "$http", "$location", "apiService", function($scope, $mdDialog, $mdToast, $http, $location, apiService) {
    var update;
    $scope.errors = null;
    $scope.close = function() {
      return $mdDialog.hide();
    };
    update = function(t) {
      var idx;
      idx = _.findIndex($scope.group.customers, {
        id: t.id
      });
      $scope.group.customers[idx] = $scope.customers.current;
      return $scope.group.customers = _.sortBy($scope.group.customers, ["name"]);
    };
    $scope.create = function() {
      return $http.post((apiService.get()) + "/groups/" + $scope.group.url + "/customers", $scope.customers.current).then(function(resp) {
        var total;
        total = resp.data.data.total;
        $scope.$emit("group.reload");
        $mdDialog.hide();
        if (total === 1) {
          return $mdToast.showSimple("1 nouveau dossier a été ajouté.");
        } else {
          return $mdToast.showSimple(total + " nouveaux dossiers ont été ajoutés.");
        }
      }, function(resp) {
        return $scope.errors = resp.data.errors;
      });
    };
    $scope.edit = function() {
      return $http.put((apiService.get()) + "/groups/" + $scope.group.url + "/customers/" + $scope.customers.selected, $scope.customers.current).then(function(resp) {
        $mdDialog.hide();
        $mdToast.showSimple("Dossier: '" + $scope.customers.current.name + "' mis a jour.");
        update($scope.customers.current);
        return $scope.customers.selected = null;
      }, function(resp) {
        return $scope.errors = resp.data.errors;
      });
    };
    return $scope["delete"] = function() {
      return $http["delete"]((apiService.get()) + "/groups/" + $scope.group.url + "/customers/" + $scope.customers.selected, $scope.customers.current).then(function(resp) {
        $mdDialog.hide();
        $mdToast.showSimple("Dossier: '" + $scope.customers.current.name + "' supprimé.");
        $scope.group.customers.splice(_.findIndex($scope.group.customers, {
          id: parseInt($scope.customers.selected)
        }), 1);
        return $scope.customers.selected = null;
      }, function(resp) {
        return $scope.errors = resp.data.errors;
      });
    };
  }
]);

angular.module("horodata").directive("appWidgetsConfigurationGuests", [
  "popupService", function(popupService) {
    var l;
    l = function(scope, elem, attr) {
      return scope.guests = {
        current: null,
        selected: null,
        edit: function(ev) {
          this.current = _.cloneDeep(_.find(scope.group.guests, {
            id: parseInt(this.selected)
          }));
          return popupService("horodata/widgets/configuration/guests_edit_form.html", "appWidgetsConfigurationGuestsDialog", scope, ev);
        },
        create: function(ev) {
          this.current = {
            email: "",
            admin: false,
            rate: 0
          };
          return popupService("horodata/widgets/configuration/guests_create_form.html", "appWidgetsConfigurationGuestsDialog", scope, ev);
        }
      };
    };
    return {
      link: l,
      restrict: "E",
      templateUrl: "horodata/widgets/configuration/guests.html"
    };
  }
]);

angular.module("horodata").controller("appWidgetsConfigurationGuestsDialog", [
  "$scope", "$mdDialog", "$mdToast", "$http", "$location", "apiService", function($scope, $mdDialog, $mdToast, $http, $location, apiService) {
    var update;
    $scope.errors = null;
    $scope.close = function() {
      return $mdDialog.hide();
    };
    update = function(t) {
      var idx;
      idx = _.findIndex($scope.group.guests, {
        id: t.id
      });
      $scope.group.guests[idx] = $scope.guests.current;
      return $scope.group.guests = _.sortBy($scope.group.guests, ["name"]);
    };
    $scope.create = function() {
      return $http.post((apiService.get()) + "/groups/" + $scope.group.url + "/guests", $scope.guests.current).then(function(resp) {
        $scope.$emit("group.reload");
        $mdDialog.hide();
        return $mdToast.showSimple("Nouvel utilisateur: '" + $scope.guests.current.email + "' ajouté.");
      }, function(resp) {
        return $scope.errors = resp.data.errors;
      });
    };
    $scope.edit = function() {
      return $http.put((apiService.get()) + "/groups/" + $scope.group.url + "/guests/" + $scope.guests.selected, $scope.guests.current).then(function(resp) {
        $mdDialog.hide();
        $mdToast.showSimple("Utilisateur: '" + $scope.guests.current.email + "' modifié.");
        update($scope.guests.current);
        return $scope.guests.selected = null;
      }, function(resp) {
        return $scope.errors = resp.data.errors;
      });
    };
    return $scope["delete"] = function() {
      return $http["delete"]((apiService.get()) + "/groups/" + $scope.group.url + "/guests/" + $scope.guests.selected, $scope.guests.current).then(function(resp) {
        $mdDialog.hide();
        $mdToast.showSimple("Utilisateur: '" + $scope.guests.current.email + "' supprimé.");
        $scope.group.guests.splice(_.findIndex($scope.group.guests, {
          id: parseInt($scope.guests.selected)
        }), 1);
        return $scope.task.selected = null;
      }, function(resp) {
        return $scope.errors = resp.data.errors;
      });
    };
  }
]);

angular.module("horodata").directive("appWidgetsConfigurationTasks", [
  "popupService", function(popupService) {
    var l;
    l = function(scope, elem, attr) {
      return scope.tasks = {
        current: null,
        selected: null,
        edit: function(ev) {
          this.current = _.cloneDeep(_.find(scope.group.tasks, {
            id: parseInt(this.selected)
          }));
          return popupService("horodata/widgets/configuration/tasks_form.html", "appWidgetsConfigurationTasksDialog", scope, ev);
        },
        create: function(ev) {
          this.current = {
            name: "",
            comment_mandatory: false
          };
          return popupService("horodata/widgets/configuration/tasks_form.html", "appWidgetsConfigurationTasksDialog", scope, ev);
        }
      };
    };
    return {
      link: l,
      restrict: "E",
      templateUrl: "horodata/widgets/configuration/tasks.html"
    };
  }
]);

angular.module("horodata").controller("appWidgetsConfigurationTasksDialog", [
  "$scope", "$mdDialog", "$mdToast", "$http", "$location", "apiService", function($scope, $mdDialog, $mdToast, $http, $location, apiService) {
    var update;
    $scope.errors = null;
    $scope.close = function() {
      return $mdDialog.hide();
    };
    update = function(t) {
      var idx;
      idx = _.findIndex($scope.group.tasks, {
        id: t.id
      });
      $scope.group.tasks[idx] = $scope.tasks.current;
      return $scope.group.tasks = _.sortBy($scope.group.tasks, ["name"]);
    };
    $scope.create = function() {
      return $http.post((apiService.get()) + "/groups/" + $scope.group.url + "/tasks", $scope.tasks.current).then(function(resp) {
        $scope.$emit("group.reload");
        $mdDialog.hide();
        return $mdToast.showSimple("Nouveau type de tâche: '" + $scope.tasks.current.name + "' ajouté.");
      }, function(resp) {
        return $scope.errors = resp.data.errors;
      });
    };
    $scope.edit = function() {
      return $http.put((apiService.get()) + "/groups/" + $scope.group.url + "/tasks/" + $scope.tasks.selected, $scope.tasks.current).then(function(resp) {
        $mdDialog.hide();
        $mdToast.showSimple("Type de tâche: '" + $scope.tasks.current.name + "' mis a jour.");
        update($scope.tasks.current);
        return $scope.tasks.selected = null;
      }, function(resp) {
        return $scope.errors = resp.data.errors;
      });
    };
    return $scope["delete"] = function() {
      return $http["delete"]((apiService.get()) + "/groups/" + $scope.group.url + "/tasks/" + $scope.tasks.selected, $scope.tasks.current).then(function(resp) {
        $mdDialog.hide();
        $mdToast.showSimple("Type de tâche: '" + $scope.tasks.current.name + "' supprimé.");
        $scope.group.tasks.splice(_.findIndex($scope.group.tasks, {
          id: parseInt($scope.tasks.selected)
        }), 1);
        return $scope.tasks.selected = null;
      }, function(resp) {
        return $scope.errors = resp.data.errors;
      });
    };
  }
]);

/**
 * HoroData Javascript Interface
 * Version: 0.0.1
 * Copyright © 2016 Hyperboloide. All rights reserved.
*/
angular.module("horodata", ["ngMaterial", "ngRoute", "ngMessages"]).config([
  "$mdDateLocaleProvider", "$mdThemingProvider", "$locationProvider", "$routeProvider", function($mdDateLocaleProvider, $mdThemingProvider, $locationProvider, $routeProvider) {
    var months;
    moment.locale('fr');
    $mdThemingProvider.theme('default').primaryPalette('blue').accentPalette('pink');
    $mdThemingProvider.setDefaultTheme('default');
    $locationProvider.html5Mode(true);
    $routeProvider.when("/", {
      templateUrl: "horodata/views/index.html",
      controller: "Index"
    }).when("/:group", {
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

angular.module("horodata").run(["$templateCache", function($templateCache) {$templateCache.put("horodata/menu/bottom_sheet.html","<md-bottom-sheet class=\"md-list md-has-header\" ng-cloak><md-subheader>Menu</md-subheader><md-list><md-list-item><md-button class=md-list-item-content><md-icon>trending_up</md-icon><span class=md-inline-list-icon-label>Plan et Quota</span></md-button></md-list-item><md-list-item><md-button class=md-list-item-content><md-icon>euro_symbol</md-icon><span class=md-inline-list-icon-label>Abonnement et Factures</span></md-button></md-list-item><md-list-item><md-button class=md-list-item-content ng-href=\"{{ home }}/account/logout\"><md-icon>directions_run</md-icon><span class=md-inline-list-icon-label>Quitter</span></md-button></md-list-item></md-list></md-bottom-sheet>");
$templateCache.put("horodata/menu/sidenav.html","<md-sidenav class=\"md-sidenav-left md-whiteframe-z2\" md-component-id=sidenav md-is-locked-open=\"$mdMedia(\'gt-sm\')\" layout=column><md-toolbar hide-gt-sm><div class=md-toolbar-tools><h1><md-icon class=md-24>access_time</md-icon>&nbsp; Horodata</h1><span flex></span><md-button ng-click=toggleSidenav() class=md-icon-button><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-content layout-padding hide show-gt-sm><div layout=column layout-align=\"center center\"><div class=md-display-3><i class=material-icons style=\"font-size: 64px;\">access_time</i></div><div class=md-headline>Horodata</div></div></md-content><md-content flex><section><md-subheader class=md-accent><div layout=row layout-align=\"space-between center\"><span class=md-headline>Groupes</span><app-widgets-new-group layout-align=end></app-widgets-new-group></div></md-subheader><md-list flex layout=column class=md-body-1><md-list-item ng-repeat=\"group in groups()\" ng-click=changeGroup(group.url)><span class=md-subhead ng-class=\"{\'group-selected\': currentGroupUrl == group.url}\">{{group.name}}</span></md-list-item></md-list><section></section></section></md-content></md-sidenav>");
$templateCache.put("horodata/menu/toolbar.html","<div><app-widgets-new-task></app-widgets-new-task><md-toolbar><div class=md-toolbar-tools><md-button hide-gt-sm class=md-icon-button aria-label=Settings ng-click=toggleSidenav()><md-icon class=md-24>menu</md-icon></md-button><h2><span>{{ MainTitle().title }}</span></h2><span flex></span><md-menu><md-button ng-click=\"openMenu($mdOpenMenu, $event)\" class=\"md-fab md-mini\" aria-label=Favorite><md-icon class=md-24>account_circle</md-icon><md-tooltip md-direction=left>Bonjour, {{ user.name }}</md-tooltip></md-button><md-menu-content width=4><md-menu-item><md-button ng-click=ctrl.redial($event)><md-icon>trending_up</md-icon>Plan et Quota</md-button></md-menu-item><md-menu-item><md-button><md-icon>euro_symbol</md-icon>Abonnement et Factures</md-button></md-menu-item><md-menu-divider></md-menu-divider><md-menu-item><md-button ng-href=\"{{ home }}/account/logout\"><md-icon><i class=material-icons>directions_run</i></md-icon>Quitter</md-button></md-menu-item></md-menu-content></md-menu></div></md-toolbar></div>");
$templateCache.put("horodata/views/group.html","<div><app-widgets-search-bar></app-widgets-search-bar><md-content><md-tabs md-dynamic-height md-border-bottom><md-tab label=Heures><app-widgets-empty-group-boxed ng-if=\"group.tasks.length == 0 || group.customers.length == 0\" group=group></app-widgets-empty-group-boxed><app-widgets-listing ng-if=\"group.tasks.length > 0 && group.customers.length > 0\"></app-widgets-listing></md-tab><md-tab label=Statistiques><app-widgets-empty-group-boxed ng-if=\"group.tasks.length == 0 || group.customers.length == 0\" group=group></app-widgets-empty-group-boxed></md-tab><md-tab ng-if=isOwner label=Configuration layout=column layout-margin><app-widgets-configuration-guests></app-widgets-configuration-guests><md-divider></md-divider><app-widgets-configuration-customers></app-widgets-configuration-customers><md-divider></md-divider><app-widgets-configuration-tasks></app-widgets-configuration-tasks></md-tab></md-tabs></md-content></div>");
$templateCache.put("horodata/views/index.html","<div><app-new-group></app-new-group></div>");
$templateCache.put("horodata/views/new_task_form.html","<md-dialog aria-label=\"Saisir une tâche\" flex=40><div ng-if=\"group.tasks.length == 0 || group.customers.length == 0\"><md-dialog-content><app-widgets-empty-group><app-widgets-empty-group></app-widgets-empty-group></app-widgets-empty-group></md-dialog-content><md-dialog-actions><md-button ng-click=close() class=md-raised>Fermer</md-button></md-dialog-actions></div><form ng-if=\"group.tasks.length > 0 && group.customers.length > 0\" name=newTaskForm><md-toolbar><div class=md-toolbar-tools><h2>Saisir une tâche</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><div layout=column><md-input-container flex><label>Sélectionnez un dossier</label><md-select ng-model=task.customer><md-option ng-repeat=\"c in group.customers\" value=\"{{ c.id }}\">{{ c.name }}</md-option></md-select><small ng-if=errors.customer class=input-error>{{ errors.customer }}</small></md-input-container><md-input-container flex><label>Sélectionnez une tâche</label><md-select ng-model=task.task><md-option ng-repeat=\"t in group.tasks\" value=\"{{ t.id }}\">{{ t.name }}</md-option></md-select><small ng-if=errors.task class=input-error>{{ errors.task }}</small></md-input-container><div layout=row layout-align=\"space-between center\"><md-input-container flex><label>Durée en heures</label><md-select ng-model=task.hours><md-option ng-repeat=\"h in hours\" value=\"{{ h }}\">{{ h }} <span ng-if=\"h > 1\">Heures</span> <span ng-if=\"h <= 1\">Heure</span></md-option></md-select><small ng-if=errors.duration class=input-error>{{ errors.duration }}</small></md-input-container><md-input-container flex><label>Durée en minutes</label><md-select ng-model=task.minutes><md-option ng-repeat=\"m in minutes\" value=\"{{ m }}\">{{ m }} <span ng-if=\"m > 1\">Minutes</span> <span ng-if=\"m <= 1\">Minute</span></md-option></md-select><small ng-if=errors.duration class=input-error>{{ errors.duration }}</small></md-input-container></div><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.comment}\"><label>Commentaire</label> <textarea md-no-autogrow ng-model=task.comment rows=3 md-select-on-focus></textarea> <small ng-if=errors.comment class=input-error>{{ errors.comment }}</small></md-input-container></div></div></md-dialog-content><md-dialog-actions><md-button ng-click=send() class=\"md-primary md-raised\">Ajouter</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/empty_group.html","<div layout-padding layout-margin layout=column layout-align=\"center center\"><div ng-if=\"group.tasks.length == 0 && group.customers.length == 0\"><div class=\"md-headline input-error\">Le groupe ne présente aucun dossier et aucun type de tâche</div><p>Vous devez ajouter des dossiers et des types de tâche pour permettre la saisie. Cliquez sur l\'onglet \"Configuration\" pour en ajouter.</p></div><div ng-if=\"group.tasks.length > 0 && group.customers.length == 0\"><div class=\"md-headline input-error\">Le groupe ne présente aucun dossier</div><p>Vous devez ajouter des dossiers pour permettre la saisie. Cliquez sur l\'onglet \"Configuration\" pour en ajouter.</p></div><div ng-if=\"group.tasks.length == 0 && group.customers.length > 0\"><div class=\"md-headline input-error\">Aucune tâches definie</div><p>Vous devez ajouter des types de tâche pour permettre la saisie. Cliquez sur l\'onglet \"Configuration\" pour en ajouter.</p></div></div>");
$templateCache.put("horodata/widgets/empty_group_boxed.html","<div layout=row flex layout-align=\"center center\"><md-whiteframe class=\"md-whiteframe-1dp md-accent\" flex=100 flex-sm=70 flex-md=60 flex-gt-md=50 layout-padding layout-margin layout=column><app-widgets-empty-group><app-widgets-empty-group></app-widgets-empty-group></app-widgets-empty-group></md-whiteframe></div>");
$templateCache.put("horodata/widgets/listing.html","<div><div ng-if=!listing.data() layout=column layout-align=\"center center\"><md-progress-circular md-mode=indeterminate md-diameter=150></md-progress-circular><p>Chargement, veuillez patienter.</p></div><div ng-if=listing.data()><md-list class=md-body-1><md-list-item class=md-body-2><div layout=row layout-align=\"space-between center\" flex><div flex=15>Jour</div><div ng-if=isOwner flex=25>Utilisateur</div><div flex=25>Dossier</div><div flex=25>Tâche</div><div flex=10>Durée</div></div><md-divider></md-divider></md-list-item><md-list-item ng-repeat=\"i in listing.data()\" class=secondary-button-padding ng-click=alert($event)><div layout=row layout-align=\"space-between center\" flex hide show-gt-sm><div flex=15>{{ i.created | Day }}</div><div flex=25 ng-if=isOwner>{{ i.creator.full_name }}</div><div flex=25>{{ customers[i.customer_id] }}</div><div flex=25>{{ tasks[i.task_id] }}</div><div flex=10>{{ i.duration | Duration }}</div></div><md-divider></md-divider></md-list-item></md-list></div></div>");
$templateCache.put("horodata/widgets/new_group.html","<md-button class=\"md-fab md-primary md-mini\" ng-click=showNewGroupDialog($event) aria-label=\"Créer un groupe\"><md-icon class=md-24>add</md-icon><md-tooltip md-direction=top>Créer un groupe</md-tooltip></md-button>");
$templateCache.put("horodata/widgets/new_group_form.html","<md-dialog aria-label=\"Créer un Groupe\" flex=40><form name=newGroupForm><md-toolbar><div class=md-toolbar-tools><h2>Créer un Groupe</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.name}\"><label>Nom du Groupe</label> <input type=text ng-model=name> <small ng-if=errors.name class=input-error>{{ errors.name }}</small></md-input-container></div></md-dialog-content><md-dialog-actions><md-button ng-click=send() class=\"md-primary md-raised\">Créer</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/new_task.html","<md-button ng-if=MainTitle().canEdit class=\"md-fab md-fab-bottom-right\" aria-label=Add ng-click=showNewTaskDialog($event)><md-tooltip md-direction=left>Saisir une tâche</md-tooltip><md-icon class=md-48 style=\"margin: -2px 0px 0px -1px\">access_time</md-icon></md-button>");
$templateCache.put("horodata/widgets/search_bar.html","<div><div layout=row layout-align=\"space-around center\" hide show-gt-md layout-padding><md-datepicker flex=25 ng-model=search.begin md-placeholder=\"Date début\" md-max-date=search.end></md-datepicker><md-datepicker flex=25 ng-model=search.end md-placeholder=\"Date fin\" md-max-date=today></md-datepicker><div flex=25><div layout=row><md-input-container flex><label>Dossier</label><md-select ng-model=search.customer><md-option ng-repeat=\"c in group.customers\" value=\"{{ c.id }}\">{{ c.name }}</md-option></md-select></md-input-container><md-button ng-click=\"search.customer = null\" ng-if=search.customer class=\"md-fab md-mini\" aria-label=désélectionner><md-icon>close</md-icon></md-button></div></div><div flex=25 ng-if=isOwner><div layout=row><md-input-container flex><label>Utilisateur</label><md-select ng-model=search.guest><md-option ng-repeat=\"u in group.guests\" ng-if=u.full_name value=\"{{ u.id }}\">{{ u.full_name }}</md-option></md-select></md-input-container><md-button ng-click=\"search.guest = null\" ng-if=search.guest class=\"md-fab md-mini\" aria-label=désélectionner><md-icon>close</md-icon></md-button></div></div></div><md-divider></md-divider></div>");
$templateCache.put("horodata/widgets/configuration/customers.html","<div layout-margin layout-padding><div layout=column flex-sm=70 flex-md=80 flex-gt-md=60><div layout=row layout-align=\"space-between center\"><div class=md-display-1>Dossiers</div><md-button ng-click=customers.create($event) class=\"md-raised md-primary\" hide show-gt-sm><md-icon class=md-18>add</md-icon>Ajouter</md-button><md-button ng-click=customers.create($event) class=\"md-fab md-primary\" hide-gt-sm aria-label=\"Ajouter un nouveau dossier\"><md-tooltip md-direction=top>Ajouter un nouveau dossier</md-tooltip><md-icon class=md-36>add</md-icon></md-button></div><p class=md-body-1>Listez l\'ensemble des dossiers sur lesquels travaillent vos collaborateurs.<br>Les dossiers peuvent représenter vos client ou diverses projets.</p><md-whiteframe ng-if=\"group.customers.length == 0\" class=\"md-whiteframe-1dp md-accent\" layout-padding layout=column layout-align=\"center center\"><div class=\"md-headline input-error\">Aucun dossier</div><div>Ajouter des dossiers pour permettre la saisie de tâches.</div></md-whiteframe><div layout=row ng-if=\"group.customers.length > 0\"><md-input-container flex><label>Sélectionnez un dossier pour le modifier</label><md-select ng-model=customers.selected><md-option ng-repeat=\"c in group.customers\" value=\"{{ c.id }}\">{{ c.name }}</md-option></md-select></md-input-container><md-button ng-if=customers.selected ng-click=customers.edit($event) class=\"md-fab md-mini md-accent\" aria-label=\"Editer le dossier\"><md-tooltip md-direction=top>Modifier le dossier</md-tooltip><md-icon class=md-24>edit</md-icon></md-button></div></div></div>");
$templateCache.put("horodata/widgets/configuration/customers_create_form.html","<md-dialog aria-label=\"Ajouter un nouveau dossier\" flex=40><form name=newCustomerForm><md-toolbar><div class=md-toolbar-tools><h2>Ajouter un ou des nouveaux dossiers</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.customers}\"><label>Nom du ou des dossiers à ajouter (pour ajouter plusieurs dossiers à la fois, saisissez un dossier par ligne, ou, copier-coller une liste depuis un tableau).</label> <textarea md-no-autogrow ng-model=customers.current.customers rows=5 md-select-on-focus></textarea> <small ng-if=errors.customers class=input-error>{{ errors.customers }}</small></md-input-container></div></md-dialog-content><md-dialog-actions><md-button ng-click=create() class=\"md-primary md-raised\">Ajouter</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/configuration/customers_edit_form.html","<md-dialog aria-label=\"Modifier le dossier\" flex=40><form name=editCustomerForm><md-toolbar><div class=md-toolbar-tools><h2>Modifier le dossier</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.name}\"><label>Nom du dossier</label> <input md-maxlength=40 type=text ng-model=customers.current.name> <small ng-if=errors.name class=input-error>{{ errors.name }}</small></md-input-container></div></md-dialog-content><md-dialog-actions><md-button ng-click=edit() class=\"md-primary md-raised\">Modifier</md-button><md-button ng-click=delete() class=\"md-warn md-raised\">Supprimer</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/configuration/guests.html","<div layout-margin layout-padding><div layout=column flex-sm=70 flex-md=80 flex-gt-md=60><div layout=row layout-align=\"space-between center\"><div class=md-display-1>Utilisateurs</div><md-button ng-click=guests.create($event) class=\"md-raised md-primary\" hide show-gt-sm><md-icon class=md-18>add</md-icon>Ajouter</md-button><md-button ng-click=guests.create($event) class=\"md-fab md-primary\" hide-gt-sm aria-label=\"Ajouter un nouvel utilisateur\"><md-tooltip md-direction=top>Ajouter un nouvel utilisateur</md-tooltip><md-icon class=md-36>add</md-icon></md-button></div><p class=md-body-1>Invitez vos collaborateurs à saisir leurs tâches dans Horodata.<br></p><div layout=row ng-if=\"group.guests.length > 0\"><md-input-container flex><label>Selectionnez un utilisateur pour le modifier</label><md-select ng-model=guests.selected><md-option ng-repeat=\"g in group.guests\" value=\"{{ g.id }}\"><span ng-if=!g.full_name>{{ g.email }}</span> <span ng-if=g.full_name>{{ g.full_name }} <small>{{ g.email }}</small></span></md-option></md-select></md-input-container><md-button ng-if=guests.selected ng-click=guests.edit($event) class=\"md-fab md-mini md-accent\" aria-label=\"Editer une tâche\"><md-tooltip md-direction=top>Modifier l\'utilisateur</md-tooltip><md-icon class=md-24>edit</md-icon></md-button></div></div></div>");
$templateCache.put("horodata/widgets/configuration/guests_create_form.html","<md-dialog aria-label=\"Ajouter un nouvel Utilisateur\" flex=40><form name=newUserForm><md-toolbar><div class=md-toolbar-tools><h2>Ajouter un nouvel Utilisateur</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.email}\"><label>Adresse email</label> <input type=text md-maxlength=100 ng-model=guests.current.email> <small ng-if=errors.email class=input-error>{{ errors.email }}</small></md-input-container><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.rate}\"><label>Taux horaire</label> <input type=number ng-model=guests.current.rate ng-value=0> <small ng-if=errors.rate class=input-error>{{ errors.rate }}</small></md-input-container><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.admin}\"><md-switch ng-model=guests.current.admin aria-label=\"Droit administrateur\">Droit administrateur</md-switch><div class=input-hint>Cochez la case, uniquement si vous souhaitez autoriser l\'utilisateur, à accéder à la configuration du groupe.</div><small ng-if=errors.admin class=input-error>{{ errors.admin }}</small></md-input-container></div></md-dialog-content><md-dialog-actions><md-button ng-click=create() class=\"md-primary md-raised\">Ajouter</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/configuration/guests_edit_form.html","<md-dialog aria-label=\"Nouvel Utilisateur\" flex=40><form name=newUserForm><md-toolbar><div class=md-toolbar-tools><h2>Modifier l\'utilisateur <span ng-if=!guests.current.full_name>{{ guests.current.email }}</span> <span ng-if=guests.current.full_name>{{ guests.current.full_name }}</span></h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.rate}\"><label>Taux horaire</label> <input type=number ng-model=guests.current.rate ng-value=0> <small ng-if=errors.rate class=input-error>{{ errors.rate }}</small></md-input-container><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.admin}\"><md-switch ng-model=guests.current.admin aria-label=Administrateur>Droit administrateur</md-switch><div class=input-hint>Cochez la case, uniquement si vous souhaitez autoriser l\'utilisateur, à accéder à la configuration du groupe.</div><small ng-if=errors.admin class=input-error>{{ errors.admin }}</small></md-input-container></div></md-dialog-content><md-dialog-actions><md-button ng-click=edit() class=\"md-primary md-raised\">Modifier</md-button><md-button ng-click=delete() class=\"md-warn md-raised\">Supprimer</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/configuration/tasks.html","<div layout-margin layout-padding><div layout=column flex-sm=70 flex-md=80 flex-gt-md=60><div layout=row layout-align=\"space-between center\"><div class=md-display-1>Types de tâche</div><md-button ng-click=tasks.create($event) class=\"md-raised md-primary\" hide show-gt-sm><md-icon class=md-18>add</md-icon>Ajouter</md-button><md-button ng-click=tasks.create($event) class=\"md-fab md-primary\" hide-gt-sm aria-label=\"Ajouter un type de tâche\"><md-tooltip md-direction=top>Ajouter un type de tâche</md-tooltip><md-icon class=md-36>add</md-icon></md-button></div><p class=md-body-1>Listez les types de tâche qu\'accomplissent vos collaborateurs.</p><md-whiteframe ng-if=\"group.tasks.length == 0\" class=\"md-whiteframe-1dp md-accent\" layout-padding layout=column layout-align=\"center center\"><div class=\"md-headline input-error\">Aucun type de tâche</div><div>Ajouter des types pour permettre la saisie de tâches.</div></md-whiteframe><div layout=row ng-if=\"group.tasks.length > 0\"><md-input-container flex><label>Sélectionnez un type de tâche pour le modifier</label><md-select ng-model=tasks.selected><md-option ng-repeat=\"t in group.tasks\" value=\"{{ t.id }}\">{{ t.name }} <small ng-if=t.comment_mandatory>(Commentaire obligatoire)</small></md-option></md-select></md-input-container><md-button ng-if=tasks.selected ng-click=tasks.edit($event) class=\"md-fab md-mini md-accent\" aria-label=\"Editer une tâche\"><md-tooltip md-direction=top>Modifier la tâche</md-tooltip><md-icon class=md-24>edit</md-icon></md-button></div></div></div>");
$templateCache.put("horodata/widgets/configuration/tasks_form.html","<md-dialog aria-label=\"Ajouter un nouveau type de tâche\" flex=40><form name=newTaskForm><md-toolbar><div class=md-toolbar-tools><h2>Ajouter un nouveau type de tâche</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.name}\"><label>Nom de la tâche</label> <input md-maxlength=30 type=text ng-model=tasks.current.name> <small ng-if=errors.name class=input-error>{{ errors.name }}</small></md-input-container><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.admin}\"><md-switch ng-model=tasks.current.comment_mandatory aria-label=Administrateur>Commentaire obligatoire</md-switch><div class=input-hint>Cochez la case, uniquement si vous souhaitez rendre les commentaires obligatoires.</div></md-input-container></div></md-dialog-content><md-dialog-actions><md-button ng-if=!tasks.current.id ng-click=create() class=\"md-primary md-raised\">Ajouter</md-button><md-button ng-if=tasks.current.id ng-click=edit() class=\"md-primary md-raised\">Modifier</md-button><md-button ng-if=tasks.current.id ng-click=delete() class=\"md-warn md-raised\">Supprimer</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");}]);
angular.module("horodata").directive("appMenuSidenav", [
  "$mdSidenav", "$http", "$location", "apiService", "$routeParams", "groupNewService", function($mdSidenav, $http, $location, apiService, $routeParams, groupNewService) {
    var l;
    l = function(scope, elem) {
      scope.toggleSidenav = function() {
        return $mdSidenav("sidenav").toggle();
      };
      groupNewService.fetch();
      scope.groups = function() {
        return groupNewService.listing();
      };
      scope.changeGroup = function(url) {
        return $location.path(url);
      };
      return scope.$on("$routeChangeSuccess", function() {
        return scope.currentGroupUrl = $routeParams.group;
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
  "titleService", "userService", "homeService", "$mdMedia", "$mdBottomSheet", function(titleService, userService, homeService, $mdMedia, $mdBottomSheet) {
    var l;
    l = function(scope, elem) {
      scope.MainTitle = titleService.get;
      userService.get(function(u) {
        return scope.user = u;
      });
      scope.home = homeService.get();
      return scope.openMenu = function($mdOpenMenu, $event) {
        if ($mdMedia('xs') || $mdMedia('sm')) {
          return $mdBottomSheet.show({
            templateUrl: "horodata/menu/bottom_sheet.html"
          });
        } else {
          return $mdOpenMenu($event);
        }
      };
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

angular.module("horodata").filter("Day", [
  function() {
    return function(input) {
      return moment(input).format('LL');
    };
  }
]);

angular.module("horodata").filter("Ago", [
  function() {
    return function(input) {
      return moment(input).fromNow();
    };
  }
]);

angular.module("horodata").filter("Duration", [
  function() {
    return function(input) {
      var d, hours, minutes;
      d = moment.duration(input, 'seconds');
      minutes = d.minutes();
      if (minutes < 10) {
        minutes = "0" + minutes;
      }
      hours = d.hours();
      return hours + "h" + minutes;
    };
  }
]);

angular.module('horodata').factory("groupNewService", [
  "apiService", "$http", function(apiService, $http) {
    var callback, fetchListing, groups;
    callback = null;
    groups = [];
    fetchListing = function() {
      return $http.get((apiService.get()) + "/groups").then(function(resp) {
        return groups = resp.data.data.results;
      });
    };
    return {
      set: function(fn) {
        return callback = fn;
      },
      open: function(ev) {
        return callback(ev);
      },
      listing: function() {
        return groups;
      },
      fetch: function() {
        return fetchListing();
      }
    };
  }
]);

angular.module('horodata').factory("homeService", [
  function() {
    var root;
    root = document.getElementsByTagName("home")[0].getAttribute("href");
    return {
      get: function() {
        return root;
      }
    };
  }
]);

var bind = function(fn, me){ return function(){ return fn.apply(me, arguments); }; };

angular.module('horodata').factory("listingService", [
  "$http", "apiService", function($http, apiService) {
    var Listing, listing;
    Listing = (function() {
      function Listing(groupUrl1, begin, end, customer, guest) {
        this.groupUrl = groupUrl1;
        this.fetch = bind(this.fetch, this);
        this.size = 50;
        this.list = null;
        this.loading = false;
        this.total = -1;
        this.params = {
          begin: moment(begin).format('YYYY-MM-DD'),
          end: moment(end).format('YYYY-MM-DD'),
          customer: customer,
          guest: guest
        };
      }

      Listing.prototype.pages = function(page) {
        if (this.total === -1) {
          return [];
        }
      };

      Listing.prototype.fetch = function(page) {
        var params;
        if (this.loading) {
          return;
        }
        this.loading = true;
        params = _.cloneDeep(this.params);
        params.offset = page * this.size;
        params.size = this.size;
        return $http.get((apiService.get()) + "/groups/" + this.groupUrl + "/jobs", {
          params: params
        }).then((function(_this) {
          return function(resp) {
            _this.list = resp.data.data.results;
            _this.loading = false;
            return _this.total = resp.data.data.total;
          };
        })(this), (function(_this) {
          return function(resp) {
            console.log(resp.error);
            return _this.loading = false;
          };
        })(this));
      };

      return Listing;

    })();
    listing = {};
    return {
      data: function() {
        if ((listing.list == null) || listing.loading) {
          null;
        }
        return listing.list;
      },
      listing: function() {
        return listing;
      },
      search: function(groupUrl, params) {
        return listing = new Listing(groupUrl, params.begin, params.end, params.customer, params.guest);
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
  "$http", "$routeParams", "$scope", "titleService", "userService", "apiService", "groupNewService", "popupService", "listingService", function($http, $routeParams, $scope, titleService, userService, apiService, groupNewService, popupService, listingService) {
    var fetchUsers, getGroup;
    $scope.isGroupView = true;
    $scope.search = {
      begin: moment().subtract(1, 'months').toDate(),
      end: new Date(),
      customer: null,
      guest: null
    };
    $scope.$watch("search", function(v) {
      if (v == null) {
        return;
      }
      listingService.search($routeParams.group, v);
      return listingService.listing().fetch(0);
    }, true);
    fetchUsers = function() {
      return $http.get((apiService.get()) + "/groups/" + $routeParams.group + "/users").then(function(resp) {
        return $scope.users = resp.data.data;
      });
    };
    $scope.isOwner = false;
    getGroup = function() {
      return $http.get((apiService.get()) + "/groups/" + $routeParams.group).then(function(resp) {
        $scope.group = resp.data.data;
        $scope.isOwner = $scope.group.owner === $scope.user.id;
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
    var x;
    $scope.task = {
      minutes: 0,
      hours: 0
    };
    $scope.errors = null;
    $scope.hours = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12];
    $scope.minutes = (function() {
      var i, results;
      results = [];
      for (x = i = 0; i <= 55; x = i += 5) {
        results.push(x);
      }
      return results;
    })();
    $scope.close = function() {
      return $mdDialog.hide();
    };
    return $scope.send = function() {
      var task;
      task = {
        duration: $scope.task.hours * 3600 + $scope.task.minutes * 60,
        task: parseInt($scope.task.task),
        customer: parseInt($scope.task.customer),
        comment: $scope.task.comment
      };
      return $http.post((apiService.get()) + "/groups/" + $scope.group.url + "/jobs", task).then(function(resp) {
        $mdDialog.hide();
        return $mdToast.showSimple("Nouvelle tâche ajoutée.");
      }, function(resp) {
        return $scope.errors = resp.data.errors;
      });
    };
  }
]);

angular.module("horodata").controller("Index", [
  "$http", "$scope", "userService", "titleService", function($http, $scope, userService, titleService) {
    return titleService.set("Bienvenu");
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

angular.module("horodata").directive("appWidgetsEmptyGroupBoxed", [
  function() {
    return {
      restrict: "E",
      templateUrl: "horodata/widgets/empty_group_boxed.html"
    };
  }
]);

angular.module("horodata").directive("appWidgetsListing", [
  "listingService", "$timeout", function(listingService, $timeout) {
    var l;
    l = function(scope) {
      var i, j, k, len, len1, ref, ref1;
      scope.tasks = {};
      ref = scope.group.tasks;
      for (j = 0, len = ref.length; j < len; j++) {
        i = ref[j];
        scope.tasks[i.id] = i.name;
      }
      scope.customers = {};
      ref1 = scope.group.customers;
      for (k = 0, len1 = ref1.length; k < len1; k++) {
        i = ref1[k];
        scope.customers[i.id] = i.name;
      }
      return scope.listing = listingService;
    };
    return {
      link: l,
      replace: true,
      restrict: "E",
      templateUrl: "horodata/widgets/listing.html"
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
  "$scope", "$mdDialog", "$mdToast", "$http", "$location", "apiService", "groupNewService", function($scope, $mdDialog, $mdToast, $http, $location, apiService, groupNewService) {
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
        $location.path("/" + group.url);
        return groupNewService.fetch();
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
    var l;
    l = function(scope) {
      return scope.today = new Date();
    };
    return {
      link: l,
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
        return $mdToast.showSimple("Nouveau type de tâche '" + $scope.tasks.current.name + "' ajouté.");
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

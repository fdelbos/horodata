/**
 * HoroData Javascript Interface
 * Version: 0.0.1
 * Copyright © 2016 Hyperboloide. All rights reserved.
*/
angular.module("horodata", ["ngMaterial", "ngRoute", "ngMessages", "gridshore.c3js.chart"]).config([
  "$mdDateLocaleProvider", "$mdThemingProvider", "$locationProvider", "$routeProvider", function($mdDateLocaleProvider, $mdThemingProvider, $locationProvider, $routeProvider) {
    var months;
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
    moment.locale('fr');
    months = ["Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre"];
    $mdDateLocaleProvider.month = months;
    $mdDateLocaleProvider.days = ['dimanche', 'lundi', 'mardi', 'mercredi', 'jeudi', 'venredi', 'samedi'];
    $mdDateLocaleProvider.shortDays = ['Di', 'Lu', 'Ma', 'Me', 'Je', 'Ve', 'Sa'];
    $mdDateLocaleProvider.firstDayOfWeek = 1;
    $mdDateLocaleProvider.msgCalendar = 'Calendrier';
    $mdDateLocaleProvider.msgOpenCalendar = 'Ouvrir le calendrier';
    $mdDateLocaleProvider.monthHeaderFormatter = function(date) {
      return months[date.getMonth()] + ' ' + date.getFullYear();
    };
    $mdDateLocaleProvider.parseDate = function(dateString) {
      if (moment(dateString, 'L', true).isValid()) {
        return m.toDate();
      } else {
        return new Date(NaN);
      }
    };
    return $mdDateLocaleProvider.formatDate = function(date) {
      return moment(date).format('L');
    };
  }
]).run([
  "$http", function($http) {
    return $http.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
  }
]);

angular.module("horodata").run(["$templateCache", function($templateCache) {$templateCache.put("horodata/menu/bottom_sheet.html","<md-bottom-sheet class=\"md-list md-has-header\" ng-cloak><div layout=row hide show-sm layout-padding layout-align=\"space-around center\"><div layout=column layout-align=\"center center\"><md-button class=\"md-fab md-primary\" ng-click=showProfile($event)><md-icon class=md-36>account_box</md-icon></md-button><div>Profile</div></div><div layout=column layout-align=\"center center\"><md-button ng-click=showQuotas($event) class=\"md-fab md-primary\"><md-icon class=md-36>trending_up</md-icon></md-button><div>Quotas</div></div><div layout=column layout-align=\"center center\"><md-button class=\"md-fab md-primary\"><md-icon class=md-36>euro_symbol</md-icon></md-button><div>Abonnement</div></div><div layout=column layout-align=\"center center\"><md-button class=\"md-fab md-warn\" ng-href=\"{{ home }}/account/logout\"><md-icon class=md-36>directions_run</md-icon></md-button><div>Quitter</div></div></div><div layout=column hide show-xs layout-padding layout-align=\"start start\"><div layout=row layout-align=\"center center\"><md-button class=\"md-fab md-mini md-primary\" ng-click=showProfile($event)><md-icon class=md-24>account_box</md-icon></md-button><div>Profile</div></div><div layout=row layout-align=\"center center\"><md-button ng-click=showQuotas($event) class=\"md-fab md-mini md-primary\"><md-icon class=md-24>trending_up</md-icon></md-button><div>Quotas</div></div><div layout=row layout-align=\"center center\"><md-button class=\"md-fab md-mini md-primary\"><md-icon class=md-24>euro_symbol</md-icon></md-button><div>Abonnement</div></div><div layout=row layout-align=\"center center\"><md-button class=\"md-fab md-mini md-warn\" ng-href=\"{{ home }}/account/logout\"><md-icon class=md-24>directions_run</md-icon></md-button><div>Quitter</div></div></div></md-bottom-sheet>");
$templateCache.put("horodata/menu/sidenav.html","<md-sidenav class=\"md-sidenav-left md-whiteframe-z2\" md-component-id=sidenav md-is-locked-open=\"$mdMedia(\'gt-md\')\" layout=column><md-toolbar hide-gt-sm><div class=md-toolbar-tools><h1><md-icon class=md-24>access_time</md-icon>&nbsp; Horodata</h1><span flex></span><md-button ng-click=toggleSidenav() class=md-icon-button><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-content layout-padding hide show-gt-sm><div layout=column layout-align=\"center center\"><div class=md-display-3><i class=material-icons style=\"font-size: 64px;\">access_time</i></div><div class=md-headline>Horodata</div></div></md-content><md-content flex><section><md-subheader class=md-accent><div layout=row layout-align=\"space-between center\"><span class=md-headline>Groupes</span><app-widgets-new-group layout-align=end></app-widgets-new-group></div></md-subheader><md-list flex layout=column class=md-body-1><md-list-item ng-repeat=\"group in groups()\" ng-click=changeGroup(group.url)><span class=md-subhead ng-class=\"{\'group-selected\': currentGroupUrl == group.url}\">{{group.name}}</span></md-list-item></md-list><section></section></section></md-content></md-sidenav>");
$templateCache.put("horodata/menu/toolbar.html","<div><app-widgets-big-button></app-widgets-big-button><md-toolbar><div class=md-toolbar-tools><md-button hide-gt-md class=md-icon-button aria-label=Settings ng-click=toggleSidenav()><md-icon class=md-24>menu</md-icon></md-button><h2><span>{{ MainTitle().title }}</span></h2><span flex></span><md-menu><md-button ng-click=\"openMenu($mdOpenMenu, $event)\" class=\"md-fab md-mini\" aria-label=Favorite><img ng-src=\"{{ user.picture | Profile }}\" class=profile-icon alt=\"{{ user.name }}\" style=\"height: 40px; width: 40px;\"><md-tooltip md-direction=left show-gt-md>Bonjour, {{ user.name }}</md-tooltip></md-button><md-menu-content width=4><md-menu-item><md-button ng-click=showProfile($event)><md-icon>account_box</md-icon>Profile</md-button></md-menu-item><md-menu-item><md-button ng-click=showQuotas($event)><md-icon>trending_up</md-icon>Quotas</md-button></md-menu-item><md-menu-item><md-button><md-icon>euro_symbol</md-icon>Abonnement</md-button></md-menu-item><md-menu-divider></md-menu-divider><md-menu-item><md-button ng-href=\"{{ home }}/account/logout\"><md-icon><i class=material-icons>directions_run</i></md-icon>Quitter</md-button></md-menu-item></md-menu-content></md-menu></div></md-toolbar></div>");
$templateCache.put("horodata/views/group.html","<div flex layout=column><md-content flex><md-tabs md-selected=selectedTab md-dynamic-height md-border-bottom><md-tab label=Saisies layout=row><md-content md-swipe-left=selectTab(1) flex><app-widgets-empty-group-boxed ng-if=\"group.tasks.length == 0 || group.customers.length == 0\" group=group></app-widgets-empty-group-boxed><app-widgets-listing ng-if=\"group.tasks.length > 0 && group.customers.length > 0\"></app-widgets-listing></md-content></md-tab><md-tab label=Statistiques><div md-swipe-left=selectTab(2) md-swipe-right=selectTab(0) layout=column><app-widgets-empty-group-boxed ng-if=\"group.tasks.length == 0 || group.customers.length == 0\" group=group></app-widgets-empty-group-boxed><app-widgets-stats ng-if=\"group.tasks.length > 0 && group.customers.length > 0\"></app-widgets-stats><div></div></div></md-tab><md-tab ng-if=isAdmin label=Configuration layout=column layout-margin><div md-swipe-right=selectTab(1)><app-widgets-configuration></app-widgets-configuration></div></md-tab></md-tabs></md-content></div>");
$templateCache.put("horodata/views/index.html","<div><app-new-group></app-new-group></div>");
$templateCache.put("horodata/views/profile.html","<md-dialog aria-label=Profile flex=40><form name=profileForm><app-widgets-common-dialog-toolbar>Profile</app-widgets-common-dialog-toolbar><md-dialog-content><div class=md-dialog-content><div layout=column><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.email}\"><label>Nom</label> <input type=text md-maxlength=50 ng-model=name> <small ng-if=errors.email class=input-error>{{ errors.email }}</small></md-input-container><div layout=column flex><span class=md-caption>Email</span><div>{{ user.email }}</div></div><br></div></div></md-dialog-content><app-widgets-common-dialog-actions><md-button ng-click=send() class=\"md-primary md-raised\">Modifier</md-button></app-widgets-common-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/views/quotas.html","<md-dialog aria-label=Quota flex=40><app-widgets-common-dialog-toolbar>Quotas</app-widgets-common-dialog-toolbar><md-dialog-content><div class=md-dialog-content><div layout=column><h3 class=md-headline>Plan <span ng-if=\"quotas.quotas.plan == \'free\'\">Gratuit</span> <span ng-if=\"quotas.quotas.plan == \'small\'\">10 Utilisateurs</span> <span ng-if=\"quotas.quotas.plan == \'medium\'\">30 Utilisateurs</span> <span ng-if=\"quotas.quotas.plan == \'large\'\">100 Utilisateurs</span></h3><app-widgets-quota label=Groupes current=quotas.usage.groups max=quotas.quotas.limits.groups></app-widgets-quota><app-widgets-quota label=Utilisateurs current=quotas.usage.guests max=quotas.quotas.limits.guests></app-widgets-quota><app-widgets-quota label=\"Saisies (aujourd\'hui)\" current=quotas.usage.jobs max=quotas.quotas.limits.jobs></app-widgets-quota><p>Pour modifier votre plan et changer vos quotas, rendez vous dans le menu <strong>Abonnement</strong>.</p></div></div></md-dialog-content><app-widgets-common-dialog-actions></app-widgets-common-dialog-actions></md-dialog>");
$templateCache.put("horodata/widgets/empty_group.html","<div layout-padding layout-margin layout=column layout-align=\"center center\"><div ng-if=\"group.tasks.length == 0 && group.customers.length == 0\"><div class=\"md-headline input-error\">Le groupe ne présente aucun dossier et type</div><p>Vous devez ajouter des dossiers et des types pour permettre la saisie de tâches. Cliquez sur l\'onglet \"Configuration\" pour en ajouter.</p></div><div ng-if=\"group.tasks.length > 0 && group.customers.length == 0\"><div class=\"md-headline input-error\">Le groupe ne présente aucun dossier</div><p>Vous devez ajouter des dossiers pour permettre la saisie de tâches. Cliquez sur l\'onglet \"Configuration\" pour en ajouter.</p></div><div ng-if=\"group.tasks.length == 0 && group.customers.length > 0\"><div class=\"md-headline input-error\">Le groupe ne présente aucun type</div><p>Vous devez ajouter des types pour permettre la saisie de tâches. Cliquez sur l\'onglet \"Configuration\" pour en ajouter.</p></div></div>");
$templateCache.put("horodata/widgets/empty_group_boxed.html","<div layout=row flex layout-align=\"center center\"><md-whiteframe class=\"md-whiteframe-1dp md-accent\" flex=100 flex-sm=70 flex-md=60 flex-gt-md=50 layout-padding layout-margin layout=column><app-widgets-empty-group><app-widgets-empty-group></app-widgets-empty-group></app-widgets-empty-group></md-whiteframe></div>");
$templateCache.put("horodata/widgets/loading.html","<div layout=column layout-align=\"center center\" flex><md-progress-circular md-mode=indeterminate md-diameter=150></md-progress-circular><p>Chargement, veuillez patienter.</p></div>");
$templateCache.put("horodata/widgets/new_group.html","<md-button class=\"md-fab md-primary md-mini\" ng-click=showNewGroupDialog($event) aria-label=\"Créer un groupe\"><md-icon class=md-24>add</md-icon><md-tooltip md-direction=top>Créer un groupe</md-tooltip></md-button>");
$templateCache.put("horodata/widgets/new_group_form.html","<md-dialog aria-label=\"Créer un groupe\" flex=40><form name=newGroupForm><app-widgets-common-dialog-toolbar>Créer un groupe</app-widgets-common-dialog-toolbar><md-dialog-content><div class=md-dialog-content><app-widgets-common-quota-error></app-widgets-common-quota-error><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.name}\"><label>Nom du groupe</label> <input type=text ng-model=name> <small ng-if=errors.name class=input-error>{{ errors.name }}</small></md-input-container></div></md-dialog-content><app-widgets-common-dialog-actions><md-button ng-click=send() class=\"md-primary md-raised\">Créer</md-button></app-widgets-common-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/quota.html","<div><div class=md-body-1 layout=row layout-align=\"space-between center\"><span flex=60>{{ label }}</span> <span flex=30 style=text-align:right>{{ current }}/{{ max }} &nbsp;</span> <span flex=10 style=text-align:right ng-class=\"{\'input-error\': percent >= 80}\">{{ percent }}%</span></div><md-progress-linear md-mode=determinate ng-class=\"{\'md-warn\': percent >= 80}\" value=\"{{ percent }}\"></md-progress-linear><br><br></div>");
$templateCache.put("horodata/widgets/big_button/export.html","<md-dialog aria-label=\"Exporter les tâches\" flex=40><div ng-if=\"group.tasks.length == 0 || group.customers.length == 0\"><app-widgets-common-dialog-toolbar>Exporter les tâches</app-widgets-common-dialog-toolbar><md-dialog-content><app-widgets-empty-group><app-widgets-empty-group></app-widgets-empty-group></app-widgets-empty-group></md-dialog-content><app-widgets-common-dialog-actions></app-widgets-common-dialog-actions></div><form ng-if=\"group.tasks.length > 0 && group.customers.length > 0\" name=newTaskForm><app-widgets-common-dialog-toolbar>Exporter les tâches</app-widgets-common-dialog-toolbar><md-dialog-content><div class=\"md-dialog-content md-body-1\"><p class=md-subhead>Saisies du {{ filter.begin | Day }} au {{ filter.end | Day }}</p>Format d\'export:<br><br><div layout=column><md-radio-group ng-model=export.fileType class=md-primary><md-radio-button value=xlsx>XLSX</md-radio-button><md-radio-button value=csv>CSV</md-radio-button></md-radio-group></div></div></md-dialog-content><app-widgets-common-dialog-actions><md-button ng-click=hide() ng-href=\"{{ url }}_{{ export.fileType }}{{ filter.urlParams() }}\" class=\"md-primary md-raised\">Telecharger</md-button></app-widgets-common-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/big_button/new_task.html","<md-dialog aria-label=\"Saisir une tâche\" flex=50><div ng-if=\"group.tasks.length == 0 || group.customers.length == 0\"><app-widgets-common-dialog-toolbar>Saisir une tâche</app-widgets-common-dialog-toolbar><md-dialog-content><app-widgets-empty-group><app-widgets-empty-group></app-widgets-empty-group></app-widgets-empty-group></md-dialog-content><app-widgets-common-dialog-actions></app-widgets-common-dialog-actions></div><form ng-if=\"group.tasks.length > 0 && group.customers.length > 0\" name=newTaskForm><app-widgets-common-dialog-toolbar>Saisir une tâche</app-widgets-common-dialog-toolbar><md-dialog-content><div class=md-dialog-content><div layout=column><app-widgets-common-quota-error></app-widgets-common-quota-error><md-input-container flex><label>Sélectionnez un dossier</label><md-select ng-model=task.customer><md-option ng-repeat=\"c in group.customers\" value=\"{{ c.id }}\">{{ c.name }}</md-option></md-select><small ng-if=errors.customer class=input-error>{{ errors.customer }}</small></md-input-container><md-input-container flex><label>Sélectionnez une tâche</label><md-select ng-model=task.task><md-option ng-repeat=\"t in group.tasks\" value=\"{{ t.id }}\">{{ t.name }}</md-option></md-select><small ng-if=errors.task class=input-error>{{ errors.task }}</small></md-input-container><div layout=row layout-align=\"space-between center\"><md-input-container flex><label>Durée en heures</label><md-select ng-model=task.hours><md-option ng-repeat=\"h in hours\" value=\"{{ h }}\">{{ h }} <span ng-if=\"h > 1\">heures</span> <span ng-if=\"h <= 1\">heure</span></md-option></md-select><small ng-if=errors.duration class=input-error>{{ errors.duration }}</small></md-input-container><md-input-container flex><label>Durée en minutes</label><md-select ng-model=task.minutes><md-option ng-repeat=\"m in minutes\" value=\"{{ m }}\">{{ m }} <span ng-if=\"m > 1\">minutes</span> <span ng-if=\"m <= 1\">minute</span></md-option></md-select><small ng-if=errors.duration class=input-error>{{ errors.duration }}</small></md-input-container></div><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.comment}\"><label>Commentaire</label> <textarea md-no-autogrow ng-model=task.comment rows=3 md-select-on-focus></textarea> <small ng-if=errors.comment class=input-error>{{ errors.comment }}</small></md-input-container></div></div></md-dialog-content><app-widgets-common-dialog-actions><md-button ng-click=send() class=\"md-primary md-raised\">Ajouter</md-button></app-widgets-common-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/big_button/root.html","<md-button class=\"md-fab md-fab-bottom-right\" ng-if=currentTab() ng-class=\"{\'md-primary\': currentTab() == \'export\'}\" aria-label=Add ng-click=newDialog($event)><md-tooltip md-direction=left><span ng-if=\"currentTab() == \'jobs\'\">Saisir une tâche</span> <span ng-if=\"currentTab() == \'export\'\">Exporter les tâches</span></md-tooltip><md-icon ng-if=\"currentTab() == \'jobs\'\" class=md-48 style=\"margin: -2px 0px 0px -1px\">access_time</md-icon><md-icon ng-if=\"currentTab() == \'export\'\" class=md-48 style=\"margin: -2px 0px 0px -1px\">file_download</md-icon></md-button>");
$templateCache.put("horodata/widgets/common/dialog_actions.html","<md-dialog-actions><ng-transclude ng-if=!loading></ng-transclude><md-button ng-if=!loading ng-click=hide() class=md-raised>Annuler</md-button><md-progress-linear ng-if=loading md-mode=indeterminate></md-progress-linear></md-dialog-actions>");
$templateCache.put("horodata/widgets/common/dialog_toolbar.html","<md-toolbar ng-class=\"{\'md-warn\': warn}\"><div class=md-toolbar-tools><h2><ng-transclude></ng-transclude></h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=hide()><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar>");
$templateCache.put("horodata/widgets/common/quota_error.html","<div layout=column layout-fill><md-whiteframe ng-if=quotaError class=md-whiteframe-1dp layout-padding flex layout-align=\"center center\" style=text-align:center><div ng-switch=quotaError.limit class=\"md-title input-error\"><span ng-switch-when=groups>Vous ne pouvez pas créer de nouveau groupe.</span> <span ng-switch-when=guests>Vous ne pouvez pas inviter un nouvel utilisateur.</span> <span ng-switch-when=jobs>Vous ne pouvez pas saisir une nouvelle tache.</span></div><p ng-if=\"quotaError.limit == \'groups\'\">Pour modifier votre plan et changer vos quotas, rendez vous dans le menu <strong>Abonnement</strong>.</p><div ng-if=\"quotaError.limit != \'groups\'\"><p ng-if=isOwner>Pour modifier votre plan et changer vos quotas, rendez vous dans le menu <strong>Abonnement</strong>.</p><p ng-if=!isOwner>Contactez le proprietaire de ce groupe pour ajouter des utilisateurs.</p></div></md-whiteframe></div>");
$templateCache.put("horodata/widgets/configuration/customers.html","<div layout-padding><div layout=column flex-sm=70 flex-md=80 flex-gt-md=60><div layout=row layout-align=\"space-between center\"><div class=md-display-1 hide show-gt-sm>Dossiers</div><div class=md-headline hide-gt-sm>Dossiers</div><md-button ng-click=customers.create($event) class=\"md-raised md-primary\" hide show-gt-sm><md-tooltip md-direction=top>Ajouter de nouveaux dossiers</md-tooltip><md-icon class=md-18>add</md-icon>Ajouter</md-button><md-button ng-click=customers.create($event) class=\"md-fab md-mini md-primary\" hide-gt-sm aria-label=\"Ajouter de nouveaux dossiers\"><md-icon class=md-24>add</md-icon></md-button></div><p class=md-body-1>Listez l\'ensemble des dossiers sur lesquels travaillent vos collaborateurs.<br>Les dossiers peuvent représenter vos clients ou divers projets.</p><md-whiteframe ng-if=\"group.customers.length == 0\" class=\"md-whiteframe-1dp md-accent\" layout-padding layout=column layout-align=\"center center\"><div class=\"md-headline input-error\">Aucun dossier</div><div>Ajouter des dossiers pour permettre la saisie de tâches.</div></md-whiteframe><md-whiteframe class=md-whiteframe-1dp flex layout-padding layout=\"space-around center\" ng-if=\"group.customers.length > 0\"><div layout=row flex><md-input-container flex><label>Sélectionnez un dossier</label><md-select ng-model=customers.selected><md-option ng-repeat=\"c in group.customers\" ng-if=c.active value=\"{{ c.id }}\">{{ c.name }}</md-option></md-select></md-input-container><md-button ng-if=customers.selected ng-click=customers.edit($event) class=\"md-fab md-mini md-accent\" aria-label=\"Modifier le dossier\"><md-tooltip md-direction=top>Modifier le dossier</md-tooltip><md-icon class=md-24>edit</md-icon></md-button></div></md-whiteframe></div></div>");
$templateCache.put("horodata/widgets/configuration/customers_create_form.html","<md-dialog aria-label=\"Ajouter de nouveaux dossiers\" flex=40><form name=newCustomerForm><app-widgets-common-dialog-toolbar>Ajouter de nouveaux dossiers</app-widgets-common-dialog-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.customers}\"><label>Nom des dossiers (un dossier par ligne)</label> <textarea md-no-autogrow ng-model=customers.current.customers rows=5 md-select-on-focus></textarea> <small ng-if=errors.customers class=input-error>{{ errors.customers }}</small></md-input-container></div></md-dialog-content><app-widgets-common-dialog-actions><md-button ng-click=create() class=\"md-primary md-raised\">Ajouter</md-button></app-widgets-common-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/configuration/customers_edit_form.html","<md-dialog aria-label=\"Modifier le dossier\" flex=40><form name=editCustomerForm><app-widgets-common-dialog-toolbar>Modifier le dossier</app-widgets-common-dialog-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.name}\"><label>Nom du dossier</label> <input md-maxlength=40 type=text ng-model=customers.current.name> <small ng-if=errors.name class=input-error>{{ errors.name }}</small></md-input-container></div></md-dialog-content><app-widgets-common-dialog-actions><md-button ng-click=edit() class=\"md-primary md-raised\">Modifier</md-button><md-button ng-click=delete() class=\"md-warn md-raised\">Supprimer</md-button></app-widgets-common-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/configuration/delete.html","<div layout-padding><div layout=column flex-sm=70 flex-md=80 flex-gt-md=60><div layout=row layout-align=\"space-between center\"><div class=\"md-display-1 input-error\" hide show-gt-sm>Supprimer le groupe</div><div class=\"md-headline input-error\" hide-gt-sm>Supprimer le groupe</div><md-button ng-click=deleteGroup($event) class=\"md-raised md-warn\" hide show-gt-sm><md-tooltip md-direction=top>Supprimer le groupe</md-tooltip><md-icon class=md-18>delete</md-icon>Supprimer</md-button><md-button ng-click=deleteGroup($event) class=\"md-fab md-mini md-warn\" hide-gt-sm aria-label=\"Supprimer le groupe\"><md-icon class=md-24>delete</md-icon></md-button></div><p class=md-body-1>Toutes les donnees seront definitivement effacees. Cette operation est irreversible.<br>Pensez a exporter les saisies avant de realiser cette operation.<br></p></div></div>");
$templateCache.put("horodata/widgets/configuration/delete_confirm.html","<md-dialog aria-label=\"Supprimer le groupe\" flex=40><app-widgets-common-dialog-toolbar warn=true>Supprimer le groupe</app-widgets-common-dialog-toolbar><md-dialog-content><div class=md-dialog-content><div class=\"md-headline input-error\" layout=row layout-align=\"center center\" flex><span>Cette operation est irreversible!</span></div><br><br><div>Toutes les donnees seront definitivement effacees!<br>Cela inclus:<ul><li>Les saisies</li><li>Les exports</li><li>Les statistiques</li><li>Les utilisateurs</li><li>Les dossiers</li><li>Les taches</li></ul><br>Pensez a exporter les saisies avant de realiser cette operation!</div></div></md-dialog-content><app-widgets-common-dialog-actions><md-button ng-click=delete() class=\"md-primary md-raised md-warn\">Supprimer</md-button></app-widgets-common-dialog-actions></md-dialog>");
$templateCache.put("horodata/widgets/configuration/guests.html","<div layout-padding><div layout=column flex-sm=70 flex-md=80 flex-gt-md=60><div layout=row layout-align=\"space-between center\"><div class=md-display-1 hide show-gt-sm>Utilisateurs</div><div class=md-headline hide-gt-sm>Utilisateurs</div><md-button ng-click=guests.create($event) class=\"md-raised md-primary\" hide show-gt-sm><md-tooltip md-direction=top>Ajouter un nouvel utilisateur</md-tooltip><md-icon class=md-18>add</md-icon>Ajouter</md-button><md-button ng-click=guests.create($event) class=\"md-fab md-mini md-primary\" hide-gt-sm aria-label=\"Ajouter un nouvel utilisateur\"><md-icon class=md-24>add</md-icon></md-button></div><p class=md-body-1>Invitez vos collaborateurs à saisir leurs tâches dans Horodata.<br></p><md-whiteframe class=md-whiteframe-1dp flex layout-padding layout=\"space-around center\" ng-if=\"group.guests.length > 0\"><div layout=row flex><md-input-container flex><label>Selectionnez un utilisateur</label><md-select ng-model=guests.selected><md-option ng-repeat=\"g in group.guests\" ng-if=g.active value=\"{{ g.id }}\"><span ng-if=g.full_name>{{ g.full_name }}&nbsp;</span> <span ng-class=\"{\'md-caption\': g.full_name}\">{{ g.email }} &nbsp;</span> <strong ng-if=\"g.admin && g.user_id != group.owner\" class=\"md-caption text-accent\">(Admin) &nbsp;</strong> <strong ng-if=\"g.user_id == group.owner\" class=\"md-caption text-accent\">(Porprietaire) &nbsp;</strong></md-option></md-select></md-input-container><md-button ng-if=guests.selected ng-click=guests.edit($event) class=\"md-fab md-mini md-accent\" aria-label=\"Editer une utilisateur\"><md-tooltip md-direction=top>Modifier l\'utilisateur</md-tooltip><md-icon class=md-24>edit</md-icon></md-button></div></md-whiteframe></div></div>");
$templateCache.put("horodata/widgets/configuration/guests_create_form.html","<md-dialog aria-label=\"Ajouter un nouvel utilisateur\" flex=40><form name=newUserForm><app-widgets-common-dialog-toolbar>Ajouter un nouvel utilisateur</app-widgets-common-dialog-toolbar><md-dialog-content><div class=md-dialog-content><app-widgets-common-quota-error></app-widgets-common-quota-error><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.email}\"><label>Adresse email</label> <input type=text md-maxlength=100 ng-model=guests.current.email> <small ng-if=errors.email class=input-error>{{ errors.email }}</small></md-input-container><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.rate}\"><label>Taux horaire</label> <input type=number ng-model=guests.current.rate ng-value=0> <small ng-if=errors.rate class=input-error>{{ errors.rate }}</small></md-input-container><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.admin}\"><md-switch ng-model=guests.current.admin aria-label=\"Droit administrateur\">Droit administrateur</md-switch><div class=input-hint>Cochez la case, si vous souhaitez autoriser cet utilisateur, à accéder à la configuration et aux données de tous les utilisateurs du groupe.</div><small ng-if=errors.admin class=input-error>{{ errors.admin }}</small></md-input-container></div></md-dialog-content><app-widgets-common-dialog-actions><md-button ng-click=create() class=\"md-primary md-raised\">Ajouter</md-button></app-widgets-common-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/configuration/guests_edit_form.html","<md-dialog aria-label=\"Modifier l\'utilisateur\" flex=40><form name=newUserForm><app-widgets-common-dialog-toolbar>Modifier l\'utilisateur <span ng-if=!guests.current.full_name>{{ guests.current.email }}</span> <span ng-if=guests.current.full_name>{{ guests.current.full_name }}</span></app-widgets-common-dialog-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.rate}\"><label>Taux horaire</label> <input type=number ng-model=guests.current.rate ng-value=0> <small ng-if=errors.rate class=input-error>{{ errors.rate }}</small></md-input-container><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.admin}\"><md-switch ng-model=guests.current.admin aria-label=\"Droit administrateur\" ng-disabled=\"guests.current.user_id == group.owner\">Droit administrateur</md-switch><div class=input-hint>Cochez la case, si vous souhaitez autoriser cet utilisateur, à accéder à la configuration et aux données de tous les utilisateurs du groupe.</div><small ng-if=errors.admin class=input-error>{{ errors.admin }}</small></md-input-container></div></md-dialog-content><app-widgets-common-dialog-actions><md-button ng-click=edit() class=\"md-primary md-raised\">Modifier</md-button><md-button ng-click=delete() ng-if=\"guests.current.user_id != group.owner\" class=\"md-warn md-raised\">Supprimer</md-button></app-widgets-common-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/configuration/root.html","<div><app-widgets-configuration-customers></app-widgets-configuration-customers><md-divider></md-divider><app-widgets-configuration-tasks></app-widgets-configuration-tasks><md-divider></md-divider><app-widgets-configuration-guests></app-widgets-configuration-guests><md-divider ng-if=\"group.owner == user.id\"></md-divider><app-widgets-configuration-delete ng-if=\"group.owner == user.id\"></app-widgets-configuration-delete><div class=free-space></div></div>");
$templateCache.put("horodata/widgets/configuration/tasks.html","<div layout-padding><div layout=column flex-sm=70 flex-md=80 flex-gt-md=60><div layout=row layout-align=\"space-between center\"><div class=md-display-1 hide show-gt-sm>Types de tâche</div><div class=md-headline hide-gt-sm>Types de tâche</div><md-button ng-click=tasks.create($event) class=\"md-raised md-primary\" hide show-gt-sm><md-tooltip md-direction=top>Ajouter un type de tâche</md-tooltip><md-icon class=md-18>add</md-icon>Ajouter</md-button><md-button ng-click=tasks.create($event) class=\"md-fab md-primary md-mini\" hide-gt-sm aria-label=\"Ajouter un type de tâche\"><md-icon class=md-24>add</md-icon></md-button></div><p class=md-body-1>Listez les types de tâche qu\'accomplissent vos collaborateurs.</p><md-whiteframe ng-if=\"group.tasks.length == 0\" class=\"md-whiteframe-1dp md-accent\" layout-padding layout=column layout-align=\"center center\"><div class=\"md-headline input-error\">Aucun type de tâche</div><div>Ajouter des types pour permettre la saisie de tâches.</div></md-whiteframe><md-whiteframe class=md-whiteframe-1dp flex layout-padding layout=\"space-around center\" ng-if=\"group.tasks.length > 0\"><div layout=row flex><md-input-container flex><label>Sélectionnez un type de tâche</label><md-select ng-model=tasks.selected><md-option ng-repeat=\"t in group.tasks\" ng-if=t.active value=\"{{ t.id }}\">{{ t.name }} <small ng-if=t.comment_mandatory>(Commentaire obligatoire)</small></md-option></md-select></md-input-container><md-button ng-if=tasks.selected ng-click=tasks.edit($event) class=\"md-fab md-mini md-accent\" aria-label=\"Editer une tâche\"><md-tooltip md-direction=top>Modifier la tâche</md-tooltip><md-icon class=md-24>edit</md-icon></md-button></div></md-whiteframe></div></div>");
$templateCache.put("horodata/widgets/configuration/tasks_form.html","<md-dialog aria-label=\"Ajouter un nouveau type de tâche\" flex=40><form name=newTaskForm><app-widgets-common-dialog-toolbar><span ng-if=!tasks.current.id>Ajouter un nouveau type de tâche</span> <span ng-if=tasks.current.id>Modifier la tâche {{ tasks.current.name }}</span></app-widgets-common-dialog-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.name}\"><label>Nom de la tâche</label> <input md-maxlength=30 type=text ng-model=tasks.current.name> <small ng-if=errors.name class=input-error>{{ errors.name }}</small></md-input-container><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.admin}\"><md-switch ng-model=tasks.current.comment_mandatory aria-label=\"Commentaire obligatoire\">Commentaire obligatoire</md-switch><div class=input-hint>Cochez la case, si vous souhaitez rendre le commentaire obligatoire.</div></md-input-container></div></md-dialog-content><app-widgets-common-dialog-actions><md-button ng-if=!tasks.current.id ng-click=create() class=\"md-primary md-raised\">Ajouter</md-button><md-button ng-if=tasks.current.id ng-click=edit() class=\"md-primary md-raised\">Modifier</md-button><md-button ng-if=tasks.current.id ng-click=delete() class=\"md-warn md-raised\">Supprimer</md-button></app-widgets-common-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/detail/dialog.html","<md-dialog aria-label=\"Detail de la tâche\" flex=40><form name=editTaskForm><md-toolbar><div class=md-toolbar-tools><h2>Detail de la tâche</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon class=md-24>close</md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><div layout=column><app-widgets-detail-meta></app-widgets-detail-meta><md-input-container flex><label>Sélectionnez un dossier</label><md-select ng-model=detailJob.customer_id><md-option ng-repeat=\"c in group.customers\" value=\"{{ c.id }}\">{{ c.name }}</md-option></md-select><small ng-if=errors.customer class=input-error>{{ errors.customer }}</small></md-input-container><md-input-container flex><label>Sélectionnez une tâche</label><md-select ng-model=detailJob.task_id><md-option ng-repeat=\"t in group.tasks\" value=\"{{ t.id }}\">{{ t.name }}</md-option></md-select><small ng-if=errors.task class=input-error>{{ errors.task }}</small></md-input-container><div layout=row layout-align=\"space-between center\"><md-input-container flex><label>Durée en heures</label><md-select ng-model=detailJob.hours><md-option ng-repeat=\"h in hours\" value=\"{{ h }}\">{{ h }} <span ng-if=\"h > 1\">heures</span> <span ng-if=\"h <= 1\">heure</span></md-option></md-select><small ng-if=errors.duration class=input-error>{{ errors.duration }}</small></md-input-container><md-input-container flex><label>Durée en minutes</label><md-select ng-model=detailJob.minutes><md-option ng-repeat=\"m in minutes\" value=\"{{ m }}\">{{ m }} <span ng-if=\"m > 1\">minutes</span> <span ng-if=\"m <= 1\">minute</span></md-option></md-select><small ng-if=errors.duration class=input-error>{{ errors.duration }}</small></md-input-container></div><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.comment}\"><label>Commentaire</label> <textarea md-no-autogrow ng-model=detailJob.comment rows=3 md-select-on-focus></textarea> <small ng-if=errors.comment class=input-error>{{ errors.comment }}</small></md-input-container></div></div></md-dialog-content><app-widgets-common-dialog-actions><md-button ng-click=send() ng-if=canEdit class=\"md-primary md-raised\">Editer</md-button><md-button ng-click=delete() ng-if=canEdit class=\"md-raised md-warn\">Supprimer</md-button></app-widgets-common-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/detail/meta.html","<div><div layout=row><div ng-if=isAdmin layout=column flex><span class=md-caption>Utilisateur</span><div>{{ guests[detailJob.creator_id].full_name }}</div></div><div layout=column flex><span class=md-caption>Cree le</span><div>{{ detailJob.created | Date }}</div></div><br><br><br></div></div>");
$templateCache.put("horodata/widgets/listing/filter.html","<div><div layout=row layout-align=\"space-around center\" hide show-gt-sm layout-padding><div layout=row flex=25 layout-align=\"center center\"><span>du</span><md-datepicker ng-model=search.begin md-placeholder=\"Date début\" md-max-date=search.end></md-datepicker></div><div layout=row flex=25 layout-align=\"center center\"><span>au</span><md-datepicker ng-model=search.end md-placeholder=\"Date fin\" md-max-date=today></md-datepicker></div><div flex=25><div layout=row layout-align=\"space-between center\"><md-input-container class=no-margin flex><label>Dossier</label><md-select ng-model=search.customer><md-option ng-repeat=\"c in group.customers\" ng-if=c.active value=\"{{ c.id }}\">{{ c.name }}</md-option></md-select></md-input-container><md-button ng-click=\"search.customer = null\" ng-if=search.customer class=\"md-fab md-mini\" aria-label=désélectionner><md-icon>close</md-icon></md-button></div></div><div flex=25 ng-if=isAdmin><div layout=row layout-align=\"space-between center\"><md-input-container class=no-margin flex><label>Utilisateur</label><md-select ng-model=search.guest><md-option ng-repeat=\"u in group.guests\" ng-if=\"u.full_name && u.active\" value=\"{{ u.id }}\">{{ u.full_name }}</md-option></md-select></md-input-container><md-button ng-click=\"search.guest = null\" ng-if=search.guest class=\"md-fab md-mini\" aria-label=désélectionner><md-icon>close</md-icon></md-button></div></div></div><md-divider></md-divider></div>");
$templateCache.put("horodata/widgets/listing/root.html","<div><app-widgets-listing-filter></app-widgets-listing-filter><div ng-if=\"listing.list.length > 0\"><md-list class=md-body-1><md-list-item class=md-body-2 hide show-gt-sm><div layout=row layout-align=\"space-between center\" flex><div flex=10>Jour</div><div ng-if=isAdmin flex=30>Utilisateur</div><div flex=25>Dossier</div><div flex=25>Tâche</div><div flex=5>Durée</div></div><md-divider></md-divider></md-list-item><md-list-item ng-repeat=\"i in listing.list\" ng-click=\"showDetail($event, i)\" hide show-gt-sm><div layout=row layout-align=\"space-between center\" flex><div flex=10>{{ i.created | Day }}</div><div flex=30 ng-if=isAdmin class=contact-item><md-list-item class=\"md-2-line contact-item\"><img ng-src=\"{{ guests[i.creator_id].picture | Profile }}\" class=md-avatar alt=\"{{ guests[i.creator_id].full_name }}\"><div class=\"md-list-item-text compact\"><h3>{{ guests[i.creator_id].full_name }}</h3><p>{{ guests[i.creator_id].email }}</p></div></md-list-item></div><div flex=25>{{ customers[i.customer_id].name }}</div><div flex=25>{{ tasks[i.task_id].name }}</div><div flex=5>{{ i.duration | Duration }}</div></div><md-divider></md-divider></md-list-item><md-list-item ng-repeat=\"i in listing.list\" class=md-2-line ng-click=\"showDetail($event, i)\" hide-gt-sm><div class=md-list-item-text layout=column><div layout=row layout-align=\"space-between center\"><h3>{{ customers[i.customer_id].name }}</h3><h4>{{ i.created | Day }} - <span class=text-primary>{{ i.duration | Duration }}</span></h4></div><h4>{{ tasks[i.task_id].name }}</h4><span class=text-accent>{{ guests[i.creator_id].full_name }}</span></div><md-divider></md-divider></md-list-item></md-list></div><div ng-if=\"!listing.loading && listing.hasMore()\" layout-padding layout=row layout-align=\"center center\"><md-button ng-click=listing.next() class=\"md-raised md-primary\"><md-icon>arrow_downward</md-icon><span>Resultats suivants</span><md-icon>arrow_downward</md-icon></md-button></div><div ng-if=listing.loading layout=column layout-align=\"center center\"><md-progress-circular md-mode=indeterminate md-diameter=150></md-progress-circular><p>Chargement, veuillez patienter.</p></div><div ng-if=\"!listing.loading && listing.list.length > 0\" layout-padding layout-align=\"center center\" layout=row flex><span class=md-body-1>resultats {{ listing.list.length }} / {{ listing.total }}</span></div></div>");
$templateCache.put("horodata/widgets/stats/customer_time.html","<div layout=column><app-widgets-loading ng-if=stats.loading()></app-widgets-loading><div ng-if=\"!stats.loading() && data.length > 0\" flex><div id=chart></div><c3chart bindto-id=chart ng-if=data><chart-column ng-repeat=\"i in data\" column-id=\"{{ i.customer_id }}\" column-name=\"{{ customers[i.customer_id].name }}\" column-values=\"{{ i.duration }}\" column-type=pie></chart-column></c3chart></div><app-widgets-stats-no-data ng-if=\"!stats.loading() && data.length == 0\"></app-widgets-stats-no-data></div>");
$templateCache.put("horodata/widgets/stats/guest_time.html","<div layout=column><app-widgets-loading ng-if=stats.loading()></app-widgets-loading><div ng-if=\"!stats.loading() && data.length > 0\" flex><div id=chart></div><c3chart bindto-id=chart ng-if=data><chart-column ng-repeat=\"i in data\" column-id=\"{{ i.guest_id }}\" column-name=\"{{ guests[i.guest_id].full_name }}\" column-values=\"{{ i.duration }}\" column-type=pie></chart-column></c3chart></div><app-widgets-stats-no-data ng-if=\"!stats.loading() && data.length == 0\"></app-widgets-stats-no-data></div>");
$templateCache.put("horodata/widgets/stats/no_data.html","<div layout=row flex layout-align=\"center center\"><md-whiteframe class=\"md-whiteframe-1dp md-accent\" flex=100 flex-sm=70 flex-md=60 flex-gt-md=50 layout-padding layout-margin layout=column><div layout-padding layout-margin layout=column layout-align=\"center center\"><div><div class=md-headline>Aucune donnee pour la periode</div><p>La periode du {{ search.begin | Day }} au {{ search.end | Day }} ne comporte aucune saisie.</p></div></div></md-whiteframe></div>");
$templateCache.put("horodata/widgets/stats/root.html","<div layout=column layout-padding><div hide-md hide-lg layout=row><md-input-container flex><md-select ng-model=selected placeholder=\"Choisissez une statistique\"><md-option ng-value=s.id ng-repeat=\"s in availableStats\">{{ s.label }}</md-option></md-select></md-input-container></div><div hide show-xs layout=row flex layout-align=\"center center\"><span>du</span><md-datepicker ng-model=filter.begin md-placeholder=\"Date début\" md-max-date=filter.end flex></md-datepicker></div><div hide show-xs layout=row flex layout-align=\"center center\"><span>au</span><md-datepicker ng-model=filter.end md-placeholder=\"Date fin\" md-max-date=today flex></md-datepicker></div><div hide-xs layout=row><md-input-container hide show-gt-sm flex><md-select ng-model=selected placeholder=\"Choisissez une statistique\"><md-option ng-value=s.id ng-repeat=\"s in availableStats\">{{ s.label }}</md-option></md-select></md-input-container><div layout=row flex=50 flex-gt-sm=25 layout-align=\"center center\"><span>du</span><md-datepicker ng-model=filter.begin md-placeholder=\"Date début\" md-max-date=filter.end></md-datepicker></div><div layout=row flex=50 flex-gt-sm=25 layout-align=\"center center\"><span>au</span><md-datepicker ng-model=filter.end md-placeholder=\"Date fin\" md-max-date=today flex></md-datepicker></div></div><div ng-switch=selected flex><app-widgets-stats-customer-time ng-switch-when=customer_time></app-widgets-stats-customer-time><app-widgets-stats-task-time ng-switch-when=task_time></app-widgets-stats-task-time><app-widgets-stats-guest-time ng-switch-when=guest_time></app-widgets-stats-guest-time></div></div>");
$templateCache.put("horodata/widgets/stats/task_time.html","<div layout=column><app-widgets-loading ng-if=stats.loading()></app-widgets-loading><div ng-if=\"!stats.loading() && data.length > 0\" flex><div id=chart></div><c3chart bindto-id=chart ng-if=data><chart-column ng-repeat=\"i in data\" column-id=\"{{ i.task_id }}\" column-name=\"{{ tasks[i.task_id].name }}\" column-values=\"{{ i.duration }}\" column-type=pie></chart-column></c3chart></div><app-widgets-stats-no-data ng-if=\"!stats.loading() && data.length == 0\"></app-widgets-stats-no-data></div>");}]);
angular.module("horodata").directive("appMenuSidenav", [
  "$mdSidenav", "$http", "$location", "apiService", "$routeParams", "groupService", function($mdSidenav, $http, $location, apiService, $routeParams, groupService) {
    var l;
    l = function(scope, elem) {
      scope.toggleSidenav = function() {
        return $mdSidenav("sidenav").toggle();
      };
      scope.closeSidenav = function() {
        return $mdSidenav("sidenav").close();
      };
      groupService.fetch();
      scope.groups = function() {
        return groupService.listing();
      };
      scope.changeGroup = function(url) {
        scope.closeSidenav();
        return $location.path(url);
      };
      return scope.$on("$routeChangeSuccess", function() {
        scope.currentGroupUrl = $routeParams.group;
        return scope.closeSidenav();
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
  "titleService", "userService", "homeService", "popupService", "$mdMedia", "$mdBottomSheet", function(titleService, userService, homeService, popupService, $mdMedia, $mdBottomSheet) {
    var l;
    l = function(scope, elem) {
      scope.MainTitle = titleService.get;
      userService.get(function(u) {
        return scope.user = u;
      });
      scope.home = homeService.get();
      scope.openMenu = function($mdOpenMenu, $event) {
        if ($mdMedia('xs') || $mdMedia('sm')) {
          return $mdBottomSheet.show({
            templateUrl: "horodata/menu/bottom_sheet.html"
          });
        } else {
          return $mdOpenMenu($event);
        }
      };
      scope.showProfile = function(ev) {
        return popupService("horodata/views/profile.html", "Profile", scope, ev);
      };
      return scope.showQuotas = function(ev) {
        return popupService("horodata/views/quotas.html", "Quotas", scope, ev);
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

angular.module("horodata").filter("Date", [
  function() {
    return function(input) {
      return moment(input).format('LLLL');
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

angular.module("horodata").filter("Profile", [
  "homeService", "staticService", function(homeService, staticService) {
    return function(id) {
      if ((id != null) && id !== "") {
        return (homeService.get()) + "/profiles/" + id + ".jpg";
      } else {
        return (staticService.get()) + "/profile-default.png";
      }
    };
  }
]);

angular.module('horodata').factory("groupService", [
  "apiService", "$http", function(apiService, $http) {
    var current, fetchListing, groups;
    current = null;
    groups = [];
    fetchListing = function() {
      return $http.get((apiService.get()) + "/groups").then(function(resp) {
        return groups = resp.data.data.results;
      });
    };
    return {
      set: function(group) {
        return current = group;
      },
      get: function() {
        return current;
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
    var Listing, listingInstance;
    Listing = (function() {
      function Listing(groupUrl1, begin, end, customer, guest) {
        this.groupUrl = groupUrl1;
        this.fetch = bind(this.fetch, this);
        this.size = 5;
        this.page = 1;
        this.list = [];
        this.loading = false;
        this.total = -1;
        this.params = {
          begin: moment(begin).format('YYYY-MM-DD'),
          end: moment(end).format('YYYY-MM-DD'),
          customer: customer,
          guest: guest
        };
      }

      Listing.prototype.hasMore = function() {
        return (this.page * this.size) < this.total;
      };

      Listing.prototype.next = function() {
        if (this.hasMore()) {
          return this.fetch(this.page + 1);
        }
      };

      Listing.prototype.fetch = function(page) {
        var params;
        if (this.loading) {
          return;
        }
        this.loading = true;
        params = _.cloneDeep(this.params);
        if (page == null) {
          page = this.page;
        }
        this.page = page;
        params.offset = (page - 1) * this.size;
        params.size = this.size;
        return $http.get((apiService.get()) + "/groups/" + this.groupUrl + "/jobs", {
          params: params
        }).then((function(_this) {
          return function(resp) {
            var i, j, len, ref;
            ref = resp.data.data.results;
            for (j = 0, len = ref.length; j < len; j++) {
              i = ref[j];
              _this.list.push(i);
            }
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
    listingInstance = {};
    return {
      data: function() {
        if ((listing.list == null) || listing.loading) {
          null;
        }
        return listing.list;
      },
      get: function() {
        return listingInstance;
      },
      search: function(groupUrl, params) {
        return listingInstance = new Listing(groupUrl, params.begin, params.end, params.customer, params.guest);
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

angular.module('horodata').factory("staticService", [
  function() {
    var root;
    root = document.getElementsByTagName("static")[0].getAttribute("href");
    return {
      get: function() {
        return root;
      }
    };
  }
]);

angular.module('horodata').factory("statsService", [
  "apiService", "$http", "statsFilterService", function(apiService, $http, statsFilterService) {
    var fetch, loading;
    loading = false;
    fetch = function(group, stat, cb) {
      var begin, end;
      begin = moment(statsFilterService.begin).format('YYYY-MM-DD');
      end = moment(statsFilterService.end).format('YYYY-MM-DD');
      loading = true;
      return $http.get((apiService.get()) + "/groups/" + group + "/stats/" + stat, {
        params: {
          begin: begin,
          end: end
        }
      }).then(function(resp) {
        loading = false;
        return cb(resp.data.data);
      }, function(resp) {
        cd(null);
        return loading = false;
      });
    };
    return {
      fetch: function(group, stat, params, cb) {
        return fetch(group, stat, params, cb);
      },
      data: function() {
        return data;
      },
      loading: function() {
        return loading;
      }
    };
  }
]);

angular.module('horodata').factory("statsFilterService", [
  function() {
    var begin, end, urlParams;
    begin = moment().subtract(1, 'months').toDate();
    end = new Date();
    urlParams = function() {
      return "?begin=" + (moment(begin).format('YYYY-MM-DD')) + "&end=" + (moment(end).format('YYYY-MM-DD'));
    };
    return {
      begin: begin,
      end: end,
      urlParams: urlParams
    };
  }
]);

angular.module('horodata').factory("tabsService", [
  "$rootScope", function($rootScope) {
    var current;
    current = null;
    $rootScope.$on("$routeChangeStart", function() {
      return current = null;
    });
    return {
      get: function() {
        return current;
      },
      set: function(tab) {
        return current = tab;
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
  "$http", "$routeParams", "$scope", "titleService", "userService", "apiService", "groupService", "popupService", "listingService", "tabsService", function($http, $routeParams, $scope, titleService, userService, apiService, groupService, popupService, listingService, tabsService) {
    var getGroup;
    $scope.isGroupView = true;
    $scope.isAdmin = false;
    getGroup = function() {
      return $http.get((apiService.get()) + "/groups/" + $routeParams.group).then(function(resp) {
        $scope.group = resp.data.data;
        groupService.set($scope.group);
        $scope.isAdmin = $scope.group.guests != null;
        $scope.isOwner = $scope.user.id === $scope.group.owner;
        $scope.tasks = _.keyBy($scope.group.tasks, 'id');
        $scope.customers = _.keyBy($scope.group.customers, 'id');
        $scope.guests = _.keyBy($scope.group.guests, 'id');
        return titleService.set($scope.group.name, true);
      });
    };
    userService.get(function(u) {
      $scope.user = u;
      return getGroup();
    });
    $scope.$on("group.reload", function(e) {
      e.stopPropagation();
      return getGroup();
    });
    $scope.selectTab = function(i) {
      $scope.selectedTab = i;
      return tabsService.set(i);
    };
    $scope.$watch("selectedTab", function(v, o) {
      if (v !== o) {
        return $scope.selectTab(v);
      }
    });
    return $scope.selectTab(0);
  }
]);

angular.module("horodata").controller("Index", [
  "$http", "$scope", "userService", "titleService", function($http, $scope, userService, titleService) {
    return titleService.set("Accueil");
  }
]);

angular.module("horodata").controller("Profile", [
  "$scope", "$mdDialog", "$mdToast", "$http", "apiService", "userService", function($scope, $mdDialog, $mdToast, $http, apiService, userService) {
    $scope.errors = null;
    $scope.loading = false;
    $scope.name = $scope.user.name;
    return $scope.send = function() {
      return $scope.loading = true;
    };
  }
]);

angular.module("horodata").controller("Quotas", [
  "$scope", "$mdDialog", "$http", "apiService", function($scope, $mdDialog, $http, apiService) {
    $scope.loading = true;
    $http.get((apiService.get()) + "/users/me/quotas").then(function(resp) {
      $scope.loading = false;
      return $scope.quotas = resp.data.data;
    }, function(resp) {
      return $scope.loading = false;
    });
    return $scope.send = function() {
      return $scope.loading = true;
    };
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

angular.module("horodata").directive("appWidgetsLoading", [
  function() {
    return {
      replace: true,
      restrict: "E",
      templateUrl: "horodata/widgets/loading.html"
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
  "$scope", "$mdDialog", "$mdToast", "$http", "$location", "apiService", "groupService", function($scope, $mdDialog, $mdToast, $http, $location, apiService, groupService) {
    $scope.name = "";
    $scope.loading = false;
    $scope.errors = null;
    $scope.quotaError = null;
    return $scope.send = function() {
      $scope.loading = true;
      return $http.post((apiService.get()) + "/groups", {
        name: $scope.name
      }).then(function(resp) {
        var group;
        group = resp.data.data;
        $mdDialog.hide();
        $mdToast.showSimple("Nouveau groupe '" + group.name + "' créé");
        $location.path("/" + group.url);
        return groupService.fetch();
      }, function(resp) {
        $scope.loading = false;
        if (resp.status === 429 && _.get(resp, "data.errors.type") === "quota") {
          return $scope.quotaError = resp.data.errors;
        } else {
          return $scope.errors = resp.data.errors;
        }
      });
    };
  }
]);

angular.module("horodata").directive("appWidgetsQuota", [
  function() {
    var l;
    l = function(scope) {
      var genPercent;
      genPercent = function() {
        return scope.percent = Math.floor(scope.current / scope.max * 100);
      };
      scope.$watch("current", function() {
        return genPercent();
      });
      return scope.$watch("max", function() {
        return genPercent();
      });
    };
    return {
      link: l,
      scope: {
        label: "@",
        current: "=",
        max: "="
      },
      restrict: "E",
      templateUrl: "horodata/widgets/quota.html"
    };
  }
]);

angular.module("horodata").directive("appWidgetsBigButton", [
  "tabsService", "popupService", function(tabsService, popupService) {
    var l;
    l = function(scope) {
      scope.currentTab = function() {
        switch (tabsService.get()) {
          case 0:
            return "jobs";
          case 1:
            return "export";
          default:
            return null;
        }
      };
      return scope.newDialog = function(ev) {
        if (scope.currentTab() === "jobs") {
          popupService("horodata/widgets/big_button/new_task.html", "newTaskDialog", scope, ev);
        }
        if (scope.currentTab() === "export") {
          return popupService("horodata/widgets/big_button/export.html", "exportDialog", scope, ev);
        }
      };
    };
    return {
      link: l,
      restrict: "E",
      templateUrl: "horodata/widgets/big_button/root.html"
    };
  }
]);

angular.module("horodata").controller("newTaskDialog", [
  "$scope", "$mdDialog", "$mdToast", "$http", "$location", "apiService", "listingService", "groupService", function($scope, $mdDialog, $mdToast, $http, $location, apiService, listingService, groupService) {
    var x;
    $scope.group = groupService.get();
    $scope.task = {
      minutes: 0,
      hours: 0
    };
    $scope.errors = null;
    $scope.loading = false;
    $scope.quotaError = null;
    $scope.hours = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12];
    $scope.minutes = (function() {
      var i, results;
      results = [];
      for (x = i = 0; i <= 55; x = i += 5) {
        results.push(x);
      }
      return results;
    })();
    return $scope.send = function() {
      var task;
      task = {
        duration: $scope.task.hours * 3600 + $scope.task.minutes * 60,
        task: parseInt($scope.task.task),
        customer: parseInt($scope.task.customer),
        comment: $scope.task.comment
      };
      $scope.loading = true;
      return $http.post((apiService.get()) + "/groups/" + $scope.group.url + "/jobs", task).then(function(resp) {
        $mdDialog.hide();
        $mdToast.showSimple("Nouvelle tâche saisie");
        return listingService.listing().fetch();
      }, function(resp) {
        $scope.loading = false;
        if (resp.status === 429 && _.get(resp, "data.errors.type") === "quota") {
          return $scope.quotaError = resp.data.errors;
        } else {
          return $scope.errors = resp.data.errors;
        }
      });
    };
  }
]);

angular.module("horodata").controller("exportDialog", [
  "$scope", "$mdDialog", "apiService", "groupService", "statsFilterService", function($scope, $mdDialog, apiService, groupService, statsFilterService) {
    $scope.filter = statsFilterService;
    $scope.group = groupService.get();
    $scope["export"] = {
      fileType: "xlsx"
    };
    return $scope.url = (apiService.get()) + "/groups/" + $scope.group.url + "/export";
  }
]);

angular.module("horodata").directive("appWidgetsCommonDialogActions", [
  "$mdDialog", function($mdDialog) {
    var l;
    l = function(scope) {
      return scope.hide = function() {
        return $mdDialog.hide();
      };
    };
    return {
      link: l,
      transclude: true,
      replace: true,
      restrict: "E",
      templateUrl: "horodata/widgets/common/dialog_actions.html"
    };
  }
]);

angular.module("horodata").directive("appWidgetsCommonDialogToolbar", [
  "$mdDialog", function($mdDialog) {
    var l;
    l = function(scope) {
      return scope.hide = function() {
        return $mdDialog.hide();
      };
    };
    return {
      link: l,
      scope: {
        warn: "="
      },
      transclude: true,
      replace: true,
      restrict: "E",
      templateUrl: "horodata/widgets/common/dialog_toolbar.html"
    };
  }
]);

angular.module("horodata").directive("appWidgetsCommonQuotaError", [
  function() {
    var l;
    l = function(scope) {};
    return {
      link: l,
      replace: true,
      restrict: "E",
      templateUrl: "horodata/widgets/common/quota_error.html"
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
    $scope.loading = false;
    update = function(t) {
      var idx;
      idx = _.findIndex($scope.group.customers, {
        id: t.id
      });
      $scope.group.customers[idx] = $scope.customers.current;
      return $scope.group.customers = _.sortBy($scope.group.customers, ["name"]);
    };
    $scope.create = function() {
      $scope.loading = true;
      return $http.post((apiService.get()) + "/groups/" + $scope.group.url + "/customers", $scope.customers.current).then(function(resp) {
        var total;
        total = resp.data.data.total;
        $scope.$emit("group.reload");
        $mdDialog.hide();
        if (total === 1) {
          return $mdToast.showSimple("1 nouveau dossier ajouté");
        } else {
          return $mdToast.showSimple(total + " nouveaux dossiers ajoutés");
        }
      }, function(resp) {
        $scope.errors = resp.data.errors;
        return $scope.loading = false;
      });
    };
    $scope.edit = function() {
      $scope.loading = true;
      return $http.put((apiService.get()) + "/groups/" + $scope.group.url + "/customers/" + $scope.customers.selected, $scope.customers.current).then(function(resp) {
        $mdDialog.hide();
        $mdToast.showSimple("Dossier '" + $scope.customers.current.name + "' modifié");
        update($scope.customers.current);
        return $scope.customers.selected = null;
      }, function(resp) {
        $scope.errors = resp.data.errors;
        return $scope.loading = false;
      });
    };
    return $scope["delete"] = function() {
      $scope.loading = true;
      return $http["delete"]((apiService.get()) + "/groups/" + $scope.group.url + "/customers/" + $scope.customers.selected).then(function(resp) {
        $mdDialog.hide();
        $mdToast.showSimple("Dossier '" + $scope.customers.current.name + "' supprimé");
        $scope.group.customers.splice(_.findIndex($scope.group.customers, {
          id: parseInt($scope.customers.selected)
        }), 1);
        return $scope.customers.selected = null;
      }, function(resp) {
        $scope.errors = resp.data.errors;
        return $scope.loading = false;
      });
    };
  }
]);

angular.module("horodata").directive("appWidgetsConfigurationDelete", [
  "popupService", function(popupService) {
    var l;
    l = function(scope, elem, attr) {
      return scope.deleteGroup = function(ev) {
        return popupService("horodata/widgets/configuration/delete_confirm.html", "appWidgetsConfigurationDeleteConfirm", scope, ev);
      };
    };
    return {
      link: l,
      restrict: "E",
      templateUrl: "horodata/widgets/configuration/delete.html"
    };
  }
]);

angular.module("horodata").controller("appWidgetsConfigurationDeleteConfirm", [
  "$scope", "$mdDialog", "$mdToast", "$http", "$location", "apiService", "groupService", function($scope, $mdDialog, $mdToast, $http, $location, apiService, groupService) {
    $scope.loading = false;
    return $scope["delete"] = function() {
      $scope.loading = true;
      return $http["delete"]((apiService.get()) + "/groups/" + $scope.group.url).then(function(resp) {
        $mdDialog.hide();
        $mdToast.showSimple("Groupe '" + $scope.group.name + "' supprimé");
        $location.path("/");
        return groupService.fetch();
      }, function(resp) {
        $scope.errors = resp.data.errors;
        return $scope.loading = false;
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
    $scope.loading = false;
    $scope.quotaError = null;
    update = function(t) {
      var idx;
      idx = _.findIndex($scope.group.guests, {
        id: t.id
      });
      $scope.group.guests[idx] = $scope.guests.current;
      return $scope.group.guests = _.sortBy($scope.group.guests, ["name"]);
    };
    $scope.create = function() {
      $scope.loading = true;
      return $http.post((apiService.get()) + "/groups/" + $scope.group.url + "/guests", $scope.guests.current).then(function(resp) {
        $scope.$emit("group.reload");
        $mdDialog.hide();
        return $mdToast.showSimple("Nouvel utilisateur '" + $scope.guests.current.email + "' ajouté");
      }, function(resp) {
        $scope.loading = false;
        if (resp.status === 429 && _.get(resp, "data.errors.type") === "quota") {
          return $scope.quotaError = resp.data.errors;
        } else {
          return $scope.errors = resp.data.errors;
        }
      });
    };
    $scope.edit = function() {
      $scope.loading = true;
      return $http.put((apiService.get()) + "/groups/" + $scope.group.url + "/guests/" + $scope.guests.selected, $scope.guests.current).then(function(resp) {
        $mdDialog.hide();
        $mdToast.showSimple("Utilisateur '" + $scope.guests.current.email + "' modifié");
        update($scope.guests.current);
        return $scope.guests.selected = null;
      }, function(resp) {
        $scope.loading = false;
        return $scope.errors = resp.data.errors;
      });
    };
    return $scope["delete"] = function() {
      $scope.loading = true;
      return $http["delete"]((apiService.get()) + "/groups/" + $scope.group.url + "/guests/" + $scope.guests.selected, $scope.guests.current).then(function(resp) {
        $mdDialog.hide();
        $mdToast.showSimple("Utilisateur '" + $scope.guests.current.email + "' supprimé");
        $scope.group.guests.splice(_.findIndex($scope.group.guests, {
          id: parseInt($scope.guests.selected)
        }), 1);
        return $scope.guests.selected = null;
      }, function(resp) {
        $scope.loading = false;
        return $scope.errors = resp.data.errors;
      });
    };
  }
]);

angular.module("horodata").directive("appWidgetsConfiguration", [
  function() {
    return {
      replace: true,
      restrict: "E",
      templateUrl: "horodata/widgets/configuration/root.html"
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
    $scope.loading = false;
    update = function(t) {
      var idx;
      idx = _.findIndex($scope.group.tasks, {
        id: t.id
      });
      $scope.group.tasks[idx] = $scope.tasks.current;
      return $scope.group.tasks = _.sortBy($scope.group.tasks, ["name"]);
    };
    $scope.create = function() {
      $scope.loading = true;
      return $http.post((apiService.get()) + "/groups/" + $scope.group.url + "/tasks", $scope.tasks.current).then(function(resp) {
        $scope.$emit("group.reload");
        $mdDialog.hide();
        return $mdToast.showSimple("Nouveau type de tâche '" + $scope.tasks.current.name + "' ajouté");
      }, function(resp) {
        $scope.loading = false;
        return $scope.errors = resp.data.errors;
      });
    };
    $scope.edit = function() {
      $scope.loading = true;
      return $http.put((apiService.get()) + "/groups/" + $scope.group.url + "/tasks/" + $scope.tasks.selected, $scope.tasks.current).then(function(resp) {
        $mdDialog.hide();
        $mdToast.showSimple("Type de tâche '" + $scope.tasks.current.name + "' modifié");
        update($scope.tasks.current);
        return $scope.tasks.selected = null;
      }, function(resp) {
        $scope.loading = false;
        return $scope.errors = resp.data.errors;
      });
    };
    return $scope["delete"] = function() {
      $scope.loading = true;
      return $http["delete"]((apiService.get()) + "/groups/" + $scope.group.url + "/tasks/" + $scope.tasks.selected, $scope.tasks.current).then(function(resp) {
        $mdDialog.hide();
        $mdToast.showSimple("Type de tâche '" + $scope.tasks.current.name + "' supprimé");
        $scope.group.tasks.splice(_.findIndex($scope.group.tasks, {
          id: parseInt($scope.tasks.selected)
        }), 1);
        return $scope.tasks.selected = null;
      }, function(resp) {
        $scope.loading = false;
        return $scope.errors = resp.data.errors;
      });
    };
  }
]);

angular.module("horodata").controller("detailDialog", [
  "$scope", "$mdDialog", "$mdToast", "$http", "$location", "apiService", "groupService", function($scope, $mdDialog, $mdToast, $http, $location, apiService, groupService) {
    var x;
    $scope.name = "";
    $scope.errors = null;
    $scope.loading = false;
    $scope.canEdit = false;
    if ($scope.isAdmin) {
      $scope.canEdit = true;
    } else if (moment($scope.detailJob.created).isSame(new Date(), "day")) {
      $scope.canEdit = true;
    }
    $scope.hours = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12];
    $scope.minutes = (function() {
      var i, results;
      results = [];
      for (x = i = 0; i <= 55; x = i += 5) {
        results.push(x);
      }
      return results;
    })();
    $scope.detailJob.hours = Math.floor($scope.detailJob.duration / 3600);
    $scope.detailJob.minutes = Math.floor(($scope.detailJob.duration % 3600) / 60);
    $scope.close = function() {
      return $mdDialog.hide();
    };
    return $scope.send = function() {
      $scope.loading = true;
      return $http.post((apiService.get()) + "/groups", {
        name: $scope.name
      }).then(function(resp) {
        var group;
        group = resp.data.data;
        $mdDialog.hide();
        $mdToast.showSimple("Nouveau groupe '" + group.name + "' créé");
        $location.path("/" + group.url);
        return groupService.fetch();
      }, function(resp) {
        $scope.loading = false;
        return $scope.errors = resp.data.errors;
      });
    };
  }
]);

angular.module("horodata").directive("appWidgetsDetailMeta", [
  function() {
    return {
      replace: true,
      restrict: "E",
      templateUrl: "horodata/widgets/detail/meta.html"
    };
  }
]);

angular.module("horodata").directive("appWidgetsListingFilter", [
  function() {
    var l;
    l = function(scope) {
      return scope.today = new Date();
    };
    return {
      link: l,
      replace: true,
      restrict: "E",
      templateUrl: "horodata/widgets/listing/filter.html"
    };
  }
]);

angular.module("horodata").directive("appWidgetsListing", [
  "listingService", "$timeout", "$location", "popupService", function(listingService, $timeout, $location, popupService) {
    var l;
    l = function(scope) {
      scope.search = {
        begin: moment().subtract(1, 'months').toDate(),
        end: new Date(),
        customer: null,
        guest: null
      };
      scope.$watch("search", function(v) {
        if (v == null) {
          return;
        }
        listingService.search(scope.group.url, v);
        scope.listing = listingService.get();
        return scope.listing.fetch(1);
      }, true);
      scope.goTo = function(page) {
        $location.search("page", page);
        return listingService.listing().fetch(page);
      };
      return scope.showDetail = function(ev, job) {
        scope.detailJob = _.cloneDeep(job);
        return popupService("horodata/widgets/detail/dialog.html", "detailDialog", scope, ev);
      };
    };
    return {
      link: l,
      replace: true,
      restrict: "E",
      templateUrl: "horodata/widgets/listing/root.html"
    };
  }
]);

angular.module("horodata").directive("appWidgetsStatsCustomerTime", [
  "statsService", "statsFilterService", function(statsService, statsFilterService) {
    var l;
    l = function(scope) {
      var update;
      scope.stats = statsService;
      scope.filter = statsFilterService;
      update = function() {
        return statsService.fetch(scope.group.url, "customer_time", (function(_this) {
          return function(data) {
            return scope.data = data;
          };
        })(this));
      };
      update();
      return scope.$watch("filter", function(v, o) {
        if (v.begin === o.begin && v.end === o.end) {
          return;
        }
        return update();
      }, true);
    };
    return {
      link: l,
      replace: true,
      restrict: "E",
      templateUrl: "horodata/widgets/stats/customer_time.html"
    };
  }
]);

angular.module("horodata").directive("appWidgetsStatsGuestTime", [
  "statsService", "statsFilterService", function(statsService, statsFilterService) {
    var l;
    l = function(scope) {
      var update;
      scope.stats = statsService;
      scope.filter = statsFilterService;
      update = function() {
        return statsService.fetch(scope.group.url, "guest_time", (function(_this) {
          return function(data) {
            return scope.data = data;
          };
        })(this));
      };
      update();
      return scope.$watch("filter", function(v, o) {
        if (v.begin === o.begin && v.end === o.end) {
          return;
        }
        return update();
      }, true);
    };
    return {
      link: l,
      replace: true,
      restrict: "E",
      templateUrl: "horodata/widgets/stats/guest_time.html"
    };
  }
]);

angular.module("horodata").directive("appWidgetsStatsNoData", [
  function() {
    return {
      replace: true,
      restrict: "E",
      templateUrl: "horodata/widgets/stats/no_data.html"
    };
  }
]);

angular.module("horodata").directive("appWidgetsStats", [
  "statsFilterService", function(statsFilterService) {
    var l;
    l = function(scope) {
      scope.filter = statsFilterService;
      scope.today = new Date();
      scope.availableStats = [
        {
          id: "customer_time",
          label: "Repartition du temps par dossier."
        }, {
          id: "task_time",
          label: "Repartition du temps par tâche."
        }, {
          id: "guest_time",
          label: "Repartition du temps par utilisateur."
        }
      ];
      return scope.selected = null;
    };
    return {
      link: l,
      replace: true,
      restrict: "E",
      templateUrl: "horodata/widgets/stats/root.html"
    };
  }
]);

angular.module("horodata").directive("appWidgetsStatsTaskTime", [
  "statsService", "statsFilterService", function(statsService, statsFilterService) {
    var l;
    l = function(scope) {
      var update;
      scope.stats = statsService;
      scope.filter = statsFilterService;
      update = function() {
        return statsService.fetch(scope.group.url, "task_time", (function(_this) {
          return function(data) {
            return scope.data = data;
          };
        })(this));
      };
      update();
      return scope.$watch("filter", function(v, o) {
        if (v.begin === o.begin && v.end === o.end) {
          return;
        }
        return update();
      }, true);
    };
    return {
      link: l,
      replace: true,
      restrict: "E",
      templateUrl: "horodata/widgets/stats/task_time.html"
    };
  }
]);

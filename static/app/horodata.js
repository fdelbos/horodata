/**
 * HoroData Javascript Interface
 * Version: 0.0.1
 * Copyright © 2016 Hyperboloide. All rights reserved.
*/
angular.module("horodata", ["ngMaterial", "ngRoute", "ngMessages"]).config([
  "$mdDateLocaleProvider", "$mdThemingProvider", "$locationProvider", "$routeProvider", function($mdDateLocaleProvider, $mdThemingProvider, $locationProvider, $routeProvider) {
    var months;
    $mdThemingProvider.theme('default').primaryPalette('blue').accentPalette('pink');
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
]);

angular.module("horodata").run(["$templateCache", function($templateCache) {$templateCache.put("horodata/menu/sidenav.html","<md-sidenav class=md-sidenav-left md-component-id=sidenav md-is-locked-open=\"$mdMedia(\'gt-sm\')\" md-disable-backdrop layout=column md-whiteframe=4><md-toolbar hide-gt-sm><div class=md-toolbar-tools><h1><i class=\"fa fa-clock-o\"></i>&nbsp; Horo Data</h1><span flex></span><md-button ng-click=toggleSidenav() class=md-icon-button><span class=\"fa fa-close\" aria-label=Fermer></span></md-button></div></md-toolbar><md-content layout-padding hide show-gt-sm><div layout=column layout-align=\"center center\"><div class=md-display-3><i class=\"fa fa-clock-o\"></i></div><div class=md-headline>Horo Data</div></div></md-content><md-content flex><section><md-subheader class=md-accent><div layout=row layout-align=\"space-between center\"><span class=md-headline>Groupes</span><app-widgets-new-group layout-align=end size=small></app-widgets-new-group></div></md-subheader><md-list flex layout=column class=md-body-1><md-list-item ng-repeat=\"group in groups\" ng-href=\"./group/{{ group.url}}\">{{group.name}}</md-list-item></md-list><section></section></section></md-content></md-sidenav>");
$templateCache.put("horodata/menu/toolbar.html","<md-toolbar><div class=md-toolbar-tools><md-button hide-gt-sm class=md-icon-button aria-label=Settings ng-click=toggleSidenav()><span class=\"fa fa-bars\"></span></md-button><h2><span>{{ MainTitle().title }}</span></h2><span flex></span><md-button class=\"md-fab md-mini\" aria-label=Favorite><span class=\"fa fa-user\"></span></md-button></div></md-toolbar>");
$templateCache.put("horodata/views/group.html","<div><app-widgets-search-bar></app-widgets-search-bar><md-content><md-tabs md-dynamic-height md-border-bottom><md-tab label=Heures></md-tab><md-tab label=Statistiques></md-tab><md-tab label=Configuration layout=column layout-margin><app-widgets-configuration-users></app-widgets-configuration-users><md-divider></md-divider><app-widgets-configuration-customers></app-widgets-configuration-customers><md-divider></md-divider><app-widgets-configuration-tasks></app-widgets-configuration-tasks></md-tab></md-tabs></md-content></div>");
$templateCache.put("horodata/views/index.html","<div><app-new-group></app-new-group></div>");
$templateCache.put("horodata/widgets/new_group.html","<md-button ng-class=\"small ? \'md-fab md-primary md-mini\' : \'md-raised md-primary\'\" ng-click=showNewGroupDialog() aria-label=\"Nouveau Groupe\"><span class=\"fa fa-plus\"></span> <span ng-if=!small>Nouveau Groupe</span></md-button>");
$templateCache.put("horodata/widgets/new_group_form.html","<md-dialog aria-label=\"Nouveau Groupe\" flex=40><form name=newGroupForm><md-toolbar><div class=md-toolbar-tools><h2>Nouveau Groupe</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon md-font-icon=\"fa fa-close\"></md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.name}\"><label>Nom du Groupe</label> <input type=text ng-model=name> <small ng-if=errors.name class=input-error>{{ errors.name }}</small></md-input-container></div></md-dialog-content><md-dialog-actions><md-button ng-click=send() class=\"md-primary md-raised\">Creer</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/search_bar.html","<div><div layout=row layout-align=\"space-around center\" hide show-gt-md layout-padding><md-datepicker flex=25 ng-model=beginDate md-placeholder=\"Date debut\" md-max-date=endDate></md-datepicker><md-datepicker flex=25 ng-model=endDate md-placeholder=\"Date fin\" md-max-date=maxDate></md-datepicker><md-input-container flex=25><label>Utilisateur</label><md-select ng-model=user><md-option ng-repeat=\"u in users\" value=\"{{ u.email }}\">{{ u.name }}</md-option></md-select></md-input-container><md-input-container flex=25><label>Dossier</label><md-select ng-model=folder><md-option ng-repeat=\"u in users\" value=\"{{ u.email }}\">{{ u.name }}</md-option></md-select></md-input-container></div><md-divider></md-divider></div>");
$templateCache.put("horodata/widgets/configuration/customers.html","<div layout-margin layout-padding><div layout=column><div layout=row layout-align=\"space-between center\"><div class=md-display-1>Dossiers</div><md-button ng-click=showConfigurationCustomersDialog() class=\"md-raised md-primary\" hide show-gt-sm><span class=\"fa fa-plus\"></span> Ajouter</md-button><md-button ng-click=showConfigurationCustomersDialog() class=\"md-fab md-primary\" hide-gt-sm aria-label=\"Ajouter un dossier\"><span class=\"fa fa-plus\"></span></md-button></div><p class=md-body-1>Un dossier peut representer un client ou un projet. Chaque nouvelle tache est attachee a un dossier.</p><div layout=row><md-input-container flex><label>Dossiers</label><md-select ng-model=selectedCustomer><md-option ng-repeat=\"u in users\" value=\"{{ u.email }}\">{{ u.name }}</md-option></md-select></md-input-container><div ng-if=selectedCustomer><md-button ng-click=showConfigurationCustomersDialog() class=\"md-raised md-accent\" hide show-gt-sm><span class=\"fa fa-pencil\"></span> Editer</md-button><md-button ng-click=showConfigurationCustomersDialog() class=\"md-fab md-accent\" hide-gt-sm aria-label=\"Editer un dossier\"><span class=\"fa fa-pencil\"></span></md-button></div></div></div></div>");
$templateCache.put("horodata/widgets/configuration/customers_form.html","<md-dialog aria-label=\"Nouveau Dossier\" flex=40><form name=newCustomerForm><md-toolbar><div class=md-toolbar-tools><h2>Nouveaux Dossiers</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon md-font-icon=\"fa fa-close\"></md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.name}\"><label>Noms des Dossiers a ajouter (saisissez un dossier par ligne).</label> <textarea md-no-autogrow md-maxlength=1500 ng-model=name rows=5 md-select-on-focus></textarea> <small ng-if=errors.name class=input-error>{{ errors.name }}</small></md-input-container></div></md-dialog-content><md-dialog-actions><md-button ng-click=send() class=\"md-primary md-raised\">Ajouter</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/configuration/tasks.html","<div layout-margin layout-padding><div layout=column><div layout=row layout-align=\"space-between center\"><div class=md-display-1>Types de Tâches</div><md-button ng-click=showConfigurationTasksDialog() class=\"md-raised md-primary\" hide show-gt-sm><span class=\"fa fa-plus\"></span> Ajouter</md-button><md-button ng-click=showConfigurationTasksDialog() class=\"md-fab md-primary\" hide-gt-sm aria-label=\"Ajouter une tâche\"><span class=\"fa fa-plus\"></span></md-button></div><p class=md-body-1>Pour categoriser chaque tâche, les utilisateurs doivent choisir un \"type de tâche\". Chaque type de tâche peut etre associe a un commentaire pour plus de precisions.</p><div layout=row><md-input-container flex><label>Type de Tâche</label><md-select ng-model=selectedTask><md-option ng-repeat=\"u in users\" value=\"{{ u.email }}\">{{ u.name }}</md-option></md-select></md-input-container><div ng-if=selectedTask><md-button ng-click=showConfigurationTasksDialog() class=\"md-raised md-accent\" hide show-gt-sm><span class=\"fa fa-pencil\"></span> Editer</md-button><md-button ng-click=showConfigurationTasksDialog() class=\"md-fab md-accent\" hide-gt-sm aria-label=\"Editer une tâche\"><span class=\"fa fa-pencil\"></span></md-button></div></div></div></div>");
$templateCache.put("horodata/widgets/configuration/tasks_form.html","<md-dialog aria-label=\"Nouveau type de tâche\" flex=40><form name=newTaskForm><md-toolbar><div class=md-toolbar-tools><h2>Nouveau type de tâche</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon md-font-icon=\"fa fa-close\"></md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.name}\"><label>Nom de la tâche</label> <input md-maxlength=40 type=text ng-model=name> <small ng-if=errors.name class=input-error>{{ errors.name }}</small></md-input-container><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.admin}\"><md-switch ng-model=admin aria-label=Administrateur>Commentaire obligatoire</md-switch><div class=input-hint>Force l\'utilisateur a associer un commentaire a la chaque tâche de ce type.</div></md-input-container></div></md-dialog-content><md-dialog-actions><md-button ng-click=send() class=\"md-primary md-raised\">Ajouter</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/configuration/user_form.html","<md-dialog aria-label=\"Nouvel Utilisateur\" flex=40><form name=newUserForm><md-toolbar><div class=md-toolbar-tools><h2>Nouvel Utilisateur</h2><span flex></span><md-button class=md-icon-button aria-label=Fermer ng-click=close()><md-icon md-font-icon=\"fa fa-close\"></md-icon></md-button></div></md-toolbar><md-dialog-content><div class=md-dialog-content><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.email}\"><label>Adresse Email</label> <input type=text ng-model=email> <small ng-if=errors.email class=input-error>{{ errors.email }}</small></md-input-container><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.rate}\"><label>Taux Horraire</label> <input type=text ng-model=rate ng-value=0> <small ng-if=errors.rate class=input-error>{{ errors.rate }}</small></md-input-container><md-input-container class=md-block ng-class=\"{\'md-input-invalid\': errors.admin}\"><md-switch ng-model=admin aria-label=Administrateur>Administrateur</md-switch><div class=input-hint>Cochez cette case si vous souhaitez deleguer la gestion de ce group a cet utilisateur.</div></md-input-container></div></md-dialog-content><md-dialog-actions><md-button ng-click=send() class=\"md-primary md-raised\">Ajouter</md-button><md-button ng-click=close() class=md-raised>Annuler</md-button></md-dialog-actions></form></md-dialog>");
$templateCache.put("horodata/widgets/configuration/users.html","<div layout-margin layout-padding><div layout=column><div layout=row layout-align=\"space-between center\"><div class=md-display-1>Utilisateurs</div><md-button ng-click=showConfigurationUsersDialog() class=\"md-raised md-primary\" hide show-gt-sm><span class=\"fa fa-plus\"></span> Ajouter</md-button><md-button ng-click=showConfigurationUsersDialog() class=\"md-fab md-primary\" hide-gt-sm aria-label=\"Ajouter un utilisateur\"><span class=\"fa fa-plus\"></span></md-button></div><p class=md-body-1>Pour que vos collaborateurs puissent enregister leur travail, vous devez les inviter.</p><div layout=row><md-input-container flex><label>Utilisateur</label><md-select ng-model=selectedTask><md-option ng-repeat=\"u in users\" value=\"{{ u.email }}\">{{ u.name }}</md-option></md-select></md-input-container><div ng-if=selectedTask><md-button ng-click=showConfigurationUsersDialog() class=\"md-raised md-accent\" hide show-gt-sm><span class=\"fa fa-pencil\"></span> Editer</md-button><md-button ng-click=showConfigurationUsersDialog() class=\"md-fab md-accent\" hide-gt-sm aria-label=\"Editer un utilisateur\"><span class=\"fa fa-pencil\"></span></md-button></div></div></div></div>");}]);
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
    root = $("api").attr("href");
    return {
      get: function() {
        return root;
      }
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
      title: ""
    };
    return {
      get: function() {
        return title;
      },
      set: function(t) {
        return title.title = t;
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
  "$http", "$routeParams", "$scope", "titleService", "userService", "apiService", function($http, $routeParams, $scope, titleService, userService, apiService) {
    var fetchUsers;
    $scope.maxDate = new Date();
    $scope.endDate = $scope.maxDate;
    $scope.beginDate = moment().subtract(1, 'months').toDate();
    fetchUsers = function() {
      return $http.get((apiService.get()) + "/groups/" + $routeParams.group + "/users").then(function(resp) {
        return $scope.users = resp.data.data;
      });
    };
    $scope.isOwner = false;
    return userService.get(function(u) {
      return $http.get((apiService.get()) + "/groups/" + $routeParams.group).then(function(resp) {
        $scope.group = resp.data.data;
        if ($scope.group.owner.login === u.login) {
          $scope.isOwner = true;
        }
        titleService.set($scope.group.name);
        return fetchUsers();
      });
    });
  }
]);

angular.module("horodata").controller("Index", [
  "$http", "$scope", "userService", "titleService", function($http, $scope, userService, titleService) {
    return titleService.set("Choisissez un Groupe");
  }
]);

angular.module("horodata").directive("appWidgetsNewGroup", [
  "$mdDialog", "$mdMedia", function($mdDialog, $mdMedia) {
    var l;
    l = function(scope, elem, attr) {
      if ((attr.size != null) && attr.size === "small") {
        scope.small = true;
      }
      return scope.showNewGroupDialog = function(ev) {
        var fullscreen;
        fullscreen = $mdMedia('xs') || $mdMedia('sm');
        return $mdDialog.show({
          controller: "newGroupDialog",
          templateUrl: "horodata/widgets/new_group_form.html",
          parent: angular.element(document.body),
          targetEvent: ev,
          preserveScope: true,
          scope: scope,
          clickOutsideToClose: true,
          escapeToClose: true,
          fullscreen: fullscreen
        });
      };
    };
    return {
      link: l,
      restrict: "E",
      scope: {
        size: "@"
      },
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
  "$mdDialog", "$mdMedia", function($mdDialog, $mdMedia) {
    var l;
    l = function(scope, elem, attr) {
      return scope.showConfigurationCustomersDialog = function(ev) {
        var fullscreen;
        fullscreen = $mdMedia('xs') || $mdMedia('sm');
        return $mdDialog.show({
          controller: "appWidgetsConfigurationCustomersDialog",
          templateUrl: "horodata/widgets/configuration/customers_form.html",
          parent: angular.element(document.body),
          targetEvent: ev,
          preserveScope: true,
          scope: scope,
          clickOutsideToClose: true,
          escapeToClose: true,
          fullscreen: fullscreen
        });
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

angular.module("horodata").directive("appWidgetsConfigurationTasks", [
  "$mdDialog", "$mdMedia", function($mdDialog, $mdMedia) {
    var l;
    l = function(scope, elem, attr) {
      return scope.showConfigurationTasksDialog = function(ev) {
        var fullscreen;
        fullscreen = $mdMedia('xs') || $mdMedia('sm');
        return $mdDialog.show({
          controller: "appWidgetsConfigurationTasksDialog",
          templateUrl: "horodata/widgets/configuration/tasks_form.html",
          parent: angular.element(document.body),
          targetEvent: ev,
          preserveScope: true,
          scope: scope,
          clickOutsideToClose: true,
          escapeToClose: true,
          fullscreen: fullscreen
        });
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

angular.module("horodata").directive("appWidgetsConfigurationUsers", [
  "$mdDialog", "$mdMedia", function($mdDialog, $mdMedia) {
    var l;
    l = function(scope, elem, attr) {
      return scope.showConfigurationUsersDialog = function(ev) {
        var fullscreen;
        fullscreen = $mdMedia('xs') || $mdMedia('sm');
        return $mdDialog.show({
          controller: "appWidgetsConfigurationUsersDialog",
          templateUrl: "horodata/widgets/configuration/user_form.html",
          parent: angular.element(document.body),
          targetEvent: ev,
          preserveScope: true,
          scope: scope,
          clickOutsideToClose: true,
          escapeToClose: true,
          fullscreen: fullscreen
        });
      };
    };
    return {
      link: l,
      restrict: "E",
      templateUrl: "horodata/widgets/configuration/users.html"
    };
  }
]);

angular.module("horodata").controller("appWidgetsConfigurationUsersDialog", [
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

(function() {

    var app = angular.module("exampleApp", ['ngResource', 'ngRoute']);

    // Array join filter
    app.filter('join', function() {
        return function(input) {
            return input.join(', ');
        }
    });

    app.config(function($routeProvider) {

        // Set up the route handlers
        $routeProvider
            .when('/companies', {
                templateUrl: '/static/partials/company-list.html',
                controller: 'CompanyListController',
                controllerAs: 'controller'
            })
            .when('/companies/create', {
                templateUrl: '/static/partials/company-create.html',
                controller: 'CompanyCreateController',
                controllerAs: 'controller'
            })
            .when('/companies/:companyId', {
                templateUrl: '/static/partials/company-detail.html',
                controller: 'CompanyDetailController',
                controllerAs: 'controller'
            })
            .otherwise({
                redirectTo: "/companies"
            });
    });

    // Company resource
    app.factory('Company', ['$resource', function($resource) {
        return $resource('/api/companies/:id', {
            id: '@id'
        });
    }]);

    app.controller('CompanyCreateController', ['$scope', 'Company', '$location', function($scope, Company, $location) {

        // Initialize company object
        $scope.company = {
            name: "",
            address: "",
            city: "",
            zipcode: "",
            country: "",
            email: "",
            phone: "",
            owners: [],
            directors: []
        };

        // Go back to companies list
        this.cancel = function() {
            $location.path("/companies");
        };

        // Add an owner and reset input
        this.addOwner = function() {
            if (!$scope.ownerInput) {
                return;
            }
            $scope.company.owners.push($scope.ownerInput);
            $scope.ownerInput = null;
        };

        // Remove an owner by array index
        this.removeOwnerByIdx = function(idx) {
            $scope.company.owners.splice(idx, 1);
        };

        // Add a director and reset input
        this.addDirector = function() {
            if (!$scope.directorInput) {
                return;
            }
            $scope.company.directors.push($scope.directorInput);
            $scope.directorInput = null;
        };

        // Remove a director by array index
        this.removeDirectorByIdx = function(idx) {
            $scope.company.directors.splice(idx, 1);
        };

        this.createCompany = function() {

            $scope.missingName = false;
            $scope.missingAddress = false;
            $scope.missingCity = false;
            $scope.missingZipcode = false;
            $scope.missingCountry = false;

            // Validate fields
            if (!$scope.company.name) {
                $scope.missingName = true;
            }
            if (!$scope.company.address) {
                $scope.missingAddress = true;
            }
            if (!$scope.company.city) {
                $scope.missingCity = true;
            }
            if (!$scope.company.zipcode) {
                $scope.missingZipcode = true;
            }
            if (!$scope.company.country) {
                $scope.missingCountry = true;
            }

            if ($scope.missingName || $scope.missingAddress || $scope.missingCity || $scope.missingZipcode || $scope.missingCountry) {
                return;
            }

            $scope.loading = true;

            var company = new Company($scope.company);
            company.$save(null, function resolve(company) {
                // Redirect to detail view if it went well
                $location.path("/companies/" + company.id);
            }, function reject(resp) {
                $scope.error = true;
                $scope.loading = false;
            });
        };
    }]);

    app.controller('CompanyDetailController', ['$scope', '$location', '$routeParams', 'Company', function($scope, $location, $routeParams, Company) {

        // Add an owner and reset input
        this.addOwner = function() {
            if (!$scope.ownerInput) {
                return;
            }
            $scope.company.owners.push($scope.ownerInput);
            $scope.ownerInput = null;
        };

        // Remove an owner by array index
        this.removeOwnerByIdx = function(idx) {
            $scope.company.owners.splice(idx, 1);
        };

        // Add a director and reset input
        this.addDirector = function() {
            if (!$scope.directorInput) {
                return;
            }
            $scope.company.directors.push($scope.directorInput);
            $scope.directorInput = null;
        };

        // Remove director by array index
        this.removeDirectorByIdx = function(idx) {
            $scope.company.directors.splice(idx, 1);
        };

        // Delete the company object and redirect to companies list view
        this.delete = function() {
            $scope.company.$delete(null, function(resp) {
                $location.path("/companies");
            }, function(err) {
                $scope.deleteError = true;
                console.log(err);
            });
        };

        // Save the object
        this.save = function() {
            $scope.company.$save(null, function(company) {
                $scope.company = company;
            }, function(err) {
                $scope.saveError = true;
                console.log(err);
            })
        }

        // Restore an previous revision
        this.restoreRevision = function(revision) {
            $scope.company.name = revision.name;
            $scope.company.address = revision.address;
            $scope.company.city = revision.city;
            $scope.company.zipcode = revision.zipcode;
            $scope.company.country = revision.country;
            $scope.company.email = revision.email;
            $scope.company.phone = revision.phone;
            $scope.company.directors = revision.directors;
            $scope.company.owners = revision.owners;
        };

        $scope.loading = true;
        Company.get({
            id: $routeParams.companyId
        }, function success(company) {
            $scope.company = company;
            $scope.loading = false;
        }, function error(err) {

            if (err.status === 404) {
                $scope.notFound = true;
            }

            $scope.loading = false;
        });

    }]);

    app.controller('CompanyListController', ['$scope', 'Company', '$location', function($scope, Company, $location) {

        // Redirect to the company create view
        this.createCompany = function() {
            $location.path("/companies/create");
        };

        // Redirect to the company detail view
        this.viewCompany = function(company) {
            $location.path("/companies/" + company.id);
        };

        // Update the company list
        this.updateCompaniesList = function() {
            $scope.loading = true;
            $scope.companies = Company.query(null, function(companies) {
                $scope.companies = companies;
                $scope.loading = false;
            }, function error() {
                $scope.error = true;
                $scope.loading = false;
            });
        };

        this.updateCompaniesList();
    }]);
})();

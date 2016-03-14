var gulp = require('gulp');
var coffee = require("gulp-coffee");
var concat = require("gulp-concat");
var minifyHtml = require("gulp-minify-html");
var coffeelint = require("gulp-coffeelint");
var templateCache = require('gulp-angular-templatecache');
var merge = require('gulp-merge')
var uglify = require("gulp-uglify");
var gettext = require('gulp-angular-gettext');
var karma = require("karma");
var less = require('gulp-less');
var path = require('path');
var connect = require('gulp-connect');
var bump = require('gulp-bump');
var header = require('gulp-header');
var minifyCss = require('gulp-minify-css');
var gzip = require('gulp-gzip');

var MODULE = "horodata";
var DEST = "./dist/";
var VERSION = require("./package.json").version;
var HEADER = "/**\n" +
  " * HoroData Javascript Interface\n" +
  " * Version: " + VERSION + "\n" +
  " * Copyright © 2016 Hyperboloide. All rights reserved.\n" +
  "*/\n"


function inc(importance) {
    return gulp.src(['./package.json', './bower.json'])
        .pipe(bump({type: importance}))
        .pipe(gulp.dest('./'));
}

gulp.task('patch', function() { return inc('patch'); })
gulp.task('feature', function() { return inc('minor'); })
gulp.task('release', function() { return inc('major'); })

/*
** Javascript
*/

gulp.task("local-js", function(){
  var main = gulp.src('./src/main.coffee')
    .pipe(coffee({bare: true}));

  var cof = gulp.src(['./src/coffee/**/*.coffee', '!./src/coffee/**/*_test.coffee'])
    .pipe(coffee({bare: true}));

  var htm = gulp.src('./src/coffee/**/*.html')
    .pipe(minifyHtml({removeComments: true, collapseWhitespace: true}))
    .pipe(templateCache({
      root: MODULE,
      module: MODULE
    }));

  merge(main, htm, cof)
    .pipe(concat(MODULE + '.js'))
    // .pipe(uglify())
    .pipe(header(HEADER))
    .pipe(gulp.dest('./'));
})

/*
** Css
*/

gulp.task('local-css', function(){
  gulp.src('./src/less/**.less')
    .pipe(less())
    .pipe(concat(MODULE + ".css"))
    .pipe(gulp.dest('./'));
})

gulp.task("dist", ["local-js", "local-css"], function(){
  gulp.src(DEST + VERSION +"/*.{png,jpg,js,css}")
    .pipe(gzip({
      append: false,
      gzipOptions: { level: 9 }
    }))
    .pipe(gulp.dest(DEST + VERSION))
})

gulp.task("watch", function(){
  gulp.watch(['./src/main.coffee'], ['local-js']);
  gulp.watch(['./src/coffee/**/*.coffee'], ['local-js']);
  gulp.watch(['./src/coffee/**/*.html'], ['local-js']);
  gulp.watch(['./src/less/*.less'], ['local-css']);

});


gulp.task("default", ["watch"])

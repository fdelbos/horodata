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
var addsrc = require('gulp-add-src');

var MODULE = "front";
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
** Css
*/

gulp.task('local-css', function(){
    var lessStream = gulp.src('./src/less/**/*.less')
      .pipe(less());

    var cssSream = gulp.src('./src/css/**/*.css');

    merge(cssSream, lessStream)
    .pipe(concat(MODULE + ".css"))
    .pipe(gulp.dest('./'));
})

gulp.task("dist", ["local-css"], function(){
  gulp.src(DEST + VERSION +"/*.{png,jpg,js,css}")
    .pipe(gzip({
      append: false,
      gzipOptions: { level: 9 }
    }))
    .pipe(gulp.dest(DEST + VERSION))
})

gulp.task("watch", function(){
  gulp.watch(['./src/less/*.less'], ['local-css']);

});


gulp.task("default", ["watch"])

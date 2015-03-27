var gulp = require('gulp'),
    concat = require('gulp-concat'),
    sourcemaps = require('gulp-sourcemaps'),
    babel = require('gulp-babel'),
    watch = require('gulp-watch'),
    plumber = require('gulp-plumber');

var webAppSource = ['js/components/**/*', 'js/application.js'];

gulp.task('scripts:vendor', function () {
  return gulp.src(['js/vendor/**/*.js'])
    .pipe(concat('vendors.js'))
    .pipe(gulp.dest('../public/js'));
});

gulp.task('scripts:app', function () {
  return gulp.src(webAppSource)
    .pipe(plumber())
    .pipe(sourcemaps.init())
    .pipe(concat('application.js'))
    .pipe(babel())
    .pipe(sourcemaps.write('.'))
    .pipe(gulp.dest('../public/js'));
});

gulp.task('watch:scripts', ['scripts'], function () {
  return gulp.watch(webAppSource, ['scripts:app']);
});

gulp.task('scripts', ['scripts:vendor', 'scripts:app']);

gulp.task('watch', ['watch:scripts']);

var gulp = require('gulp'),
    concat = require('gulp-concat');

gulp.task('scripts:vendor', function () {
  return gulp.src(['js/vendor/**/*.js'])
    .pipe(concat('vendors.js'))
    .pipe(gulp.dest('../public/js'));
});

gulp.task('scripts:app', function () {
  return gulp.src(['js/components/**/*'])
    .pipe(concat('application.js'))
    .pipe(gulp.dest('../public/js'));
});

gulp.task('scripts', ['scripts:vendor', 'scripts:app']);

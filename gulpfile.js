const gulp = require('gulp');
const cssnano = require('gulp-cssnano');
const uglify = require('gulp-uglify');

const paths = {
  css: './assets/css/**/*.css',
  js: './assets/js/**/*.js',
};

// CSS task
gulp.task('css', () => {
  return gulp.src(paths.css)
    .pipe(cssnano())
    .pipe(gulp.dest('./public/css'));
});

// JS task
gulp.task('js', () => {
  return gulp.src(paths.js)
    .pipe(uglify())
    .pipe(gulp.dest('./public/js'));
});

// Watch task
gulp.task('watch', () => {
  gulp.watch(paths.css, gulp.series('css'));
  gulp.watch(paths.js, gulp.series('js'));
});

// Default task
gulp.task('default', gulp.parallel('css', 'js', 'watch'));
gulp.task('build', gulp.parallel('css', 'js'));

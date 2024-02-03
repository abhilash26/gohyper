/**
 * pnpm i -D gulp, gulp-cached, gulp-cssnano, gulp-uglify
*/

const gulp = require('gulp');
const cache = require('gulp-cached');
const cssnano = require('gulp-cssnano');
const uglify = require('gulp-uglify');

const paths = {
  css: './assets/css/**/*.css',
  js: './assets/js/**/*.js',
  img: './assets/img/**/*.{jpg,jpeg,png,bmp,svg,gif,ico}',
};

// CSS task
gulp.task('css', () => {
  return gulp.src(paths.css)
    .pipe(cache('css'))
    .pipe(cssnano())
    .pipe(gulp.dest('./public/css'));
});

// JS task
gulp.task('js', () => {
  return gulp.src(paths.js)
    .pipe(cache('js'))
    .pipe(uglify())
    .pipe(gulp.dest('./public/js'));
});

// Image task
gulp.task('img', () => {
  return gulp.src(paths.img)
    .pipe(cache('img'))
    .pipe(dest('./public/img'));
});

// Watch task
gulp.task('watch', () => {
  gulp.watch(paths.css, gulp.series('css'));
  gulp.watch(paths.js, gulp.series('js'));
  gulp.watch(paths.js, gulp.series('img'));
});

// Default task
gulp.task('default', gulp.series(gulp.parallel('css', 'js', 'img'), 'watch'));
gulp.task('build', gulp.parallel('css', 'js', 'img'));

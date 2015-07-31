var gulp = require("gulp");
var shell = require("gulp-shell");

gulp.task('run script', shell.task([
'sh run.sh'
]));


gulp.task('default', function() {
    var watcher = gulp.watch('*',['run script']);
});

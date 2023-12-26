var gulp = require("gulp")
var shell = require("gulp-shell")
var { exec } = require("child_process")

function compile(cb) {
	exec("go build -o ./bin/romainserver ./src")

	cb()
}

function restart_supervisor(cb) {
	exec("supervisorctl restart myserver")

	cb()
}
     
exports.default = function (cb) {
	gulp.watch("src/*", gulp.series(compile, restart_supervisor))
	cb()
}

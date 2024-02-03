# Makefile

.PHONY: run build clean

default: run

run:
	echo "Running Pollen.."
	@pollen
	echo "Running Gulp.."
	@gulp &
	echo "Running Go.."
	@air

build:
	@pollen && gulp build && go build -o gohyper cmd/main.go

clean:
	@rm -rf tmp/

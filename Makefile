# Makefile

.PHONY: run build clean

default: run

run:
	echo "Running Pollen.."
	@pollen
	echo "Running NPM.."
	@npm run dev &
	echo "Running Go.."
	@air

build:
	@pollen && npm run build && go build -o gohyper cmd/main.go

clean:
	@rm -rf tmp/

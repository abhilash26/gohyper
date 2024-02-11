# Makefile

.PHONY: run build clean

default: run

run:
	@echo "Running Pollen.."
	@pollen
	@echo "Running NPM.."
	@pnpm run dev &
	@echo "Running Go.."
	@air

build:
	@echo "Building Pollen.."
	@pollen
	@echo "Building pnpm tasks.."
	@pnpm run build
	@echo "Copying directories.."
	@mkdir -p "cmd/build"
	@cp -r public cmd/build/
	@cp -r view cmd/build/
	@cp .env.example cmd/build/.env
	@echo "Building Go app.."
	@go build -o cmd/build/app cmd/main.go

clean:
	@echo "Cleaning tmp/ directory.."
	@rm -rf tmp/
	@echo "Cleaning building tasks.."
	@rm -rf cmd/build/

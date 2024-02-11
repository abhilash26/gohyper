# Makefile

APP_NAME="gohyper"
BUILD_DIR="cmd/${APP_NAME}/"
TEMP_DIR="tmp/"
DB_NAME="database.db"

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
	@mkdir -p "cmd/"
	@cp -r public "$BUILD_DIR"
	@cp -r view "${BUILD_DIR}"
	@cp .env "${BUILD_DIR}.env"
	@echo "Copying database.."
	@cp "${DB_NAME}" "${BUILD_DIR}${DB_NAME}"
	@echo "Building Go app.."
	@go build -o "${BUILD_DIR}app" cmd/main.go

clean:
	@echo "Cleaning ${TEMP_DIR} directory.."
	@rm -rf "$TEMP_DIR"
	@echo "Cleaning building tasks.."
	@rm -rf "$BUILD_DIR"

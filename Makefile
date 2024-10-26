build: setup-env
	docker build --no-cache -t mhrs-cli .
run: setup-env
	docker run --rm -it mhrs-cli

setup-env:
	@if [ ! -f .env ]; then \
    		echo ".env file not found. Copying .env.example to .env"; \
    		cp .env.example .env; \
    	fi
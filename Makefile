run:
	node_modules/.bin/concurrently "npm:css:build" "npm:run:go" "npm:components:watch"

clean_css:
	yarn clean


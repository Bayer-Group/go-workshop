# ----------------------
#  Go Workshop Exercises
# ----------------------
helloworld:
		go run ./basics/00helloworld.go

01:
		go run ./basics/01fmt.go # fmt-y Dumpty

02:
		go run ./basics/02conditionals.go # On one conditional

03:
		go run ./basics/03slices.go # A healthy slice of arrays

04:
		go run ./basics/04loops.go # Stuck in the loop with you

05:
		go run ./basics/05functions.go # Becoming a functioning Gopher

06:
		go run ./basics/06errors.go # I am Error

07:
		go run ./intermediate/07panics.go # Don't panic

08:
		go run ./intermediate/08pointers.go # Get to the point

09:
		go run ./intermediate/09structs.go # Hardwired to self de-struct

10:
		go run ./intermediate/10interfaces.go # Unlike a user interface, these take a bit of explanation

11:
		go run ./intermediate/11goroutines.go # Getting into the routine

12:
		go run ./intermediate/12channels.go # Now you're thinking with channels

13:
		go run ./advanced/13currying.go # Staying anonymous

14:
		go run ./advanced/14async.go # Getting out of sync

15:
		go run ./advanced/15packages.go # Here's your package

# ---------------------------
#  Documentation/Presentation
# ---------------------------

local-docs: mkdocs
		mkdocs serve

publish: mkdocs
		mkdocs gh-deploy

mkdocs:
		if [ -z "$(which mkdocs)" ]; then pip install mkdocs; pip install mkdocs-material; fi

package main

func main() {
	inputString := "{\"make\": \"Gourd\", \"model\": \"Horus\", \"color\": \"Tan\", \"year\": 2004}"

	// TODO use `go get -u` to import the package `github.com/sirupsen/logrus`, a useful library for generating log messages.
	//    Notice that that dependency is automatically added to `go.mod` and `go.sum`

	// TODO initialize a logger with logrus.New()

	// TODO set the formatter on the logger as a *logrus.JSONFormatter

	// TODO call `car.FromJSON()` to translate the json string above to a car struct
	//    if your editor is being weird about the import, try `github.com/MonsantoCo/go-workshop/internal/car`.
	//    the `internal` folder is a special package in Go. Packages within `internal` can only be imported by the parent directory packages and not programs that import your library.

	// TODO add an error check for the possible error returned by `car.FromJSON()` and pass the error created by
	//    invoking `fmt.Errorf("car.FromJSON: %w", err)` and passing the result to `logger.Println()` to print the error trace.
	//    NOTE: %w is a special fmt verb only available for fmt.Errorf. It is used to create something analogous to a stack trace.

	// TODO log the car struct returned by `car.FromJSON()`

	// TODO explore the car package (note that the returned `car` struct is private)

	// TODO try to access one of the properties from car from this file

	// TODO invoke `ToJSON()` on your car and log the resulting byte slice
}

package logger

// APILoggerSingleton is a log singleton used to
// register log entries trhoughout the application.
//
// It must be restarted at main.go to get the updated
// environment variables needed by the logger engine.
//
// ** I didn't think of a better way of implementing
// this, but I'm sure there are many! **
var APILoggerSingleton *APILogger = NewWithLogrus("tt-go-sample-api")

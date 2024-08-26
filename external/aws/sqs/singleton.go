package sqs

import "sync"

var lock = &sync.Mutex{}

// apiSQSSingleton is a singleton of the
// API's AWS SQS.
var apiSQSSingleton *APISQS

// GetAPISQSSingletonSingleton returns the current
// value of the application's SQS runner.
func GetAPISQSSingletonSingleton() *APISQS {
	if apiSQSSingleton == nil {
		lock.Lock()
		defer lock.Unlock()

		if apiSQSSingleton == nil {
			apiSQSSingleton = NewAPISQS()
		}
	}

	return apiSQSSingleton
}

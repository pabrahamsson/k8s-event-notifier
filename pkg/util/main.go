package util

import (
	eventv1 "github.com/redhat-cop/events-notifier/pkg/apis/event/v1"
	notifyv1 "github.com/redhat-cop/events-notifier/pkg/apis/notify/v1"
)

type SharedResources struct {
	Subscriptions *[]eventv1.EventSubscription
	Notifiers     *[]notifyv1.Notifier
}

func NewSharedResources() SharedResources {
	return SharedResources{
		Subscriptions: &[]eventv1.EventSubscription{},
		Notifiers:     &[]notifyv1.Notifier{},
	}
}

//
// Helper functions to check and remove string from a slice of strings.
//
func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func RemoveString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

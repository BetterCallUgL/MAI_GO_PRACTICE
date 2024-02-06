package retryupdate

import "gitlab.com/sophistik1/mai-practice-hw3/retryupdate/kvapi"

func UpdateValue(c kvapi.Client, key string, updateFn func(oldValue *string) (newValue string, err error)) error {
	panic("implement me")
}

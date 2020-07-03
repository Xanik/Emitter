package test

import "testing"

func TestStoreEvents(t *testing.T) {
	tests := []TestDBEvent{
		{
			ID:        "01EBP4DP4VECW8PHDJJFNEDVKE",
			Message:   "some message to send",
			CreatedOn: "2020-06-25T16:23:37.720Z",
		},
	}
	dB := fakeDB{t}
	for _, test := range tests {
		result, err := dB.StoreEvents(test)
		if err != nil {
			t.Errorf("error occured: %s", err)
		}
		if result != 1 {
			t.Errorf("invalid result: %v", result)
		}
	}
}

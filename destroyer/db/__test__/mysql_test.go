package test

import "testing"

func TestListTarget(t *testing.T) {
	tests := map[string]TestDBListTarget{
		"1": {
			ID:        "01EBP4DP4VECW8PHDJJFNEDVKE",
			Message:   "some message to send",
			CreatedOn: "2020-06-25T16:23:37.720Z",
		},
		"2": {
			ID:        "20EBP4DQBVECW8PHDJJFNEDVKE",
			Message:   "some other message to send",
			CreatedOn: "2020-06-25T16:23:37.720Z",
		},
	}
	dB := fakeDB{t}

	for id, test := range tests {
		target, err := dB.ListTarget(id)
		if err != nil {
			t.Errorf("error occured: %s", err)
		}
		if target[0].ID != test.ID {
			t.Errorf("%s not equal to client name %s", target[0].ID, test.ID)
		}
	}
}

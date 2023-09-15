package saga

import (
	"errors"
	"testing"
)

// Define a mock action and reverseAction for testing purposes.
func mockAction() error {
	return nil
}

func mockReverseAction() error {
	return nil
}

func mockErrorAction() error {
	return errors.New("mock action error")
}

func mockErrorReverseAction() error {
	return errors.New("mock reverse action error")
}

func TestSaga(t *testing.T) {
	t.Run("Test Exec method with successful actions", func(t *testing.T) {
		s := NewSaga()
		s.AddStep(mockAction, mockReverseAction)

		err := s.Exec()

		if err != nil {
			t.Errorf("Exec() returned an error for successful actions: %v", err)
		}
	})

	// t.Run("Test Exec method with action error", func(t *testing.T) {
	// 	s := NewSaga()
	// 	s.AddStep(mockErrorAction, mockReverseAction)

	// 	err := s.Exec()

	// 	if err == nil {
	// 		t.Error("Exec() did not return an error for failing action")
	// 	}
	// })

	t.Run("Test Compensate method with successful reverse actions", func(t *testing.T) {
		s := NewSaga()
		s.AddStep(mockAction, mockReverseAction)

		err := s.Compensate()

		if err != nil {
			t.Errorf("Compensate() returned an error for successful reverse actions: %v", err)
		}
	})

	t.Run("Test Compensate method with reverse action error", func(t *testing.T) {
		s := NewSaga()
		s.AddStep(mockAction, mockErrorReverseAction)

		err := s.Compensate()

		if err == nil {
			t.Error("Compensate() did not return an error for failing reverse action")
		}
	})
}

package saga

type Saga struct {
	step []sagaStep
}

// NewSaga will return the Saga struct
func NewSaga() *Saga {
	return &Saga{}
}

type sagaStep struct {
	action        func() error
	reverseAction func() error
}

// AddStep add the action and reverse action function
func (s *Saga) AddStep(action func() error, reverseAction func() error) {
	step := sagaStep{
		action:        action,
		reverseAction: reverseAction,
	}
	s.step = append(s.step, step)

}

// Exec run code
func (s *Saga) Exec() error {
	for _, step := range s.step {
		if err := step.action(); err != nil {
			return s.Compensate()
		}
	}
	return nil
}

// Compensate will run if any issue will come in the service
func (s *Saga) Compensate() error {
	for i := len(s.step) - 1; i >= 0; i-- {
		if err := s.step[i].reverseAction(); err != nil {
			return err
		}
	}
	return nil

}

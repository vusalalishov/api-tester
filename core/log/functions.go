package log

func NewSuite(title string) *TestSuite {
	return &TestSuite{
		Title: title,
		Cases: make([]*TestCase, 0),
		testInfo: testInfo{
			Status:   SUCCESS,
			Messages: make([]*string, 0),
		},
	}
}

func (s *TestSuite) AddCase(title string) *TestCase {
	tc := &TestCase{
		Title:     title,
		Scenarios: make([]TestScenario, 0),
		testInfo: testInfo{
			Status: SUCCESS,
		},
	}
	s.Cases = append(s.Cases, tc)
	return tc
}

func (s *TestSuite) AddMessage(message string) {
	s.Messages = append(s.Messages, &message)
}

func (s *TestSuite) SetStatus(status Status) {
	s.Status = status
}

func (s *TestScenario) SetStatus(status Status) {
	s.Status = status
}

func (c *TestCase) AddScenario(title string) *TestScenario {
	ts := &TestScenario{
		Title: title,
		testInfo: testInfo{
			Status:   SUCCESS,
			Messages: make([]*string, 0),
		},
	}
	c.Scenarios = append(c.Scenarios, ts)
	return ts
}

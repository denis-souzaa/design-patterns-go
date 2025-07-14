package average

type Average struct {
	StudentId int
	Value     float64
}

func New(studentId int, value float64) *Average {
	return &Average{StudentId: studentId, Value: value}
}

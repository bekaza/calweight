package calweight

type Init struct {
	width  float64
	height float64
	length float64
}

func GetCalculateWeight(e *Init) float64 {
	return (e.width * e.height * e.length) / 5000
}

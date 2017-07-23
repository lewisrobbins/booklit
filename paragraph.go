package booklit

type Paragraph []Content

func (con Paragraph) String() string {
	str := ""

	for _, seq := range con {
		str += seq.String()
	}

	return str
}

func (con Paragraph) IsFlow() bool {
	return false
}

func (con Paragraph) Visit(visitor Visitor) error {
	return visitor.VisitParagraph(con)
}

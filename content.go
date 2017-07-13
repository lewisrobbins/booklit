package booklit

type Content interface {
	String() string

	IsSentence() bool

	Visit(Visitor) error
}

type Visitor interface {
	VisitString(String) error
	VisitSequence(Sequence) error
	VisitReference(*Reference) error
	VisitLink(Link) error
	VisitSection(*Section) error
	VisitParagraph(Paragraph) error
	VisitTableOfContents(TableOfContents) error
	VisitPreformatted(Preformatted) error
	VisitStyled(Styled) error
	VisitTarget(Target) error
	VisitBlock(Block) error
	VisitElement(Element) error
	VisitImage(Image) error
	VisitList(List) error
}
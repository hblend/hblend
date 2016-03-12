package gotreescript

type TokenType int

const (
	TEXT TokenType = iota
	TAG
	COMMENT
	NOPARSE
)

type Token struct {
	Type    TokenType
	Partial string
	Name    string
	Flags   []string
	Args    map[string]string
}

func NewToken(token_type TokenType) *Token {

	return &Token{
		Type:    token_type,
		Partial: "",
		Name:    "",
		Flags:   nil,
		Args:    nil,
	}
}

func (this *Token) PrettyPrint() string {

	type2text := map[TokenType]string{
		TEXT:    "TEXT",
		TAG:     "TAG",
		COMMENT: "COMMENT",
		NOPARSE: "NOPARSE",
	}

	pp := ""

	pp += type2text[this.Type] + "\t{"

	if TAG == this.Type {
		pp += "name: '" + this.Name + "', "
	}
	pp += "partial: '" + this.Partial + "', "

	if nil != this.Flags {
		pp += "flags: ["
		for _, v := range this.Flags {
			pp += "'" + v + "', "
		}
		pp += "], "
	}

	if nil != this.Args {
		pp += "args: ["
		for k, v := range this.Args {
			pp += "'" + k + "'='" + v + "', "
		}
		pp += "], "
	}

	pp += "}"

	return pp
}

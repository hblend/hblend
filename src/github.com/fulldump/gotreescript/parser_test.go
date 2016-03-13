package gotreescript

import (
	"reflect"
	"testing"
)

func check(t *testing.T, code string, reference *[]*Token) {

	result := Parse(code)

	if !reflect.DeepEqual(reference, result) {

		msg := "\n\nCODE: '" + code + "'"

		msg += "\n\nThis tokens:\n"
		for _, v := range *result {
			msg += v.PrettyPrint() + "\n"
		}
		msg += "\nShould be like this:\n"
		for _, v := range *reference {
			msg += v.PrettyPrint() + "\n"
		}
		t.Errorf(msg)
	}
}

func Test_Empty_code(t *testing.T) {

	code := ""
	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: "",
		},
	}
	check(t, code, reference)
}

func Test_One_Bracket(t *testing.T) {

	code := "one [ bracket"
	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: "one [ bracket",
		},
	}
	check(t, code, reference)
}

func Test_Comment_one_line(t *testing.T) {

	code := " text one [[ this is a comment ]] text two "
	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    COMMENT,
			Partial: "[[ this is a comment ]]",
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}
	check(t, code, reference)
}

func Test_Comment_with_tab_char(t *testing.T) {
	code := " text one [[\tthis is a comment ]] text two "
	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    COMMENT,
			Partial: "[[\tthis is a comment ]]",
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}
	check(t, code, reference)
}

func Test_Comment_multiple_lines(t *testing.T) {
	code := " text one [[\nthis\nis\na\ncomment\n]] text two "
	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    COMMENT,
			Partial: "[[\nthis\nis\na\ncomment\n]]",
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}
	check(t, code, reference)
}

func Test_Comment_empty(t *testing.T) {
	code := " text one [[]] text two "
	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    COMMENT,
			Partial: "[[]]",
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}
	check(t, code, reference)
}

func Test_Comment_unexpected_end(t *testing.T) {
	code := " text one [[ unexpected end of file"
	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    COMMENT,
			Partial: "[[ unexpected end of file",
		},
	}
	check(t, code, reference)
}

func Test_One_tag(t *testing.T) {

	code := " text one [[MY_ITEM]] text two "

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM]]",
			Name:    "MY_ITEM",
			Flags:   []string{},
			Args:    map[string]string{},
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}

	check(t, code, reference)
}

func Test_One_tag_blanks(t *testing.T) {

	code := " text one [[MY_ITEM \t\n]] text two "

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM \t\n]]",
			Name:    "MY_ITEM",
			Flags:   []string{},
			Args:    map[string]string{},
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}

	check(t, code, reference)
}

func Test_One_tag_unexpected_end(t *testing.T) {

	code := " text one [[MY_ITEM"

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM",
			Name:    "MY_ITEM",
			Flags:   []string{},
			Args:    map[string]string{},
		},
	}

	check(t, code, reference)
}

func Test_Flags_no_end(t *testing.T) {

	code := " text one [[MY_ITEM flag]] text two "

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM flag]]",
			Name:    "MY_ITEM",
			Flags:   []string{"flag"},
			Args:    map[string]string{},
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}

	check(t, code, reference)
}

func Test_Flags_with_ending_separator(t *testing.T) {

	code := " text one [[MY_ITEM flag ]] text two "

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM flag ]]",
			Name:    "MY_ITEM",
			Flags:   []string{"flag"},
			Args:    map[string]string{},
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}

	check(t, code, reference)
}

func Test_Flags_multiple(t *testing.T) {

	code := " text one [[MY_ITEM one two three ]] text two "

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM one two three ]]",
			Name:    "MY_ITEM",
			Flags:   []string{"one", "two", "three"},
			Args:    map[string]string{},
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}

	check(t, code, reference)
}

func Test_Flags_diverse(t *testing.T) {

	code := " text one [[MY_ITEM :one two three\" \"four $five ]] text two "

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM :one two three\" \"four $five ]]",
			Name:    "MY_ITEM",
			Flags:   []string{":one", "two", "three\"", "\"four", "$five"},
			Args:    map[string]string{},
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}

	check(t, code, reference)
}

func Test_Attribute_equal(t *testing.T) {

	code := " text one [[MY_ITEM attribute = value ]] text two "

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM attribute = value ]]",
			Name:    "MY_ITEM",
			Flags:   []string{},
			Args:    map[string]string{"attribute": "value"},
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}

	check(t, code, reference)
}

func Test_Attribute_colon(t *testing.T) {

	code := " text one [[MY_ITEM attribute : value ]] text two "

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM attribute : value ]]",
			Name:    "MY_ITEM",
			Flags:   []string{},
			Args:    map[string]string{"attribute": "value"},
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}

	check(t, code, reference)
}

func Test_Attribute_no_separator_before_assignment(t *testing.T) {

	code := " text one [[MY_ITEM attribute= value ]] text two "

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM attribute= value ]]",
			Name:    "MY_ITEM",
			Flags:   []string{},
			Args:    map[string]string{"attribute": "value"},
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}

	check(t, code, reference)
}

func Test_Attribute_no_separator_after_assignment(t *testing.T) {

	code := " text one [[MY_ITEM attribute =value ]] text two "

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM attribute =value ]]",
			Name:    "MY_ITEM",
			Flags:   []string{},
			Args:    map[string]string{"attribute": "value"},
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}

	check(t, code, reference)
}

func Test_Attribute_no_end_after_value(t *testing.T) {

	code := " text one [[MY_ITEM attribute = value]] text two "

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM attribute = value]]",
			Name:    "MY_ITEM",
			Flags:   []string{},
			Args:    map[string]string{"attribute": "value"},
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}

	check(t, code, reference)
}

func Test_Attribute_and_after_assignment(t *testing.T) {

	code := " text one [[MY_ITEM attribute = ]] text two "

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM attribute = ]]",
			Name:    "MY_ITEM",
			Flags:   []string{},
			Args:    map[string]string{"attribute": ""},
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}

	check(t, code, reference)
}

func Test_value_single_quotes(t *testing.T) {

	code := " text one [[MY_ITEM attribute = 'Once \"uppon\" a time' ]] text two "

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM attribute = 'Once \"uppon\" a time' ]]",
			Name:    "MY_ITEM",
			Flags:   []string{},
			Args:    map[string]string{"attribute": "Once \"uppon\" a time"},
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}

	check(t, code, reference)
}

func Test_value_double_quotes(t *testing.T) {

	code := " text one [[MY_ITEM attribute = \"Once 'uppon' a time\" ]] text two "

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM attribute = \"Once 'uppon' a time\" ]]",
			Name:    "MY_ITEM",
			Flags:   []string{},
			Args:    map[string]string{"attribute": "Once 'uppon' a time"},
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}

	check(t, code, reference)
}

func Test_escaped_chars_inside_simple_quotes(t *testing.T) {

	code := " text one [[MY_ITEM attribute = ' a \\\\ b \\' c \\\" d ' ]] text two "

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM attribute = ' a \\\\ b \\' c \\\" d ' ]]",
			Name:    "MY_ITEM",
			Flags:   []string{},
			Args:    map[string]string{"attribute": " a \\ b ' c \" d "},
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}

	check(t, code, reference)
}

func Test_escaped_chars_inside_double_quotes(t *testing.T) {

	code := " text one [[MY_ITEM attribute = \" a \\\\ b \\' c \\\" d \" ]] text two "

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " text one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM attribute = \" a \\\\ b \\' c \\\" d \" ]]",
			Name:    "MY_ITEM",
			Flags:   []string{},
			Args:    map[string]string{"attribute": " a \\ b ' c \" d "},
		},
		&Token{
			Type:    TEXT,
			Partial: " text two ",
		},
	}

	check(t, code, reference)
}

func Test_combined_1(t *testing.T) {

	code := " one [[MY_ITEM flag_1 a=b flag_2 flag_3 c=d flag_4 e=f ]] two "

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " one ",
		},
		&Token{
			Type:    TAG,
			Partial: "[[MY_ITEM flag_1 a=b flag_2 flag_3 c=d flag_4 e=f ]]",
			Name:    "MY_ITEM",
			Flags:   []string{"flag_1", "flag_2", "flag_3", "flag_4"},
			Args:    map[string]string{"a": "b", "c": "d", "e": "f"},
		},
		&Token{
			Type:    TEXT,
			Partial: " two ",
		},
	}

	check(t, code, reference)
}

func Test_combined_2(t *testing.T) {

	code := ` one [[MY_ITEM 
		a=''
		b="'"
		c="function(var='', ver='') {
			return var+ver;
		}"
	]] two `

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " one ",
		},
		&Token{
			Type: TAG,
			Partial: `[[MY_ITEM 
		a=''
		b="'"
		c="function(var='', ver='') {
			return var+ver;
		}"
	]]`,
			Name:  "MY_ITEM",
			Flags: []string{},
			Args: map[string]string{"a": "", "b": "'", "c": `function(var='', ver='') {
			return var+ver;
		}`},
		},
		&Token{
			Type:    TEXT,
			Partial: " two ",
		},
	}

	check(t, code, reference)
}

func Test_attributes_repeated_keys(t *testing.T) {

	code := ` one [[MY_ITEM a="1" a="2"]] two `

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: " one ",
		},
		&Token{
			Type:    TAG,
			Partial: `[[MY_ITEM a="1" a="2"]]`,
			Name:    "MY_ITEM",
			Flags:   []string{},
			Args:    map[string]string{"a": "2"},
		},
		&Token{
			Type:    TEXT,
			Partial: " two ",
		},
	}

	check(t, code, reference)
}

func Test_multibyte(t *testing.T) {

	code := ` ú [[MY_ITEM]] € `

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: ` ú `,
		},
		&Token{
			Type:    TAG,
			Partial: `[[MY_ITEM]]`,
			Name:    "MY_ITEM",
			Flags:   []string{},
			Args:    map[string]string{},
		},
		&Token{
			Type:    TEXT,
			Partial: ` € `,
		},
	}

	check(t, code, reference)
}

func Test_noparse_1(t *testing.T) {

	code := ` a [[NAME1]] b [[Noparse]] c [[NAME2]] d `

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: ` a `,
		},
		&Token{
			Type:    TAG,
			Partial: `[[NAME1]]`,
			Name:    "NAME1",
			Flags:   []string{},
			Args:    map[string]string{},
		},
		&Token{
			Type:    TEXT,
			Partial: ` b `,
		},
		&Token{
			Type:    NOPARSE,
			Partial: `[[Noparse]]`,
			Name:    "Noparse",
			Flags:   []string{},
			Args:    map[string]string{},
		},
		&Token{
			Type:    TEXT,
			Partial: ` c [[NAME2]] d `,
		},
	}

	check(t, code, reference)
}

func Test_noparse_2(t *testing.T) {

	code := ` a [[NAME1]] b [[noparse a:b]] c [[NAME2]] d `

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: ` a `,
		},
		&Token{
			Type:    TAG,
			Partial: `[[NAME1]]`,
			Name:    "NAME1",
			Flags:   []string{},
			Args:    map[string]string{},
		},
		&Token{
			Type:    TEXT,
			Partial: ` b `,
		},
		&Token{
			Type:    NOPARSE,
			Partial: `[[noparse a:b]]`,
			Name:    "noparse",
			Flags:   []string{},
			Args:    map[string]string{"a": "b"},
		},
		&Token{
			Type:    TEXT,
			Partial: ` c [[NAME2]] d `,
		},
	}

	check(t, code, reference)
}

func Test_noparse_3(t *testing.T) {

	code := ` a [[NAME1]] b [[noparse a:b c]] c [[NAME2]] d `

	reference := &[]*Token{
		&Token{
			Type:    TEXT,
			Partial: ` a `,
		},
		&Token{
			Type:    TAG,
			Partial: `[[NAME1]]`,
			Name:    "NAME1",
			Flags:   []string{},
			Args:    map[string]string{},
		},
		&Token{
			Type:    TEXT,
			Partial: ` b `,
		},
		&Token{
			Type:    NOPARSE,
			Partial: `[[noparse a:b c]]`,
			Name:    "noparse",
			Flags:   []string{"c"},
			Args:    map[string]string{"a": "b"},
		},
		&Token{
			Type:    TEXT,
			Partial: ` c [[NAME2]] d `,
		},
	}

	check(t, code, reference)
}

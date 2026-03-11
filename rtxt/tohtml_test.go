package rtxt

import (
	"testing"
)

func TestToHTML(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name:   "empty string",
			input:  "",
			expect: "",
		},
		{
			name:   "plain text",
			input:  "Hello world",
			expect: "<p>Hello world</p>",
		},
		{
			name:   "two paragraphs",
			input:  "First paragraph\n\nSecond paragraph",
			expect: "<p>First paragraph</p><p>Second paragraph</p>",
		},
		{
			name:   "paragraphs separated by whitespace between newlines",
			input:  "First paragraph\n \t\nSecond paragraph",
			expect: "<p>First paragraph</p><p>Second paragraph</p>",
		},
		{
			name:   "bold",
			input:  "**bold**",
			expect: "<p><strong>bold</strong></p>",
		},
		{
			name:   "italic",
			input:  "//italic//",
			expect: "<p><em>italic</em></p>",
		},
		{
			name:   "underline",
			input:  "__underline__",
			expect: "<p><u>underline</u></p>",
		},
		{
			name:   "highlight",
			input:  "||highlight||",
			expect: "<p><mark>highlight</mark></p>",
		},
		{
			name:   "mixed formatting in one paragraph",
			input:  "This is **bold** and //italic// and __underlined__ and ||highlighted||",
			expect: "<p>This is <strong>bold</strong> and <em>italic</em> and <u>underlined</u> and <mark>highlighted</mark></p>",
		},
		{
			name:   "single newline within paragraph becomes br",
			input:  "Line one\nLine two",
			expect: "<p>Line one<br>Line two</p>",
		},
		{
			name:   "HTML special chars escaped",
			input:  "<script>alert('xss')</script>",
			expect: "<p>&lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;</p>",
		},
		{
			name:   "ampersand escaped",
			input:  "AT&T",
			expect: "<p>AT&amp;T</p>",
		},
		{
			name:   "three paragraphs",
			input:  "One\n\nTwo\n\nThree",
			expect: "<p>One</p><p>Two</p><p>Three</p>",
		},
		{
			name:   "formatting across multiple paragraphs",
			input:  "**bold text**\n\n//italic text//",
			expect: "<p><strong>bold text</strong></p><p><em>italic text</em></p>",
		},
		{
			name:   "single newline and double newline mixed",
			input:  "Line one\nLine two\n\nNew paragraph",
			expect: "<p>Line one<br>Line two</p><p>New paragraph</p>",
		},

		// Persian numbers
		{
			name:   "persian digits",
			input:  "۰۱۲۳۴۵۶۷۸۹",
			expect: "<p>۰۱۲۳۴۵۶۷۸۹</p>",
		},
		{
			name:   "persian number in sentence",
			input:  "قیمت ۱۲۳ تومان است",
			expect: "<p>قیمت ۱۲۳ تومان است</p>",
		},
		{
			name:   "bold persian number",
			input:  "**۴۲**",
			expect: "<p><strong>۴۲</strong></p>",
		},
		{
			name:   "persian numbers in two paragraphs",
			input:  "شماره ۱\n\nشماره ۲",
			expect: "<p>شماره ۱</p><p>شماره ۲</p>",
		},

		// Persian words
		{
			name:   "persian dorood",
			input:  "درود",
			expect: "<p>درود</p>",
		},
		{
			name:   "persian dorood bold",
			input:  "**درود**",
			expect: "<p><strong>درود</strong></p>",
		},
		{
			name:   "persian sentence",
			input:  "درود بر شما",
			expect: "<p>درود بر شما</p>",
		},
		{
			name:   "persian sentence with formatting",
			input:  "درود **دوست** //عزیز//",
			expect: "<p>درود <strong>دوست</strong> <em>عزیز</em></p>",
		},
		{
			name:   "persian two paragraphs",
			input:  "درود بر شما\n\nخوش آمدید",
			expect: "<p>درود بر شما</p><p>خوش آمدید</p>",
		},
		{
			name:   "persian with line break",
			input:  "درود\nخداحافظ",
			expect: "<p>درود<br>خداحافظ</p>",
		},

		// Korean words
		{
			name:   "korean plain text",
			input:  "안녕하세요",
			expect: "<p>안녕하세요</p>",
		},
		{
			name:   "korean bold",
			input:  "**안녕하세요**",
			expect: "<p><strong>안녕하세요</strong></p>",
		},
		{
			name:   "korean sentence with formatting",
			input:  "이것은 **굵은** 글씨와 //기울임// 글씨입니다",
			expect: "<p>이것은 <strong>굵은</strong> 글씨와 <em>기울임</em> 글씨입니다</p>",
		},
		{
			name:   "korean two paragraphs",
			input:  "첫 번째 단락\n\n두 번째 단락",
			expect: "<p>첫 번째 단락</p><p>두 번째 단락</p>",
		},
		{
			name:   "korean with line break",
			input:  "안녕하세요\n감사합니다",
			expect: "<p>안녕하세요<br>감사합니다</p>",
		},
		{
			name:   "korean with highlight",
			input:  "||중요한|| 내용",
			expect: "<p><mark>중요한</mark> 내용</p>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToHTML(tt.input)
			if got != tt.expect {
				t.Errorf("ToHTML(%q)\n got: %q\nwant: %q", tt.input, got, tt.expect)
			}
		})
	}
}

package graphql

import (
	"regexp" //nolint:depguard // bluemonday requires this pkg
	"sync"

	"github.com/microcosm-cc/bluemonday"
	gfm "github.com/shurcooL/github_flavored_markdown"
)

type Markdown string

func (m Markdown) Text() string {
	return string(m)
}

func (m Markdown) HTML() string {
	return render(string(m))
}

var (
	once   sync.Once
	policy *bluemonday.Policy
)

// Render renders Markdown content into sanitized HTML that is safe to render anywhere.
func render(content string) string {
	once.Do(func() {
		policy = bluemonday.UGCPolicy()
		policy.AllowAttrs("name").Matching(bluemonday.SpaceSeparatedTokens).OnElements("a")
		policy.AllowAttrs("rel").Matching(regexp.MustCompile(`^nofollow$`)).OnElements("a")
		policy.AllowAttrs("class").Matching(regexp.MustCompile(`^anchor$`)).OnElements("a")
		policy.AllowAttrs("aria-hidden").Matching(regexp.MustCompile(`^true$`)).OnElements("a")
		policy.AllowAttrs("type").Matching(regexp.MustCompile(`^checkbox$`)).OnElements("input")
		policy.AllowAttrs("checked", "disabled").Matching(regexp.MustCompile(`^$`)).OnElements("input")
		policy.AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("code")
	})

	unsafeHTML := gfm.Markdown([]byte(content))
	return string(policy.SanitizeBytes(unsafeHTML))
}

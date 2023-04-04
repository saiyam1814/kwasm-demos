package template

import "fmt"

const (
	detailsTypeReplaceTarget      = "{TYPE}"
	detailsPagesReplaceTarget     = "{PAGES}"
	detailsPublisherReplaceTarget = "{PUBLISHER}"
	detailsLanguageReplaceTarget  = "{LANGUAGE}"
	detailsISBN10ReplaceTarget    = "{ISBN10}"
	detailsISBN13ReplaceTarget    = "{ISBN13}"
)

var detailsHTML = fmt.Sprintf(`
	<dl>
		<dt>Type:</dt>
		<dd>%s</dd>
		<dt>Pages:</dt>
		<dd>%s</dd>
		<dt>Publisher:</dt>
		<dd>%s</dd>
		<dt>Language:</dt>
		<dd>%s</dd>
		<dt>ISBN-10:</dt>
		<dd>%s</dd>
		<dt>ISBN-13:</dt>
		<dd>%s</dd>
	</dl>
	`,
	detailsTypeReplaceTarget,
	detailsPagesReplaceTarget,
	detailsPublisherReplaceTarget,
	detailsLanguageReplaceTarget,
	detailsISBN10ReplaceTarget,
	detailsISBN13ReplaceTarget,
)

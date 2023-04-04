package template

import "fmt"

const (
	titleReplaceTarget       = "{TITLE}"
	descriptionReplaceTarget = "{DESCRIPTION_HTML}"
)

var summaryHTML = fmt.Sprintf(`
	<h3>%s</h3>
	<p>%s</p>`,
	titleReplaceTarget,
	descriptionReplaceTarget,
)

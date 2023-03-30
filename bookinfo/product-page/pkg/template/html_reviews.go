package template

import "fmt"

const (
	reviewsTextReplaceTarget     = "{TEXT}"
	reviewsReviewerReplaceTarget = "{REVIEWER}"
	reviewsServedByReplaceTarget = "{SERVED_BY}"
	reviewsReviewsReplaceTarget  = "{REVIEWS}"
	reviewsStarsReplaceTarget    = "{STARS}"
	filledStar                   = `<span class="star">&#9733;</span>`
	emptyStar                    = `<span class="empty-star">&#9734;</span>`
)

var reviewsHTML = fmt.Sprintf(
	`
	%s
	<dl>
	<dt>Reviews served by:</dt>
	<dd><u>%s</u></dd>
	</dl>
	`,
	reviewsReviewsReplaceTarget,
	reviewsServedByReplaceTarget,
)

var reviewHTML = fmt.Sprintf(
	`
	<blockquote>
		<p>%s</p>
		<small>%s</small>
		<div>
		%s
		</div>
	</blockquote>
	`,
	reviewsTextReplaceTarget,
	reviewsReviewerReplaceTarget,
	reviewsStarsReplaceTarget,
)

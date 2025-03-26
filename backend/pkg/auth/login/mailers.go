package login

import (
	"context"

	"github.com/pecet3/las-test-pdf/utils"
)

type LoginMailer struct{}

func (LoginMailer) Send(ctx context.Context, to, code, userName string) error {
	subject := "LAS-PDF Magic Code (noreply)"
	body := `
    <html>
    	<body>
    		<h2>Hello ` + userName + `,</h2>
    			<p>This is a magic code:</p>
				<h1>
					<i>` + code + `</i>
				</h1>
				<i>Please, copy and then paste it to the App.</i>
    	</body>
    </html>
    `
	if err := utils.SendEmail(ctx, to, subject, body); err != nil {
		return err
	}
	return nil
}

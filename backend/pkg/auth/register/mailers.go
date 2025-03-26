package register

import (
	"context"

	"github.com/pecet3/las-test-pdf/utils"
)

type RegisterMailer struct{}

func (RegisterMailer) Send(ctx context.Context, to, code, userName string) error {
	subject := "LAS-PDF Welcome! (noreply)"
	body := `
    <html>
    	<body>
    		<h2>Hello ` + userName + `,</h2>
				<p>We are happy You joined us!</p>
				</br>
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

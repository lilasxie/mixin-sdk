package messenger

import (
	"context"
	"encoding/json"

	"github.com/fox-one/mixin-sdk/utils"
)

type Attachment struct {
	AttachmentId string `json:"attachment_id"`
	UploadUrl    string `json:"upload_url"`
	ViewUrl      string `json:"view_url"`
}

func (m Messenger) CreateAttachment(ctx context.Context) (*Attachment, error) {
	data, err := m.Request(ctx, "POST", "/attachments", nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Attachment Attachment `json:"data"`
	}
	err = json.Unmarshal(data, &resp)
	return &resp.Attachment, err
}

func (m Messenger) Upload(ctx context.Context, file []byte) (string, string, error) {
	attachment, err := m.CreateAttachment(ctx)
	if err != nil {
		return "", "", err
	}

	headers := []string{"x-amz-acl", "public-read"}
	req, err := utils.NewRequest(attachment.UploadUrl, "PUT", string(file), headers...)
	if err != nil {
		return "", "", err
	}

	_, err = utils.DoRequest(req)
	return attachment.AttachmentId, attachment.ViewUrl, nil
}

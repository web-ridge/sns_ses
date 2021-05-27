# sns_ses
Amazon SQS (Message Queuing Service) golang convert and parse to json structs of notification-contents of the Amazon SNS message of the Amazon SES service (Amazon Simple Email Service)
[amazon/notification-contents.html](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notification-contents.html#top-level-json-object)

# Bindings for
- Received Mail
- Bounce
- s3 Receipt
- Complaint
- Delivery

## Example
```golang
func handlePubSubMailMessage(ctx context.Context, msg *sqs.Message) error {
	body := &sns_ses.Body{}
	err := json.Unmarshal([]byte(*msg.Body), &body)
	if err != nil {
		return err
	}

	message := &sns_ses.Message{}
	err = json.Unmarshal([]byte(body.Message), &message)
	if err != nil {
		return err
	}
  	switch message.NotificationType {
	case sns_ses.NotificationTypeBounce:
	case sns_ses.NotificationTypeComplaint:
	case sns_ses.NotificationTypeDelivery:
	case sns_ses.NotificationTypeReceived:
		messageContent, err := incomingMailS3.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(message.Receipt.Action.BucketName),
			Key:    aws.String(message.Receipt.Action.ObjectKey),
		})

	}
}
```

## Extra
If you want to strip signature and quoted-replies from the received message you could use another library we made: [web-ridge/email-reply-parser](https://github.com/web-ridge/email-reply-parser)

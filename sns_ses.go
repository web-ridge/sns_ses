package sns_ses

import "time"

type NotificationType string

const (
	NotificationTypeBounce    NotificationType = "Bounce"
	NotificationTypeComplaint NotificationType = "Complaint"
	NotificationTypeDelivery  NotificationType = "Delivery"
	NotificationTypeReceived  NotificationType = "Received"
)

type BounceType string

const (
	BounceTypeUndetermined BounceType = "Undetermined"
	BounceTypePermanent    BounceType = "Permanent"
	BounceTypeTransient    BounceType = "Transient"
)

type BounceSubType string

const (
	BounceSubTypeUndetermined BounceSubType = "Undetermined"
	BounceSubTypeGeneral      BounceSubType = "General"
	BounceSubTypeNoEmail      BounceSubType = "NoEmail"
)

type Body struct {
	Type        string `json:"type"`
	Subject     string `json:"subject"`
	MessageId   string `json:"messageId"`
	Message     string `json:"message"`
	PublishTime int64  `json:"publishTime"`
}

type Message struct {
	NotificationType NotificationType `json:"notificationType"`
	Mail             Mail             `json:"mail"`
	Bounce           *Bounce          `json:"bounce"`
	Complaint        *Complaint       `json:"complaint"`
	Delivery         *Delivery        `json:"delivery"`
	Receipt          *Receipt         `json:"receipt"`
}

type Receipt struct {
	Timestamp            time.Time `json:"timestamp"`
	ProcessingTimeMillis int       `json:"processingTimeMillis"`
	Recipients           []string  `json:"recipients"`
	SpamVerdict          struct {
		Status string `json:"status"`
	} `json:"spamVerdict"`
	VirusVerdict struct {
		Status string `json:"status"`
	} `json:"virusVerdict"`
	SpfVerdict struct {
		Status string `json:"status"`
	} `json:"spfVerdict"`
	DkimVerdict struct {
		Status string `json:"status"`
	} `json:"dkimVerdict"`
	DmarcVerdict struct {
		Status string `json:"status"`
	} `json:"dmarcVerdict"`
	Action struct {
		Type            string `json:"type"`
		TopicArn        string `json:"topicArn"`
		BucketName      string `json:"bucketName"`
		ObjectKeyPrefix string `json:"objectKeyPrefix"`
		ObjectKey       string `json:"objectKey"`
	} `json:"action"`
}

type Mail struct {
	Timestamp        string   `json:"timestamp"`
	MessageId        string   `json:"messageId"`
	Source           string   `json:"source"`
	SourceArn        string   `json:"sourceArn"`
	SourceIp         string   `json:"sourceIp"`
	SendingAccountId string   `json:"sendingAccountId"`
	Destination      []string `json:"destination"`
	HeadersTruncated bool     `json:"headersTruncated"`
	Headers          []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"headers"`
	CommonHeaders struct {
		From      []string `json:"from"`
		Date      string   `json:"date"`
		To        []string `json:"to"`
		MessageId string   `json:"messageId"`
		Subject   string   `json:"subject"`
	} `json:"commonHeaders"`
}

type Bounce struct {
	BounceType        BounceType    `json:"bounceType"`
	BounceSubType     BounceSubType `json:"bounceSubType"`
	BouncedRecipients []struct {
		Status         string `json:"status"`
		Action         string `json:"action"`
		DiagnosticCode string `json:"diagnosticCode,omitempty"`
		EmailAddress   string `json:"emailAddress"`
	} `json:"bouncedRecipients"`
	ReportingMTA string    `json:"reportingMTA"`
	Timestamp    time.Time `json:"timestamp"`
	FeedbackId   string    `json:"feedbackId"`
	RemoteMtaIp  string    `json:"remoteMtaIp"`
}

type Complaint struct {
	UserAgent            string `json:"userAgent"`
	ComplainedRecipients []struct {
		EmailAddress string `json:"emailAddress"`
	} `json:"complainedRecipients"`
	ComplaintFeedbackType string    `json:"complaintFeedbackType"`
	ArrivalDate           time.Time `json:"arrivalDate"`
	Timestamp             time.Time `json:"timestamp"`
	FeedbackId            string    `json:"feedbackId"`
}

type Delivery struct {
	Timestamp            time.Time `json:"timestamp"`
	ProcessingTimeMillis int       `json:"processingTimeMillis"`
	Recipients           []string  `json:"recipients"`
	SmtpResponse         string    `json:"smtpResponse"`
	ReportingMTA         string    `json:"reportingMTA"`
	RemoteMtaIp          string    `json:"remoteMtaIp"`
}

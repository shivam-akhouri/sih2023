package utils

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func GetClient() *firestore.Client {
	ctx := context.Background()
	const ProjectId = "sih138"
	cred := []byte(`{
		"type": "service_account",
		"project_id": "sih138",
		"private_key_id": "88a7c23fa41ade7102f5db3d62ad65107a155880",
		"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDOgcdRj7H/FYR4\nqL92BQOULQfV9gLdVN5h1Lo2YTJdzGMMBfYexNpA2dlT7utJJd4tgzO1uNyjjGzD\n99FR84e0mQjDizDQhxk8rBcbSpNu7Lfs66pSWg+ojQdDmZ61/JZn61X9UclFjKAa\n1I2kTxVN1xLOtHC0dwNocvQQxjPnu4wPqY9c4Umq9CZK0g1sm2l68lURdD3l0L+M\nLdCI86SKNAMknOnJE26yJSo9hXJV5t7wZ2rRekLTI33TwOVYl2wiCETYa+OKpMbU\n3A7wSEsflq+wGo3hQQtla79FPQH87erb84aQGLUEzVclmkqV1mo1Qhs6yNU5KwNw\nCWOGH3JVAgMBAAECggEAC92prA5XcjqUitFZdQjbOImw3FK/PeQJRjJQw9yoZfrf\nSxrRllmF1/mckwXUt7aJEnDqSPwXE3+3ZThJaw3VA6PSfNfz5VuJ7/JMfpk1qQdJ\nlz/AYOyDX1JBM1lxpO/fXZfIcF6hc9FdhsUlYWZY6UjdMAruufPnunn2e7tD0sIZ\n3r3+cOe+WKRDDqeYgl8tfSlGRKQZCx9M6jKpqWgTn52iozpntEjb2fjqNnRyRyts\ndTExkp4B2RaxFWgeDzXK4comrTnMTMXzPTQ4Lz58obiEeEwqUANQO910r2N7rXpH\nB7Gec3KE7ZAqkJIJzm7U8qCTq6yWJfeE9IuaM0CJoQKBgQDsAt8X3kmJE5jZDvqw\nUrryXlvKyV+2hs0xJILbYbXtD58FrbkHGUSHFOGfN2tH2kE5gmqsY/25spXUyDhi\noYFiRyd18C+gBn5+dmKVt2CA+pjppe8CKcanaQHSRLs7+EtyoVIAp6+ghmmdQieG\nUuzrnyAWyqBARgA5BpatnPEJ9QKBgQDf/zQy8T9XvPW5bnFdIyS+rDs2WPuM2H1k\nmL6y3fDybCKpnnPKB9tEn755cUwzhD/9HhCtpjG95BkyaUhWtYhXsBTvYGTQD/hC\nNM4MtbliSgfoJRTZffUdHtmu8h4THIQdk//AiX5q6L1Zvz3IoZECFO7bt6PlC4bp\nESdWOGWq4QKBgBx/CELQgx568ESdC2XIx2vGOt0UTWszgYkCeI5WkOnLrg64eBPE\nWELbIZbg0SI+1wtDlO+1rwgB6RdlCXHXQFEGeFxEIOv4fMN2NdcP6iYIm/fyIhan\nxqWmByNFSzolc66mjckodgsyz4O+XhlY3+RfNTerRYgdK5+3DPBvd6cVAoGBAIDY\nyevrvaX4edhaYTdT+dT8OWs/sCMKfMaa3kcBR5ILBb1P+gbikgxdXzF4LH8Nuckd\nvDZG/SHfzWakDQpjjbPzA3fWmObd6M4cldHVzndtYjL08ZNR0ucXUdxCRW503sIs\n2dakwGU6z4CKcz2vQjxSUjts3Qfpo4b5aQk9k63BAoGBAKwSXDQlkqFvqWQPELQv\nhpx0KZfCwGHeMitU7bgnO4IRecHRSFyFC+qiNLPqNJQg0zkNJ0rU3peIuubv6AdR\njq/wycOZEXB5gRfC61gCba4k7HdHU23Er+XHH5VBhlzUVAwE++rdLpe8VFzlkSjb\nhZrQN2ZMKPT0t63THW7r8Y9i\n-----END PRIVATE KEY-----\n",
		"client_email": "firebase-adminsdk-rqyoe@sih138.iam.gserviceaccount.com",
		"client_id": "112634079591657372135",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-rqyoe%40sih138.iam.gserviceaccount.com",
		"universe_domain": "googleapis.com"
	  }`)

	opt := option.WithCredentialsJSON(cred)
	// opt := option.WithCredentialsFile("E:\\hackathon\\Smart India Hackathon\\backend\\src\\utils\\sih138-firebase-adminsdk-rqyoe-88a7c23fa4.json")
	client, err := firestore.NewClient(ctx, ProjectId, opt)
	if err != nil {
		return nil
	}
	return client
}

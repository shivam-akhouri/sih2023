package utils

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func GetClient() *firestore.Client {
	ctx := context.Background()
	const ProjectId = "sih138"
	opt := option.WithCredentialsFile("E:\\hackathon\\Smart India Hackathon\\backend\\src\\sih138-firebase-adminsdk-rqyoe-88a7c23fa4.json")
	client, err := firestore.NewClient(ctx, ProjectId, opt)
	if err != nil {
		return nil
	}
	return client
}

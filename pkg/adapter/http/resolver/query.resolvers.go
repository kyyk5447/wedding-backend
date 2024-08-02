package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"log"

	"github.com/chaki8923/wedding-backend/pkg/domain/model"
	"github.com/chaki8923/wedding-backend/pkg/lib/graph/generated"
	sentry "github.com/getsentry/sentry-go"
)

// GetMessages is the resolver for the getMessages field.
func (r *queryResolver) GetMessages(ctx context.Context) ([]*model.Message, error) {
	todos, err := r.MsgUseCase.GetMessages()
	if err != nil {
		err = fmt.Errorf("resolver Todos() err %w", err)
		sentry.CaptureException(err)
		return nil, err
	}

	return todos, nil
}

// GetInvitation is the resolver for the getInvitation field.
func (r *queryResolver) GetInvitation(ctx context.Context) ([]*model.Invitation, error) {
	invitation, err := r.IvtUseCase.GetInvitation()
	log.Printf("queryResolver----------------------")
	if err != nil {
		err = fmt.Errorf("resolver Todos() err %w", err)
		sentry.CaptureException(err)
		return nil, err
	}
	return invitation, nil
}

// GetAllergy is the resolver for the getAllergy field.
func (r *queryResolver) GetAllergy(ctx context.Context) ([]*model.Allergy, error) {
	allergies, err := r.AgyUseCase.GetAllergy()

	if err != nil {
		err = fmt.Errorf("resolver Invitee() err %w", err)
		sentry.CaptureException(err)
		return nil, err
	}

	return allergies, nil
}

// GetInvitee is the resolver for the getInvitee field.
func (r *queryResolver) GetInvitee(ctx context.Context) ([]*model.Invitee, error) {
	invitee, err := r.IvteeUseCase.GetInvitee()

	if err != nil {
		err = fmt.Errorf("resolver Invitee() err %w", err)
		sentry.CaptureException(err)
		return nil, err
	}

	return invitee, nil
}

// GetImages is the resolver for the getImages field.
func (r *queryResolver) GetImages(ctx context.Context) ([]*model.UploadImage, error) {
	images, err := r.UpdUseCase.GetImages()

	if err != nil {
		err = fmt.Errorf("画像取得エラー err %w", err)
		sentry.CaptureException(err)
		return nil, err
	}

	return images, nil
}

// ShowInvitation is the resolver for the showInvitation field.
func (r *queryResolver) ShowInvitation(ctx context.Context, uuid string) (*model.Invitation, error) {
	invitation, err := r.IvtUseCase.ShowInvitation(uuid)
	if err != nil {
		err = fmt.Errorf("resolver 招待状詳細() err %w", err)
		sentry.CaptureException(err)
		return nil, err
	}
	return invitation, nil
}

// ShowInvitee is the resolver for the showInvitee field.
func (r *queryResolver) ShowInvitee(ctx context.Context, uuid string) (*model.Invitee, error) {
	invitee, err := r.IvteeUseCase.ShowInvitee(uuid)
	if err != nil {
		err = fmt.Errorf("resolver 招待者詳細() err %w", err)
		sentry.CaptureException(err)
		return nil, err
	}
	return invitee, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

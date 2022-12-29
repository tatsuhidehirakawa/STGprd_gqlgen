package persistence

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/domain/entity"
)

func TestOfferListRepo_GetOfferList(t *testing.T) {
	tests := []struct {
		name    string
		offerlistID  string
		want    *entity.OfferList
		wantErr error
	}{
		{
			name:   "ok",
			offerlistID: "1",
			want: &entity.OfferList{
				ID:   "1",
				Name: "user1",
			},
			wantErr: nil,
		},
		{
			name:   "ok",
			offerlistID: "2",
			want: &entity.OfferList{
				ID:   "2",
				Name: "user2",
			},
			wantErr: nil,
		},
		{
			name:    "notExistOfferList",
			offerlistID:  "999",
			want:    nil,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := offerlistRepo.GetOfferList(context.Background(), tt.offerlistID)
			if diff := cmp.Diff(tt.wantErr, err); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want, got); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}

		})
	}

}

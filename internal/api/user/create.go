package user

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/olezhek28/clean-architecture/internal/converter"
	desc "github.com/olezhek28/clean-architecture/pkg/user_v1"
	"io"
	"net/http"
)

type responceClass struct {
	responce desc.CreateResponse
}

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	uuid, err := i.userService.Create(ctx, converter.ToUserInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Uuid: uuid,
	}, nil
}

func (i *Implementation) CreateREST(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	req := &desc.UserInfo{}
	data, err := io.ReadAll(r.Body)
	proto.Unmarshal(data, req)
	w.Write([]byte(req.LastName))
	ff := converter.ToUserInfoFromDesc(&desc.UserInfo{LastName: "sad", FirstName: "assd", Age: 12})
	uuid, err := i.userService.Create(ctx, ff)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	Uuid := desc.CreateResponse{
		Uuid: uuid,
	}
	w.Write([]byte(Uuid.Uuid))
	return

}

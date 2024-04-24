package tests

import (
	"mime/multipart"
	"reflect"
	"testing"

	"github.com/Leonargo404-code/e-commerce/pkg/products"
	productsPkg "github.com/Leonargo404-code/e-commerce/pkg/products"
	"github.com/Leonargo404-code/e-commerce/pkg/products/service"
)

func Test_service_Create(t *testing.T) {
	type (
		fields struct {
			repo func() products.Repository
		}

		args struct {
			product      *productsPkg.Product
			productImage *multipart.FileHeader
		}
	)

	// succsessProductExample := &products.Product{
	// 	ID:    uuid.New(),
	// 	Name:  "product",
	// 	Value: 50.10,
	// 	Units: 10,
	// }

	// fileHeader := &multipart.FileHeader{
	// 	Filename: "iphone.jpg",
	// 	Size:     int64(1024),
	// 	Header:   make(textproto.MIMEHeader),
	// }

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *productsPkg.Product
		wantErr bool
	}{
		{
			name: "should return a missing name error",
			fields: fields{
				repo: func() productsPkg.Repository {
					return nil
				},
			},
			args: args{
				product: &productsPkg.Product{
					Name:  "",
					Value: 10,
					Units: 1,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return a missing value error",
			fields: fields{
				repo: func() productsPkg.Repository {
					return nil
				},
			},
			args: args{
				&productsPkg.Product{
					Name:  "product",
					Value: 0.0,
					Units: 1,
				},
				&multipart.FileHeader{
					Filename: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return a missing unit error",
			fields: fields{
				repo: func() productsPkg.Repository {
					return nil
				},
			},
			args: args{
				&productsPkg.Product{
					Name:  "product",
					Value: 10.20,
					Units: 0,
				},
				&multipart.FileHeader{
					Filename: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return a missing unit error",
			fields: fields{
				repo: func() productsPkg.Repository {
					return nil
				},
			},
			args: args{
				&productsPkg.Product{
					Name:  "product",
					Value: 10.20,
					Units: 0,
				},
				&multipart.FileHeader{
					Filename: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service.Service{
				Repo: tt.fields.repo(),
			}

			got, err := s.Create(tt.args.product, tt.args.productImage)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

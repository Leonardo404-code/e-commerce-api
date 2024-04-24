package tests

import (
	"errors"
	"testing"

	"github.com/Leonargo404-code/e-commerce/pkg/products"
	productsPkg "github.com/Leonargo404-code/e-commerce/pkg/products"
	"github.com/Leonargo404-code/e-commerce/pkg/products/repository"
	"github.com/Leonargo404-code/e-commerce/pkg/products/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func TestService_Delete(t *testing.T) {
	type (
		fields struct {
			Repo func() products.Repository
		}

		args struct {
			filter *productsPkg.Filter
		}
	)

	productID := uuid.New().String()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should return a error when search by product",
			fields: fields{
				Repo: func() productsPkg.Repository {
					m := &repository.Mock{}
					m.On("Get", mock.Anything).Return(nil, errors.New("error"))
					return m
				},
			},
			args: args{
				filter: &productsPkg.Filter{},
			},
			wantErr: true,
		},
		{
			name: "should return a product not found error",
			fields: fields{
				Repo: func() productsPkg.Repository {
					m := &repository.Mock{}
					m.On("Get", mock.Anything).Return(&productsPkg.Result{
						Products: []*productsPkg.Product{},
						Total:    1,
						Page:     1,
					}, nil)

					return m
				},
			},
			args: args{
				filter: &productsPkg.Filter{},
			},
			wantErr: true,
		},
		{
			name: "should return when find more than one product",
			fields: fields{
				Repo: func() productsPkg.Repository {
					m := &repository.Mock{}
					m.On("Get", mock.Anything).Return(&productsPkg.Result{
						Products: []*productsPkg.Product{
							{ID: uuid.New()},
							{ID: uuid.New()},
						},
						Total: 1,
						Page:  1,
					}, nil)
					return m
				},
			},
			args: args{
				filter: &productsPkg.Filter{},
			},
			wantErr: true,
		},
		{
			name: "should return error in delete product",
			fields: fields{
				Repo: func() productsPkg.Repository {
					m := &repository.Mock{}
					m.On("Get", mock.Anything).Return(&productsPkg.Result{
						Products: []*productsPkg.Product{
							{ID: uuid.New()},
						},
						Total: 1,
						Page:  1,
					}, nil)
					m.On("Delete", mock.Anything, mock.Anything).Return(errors.New("error"))
					return m
				},
			},
			args: args{
				filter: &productsPkg.Filter{
					ID: []string{productID},
					Pagination: &productsPkg.Pagination{
						Page:       1,
						MaxPerPage: 1,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "should return error in delete product",
			fields: fields{
				Repo: func() productsPkg.Repository {
					m := &repository.Mock{}
					m.On("Get", mock.Anything).Return(&productsPkg.Result{
						Products: []*productsPkg.Product{
							{ID: uuid.New()},
						},
						Total: 1,
						Page:  1,
					}, nil)
					m.On("Delete", mock.Anything, mock.Anything).Return(nil)
					return m
				},
			},
			args: args{
				filter: &productsPkg.Filter{
					ID: []string{productID},
					Pagination: &productsPkg.Pagination{
						Page:       1,
						MaxPerPage: 1,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service.Service{
				Repo: tt.fields.Repo(),
			}

			if err := s.Delete(tt.args.filter); (err != nil) != tt.wantErr {
				t.Errorf("Service.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

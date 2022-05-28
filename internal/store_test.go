package internal

import (
	"context"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"reflect"
	"testing"
)

func TestStore_GetUnusedConfigMaps(t *testing.T) {
	type fields struct {
		Context                    context.Context
		Namespaces                 []string
		CoreV1                     v1.CoreV1Interface
		usedSecrets                []Item
		usedConfigMaps             []Item
		usedServiceAccounts        []Item
		usedPersistentVolumeClaims []Item
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Store{
				Context:                    tt.fields.Context,
				Namespaces:                 tt.fields.Namespaces,
				CoreV1:                     tt.fields.CoreV1,
				usedSecrets:                tt.fields.usedSecrets,
				usedConfigMaps:             tt.fields.usedConfigMaps,
				usedServiceAccounts:        tt.fields.usedServiceAccounts,
				usedPersistentVolumeClaims: tt.fields.usedPersistentVolumeClaims,
			}
			got, err := k.GetUnusedConfigMaps()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUnusedConfigMaps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnusedConfigMaps() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_GetUnusedPersistentVolumeClaims(t *testing.T) {
	type fields struct {
		Context                    context.Context
		Namespaces                 []string
		CoreV1                     v1.CoreV1Interface
		usedSecrets                []Item
		usedConfigMaps             []Item
		usedServiceAccounts        []Item
		usedPersistentVolumeClaims []Item
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Store{
				Context:                    tt.fields.Context,
				Namespaces:                 tt.fields.Namespaces,
				CoreV1:                     tt.fields.CoreV1,
				usedSecrets:                tt.fields.usedSecrets,
				usedConfigMaps:             tt.fields.usedConfigMaps,
				usedServiceAccounts:        tt.fields.usedServiceAccounts,
				usedPersistentVolumeClaims: tt.fields.usedPersistentVolumeClaims,
			}
			got, err := k.GetUnusedPersistentVolumeClaims()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUnusedPersistentVolumeClaims() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnusedPersistentVolumeClaims() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_GetUnusedSecrets(t *testing.T) {
	type fields struct {
		Context                    context.Context
		Namespaces                 []string
		CoreV1                     v1.CoreV1Interface
		usedSecrets                []Item
		usedConfigMaps             []Item
		usedServiceAccounts        []Item
		usedPersistentVolumeClaims []Item
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Store{
				Context:                    tt.fields.Context,
				Namespaces:                 tt.fields.Namespaces,
				CoreV1:                     tt.fields.CoreV1,
				usedSecrets:                tt.fields.usedSecrets,
				usedConfigMaps:             tt.fields.usedConfigMaps,
				usedServiceAccounts:        tt.fields.usedServiceAccounts,
				usedPersistentVolumeClaims: tt.fields.usedPersistentVolumeClaims,
			}
			got, err := k.GetUnusedSecrets()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUnusedSecrets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnusedSecrets() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_GetUnusedServiceAccounts(t *testing.T) {
	type fields struct {
		Context                    context.Context
		Namespaces                 []string
		CoreV1                     v1.CoreV1Interface
		usedSecrets                []Item
		usedConfigMaps             []Item
		usedServiceAccounts        []Item
		usedPersistentVolumeClaims []Item
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Store{
				Context:                    tt.fields.Context,
				Namespaces:                 tt.fields.Namespaces,
				CoreV1:                     tt.fields.CoreV1,
				usedSecrets:                tt.fields.usedSecrets,
				usedConfigMaps:             tt.fields.usedConfigMaps,
				usedServiceAccounts:        tt.fields.usedServiceAccounts,
				usedPersistentVolumeClaims: tt.fields.usedPersistentVolumeClaims,
			}
			got, err := k.GetUnusedServiceAccounts()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUnusedServiceAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnusedServiceAccounts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_GetUsedResources(t *testing.T) {
	type fields struct {
		Context                    context.Context
		Namespaces                 []string
		CoreV1                     v1.CoreV1Interface
		usedSecrets                []Item
		usedConfigMaps             []Item
		usedServiceAccounts        []Item
		usedPersistentVolumeClaims []Item
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Store{
				Context:                    tt.fields.Context,
				Namespaces:                 tt.fields.Namespaces,
				CoreV1:                     tt.fields.CoreV1,
				usedSecrets:                tt.fields.usedSecrets,
				usedConfigMaps:             tt.fields.usedConfigMaps,
				usedServiceAccounts:        tt.fields.usedServiceAccounts,
				usedPersistentVolumeClaims: tt.fields.usedPersistentVolumeClaims,
			}
			if err := k.GetUsedResources(); (err != nil) != tt.wantErr {
				t.Errorf("GetUsedResources() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_missingItems(t *testing.T) {
	type args struct {
		a []Item
		b []Item
	}
	tests := []struct {
		name string
		args args
		want []Item
	}{
		{
			name: "test-equal",
			args: args{
				a: []Item{
					{
						Name:       "equal",
						Namespaces: "test",
					}},
				b: []Item{
					{
						Name:       "equal",
						Namespaces: "test",
					},
				},
			},
			want: nil,
		},
		{
			name: "test-different",
			args: args{
				a: []Item{
					{
						Name:       "equal",
						Namespaces: "test",
					}},
				b: []Item{
					{
						Name:       "different",
						Namespaces: "test",
					},
				},
			},
			want: []Item{
				{
					Name:       "different",
					Namespaces: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingItems(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("missingItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

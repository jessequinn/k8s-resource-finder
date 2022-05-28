package internal

import (
	"context"
	"reflect"
	"testing"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func TestStore_GetUnusedConfigMaps(t *testing.T) {
	fakeClientSet := fake.NewSimpleClientset(
		&apiv1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test",
				Namespace: "test",
			},
		},
		&apiv1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-different",
				Namespace: "test",
			},
		},
	)
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
		{
			name: "test-different",
			fields: fields{
				Context:     context.Background(),
				Namespaces:  []string{"test"},
				CoreV1:      fakeClientSet.CoreV1(),
				usedSecrets: nil,
				usedConfigMaps: []Item{{
					Name:       "test",
					Namespaces: "test",
				}},
				usedServiceAccounts:        nil,
				usedPersistentVolumeClaims: nil,
			},
			want: []Item{{
				Name:       "test-different",
				Namespaces: "test",
			}},
			wantErr: false,
		},
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
	fakeClientSet := fake.NewSimpleClientset(
		&apiv1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test",
				Namespace: "test",
			},
		},
		&apiv1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-different",
				Namespace: "test",
			},
		},
	)
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
		{
			name: "test-different",
			fields: fields{
				Context:             context.Background(),
				Namespaces:          []string{"test"},
				CoreV1:              fakeClientSet.CoreV1(),
				usedSecrets:         nil,
				usedConfigMaps:      nil,
				usedServiceAccounts: nil,
				usedPersistentVolumeClaims: []Item{{
					Name:       "test",
					Namespaces: "test",
				}},
			},
			want: []Item{{
				Name:       "test-different",
				Namespaces: "test",
			}},
			wantErr: false,
		},
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
	fakeClientSet := fake.NewSimpleClientset(
		&apiv1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test",
				Namespace: "test",
			},
		},
		&apiv1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-different",
				Namespace: "test",
			},
		},
	)
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
		{
			name: "test-different",
			fields: fields{
				Context:    context.Background(),
				Namespaces: []string{"test"},
				CoreV1:     fakeClientSet.CoreV1(),
				usedSecrets: []Item{{
					Name:       "test",
					Namespaces: "test",
				}},
				usedConfigMaps:             nil,
				usedServiceAccounts:        nil,
				usedPersistentVolumeClaims: nil,
			},
			want: []Item{{
				Name:       "test-different",
				Namespaces: "test",
			}},
			wantErr: false,
		},
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
	fakeClientSet := fake.NewSimpleClientset(
		&apiv1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test",
				Namespace: "test",
			},
		},
		&apiv1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-different",
				Namespace: "test",
			},
		},
	)
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
		{
			name: "",
			fields: fields{
				Context:        context.Background(),
				Namespaces:     []string{"test"},
				CoreV1:         fakeClientSet.CoreV1(),
				usedSecrets:    nil,
				usedConfigMaps: nil,
				usedServiceAccounts: []Item{{
					Name:       "test",
					Namespaces: "test",
				}},
				usedPersistentVolumeClaims: nil,
			},
			want: []Item{{
				Name:       "test-different",
				Namespaces: "test",
			}},
			wantErr: false,
		},
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

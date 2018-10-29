package sign

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/cloudflare/cfssl/config"
	"github.com/massiveco/serverlessl/store"
)

type FakeStore struct {
	file []byte
	store.Store
}

func (fs FakeStore) FetchFile(name string, contents *bytes.Buffer) error {

	contents.Write(fs.file)

	return nil
}

func Test_fetchProfiles(t *testing.T) {
	fstore := FakeStore{
		file: []byte("{\"signing\": {}}"),
	}
	wantCfg := config.Config{
		Signing: &config.Signing{},
	}
	type args struct {
		store store.Store
	}
	tests := []struct {
		name    string
		args    args
		wantCfg config.Config
		wantErr bool
	}{
		{
			name: "ok",
			args: args{store: fstore},
			wantCfg: wantCfg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCfg, err := fetchProfiles(tt.args.store)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchProfiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("fetchProfiles() = %v, want %v", gotCfg, tt.wantCfg)
			}
		})
	}
}

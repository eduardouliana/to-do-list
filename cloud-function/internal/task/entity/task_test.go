package entity

import (
	"reflect"
	"testing"
)

func TestNewTask(t *testing.T) {
	type args struct {
		id          string
		description string
		done        bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Task
		wantErr bool
	}{
		{
			name: "Test 001 - Trying to create task already done",
			args: args{
				id:          "123",
				description: "Test",
				done:        true,
			},
			wantErr: true,
		},
		{
			name: "Test 002 - Create task",
			args: args{
				id:          "123",
				description: "Test",
				done:        false,
			},
			want: &Task{
				Id:          "123",
				Description: "Test",
				Done:        false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTask(tt.args.id, tt.args.description, tt.args.done)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

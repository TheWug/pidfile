package pidfile

import (
	"testing"
	"io/ioutil"
	"os"
)

func cleanup(p *PIDFile) {
	if p != nil {
		p.Remove()
	}
}

func Test_Pidfile(t *testing.T) {
	testcases := map[string]struct{
		filename string
		expects_error bool
	}{
		"normal": {"./pidfile_test.pid", false},
		"nonexistent": {"/nonexistent/directory/pidfile_test.pid", true},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			p, e := Open(v.filename)
			defer cleanup(p)
			if v.expects_error != (e != nil) {
				t.Errorf("Unexpected error state: %v (%t, %t)", e, e != nil, v.expects_error)
			}

			data, err := ioutil.ReadFile(v.filename)
			if v.expects_error != (err != nil) { t.Errorf("Error reading pidfile: %v", err) }
			if len(data) != 0 { t.Errorf("Unexpected contents of pidfile: %s", data) }
		})
	}

	v := testcases["normal"]
	t.Run("write", func(t *testing.T) {
		p, err := Open(v.filename)
		defer cleanup(p)
		if err != nil { t.Errorf("Unexpected error: %v", err) }

		err = p.Write()
		if err != nil { t.Errorf("Error writing pidfile: %v", err) }

		data, err := ioutil.ReadFile(v.filename)
		if err != nil { t.Errorf("Error reading pidfile: %v", err) }
		if len(data) == 0 { t.Errorf("Unexpected empty pidfile: %s", data) }
	})

	t.Run("close", func(t *testing.T) {
		p, err := Open(v.filename)
		p.Write()
		err = p.Close()
		if err != nil { t.Errorf("Error closing pidfile: %v", err) }

		data, err := ioutil.ReadFile(v.filename)
		if err != nil { t.Errorf("Error reading pidfile: %v", err) }
		if len(data) == 0 { t.Errorf("Unexpected empty pidfile: %s", data) }

		err = os.Remove(v.filename)
		if err != nil { t.Errorf("Unexpected error removing pidfile: %v", err) }
	})

	t.Run("remove", func(t *testing.T) {
		p, err := Open(v.filename)
		p.Write()
		err = p.Remove()
		if err != nil { t.Errorf("Error removing pidfile: %v", err) }

		data, err := ioutil.ReadFile(v.filename)
		if err == nil { t.Errorf("Unexpected success reading pidfile: %v", data) }

		err = os.Remove(v.filename)
		if err == nil { t.Errorf("Unexpected success removing pidfile: %v", err) }
	})
}


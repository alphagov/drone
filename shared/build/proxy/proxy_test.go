package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	// test creating a proxy with a few different
	// addresses, and our ability to create the
	// proxy shell script.
	p := Proxy{}
	p.Set("8080", "172.1.4.5")
	b := p.Bytes()

	expected := `#!/bin/bash
[ -x /usr/bin/socat ] && socat TCP-LISTEN:8080,fork TCP:172.1.4.5:8080 &
`
	if string(b) != expected {
		t.Errorf("Invalid proxy \n%s", expected)
	}

	// test creating a proxy script when there
	// are no proxy addresses added to the map
	p = Proxy{}
	b = p.Bytes()
	expected = "#!/bin/bash\n"
	if string(b) != expected {
		t.Errorf("Invalid proxy \n%s", expected)
	}
}

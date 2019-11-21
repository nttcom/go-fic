package testing

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"sync"
	"testing"
	"time"

	fic "github.com/nttcom/go-fic"
	th "github.com/nttcom/go-fic/testhelper"
	"github.com/nttcom/go-fic/testhelper/client"
)

func TestAuthenticatedHeaders(t *testing.T) {
	p := &fic.ProviderClient{
		TokenID: "1234",
	}
	expected := map[string]string{"X-Auth-Token": "1234"}
	actual := p.AuthenticatedHeaders()
	th.CheckDeepEquals(t, expected, actual)
}

func TestUserAgent(t *testing.T) {
	p := &fic.ProviderClient{}

	p.UserAgent.Prepend("custom-user-agent/2.4.0")
	expected := "custom-user-agent/2.4.0 go-fic/1.0.0"
	actual := p.UserAgent.Join()
	th.CheckEquals(t, expected, actual)

	p.UserAgent.Prepend("another-custom-user-agent/0.3.0", "a-third-ua/5.9.0")
	expected = "another-custom-user-agent/0.3.0 a-third-ua/5.9.0 custom-user-agent/2.4.0 go-fic/1.0.0"
	actual = p.UserAgent.Join()
	th.CheckEquals(t, expected, actual)

	p.UserAgent = fic.UserAgent{}
	expected = "go-fic/1.0.0"
	actual = p.UserAgent.Join()
	th.CheckEquals(t, expected, actual)
}

func TestConcurrentReauth(t *testing.T) {
	var info = struct {
		numreauths int
		mut        *sync.RWMutex
	}{
		0,
		new(sync.RWMutex),
	}

	numconc := 20

	prereauthTok := client.TokenID
	postreauthTok := "12345678"

	p := new(fic.ProviderClient)
	p.UseTokenLock()
	p.SetToken(prereauthTok)
	p.ReauthFunc = func() error {
		time.Sleep(1 * time.Second)
		p.AuthenticatedHeaders()
		info.mut.Lock()
		info.numreauths++
		info.mut.Unlock()
		p.TokenID = postreauthTok
		return nil
	}

	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Auth-Token") != postreauthTok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		info.mut.RLock()
		hasReauthed := info.numreauths != 0
		info.mut.RUnlock()

		if hasReauthed {
			th.CheckEquals(t, p.Token(), postreauthTok)
		}

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{}`)
	})

	wg := new(sync.WaitGroup)
	reqopts := new(fic.RequestOpts)
	reqopts.MoreHeaders = map[string]string{
		"X-Auth-Token": prereauthTok,
	}

	for i := 0; i < numconc; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := p.Request("GET", fmt.Sprintf("%s/route", th.Endpoint()), reqopts)
			th.CheckNoErr(t, err)
			if resp == nil {
				t.Errorf("got a nil response")
				return
			}
			if resp.Body == nil {
				t.Errorf("response body was nil")
				return
			}
			defer resp.Body.Close()
			actual, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("error reading response body: %s", err)
				return
			}
			th.CheckByteArrayEquals(t, []byte(`{}`), actual)
		}()
	}

	wg.Wait()

	th.AssertEquals(t, 1, info.numreauths)
}

func TestReauthEndLoop(t *testing.T) {

	p := new(fic.ProviderClient)
	p.UseTokenLock()
	p.SetToken(client.TokenID)
	p.ReauthFunc = func() error {
		// Reauth func is working and returns no error
		return nil
	}

	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
		// route always return 401
		w.WriteHeader(http.StatusUnauthorized)
		return
	})

	reqopts := new(fic.RequestOpts)
	_, err := p.Request("GET", fmt.Sprintf("%s/route", th.Endpoint()), reqopts)
	if err == nil {
		t.Errorf("request ends with a nil error")
		return
	}

	if reflect.TypeOf(err) != reflect.TypeOf(&fic.ErrErrorAfterReauthentication{}) {
		t.Errorf("error is not an ErrErrorAfterReauthentication")
	}
}

package consent_test

import (
	"github.com/ory/hydra/driver/config"
	"github.com/ory/hydra/internal"
	hydra "github.com/ory/hydra/internal/httpclient/client"
	"github.com/ory/hydra/internal/testhelpers"
	"github.com/ory/x/urlx"
	"testing"
	"time"
)

func TestLogoutFlows(t *testing.T) {
	reg := internal.NewMockedRegistry(t)
	reg.Config().Set(config.KeyAccessTokenStrategy, "opaque")
	reg.Config().Set(config.KeyConsentRequestMaxAge, time.Hour)
	publicTS, adminTS := testhelpers.NewOAuth2Server(t, reg)

	t.Logf("Public URL: %s", publicTS.URL)
	t.Logf("Admin URL: %s", adminTS.URL)

	adminClient := hydra.NewHTTPClientWithConfig(nil, &hydra.TransportConfig{Schemes: []string{"http"}, Host: urlx.ParseOrPanic(adminTS.URL).Host})

}

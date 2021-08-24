package common

import (
	"fmt"
	tcerr "github.com/iotopo/topstack-sdk-go/topstack/common/errors"
	"github.com/iotopo/topstack-sdk-go/topstack/common/gobreaker"
	"github.com/iotopo/topstack-sdk-go/topstack/common/wsse"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	tchttp "github.com/iotopo/topstack-sdk-go/topstack/common/http"
	"github.com/iotopo/topstack-sdk-go/topstack/common/profile"
)

type Client struct {
	region      string
	httpClient  *http.Client
	httpProfile *profile.HttpProfile
	profile     *profile.ClientProfile
	credential  CredentialIface
	//signMethod      string
	//unsignedPayload bool
	debug bool

	cb    *gobreaker.CircuitBreaker
	//rb              *circuitBreaker
}

func (c *Client) Send(request tchttp.Request, response tchttp.Response) (err error) {
	if request.GetScheme() == "" {
		request.SetScheme(c.httpProfile.Scheme)
	}

	if request.GetRootDomain() == "" {
		request.SetRootDomain(c.httpProfile.RootDomain)
	}

	if request.GetDomain() == "" {
		domain := c.httpProfile.Endpoint
		if domain == "" {
			domain = request.GetServiceDomain(request.GetService())
		}
		request.SetDomain(domain)
	}

	if request.GetHttpMethod() == "" {
		request.SetHttpMethod(c.httpProfile.ReqMethod)
	}

	//// reflect to inject client if field ClientToken exists and retry feature is enabled
	//if c.profile.NetworkFailureMaxRetries > 0 || c.profile.RateLimitExceededMaxRetries > 0 {
	//	safeInjectClientToken(request)
	//}

	if c.profile.DisableBreaker || c.cb == nil {
		return c.sendWithSignature(request, response)
	} else {
		return c.sendWithRegionBreaker(request, response)
	}
}

func (c *Client) sendWithRegionBreaker(request tchttp.Request, response tchttp.Response) (err error) {
	defer func() {
		e := recover()
		if e != nil {
			msg := fmt.Sprintf("%s", e)
			err = tcerr.NewTopStackSDKError("ClientError.CircuitBreakerError", msg, "")
		}
	}()

	_, err = c.cb.Execute(func() (interface{}, error) {
		err := c.sendWithSignature(request, response)
		return nil, err
	})

	return err
}

func (c *Client) sendWithSignature(request tchttp.Request, response tchttp.Response) (err error) {

	httpRequest, err := http.NewRequest(request.GetHttpMethod(), request.GetUrl(), request.GetBodyReader())
	if err != nil {
		return err
	}
	if request.GetHttpMethod() == "GET" {
		httpRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else { // POST
		httpRequest.Header.Set("Content-Type", "application/json")
	}

	httpRequest.Header.Set("Host", request.GetDomain())
	httpRequest.Header.Set("X-TP-Action", request.GetAction())
	httpRequest.Header.Set("X-TP-Version", request.GetVersion())
	httpRequest.Header.Set("X-TP-Language", c.profile.Language)

	httpResponse, err := c.sendWithRateLimitRetry(httpRequest, isRetryable(request))
	if err != nil {
		return err
	}
	err = tchttp.ParseFromHttpResponse(httpResponse, response)
	return err
}

// send http request
func (c *Client) sendHttp(request *http.Request) (response *http.Response, err error) {
	if c.debug {
		outBytes, err := httputil.DumpRequest(request, true)
		if err != nil {
			log.Printf("[ERROR] dump request failed because %s", err)
			return nil, err
		}
		log.Printf("[DEBUG] http request = %s", outBytes)
	}

	response, err = c.httpClient.Do(request)
	return response, err
}

func (c *Client) GetRegion() string {
	return c.region
}

func (c *Client) Init() *Client {
	c.httpClient = &http.Client{
		Transport: &wsse.Transport{
			Username: c.credential.GetSecretId(),
			Password: c.credential.GetSecretKey(),
		},
	}
	//c.region = region
	//c.signMethod = SHA256
	c.debug = false
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return c
}

func (c *Client) WithSecretId(secretId, secretKey string) *Client {
	c.credential = NewCredential(secretId, secretKey)
	return c
}

func (c *Client) WithCredential(cred CredentialIface) *Client {
	c.credential = cred
	return c
}

func (c *Client) WithProfile(clientProfile *profile.ClientProfile) *Client {
	c.profile = clientProfile
	if !c.profile.DisableBreaker {
		c.withRegionBreaker()
	}
	//c.signMethod = clientProfile.SignMethod
	//c.unsignedPayload = clientProfile.UnsignedPayload
	c.httpProfile = clientProfile.HttpProfile
	c.debug = clientProfile.Debug
	c.httpClient.Timeout = time.Duration(c.httpProfile.ReqTimeout) * time.Second
	return c
}

//
//func (c *Client) WithSignatureMethod(method string) *Client {
//	c.signMethod = method
//	return c
//}

//func (c *Client) WithHttpTransport(transport http.RoundTripper) *Client {
//	c.httpClient.Transport = transport
//	return c
//}

func (c *Client) WithDebug(flag bool) *Client {
	c.debug = flag
	return c
}

//// WithProvider use specify provider to get a credential and use it to build a client
//func (c *Client) WithProvider(provider Provider) (*Client, error) {
//	cred, err := provider.GetCredential()
//	if err != nil {
//		return nil, err
//	}
//	return c.WithCredential(cred), nil
//}

func (c *Client) withRegionBreaker() *Client {
	if c.profile.BreakerInterval <= 0 {
		c.profile.BreakerInterval = 5
	}
	if c.profile.BreakerInterval <= 0 {
		c.profile.BreakerInterval = 60 * time.Second
	}
	if c.profile.BreakerTimeout <= 0 {
		c.profile.BreakerTimeout = 60 * time.Second
	}
	set := gobreaker.Settings {
		MaxRequests: uint32(c.profile.BreakerMaxRequest),
		Interval: c.profile.BreakerInterval,
		Timeout: c.profile.BreakerTimeout,
		IsSuccessful: func(err error) bool {
			isSuccess := false
			// Success is considered only when the server returns an effective response (have requestId and the code is not InternalError )
			if e, ok := err.(*tcerr.TopStackSDKError); ok {
				if e.GetRequestId() != "" && e.GetCode() != "InternalError" {
					isSuccess = true
				}
			}
			return isSuccess
		},
	}
	c.cb = gobreaker.NewCircuitBreaker(set)
	return c
}

func NewClient(credential CredentialIface, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{
	}
	client.WithCredential(credential).WithProfile(clientProfile).Init()
	return
}

//// NewClientWithProviders build client with your custom providers;
//// If you don't specify the providers, it will use the DefaultProviderChain to find credential
//func NewClientWithProviders(region string, providers ...Provider) (client *Client, err error) {
//	client = (&Client{}).Init(region)
//	var pc Provider
//	if len(providers) == 0 {
//		pc = DefaultProviderChain()
//	} else {
//		pc = NewProviderChain(providers)
//	}
//	return client.WithProvider(pc)
//}
//

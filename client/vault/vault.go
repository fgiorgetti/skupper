package vault

import (
	"bytes"
	"context"
	jsonencoding "encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"
)

type Client struct {
	token   string
	address string
	siteId  string
}

func NewClient(address, token, siteId string) *Client {
	return &Client{
		address: address,
		token:   token,
		siteId:  siteId,
	}
}

// secretList type used to retrieve list of tokens available for a given site
type secretList struct {
	Data dataKeys `json:"data"`
}
type dataKeys struct {
	Keys []string `json:"keys"`
}

// secretPublish is used to push a token to a destination site
type secretPublish struct {
	Data dataSecret `json:"data"`
}
type dataSecret struct {
	Secret string `json:"secret"`
}

// secretToken is used to retrieve a particular token
type secretToken struct {
	Data dataData `json:"data"`
}
type dataData struct {
	Data dataSecret `json:"data"`
}

func request[T any](ctx context.Context, c *Client, method, url, data string) (T, string, *http.Response, error) {
	var retType T
	cli := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, url, strings.NewReader(data))
	if err != nil {
		return retType, "", nil, fmt.Errorf("error preparing http request - %v", err)
	}
	req.Header.Add("X-Vault-Token", c.token)
	resp, err := cli.Do(req)
	if err != nil {
		return retType, "", nil, fmt.Errorf("error retrieving available tokens - %v", err)
	}
	defer resp.Body.Close()
	bodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		return retType, "", nil, fmt.Errorf("error reading response data - %v", err)
	}

	err = jsonencoding.Unmarshal(bodyData, &retType)
	bodyDataStr := string(bodyData)
	if err != nil {
		return retType, bodyDataStr, resp, fmt.Errorf("error unmarshalling json response - %v", err)
	}

	if resp.StatusCode != 200 {
		return retType, bodyDataStr, resp, fmt.Errorf("error posting to %s - code: %d - response: %s", url, resp.StatusCode, bodyDataStr)
	}

	return retType, bodyDataStr, resp, nil
}

func (c *Client) PublishToken(ctx context.Context, targetSiteId string, secret *corev1.Secret) error {
	var b bytes.Buffer
	s := json.NewSerializerWithOptions(json.DefaultMetaFactory, scheme.Scheme, scheme.Scheme, json.SerializerOptions{true, false, false})
	err := s.Encode(secret, &b)
	if err != nil {
		return fmt.Errorf("error encoding secret - %v", err)
	}
	url := fmt.Sprintf(`%s/v1/secret/data/skupper/token/%s/%s`, c.address, targetSiteId, c.siteId)
	var dataSecret secretPublish
	dataSecret.Data.Secret = b.String()
	dataJson, err := jsonencoding.MarshalIndent(dataSecret, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshalling secret data - %v", err)
	}
	_, respBody, resp, err := request[interface{}](ctx, c, "POST", url, string(dataJson))
	if err != nil || resp.StatusCode != 200 {
		fmt.Printf("error creating secret - %v - code: %d - response: %s", err, resp.StatusCode, respBody)
	}
	return nil
}

func (c *Client) retrieveToken(ctx context.Context, destSiteId string) (*corev1.Secret, error) {
	var secret secretToken
	url := fmt.Sprintf(`%s/v1/secret/data/skupper/token/%s/%s`, c.address, c.siteId, destSiteId)
	secret, respBody, resp, err := request[secretToken](ctx, c, "GET", url, "")
	if err != nil {
		return nil, fmt.Errorf("error retrieving token info - code: %d - error: %v - response: %s", resp.StatusCode, err, respBody)
	}
	s := json.NewSerializerWithOptions(json.DefaultMetaFactory, scheme.Scheme, scheme.Scheme, json.SerializerOptions{true, false, false})
	var k8sSecret corev1.Secret
	_, _, err = s.Decode([]byte(secret.Data.Data.Secret), nil, &k8sSecret)
	return &k8sSecret, err
}

func (c *Client) RetrieveTokens(ctx context.Context) ([]*corev1.Secret, error) {
	var siteList secretList
	url := fmt.Sprintf(`%s/v1/secret/metadata/skupper/token/%s?list=true`, c.address, c.siteId)
	siteList, body, resp, err := request[secretList](ctx, c, "GET", url, "")
	if err != nil {
		return nil, fmt.Errorf("error retrieving tokens list - code: %d - resp: %s - error: %v", resp.StatusCode, body, err)
	}
	destSiteIds := siteList.Data.Keys
	var tokens []*corev1.Secret
	for _, destSiteId := range destSiteIds {
		token, err := c.retrieveToken(ctx, destSiteId)
		if err != nil {
			return nil, fmt.Errorf("error retrieving token - %v", err)
		}
		token.Namespace = ""
		token.ResourceVersion = ""
		tokens = append(tokens, token)
	}
	return tokens, nil
}

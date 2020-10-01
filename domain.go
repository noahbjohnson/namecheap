package namecheap

import (
	"fmt"
	"strconv"
)

// GetDomains retrieves all the domains avaialble on account.
func (c *Client) GetDomains() ([]Domain, error) {
	var domainsResponse DomainsResponse
	params := map[string]string{
		"Command": "namecheap.domains.getList",
	}
	req, err := c.NewRequest(params)
	if err != nil {
		return nil, err
	}
	resp, err := c.Http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = c.decode(resp.Body, &domainsResponse)
	if err != nil {
		return nil, err
	}
	if len(domainsResponse.Errors) > 0 {
		apiErr := domainsResponse.Errors[0]
		return nil, fmt.Errorf("%s (%s)", apiErr.Message, apiErr.Number)
	}
	return domainsResponse.CommandResponse.Domains, nil
}

// GetDomainsWithOptions retrieves paginated domains available on account.
func (c *Client) GetDomainsWithOptions(pageSize, page int) ([]Domain, error) {
	var domainsResponse DomainsResponse
	pageSizeString := strconv.Itoa(pageSize)
	pageString := strconv.Itoa(page)
	params := map[string]string{
		"Command":  "namecheap.domains.getList",
		"PageSize": pageSizeString,
		"Page":     pageString,
	}
	req, err := c.NewRequest(params)
	if err != nil {
		return nil, err
	}
	resp, err := c.Http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = c.decode(resp.Body, &domainsResponse)
	if err != nil {
		return nil, err
	}
	if len(domainsResponse.Errors) > 0 {
		apiErr := domainsResponse.Errors[0]
		return nil, fmt.Errorf("%s (%s)", apiErr.Message, apiErr.Number)
	}
	return domainsResponse.CommandResponse.Domains, nil
}

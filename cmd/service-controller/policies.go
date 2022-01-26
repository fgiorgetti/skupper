package main

import (
	"fmt"
	"net/http"
	"text/tabwriter"

	"github.com/gorilla/mux"
	"github.com/skupperproject/skupper/client"
)

type PolicyManager struct {
	validator *client.ClusterPolicyValidator
}

func newPolicyManager(cli *client.VanClient) *PolicyManager {
	p := &PolicyManager{}
	p.validator = client.NewClusterPolicyValidator(cli)
	return p
}

func fromPolicyValidationResult(res *client.PolicyValidationResult) client.PolicyAPIResult {
	err := ""
	if res.Error() != nil {
		err = res.Error().Error()
	}
	return client.PolicyAPIResult{
		Allowed:   res.Allowed(),
		AllowedBy: res.AllowPolicyNames(),
		Enabled:   res.Enabled(),
		Error:     err,
	}
}

func (p *PolicyManager) response(pr client.PolicyAPIResult, w http.ResponseWriter) {
	tw := tabwriter.NewWriter(w, 0, 4, 1, ' ', 0)
	_, _ = fmt.Fprintln(tw, fmt.Sprintf("%s\t%s\t%s\t%s\t", "ALLOWED", "POLICY ENABLED", "ERROR", "ALLOWED BY"))
	_, _ = fmt.Fprintln(tw, fmt.Sprintf("%v\t%v\t%s\t%s\t", pr.Allowed, pr.Enabled, pr.Error, ""))
	for _, policy := range pr.AllowedBy {
		_, _ = fmt.Fprintln(tw, fmt.Sprintf("%s\t%s\t%s\t%s\t", "", "", "", policy))
	}
	_ = tw.Flush()
}

func (p *PolicyManager) gateway() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			res := p.validator.ValidateCreateGateway()
			pr := fromPolicyValidationResult(res)
			if wantsJsonOutput(r) {
				writeJson(pr, w)
			} else {
				p.response(pr, w)
			}
		}
	})
}

func (p *PolicyManager) expose() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if r.Method == http.MethodGet {
			resourceType, okType := vars["resourceType"]
			resourceName, okName := vars["resourceName"]
			if !okType || !okName {
				http.Error(w, "Invalid parameters", http.StatusInternalServerError)
				return
			}
			res := p.validator.ValidateExpose(resourceType, resourceName)
			pr := fromPolicyValidationResult(res)
			if wantsJsonOutput(r) {
				writeJson(pr, w)
			} else {
				p.response(pr, w)
			}
		}
	})
}

func (p *PolicyManager) service() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if r.Method == http.MethodGet {
			name, ok := vars["name"]
			if !ok {
				http.Error(w, "Invalid parameters", http.StatusInternalServerError)
				return
			}
			res := p.validator.ValidateImportService(name)
			pr := fromPolicyValidationResult(res)
			if wantsJsonOutput(r) {
				writeJson(pr, w)
			} else {
				p.response(pr, w)
			}
		}
	})
}

func (p *PolicyManager) incomingLink() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			res := p.validator.ValidateIncomingLink()
			pr := fromPolicyValidationResult(res)
			if wantsJsonOutput(r) {
				writeJson(pr, w)
			} else {
				p.response(pr, w)
			}
		}
	})
}

func (p *PolicyManager) outgoingLink() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if r.Method == http.MethodGet {
			hostname, ok := vars["hostname"]
			if !ok {
				http.Error(w, "Invalid parameters", http.StatusInternalServerError)
				return
			}
			res := p.validator.ValidateOutgoingLink(hostname)
			pr := fromPolicyValidationResult(res)
			if wantsJsonOutput(r) {
				writeJson(pr, w)
			} else {
				p.response(pr, w)
			}
		}
	})
}
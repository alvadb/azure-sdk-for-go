package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.12.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// ExpressRouteCircuitPeeringsResourceProviderClient is the the Windows Azure
// Network management API provides a RESTful set of web services that
// interact with Windows Azure Networks service to manage your network
// resrources. The API has entities that capture the relationship between an
// end user and the Windows Azure Networks service.
type ExpressRouteCircuitPeeringsResourceProviderClient struct {
	ResourceProviderClient
}

// NewExpressRouteCircuitPeeringsResourceProviderClient creates an instance of
// the ExpressRouteCircuitPeeringsResourceProviderClient client.
func NewExpressRouteCircuitPeeringsResourceProviderClient(subscriptionID string) ExpressRouteCircuitPeeringsResourceProviderClient {
	return NewExpressRouteCircuitPeeringsResourceProviderClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExpressRouteCircuitPeeringsResourceProviderClientWithBaseURI creates an
// instance of the ExpressRouteCircuitPeeringsResourceProviderClient client.
func NewExpressRouteCircuitPeeringsResourceProviderClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteCircuitPeeringsResourceProviderClient {
	return ExpressRouteCircuitPeeringsResourceProviderClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put Pering operation creates/updates an peering in the
// specified ExpressRouteCircuits
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the express route circuit. peeringName is the name of the peering.
// peeringParameters is parameters supplied to the create/update
// ExpressRouteCircuit Peering operation
func (client ExpressRouteCircuitPeeringsResourceProviderClient) CreateOrUpdate(resourceGroupName string, circuitName string, peeringName string, peeringParameters ExpressRouteCircuitPeering) (result ExpressRouteCircuitPeering, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, circuitName, peeringName, peeringParameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitPeeringsResourceProviderClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitPeeringsResourceProviderClient", "CreateOrUpdate", "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitPeeringsResourceProviderClient", "CreateOrUpdate", "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ExpressRouteCircuitPeeringsResourceProviderClient) CreateOrUpdatePreparer(resourceGroupName string, circuitName string, peeringName string, peeringParameters ExpressRouteCircuitPeering) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       url.QueryEscape(circuitName),
		"peeringName":       url.QueryEscape(peeringName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}"),
		autorest.WithJSON(peeringParameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitPeeringsResourceProviderClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusCreated, http.StatusOK)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitPeeringsResourceProviderClient) CreateOrUpdateResponder(resp *http.Response) (result ExpressRouteCircuitPeering, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the delete peering operation deletes the specified peering from the
// ExpressRouteCircuit.
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the express route circuit. peeringName is the name of the peering.
func (client ExpressRouteCircuitPeeringsResourceProviderClient) Delete(resourceGroupName string, circuitName string, peeringName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, circuitName, peeringName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitPeeringsResourceProviderClient", "Delete", "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitPeeringsResourceProviderClient", "Delete", "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitPeeringsResourceProviderClient", "Delete", "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExpressRouteCircuitPeeringsResourceProviderClient) DeletePreparer(resourceGroupName string, circuitName string, peeringName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       url.QueryEscape(circuitName),
		"peeringName":       url.QueryEscape(peeringName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitPeeringsResourceProviderClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusNoContent, http.StatusAccepted, http.StatusOK)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitPeeringsResourceProviderClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusAccepted, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the GET peering operation retrieves the specified authorization from
// the ExpressRouteCircuit.
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the express route circuit. peeringName is the name of the peering.
func (client ExpressRouteCircuitPeeringsResourceProviderClient) Get(resourceGroupName string, circuitName string, peeringName string) (result ExpressRouteCircuitPeering, ae error) {
	req, err := client.GetPreparer(resourceGroupName, circuitName, peeringName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitPeeringsResourceProviderClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitPeeringsResourceProviderClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitPeeringsResourceProviderClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExpressRouteCircuitPeeringsResourceProviderClient) GetPreparer(resourceGroupName string, circuitName string, peeringName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       url.QueryEscape(circuitName),
		"peeringName":       url.QueryEscape(peeringName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitPeeringsResourceProviderClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitPeeringsResourceProviderClient) GetResponder(resp *http.Response) (result ExpressRouteCircuitPeering, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List peering operation retrieves all the peerings in an
// ExpressRouteCircuit.
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the curcuit.
func (client ExpressRouteCircuitPeeringsResourceProviderClient) List(resourceGroupName string, circuitName string) (result ExpressRouteCircuitPeeringListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName, circuitName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitPeeringsResourceProviderClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitPeeringsResourceProviderClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitPeeringsResourceProviderClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ExpressRouteCircuitPeeringsResourceProviderClient) ListPreparer(resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       url.QueryEscape(circuitName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitPeeringsResourceProviderClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitPeeringsResourceProviderClient) ListResponder(resp *http.Response) (result ExpressRouteCircuitPeeringListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client ExpressRouteCircuitPeeringsResourceProviderClient) ListNextResults(lastResults ExpressRouteCircuitPeeringListResult) (result ExpressRouteCircuitPeeringListResult, ae error) {
	req, err := lastResults.ExpressRouteCircuitPeeringListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitPeeringsResourceProviderClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitPeeringsResourceProviderClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitPeeringsResourceProviderClient", "List", "Failure responding to next results request request")
	}

	return
}
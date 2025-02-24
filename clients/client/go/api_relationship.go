/*
Ory APIs

# Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | -------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------ | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 

API version: v1.17.1
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


type RelationshipAPI interface {

	/*
	CheckOplSyntax Check the syntax of an OPL file

	The OPL file is expected in the body of the request.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RelationshipAPICheckOplSyntaxRequest
	*/
	CheckOplSyntax(ctx context.Context) RelationshipAPICheckOplSyntaxRequest

	// CheckOplSyntaxExecute executes the request
	//  @return CheckOplSyntaxResult
	CheckOplSyntaxExecute(r RelationshipAPICheckOplSyntaxRequest) (*CheckOplSyntaxResult, *http.Response, error)

	/*
	CreateRelationship Create a Relationship

	Use this endpoint to create a relationship.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RelationshipAPICreateRelationshipRequest
	*/
	CreateRelationship(ctx context.Context) RelationshipAPICreateRelationshipRequest

	// CreateRelationshipExecute executes the request
	//  @return Relationship
	CreateRelationshipExecute(r RelationshipAPICreateRelationshipRequest) (*Relationship, *http.Response, error)

	/*
	DeleteRelationships Delete Relationships

	Use this endpoint to delete relationships

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RelationshipAPIDeleteRelationshipsRequest
	*/
	DeleteRelationships(ctx context.Context) RelationshipAPIDeleteRelationshipsRequest

	// DeleteRelationshipsExecute executes the request
	DeleteRelationshipsExecute(r RelationshipAPIDeleteRelationshipsRequest) (*http.Response, error)

	/*
	GetRelationships Query relationships

	Get all relationships that match the query. Only the namespace field is required.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RelationshipAPIGetRelationshipsRequest
	*/
	GetRelationships(ctx context.Context) RelationshipAPIGetRelationshipsRequest

	// GetRelationshipsExecute executes the request
	//  @return Relationships
	GetRelationshipsExecute(r RelationshipAPIGetRelationshipsRequest) (*Relationships, *http.Response, error)

	/*
	ListRelationshipNamespaces Query namespaces

	Get all namespaces

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RelationshipAPIListRelationshipNamespacesRequest
	*/
	ListRelationshipNamespaces(ctx context.Context) RelationshipAPIListRelationshipNamespacesRequest

	// ListRelationshipNamespacesExecute executes the request
	//  @return RelationshipNamespaces
	ListRelationshipNamespacesExecute(r RelationshipAPIListRelationshipNamespacesRequest) (*RelationshipNamespaces, *http.Response, error)

	/*
	PatchRelationships Patch Multiple Relationships

	Use this endpoint to patch one or more relationships.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RelationshipAPIPatchRelationshipsRequest
	*/
	PatchRelationships(ctx context.Context) RelationshipAPIPatchRelationshipsRequest

	// PatchRelationshipsExecute executes the request
	PatchRelationshipsExecute(r RelationshipAPIPatchRelationshipsRequest) (*http.Response, error)
}

// RelationshipAPIService RelationshipAPI service
type RelationshipAPIService service

type RelationshipAPICheckOplSyntaxRequest struct {
	ctx context.Context
	ApiService RelationshipAPI
	body *string
}

func (r RelationshipAPICheckOplSyntaxRequest) Body(body string) RelationshipAPICheckOplSyntaxRequest {
	r.body = &body
	return r
}

func (r RelationshipAPICheckOplSyntaxRequest) Execute() (*CheckOplSyntaxResult, *http.Response, error) {
	return r.ApiService.CheckOplSyntaxExecute(r)
}

/*
CheckOplSyntax Check the syntax of an OPL file

The OPL file is expected in the body of the request.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RelationshipAPICheckOplSyntaxRequest
*/
func (a *RelationshipAPIService) CheckOplSyntax(ctx context.Context) RelationshipAPICheckOplSyntaxRequest {
	return RelationshipAPICheckOplSyntaxRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CheckOplSyntaxResult
func (a *RelationshipAPIService) CheckOplSyntaxExecute(r RelationshipAPICheckOplSyntaxRequest) (*CheckOplSyntaxResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CheckOplSyntaxResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipAPIService.CheckOplSyntax")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/opl/syntax/check"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RelationshipAPICreateRelationshipRequest struct {
	ctx context.Context
	ApiService RelationshipAPI
	createRelationshipBody *CreateRelationshipBody
}

func (r RelationshipAPICreateRelationshipRequest) CreateRelationshipBody(createRelationshipBody CreateRelationshipBody) RelationshipAPICreateRelationshipRequest {
	r.createRelationshipBody = &createRelationshipBody
	return r
}

func (r RelationshipAPICreateRelationshipRequest) Execute() (*Relationship, *http.Response, error) {
	return r.ApiService.CreateRelationshipExecute(r)
}

/*
CreateRelationship Create a Relationship

Use this endpoint to create a relationship.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RelationshipAPICreateRelationshipRequest
*/
func (a *RelationshipAPIService) CreateRelationship(ctx context.Context) RelationshipAPICreateRelationshipRequest {
	return RelationshipAPICreateRelationshipRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Relationship
func (a *RelationshipAPIService) CreateRelationshipExecute(r RelationshipAPICreateRelationshipRequest) (*Relationship, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Relationship
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipAPIService.CreateRelationship")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/relation-tuples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createRelationshipBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RelationshipAPIDeleteRelationshipsRequest struct {
	ctx context.Context
	ApiService RelationshipAPI
	namespace *string
	object *string
	relation *string
	subjectId *string
	subjectSetNamespace *string
	subjectSetObject *string
	subjectSetRelation *string
}

// Namespace of the Relationship
func (r RelationshipAPIDeleteRelationshipsRequest) Namespace(namespace string) RelationshipAPIDeleteRelationshipsRequest {
	r.namespace = &namespace
	return r
}

// Object of the Relationship
func (r RelationshipAPIDeleteRelationshipsRequest) Object(object string) RelationshipAPIDeleteRelationshipsRequest {
	r.object = &object
	return r
}

// Relation of the Relationship
func (r RelationshipAPIDeleteRelationshipsRequest) Relation(relation string) RelationshipAPIDeleteRelationshipsRequest {
	r.relation = &relation
	return r
}

// SubjectID of the Relationship
func (r RelationshipAPIDeleteRelationshipsRequest) SubjectId(subjectId string) RelationshipAPIDeleteRelationshipsRequest {
	r.subjectId = &subjectId
	return r
}

// Namespace of the Subject Set
func (r RelationshipAPIDeleteRelationshipsRequest) SubjectSetNamespace(subjectSetNamespace string) RelationshipAPIDeleteRelationshipsRequest {
	r.subjectSetNamespace = &subjectSetNamespace
	return r
}

// Object of the Subject Set
func (r RelationshipAPIDeleteRelationshipsRequest) SubjectSetObject(subjectSetObject string) RelationshipAPIDeleteRelationshipsRequest {
	r.subjectSetObject = &subjectSetObject
	return r
}

// Relation of the Subject Set
func (r RelationshipAPIDeleteRelationshipsRequest) SubjectSetRelation(subjectSetRelation string) RelationshipAPIDeleteRelationshipsRequest {
	r.subjectSetRelation = &subjectSetRelation
	return r
}

func (r RelationshipAPIDeleteRelationshipsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteRelationshipsExecute(r)
}

/*
DeleteRelationships Delete Relationships

Use this endpoint to delete relationships

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RelationshipAPIDeleteRelationshipsRequest
*/
func (a *RelationshipAPIService) DeleteRelationships(ctx context.Context) RelationshipAPIDeleteRelationshipsRequest {
	return RelationshipAPIDeleteRelationshipsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RelationshipAPIService) DeleteRelationshipsExecute(r RelationshipAPIDeleteRelationshipsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipAPIService.DeleteRelationships")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/relation-tuples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.namespace != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespace", r.namespace, "")
	}
	if r.object != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "object", r.object, "")
	}
	if r.relation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "relation", r.relation, "")
	}
	if r.subjectId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subject_id", r.subjectId, "")
	}
	if r.subjectSetNamespace != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subject_set.namespace", r.subjectSetNamespace, "")
	}
	if r.subjectSetObject != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subject_set.object", r.subjectSetObject, "")
	}
	if r.subjectSetRelation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subject_set.relation", r.subjectSetRelation, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type RelationshipAPIGetRelationshipsRequest struct {
	ctx context.Context
	ApiService RelationshipAPI
	pageToken *string
	pageSize *int64
	namespace *string
	object *string
	relation *string
	subjectId *string
	subjectSetNamespace *string
	subjectSetObject *string
	subjectSetRelation *string
}

func (r RelationshipAPIGetRelationshipsRequest) PageToken(pageToken string) RelationshipAPIGetRelationshipsRequest {
	r.pageToken = &pageToken
	return r
}

func (r RelationshipAPIGetRelationshipsRequest) PageSize(pageSize int64) RelationshipAPIGetRelationshipsRequest {
	r.pageSize = &pageSize
	return r
}

// Namespace of the Relationship
func (r RelationshipAPIGetRelationshipsRequest) Namespace(namespace string) RelationshipAPIGetRelationshipsRequest {
	r.namespace = &namespace
	return r
}

// Object of the Relationship
func (r RelationshipAPIGetRelationshipsRequest) Object(object string) RelationshipAPIGetRelationshipsRequest {
	r.object = &object
	return r
}

// Relation of the Relationship
func (r RelationshipAPIGetRelationshipsRequest) Relation(relation string) RelationshipAPIGetRelationshipsRequest {
	r.relation = &relation
	return r
}

// SubjectID of the Relationship
func (r RelationshipAPIGetRelationshipsRequest) SubjectId(subjectId string) RelationshipAPIGetRelationshipsRequest {
	r.subjectId = &subjectId
	return r
}

// Namespace of the Subject Set
func (r RelationshipAPIGetRelationshipsRequest) SubjectSetNamespace(subjectSetNamespace string) RelationshipAPIGetRelationshipsRequest {
	r.subjectSetNamespace = &subjectSetNamespace
	return r
}

// Object of the Subject Set
func (r RelationshipAPIGetRelationshipsRequest) SubjectSetObject(subjectSetObject string) RelationshipAPIGetRelationshipsRequest {
	r.subjectSetObject = &subjectSetObject
	return r
}

// Relation of the Subject Set
func (r RelationshipAPIGetRelationshipsRequest) SubjectSetRelation(subjectSetRelation string) RelationshipAPIGetRelationshipsRequest {
	r.subjectSetRelation = &subjectSetRelation
	return r
}

func (r RelationshipAPIGetRelationshipsRequest) Execute() (*Relationships, *http.Response, error) {
	return r.ApiService.GetRelationshipsExecute(r)
}

/*
GetRelationships Query relationships

Get all relationships that match the query. Only the namespace field is required.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RelationshipAPIGetRelationshipsRequest
*/
func (a *RelationshipAPIService) GetRelationships(ctx context.Context) RelationshipAPIGetRelationshipsRequest {
	return RelationshipAPIGetRelationshipsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Relationships
func (a *RelationshipAPIService) GetRelationshipsExecute(r RelationshipAPIGetRelationshipsRequest) (*Relationships, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Relationships
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipAPIService.GetRelationships")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_token", r.pageToken, "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "")
	}
	if r.namespace != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespace", r.namespace, "")
	}
	if r.object != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "object", r.object, "")
	}
	if r.relation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "relation", r.relation, "")
	}
	if r.subjectId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subject_id", r.subjectId, "")
	}
	if r.subjectSetNamespace != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subject_set.namespace", r.subjectSetNamespace, "")
	}
	if r.subjectSetObject != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subject_set.object", r.subjectSetObject, "")
	}
	if r.subjectSetRelation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subject_set.relation", r.subjectSetRelation, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RelationshipAPIListRelationshipNamespacesRequest struct {
	ctx context.Context
	ApiService RelationshipAPI
}

func (r RelationshipAPIListRelationshipNamespacesRequest) Execute() (*RelationshipNamespaces, *http.Response, error) {
	return r.ApiService.ListRelationshipNamespacesExecute(r)
}

/*
ListRelationshipNamespaces Query namespaces

Get all namespaces

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RelationshipAPIListRelationshipNamespacesRequest
*/
func (a *RelationshipAPIService) ListRelationshipNamespaces(ctx context.Context) RelationshipAPIListRelationshipNamespacesRequest {
	return RelationshipAPIListRelationshipNamespacesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RelationshipNamespaces
func (a *RelationshipAPIService) ListRelationshipNamespacesExecute(r RelationshipAPIListRelationshipNamespacesRequest) (*RelationshipNamespaces, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RelationshipNamespaces
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipAPIService.ListRelationshipNamespaces")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/namespaces"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RelationshipAPIPatchRelationshipsRequest struct {
	ctx context.Context
	ApiService RelationshipAPI
	relationshipPatch *[]RelationshipPatch
}

func (r RelationshipAPIPatchRelationshipsRequest) RelationshipPatch(relationshipPatch []RelationshipPatch) RelationshipAPIPatchRelationshipsRequest {
	r.relationshipPatch = &relationshipPatch
	return r
}

func (r RelationshipAPIPatchRelationshipsRequest) Execute() (*http.Response, error) {
	return r.ApiService.PatchRelationshipsExecute(r)
}

/*
PatchRelationships Patch Multiple Relationships

Use this endpoint to patch one or more relationships.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RelationshipAPIPatchRelationshipsRequest
*/
func (a *RelationshipAPIService) PatchRelationships(ctx context.Context) RelationshipAPIPatchRelationshipsRequest {
	return RelationshipAPIPatchRelationshipsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RelationshipAPIService) PatchRelationshipsExecute(r RelationshipAPIPatchRelationshipsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipAPIService.PatchRelationships")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/relation-tuples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.relationshipPatch
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

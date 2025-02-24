/*
 * Ory APIs
 * # Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | -------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------ | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 
 *
 * The version of the OpenAPI document: v1.16.10
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.api;

import sh.ory.ApiException;
import sh.ory.model.CreateProjectApiKeyRequest;
import sh.ory.model.CreateProjectBody;
import sh.ory.model.ErrorGeneric;
import sh.ory.model.GenericError;
import sh.ory.model.GetOrganizationResponse;
import sh.ory.model.JsonPatch;
import sh.ory.model.ListOrganizationsResponse;
import sh.ory.model.Organization;
import sh.ory.model.OrganizationBody;
import sh.ory.model.Project;
import sh.ory.model.ProjectApiKey;
import sh.ory.model.ProjectMember;
import sh.ory.model.ProjectMetadata;
import sh.ory.model.SetProject;
import sh.ory.model.SuccessfulProjectUpdate;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for ProjectApi
 */
@Disabled
public class ProjectApiTest {

    private final ProjectApi api = new ProjectApi();

    /**
     * Create an Enterprise SSO Organization
     *
     * Creates an Enterprise SSO Organization in a project.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void createOrganizationTest() throws ApiException {
        String projectId = null;
        OrganizationBody organizationBody = null;
        Organization response = api.createOrganization(projectId, organizationBody);
        // TODO: test validations
    }

    /**
     * Create a Project
     *
     * Creates a new project.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void createProjectTest() throws ApiException {
        CreateProjectBody createProjectBody = null;
        Project response = api.createProject(createProjectBody);
        // TODO: test validations
    }

    /**
     * Create project API key
     *
     * Create an API key for a project.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void createProjectApiKeyTest() throws ApiException {
        String project = null;
        CreateProjectApiKeyRequest createProjectApiKeyRequest = null;
        ProjectApiKey response = api.createProjectApiKey(project, createProjectApiKeyRequest);
        // TODO: test validations
    }

    /**
     * Delete Enterprise SSO Organization
     *
     * Irrecoverably deletes an Enterprise SSO Organization in a project by its ID.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deleteOrganizationTest() throws ApiException {
        String projectId = null;
        String organizationId = null;
        api.deleteOrganization(projectId, organizationId);
        // TODO: test validations
    }

    /**
     * Delete project API key
     *
     * Deletes an API key and immediately removes it.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deleteProjectApiKeyTest() throws ApiException {
        String project = null;
        String tokenId = null;
        api.deleteProjectApiKey(project, tokenId);
        // TODO: test validations
    }

    /**
     * Get Enterprise SSO Organization by ID
     *
     * Retrieves an Enterprise SSO Organization for a project by its ID
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getOrganizationTest() throws ApiException {
        String projectId = null;
        String organizationId = null;
        GetOrganizationResponse response = api.getOrganization(projectId, organizationId);
        // TODO: test validations
    }

    /**
     * Get a Project
     *
     * Get a projects you have access to by its ID.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getProjectTest() throws ApiException {
        String projectId = null;
        Project response = api.getProject(projectId);
        // TODO: test validations
    }

    /**
     * Get all members associated with this project
     *
     * This endpoint requires the user to be a member of the project with the role &#x60;OWNER&#x60; or &#x60;DEVELOPER&#x60;.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getProjectMembersTest() throws ApiException {
        String project = null;
        List<ProjectMember> response = api.getProjectMembers(project);
        // TODO: test validations
    }

    /**
     * List all Enterprise SSO organizations
     *
     * Lists all Enterprise SSO organizations in a project.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listOrganizationsTest() throws ApiException {
        String projectId = null;
        Long pageSize = null;
        String pageToken = null;
        String domain = null;
        ListOrganizationsResponse response = api.listOrganizations(projectId, pageSize, pageToken, domain);
        // TODO: test validations
    }

    /**
     * List a project&#39;s API keys
     *
     * A list of all the project&#39;s API keys.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listProjectApiKeysTest() throws ApiException {
        String project = null;
        List<ProjectApiKey> response = api.listProjectApiKeys(project);
        // TODO: test validations
    }

    /**
     * List All Projects
     *
     * Lists all projects you have access to.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listProjectsTest() throws ApiException {
        List<ProjectMetadata> response = api.listProjects();
        // TODO: test validations
    }

    /**
     * Patch an Ory Network Project Configuration
     *
     * Deprecated: Use the &#x60;patchProjectWithRevision&#x60; endpoint instead to specify the exact revision the patch was generated for.  This endpoints allows you to patch individual Ory Network project configuration keys for Ory&#39;s services (identity, permission, ...). The configuration format is fully compatible with the open source projects for the respective services (e.g. Ory Kratos for Identity, Ory Keto for Permissions).  This endpoint expects the &#x60;version&#x60; key to be set in the payload. If it is unset, it will try to import the config as if it is from the most recent version.  If you have an older version of a configuration, you should set the version key in the payload!  While this endpoint is able to process all configuration items related to features (e.g. password reset), it does not support operational configuration items (e.g. port, tracing, logging) otherwise available in the open source.  For configuration items that can not be translated to the Ory Network, this endpoint will return a list of warnings to help you understand which parts of your config could not be processed.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void patchProjectTest() throws ApiException {
        String projectId = null;
        List<JsonPatch> jsonPatch = null;
        SuccessfulProjectUpdate response = api.patchProject(projectId, jsonPatch);
        // TODO: test validations
    }

    /**
     * Patch an Ory Network Project Configuration based on a revision ID
     *
     * This endpoints allows you to patch individual Ory Network Project configuration keys for Ory&#39;s services (identity, permission, ...). The configuration format is fully compatible with the open source projects for the respective services (e.g. Ory Kratos for Identity, Ory Keto for Permissions).  This endpoint expects the &#x60;version&#x60; key to be set in the payload. If it is unset, it will try to import the config as if it is from the most recent version.  If you have an older version of a configuration, you should set the version key in the payload!  While this endpoint is able to process all configuration items related to features (e.g. password reset), it does not support operational configuration items (e.g. port, tracing, logging) otherwise available in the open source.  For configuration items that can not be translated to the Ory Network, this endpoint will return a list of warnings to help you understand which parts of your config could not be processed.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void patchProjectWithRevisionTest() throws ApiException {
        String projectId = null;
        String revisionId = null;
        List<JsonPatch> jsonPatch = null;
        SuccessfulProjectUpdate response = api.patchProjectWithRevision(projectId, revisionId, jsonPatch);
        // TODO: test validations
    }

    /**
     * Irrecoverably purge a project
     *
     * !! Use with extreme caution !!  Using this API endpoint you can purge (completely delete) a project and its data. This action can not be undone and will delete ALL your data.  Calling this endpoint will additionally delete custom domains and other related data.  If the project is linked to a subscription, the subscription needs to be unlinked first.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void purgeProjectTest() throws ApiException {
        String projectId = null;
        api.purgeProject(projectId);
        // TODO: test validations
    }

    /**
     * Remove a member associated with this project
     *
     * This also sets their invite status to &#x60;REMOVED&#x60;. This endpoint requires the user to be a member of the project with the role &#x60;OWNER&#x60;.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void removeProjectMemberTest() throws ApiException {
        String project = null;
        String member = null;
        api.removeProjectMember(project, member);
        // TODO: test validations
    }

    /**
     * Update an Ory Network Project Configuration
     *
     * This endpoints allows you to update the Ory Network project configuration for individual services (identity, permission, ...). The configuration is fully compatible with the open source projects for the respective services (e.g. Ory Kratos for Identity, Ory Keto for Permissions).  This endpoint expects the &#x60;version&#x60; key to be set in the payload. If it is unset, it will try to import the config as if it is from the most recent version.  If you have an older version of a configuration, you should set the version key in the payload!  While this endpoint is able to process all configuration items related to features (e.g. password reset), it does not support operational configuration items (e.g. port, tracing, logging) otherwise available in the open source.  For configuration items that can not be translated to the Ory Network, this endpoint will return a list of warnings to help you understand which parts of your config could not be processed.  Be aware that updating any service&#39;s configuration will completely override your current configuration for that service!
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void setProjectTest() throws ApiException {
        String projectId = null;
        SetProject setProject = null;
        SuccessfulProjectUpdate response = api.setProject(projectId, setProject);
        // TODO: test validations
    }

    /**
     * Update an Enterprise SSO Organization
     *
     * Updates an Enterprise SSO Organization in a project by its ID.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updateOrganizationTest() throws ApiException {
        String projectId = null;
        String organizationId = null;
        OrganizationBody organizationBody = null;
        Organization response = api.updateOrganization(projectId, organizationId, organizationBody);
        // TODO: test validations
    }

}

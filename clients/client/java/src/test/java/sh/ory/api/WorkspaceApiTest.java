/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.13.6
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.api;

import sh.ory.ApiException;
import sh.ory.model.CreateWorkspaceBody;
import sh.ory.model.ErrorGeneric;
import sh.ory.model.ListWorkspaceProjects;
import sh.ory.model.ListWorkspaces;
import sh.ory.model.UpdateWorkspaceBody;
import sh.ory.model.Workspace;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for WorkspaceApi
 */
@Disabled
public class WorkspaceApiTest {

    private final WorkspaceApi api = new WorkspaceApi();

    /**
     * Create a new workspace
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void createWorkspaceTest() throws ApiException {
        CreateWorkspaceBody createWorkspaceBody = null;
        Workspace response = api.createWorkspace(createWorkspaceBody);
        // TODO: test validations
    }

    /**
     * Get a workspace
     *
     * Any workspace member can access this endpoint.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getWorkspaceTest() throws ApiException {
        String workspace = null;
        Workspace response = api.getWorkspace(workspace);
        // TODO: test validations
    }

    /**
     * List all projects of a workspace
     *
     * Any workspace member can access this endpoint.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listWorkspaceProjectsTest() throws ApiException {
        String workspace = null;
        ListWorkspaceProjects response = api.listWorkspaceProjects(workspace);
        // TODO: test validations
    }

    /**
     * List workspaces the user is a member of
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listWorkspacesTest() throws ApiException {
        Long pageSize = null;
        String pageToken = null;
        ListWorkspaces response = api.listWorkspaces(pageSize, pageToken);
        // TODO: test validations
    }

    /**
     * Update an workspace
     *
     * Workspace members with the role &#x60;OWNER&#x60; can access this endpoint.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updateWorkspaceTest() throws ApiException {
        String workspace = null;
        UpdateWorkspaceBody updateWorkspaceBody = null;
        Workspace response = api.updateWorkspace(workspace, updateWorkspaceBody);
        // TODO: test validations
    }

}

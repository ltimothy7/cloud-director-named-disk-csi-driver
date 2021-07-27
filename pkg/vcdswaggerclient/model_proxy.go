/*
 * VMware Cloud Director OpenAPI
 *
 * VMware Cloud Director OpenAPI is a new API that is defined using the OpenAPI standards.<br/> This ReSTful API borrows some elements of the legacy VMware Cloud Director API and establishes new patterns for use as described below. <h4>Authentication</h4> Authentication and Authorization schemes are the same as those for the legacy APIs. You can authenticate using the JWT token via the <code>Authorization</code> header or specifying a session using <code>x-vcloud-authorization</code> (The latter form is deprecated). <h4>Operation Patterns</h4> This API follows the following general guidelines to establish a consistent CRUD pattern: <table> <tr>   <th>Operation</th><th>Description</th><th>Response Code</th><th>Response Content</th> </tr><tr>   <td>GET /items<td>Returns a paginated list of items<td>200<td>Response will include Navigational links to the items in the list. </tr><tr>   <td>POST /items<td>Returns newly created item<td>201<td>Content-Location header links to the newly created item </tr><tr>   <td>GET /items/urn<td>Returns an individual item<td>200<td>A single item using same data type as that included in list above </tr><tr>   <td>PUT /items/urn<td>Updates an individual item<td>200<td>Updated view of the item is returned </tr><tr>   <td>DELETE /items/urn<td>Deletes the item<td>204<td>No content is returned. </tr> </table> <h5>Asynchronous operations</h5> Asynchronous operations are determined by the server. In those cases, instead of responding as described above, the server responds with an HTTP Response code 202 and an empty body. The tracking task (which is the same task as all legacy API operations use) is linked via the URI provided in the <code>Location</code> header.<br/> All API calls can choose to service a request asynchronously or synchronously as determined by the server upon interpreting the request. Operations that choose to exhibit this dual behavior will have both options documented by specifying both response code(s) below. The caller must be prepared to handle responses to such API calls by inspecting the HTTP Response code. <h5>Error Conditions</h5> <b>All</b> operations report errors using the following error reporting rules: <ul>   <li>400: Bad Request - In event of bad request due to incorrect data or other user error</li>   <li>401: Bad Request - If user is unauthenticated or their session has expired</li>   <li>403: Forbidden - If the user is not authorized or the entity does not exist</li> </ul> <h4>OpenAPI Design Concepts and Principles</h4> <ul>   <li>IDs are full Uniform Resource Names (URNs).</li>   <li>OpenAPI's <code>Content-Type</code> is always <code>application/json</code></li>   <li>REST links are in the Link header.</li>   <ul>     <li>Multiple relationships for any link are represented by multiple values in a space-separated list.</li>     <li>Links have a custom VMware Cloud Director-specific &quot;model&quot; attribute that hints at the applicable data         type for the links.</li>     <li>title + rel + model attributes evaluates to a unique link.</li>     <li>Links follow Hypermedia as the Engine of Application State (HATEOAS) principles. Links are present if         certain operations are present and permitted for the user&quot;s current role and the state of the         referred entities.</li>   </ul>   <li>APIs follow a flat structure relying on cross-referencing other entities instead of the navigational style       used by the legacy VMware Cloud Director APIs.</li>   <li>Most endpoints that return a list support filtering and sorting similar to the query service in the legacy       VMware Cloud Director APIs.</li>   <li>Accept header must be included to specify the API version for the request similar to calls to existing legacy       VMware Cloud Director APIs.</li>   <li>Each feature has a version in the path element present in its URL.<br/>       <b>Note</b> API URL's without a version in their paths must be considered experimental.</li> </ul> 
 *
 * API version: 35.0
 * Contact: https://code.vmware.com/support
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// proxy configuration. This configures the vCD proxying capability for one specific server within the vSphere/NSX/ESX estate known to vCD. 
type Proxy struct {
	Name string `json:"name"`
	Id string `json:"id,omitempty"`
	// The EntityReference of the parent entity (e.g. vCenter Server, SDDC) . This is not editable once the proxy has been created.
	Parent *EntityReference `json:"parent"`
	ProxyType string `json:"proxyType"`
	// True if the proxy is enabled. Proxy can only be enabled/disabled by privileged users. A disabled proxy cannot be activated and thus, cannot be used. When a proxy is disabled, all active sessions are terminated. 
	Enabled bool `json:"enabled,omitempty"`
	// Whether this proxy has been published to tenants.
	TenantVisible bool `json:"tenantVisible,omitempty"`
	// True if the proxy is currently active for the user session associated with the request made to get the proxy. An inactive proxy cannot be used. 
	Active bool `json:"active,omitempty"`
	// The generated read-only token that should be used as the password when using this proxy. To generate a new token, activate the proxy. The token is tied to the user session that activated the proxy. If the proxy is inactive, this value will be null. 
	Token string `json:"token,omitempty"`
	// The EntityReference of the parent proxy. If a proxy has a parent, the proxy is activated along with its parent and shares the token with its parent. Each proxy may only have one parent. A parent proxy cannot have a parent of its own. 
	ParentProxy *EntityReference `json:"parentProxy,omitempty"`
	// IP address or FQDN of the host being proxied. Lower case formatting will be applied to the value of the property. This is not editable once the proxy has been created. 
	TargetHost string `json:"targetHost"`
	// The URL of the proxied component's UI endpoint. This is the URL that the browser tab  will be pointed to when the proxy is launched via the H5 UI of VCD. 
	UiUrl string `json:"uiUrl,omitempty"`
}

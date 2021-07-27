# SslSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledSslProtocols** | **[]string** | SSL protocols | [optional] [default to null]
**EnabledSslCiphers** | **[]string** | SSL ciphers | [optional] [default to null]
**KeySize** | **int32** | Size of keys generated | [optional] [default to null]
**CertificateValidityDays** | **int32** | Number of days generated certificates are valid for | [optional] [default to null]
**CertificateSignatureAlgorithm** | **string** | Algorithm used to sign generated certificates | [optional] [default to null]
**FipsMode** | **string** | The current FIPS mode of this server group. &lt;ul&gt;   &lt;li&gt;     &lt;em&gt;ON&lt;/em&gt;: Indicates FIPS mode is enabled for this server group.   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;OFF&lt;/em&gt;: Indicates FIPS mode is disabled for this server group.   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;ON_PENDING&lt;/em&gt;: FIPS mode has been requested to be toggled from OFF to ON.   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;OFF_PENDING&lt;/em&gt;: FIPS mode has been requested to be toggled from ON to OFF.   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;ON_PENDING_RESTART&lt;/em&gt;: FIPS mode has been requested to be toggled from OFF to ON,     and only requires a restart of all cells to make the toggle.   &lt;/li&gt;   &lt;li&gt;     &lt;em&gt;OFF_PENDING_RESTART&lt;/em&gt;: FIPS mode has been requested to be toggled from ON to OFF,     and only requires a restart of all cells to make the toggle.   &lt;/li&gt; &lt;/ul&gt;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



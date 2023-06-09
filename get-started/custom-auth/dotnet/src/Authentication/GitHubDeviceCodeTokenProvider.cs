// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

// <TokenProviderSnippet>
using IdentityModel.Client;
using Microsoft.Kiota.Abstractions.Authentication;

namespace RepoStars.Authentication;

public class GitHubDeviceCodeTokenProvider : IAccessTokenProvider
{
    private static string[] GitHubApiHosts = new[] {"api.github.com"};
    private static string AuthorizationEndpoint = "https://github.com/login/device/code";
    private static string TokenEndpoint = "https://github.com/login/oauth/access_token";

    private string clientId;
    private string[] scopes;
    private Func<DeviceAuthorizationResponse, Task> devicePrompt;
    private string? cachedToken;

    public AllowedHostsValidator AllowedHostsValidator => new AllowedHostsValidator(GitHubApiHosts);

    public GitHubDeviceCodeTokenProvider(
        string clientId,
        string[] scopes,
        Func<DeviceAuthorizationResponse, Task> devicePrompt)
    {
        this.clientId = !string.IsNullOrEmpty(clientId) ? clientId :
            throw new ArgumentException("Client ID must not be empty or null");
        this.scopes = (scopes != null && scopes.Count() > 0) ? scopes :
            throw new ArgumentException("Scopes must not be empty or null");
        this.devicePrompt = devicePrompt ??
            throw new ArgumentException("Device prompt must not be null");
    }

    public async Task<string> GetAuthorizationTokenAsync(
        Uri uri,
        Dictionary<string, object>? additionalAuthenticationContext = null,
        CancellationToken cancellationToken = default)
    {
        if (!AllowedHostsValidator.IsUrlHostValid(uri))
        {
            return string.Empty;
        }

        if(!uri.Scheme.Equals("https", StringComparison.OrdinalIgnoreCase))
        {
            throw new ArgumentException("Only https is supported");
        }

        if (!string.IsNullOrEmpty(this.cachedToken))
        {
            return this.cachedToken;
        }

        // Implement GitHub's device flow to get a token
        // See: https://docs.github.com/en/apps/oauth-apps/building-oauth-apps/authorizing-oauth-apps

        // Step 1: Request device and user verification codes
        var client = new HttpClient();
        var response = await client.RequestDeviceAuthorizationAsync(
            new DeviceAuthorizationRequest
            {
                Address = AuthorizationEndpoint,
                ClientId = this.clientId,
                Scope = "user repo",
            },
            cancellationToken);

        // Step 2: Prompt the user
        // Pass response to the callback
        // Callback will display instructions to the user
        await this.devicePrompt(response);

        var interval = response.Interval * 1000;
        var tokenRequest = new DeviceTokenRequest
        {
            Address = TokenEndpoint,
            ClientId = this.clientId,
            DeviceCode = response.DeviceCode,
        };

        // Step 3: Poll GitHub for completed authorization
        while (true)
        {
            // Wait the indicated amount of seconds between polls
            await Task.Delay(interval);
            var tokenResponse = await client.RequestDeviceTokenAsync(
                    tokenRequest, cancellationToken);
            if (tokenResponse.IsError)
            {
                // If the user hasn't completed auth yet, GitHub returns "authorization_pending"
                if (string.Compare(tokenResponse.Error, "authorization_pending",
                    StringComparison.InvariantCultureIgnoreCase) != 0)
                {
                    throw tokenResponse.Exception ?? new Exception(tokenResponse.ErrorDescription);
                }
                continue;
            }

            // Cache the token (in memory) for subsequent calls
            this.cachedToken = tokenResponse.AccessToken;
            return tokenResponse.AccessToken;
        }
    }
}
// </TokenProviderSnippet>

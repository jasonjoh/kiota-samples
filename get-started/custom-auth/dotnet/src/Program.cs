// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

// <ProgramSnippet>
using Microsoft.Kiota.Abstractions.Authentication;
using Microsoft.Kiota.Http.HttpClientLibrary;
using RepoStars.Authentication;
using RepoStars.Client;

var clientId = string.Empty;
var githubScopes = new[] { "user", "repo" };

try
{
    var settings = Settings.LoadSettings();
    clientId = settings.ClientId ?? throw new Exception(
        "Client ID cannot be empty or null");
}
catch (Exception ex)
{
    Console.WriteLine(ex.Message);
    return;
}

var tokenProvider = new GitHubDeviceCodeTokenProvider(clientId, githubScopes, (response) => {
    Console.WriteLine(
        $"Please go to {response.VerificationUri} and enter code {response.UserCode} to authenticate.");
    return Task.CompletedTask;
});

var authProvider = new BaseBearerTokenAuthenticationProvider(tokenProvider);

// Create request adapter using the HttpClient-based implementation
var adapter = new HttpClientRequestAdapter(authProvider);
// Create the API client
var client = new ReposClient(adapter);

try
{
    var repos = await client.User.Repos.GetAsync();
    if (repos != null)
    {
        Console.WriteLine($"Found {repos.Count} repositories.");
        foreach (var repo in repos)
        {
            Console.WriteLine($"{repo.Name} has {repo.StargazersCount} stars");
        }
    }
}
catch (Exception ex)
{
    Console.WriteLine($"ERROR: {ex.Message}");
    Console.WriteLine(ex.StackTrace);
}
// </ProgramSnippet>

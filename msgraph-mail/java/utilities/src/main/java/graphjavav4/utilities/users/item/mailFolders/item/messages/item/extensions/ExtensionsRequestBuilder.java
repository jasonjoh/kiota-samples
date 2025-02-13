package graphjavav4.utilities.users.item.mailfolders.item.messages.item.extensions;

import com.microsoft.kiota.BaseRequestBuilder;
import com.microsoft.kiota.BaseRequestConfiguration;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.QueryParameter;
import com.microsoft.kiota.RequestAdapter;
import com.microsoft.kiota.RequestInformation;
import com.microsoft.kiota.RequestOption;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParsableFactory;
import graphjavav4.utilities.models.Extension;
import graphjavav4.utilities.models.ExtensionCollectionResponse;
import graphjavav4.utilities.users.item.mailfolders.item.messages.item.extensions.item.ExtensionItemRequestBuilder;
import java.net.URISyntaxException;
import java.util.Collection;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
/**
 * Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/extensions
 */
public class ExtensionsRequestBuilder extends BaseRequestBuilder {
    /**
     * Instantiates a new ExtensionsRequestBuilder and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    @javax.annotation.Nullable
    public ExtensionsRequestBuilder(@javax.annotation.Nonnull final HashMap<String, Object> pathParameters, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/{message%2Did}/extensions{?%24top,%24skip,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters);
    }
    /**
     * Instantiates a new ExtensionsRequestBuilder and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    @javax.annotation.Nullable
    public ExtensionsRequestBuilder(@javax.annotation.Nonnull final String rawUrl, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/{message%2Did}/extensions{?%24top,%24skip,%24filter,%24count,%24orderby,%24select,%24expand}", rawUrl);
    }
    /**
     * The collection of open extensions defined for the message. Nullable.
     * @return a CompletableFuture of ExtensionCollectionResponse
     */
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<ExtensionCollectionResponse> get() {
        try {
            final RequestInformation requestInfo = toGetRequestInformation(null);
            return this.requestAdapter.sendAsync(requestInfo, ExtensionCollectionResponse::createFromDiscriminatorValue, null);
        } catch (URISyntaxException ex) {
            final java.util.concurrent.CompletableFuture<ExtensionCollectionResponse> executionException = new java.util.concurrent.CompletableFuture<ExtensionCollectionResponse>();
            executionException.completeExceptionally(ex);
            return executionException;
        }
    }
    /**
     * The collection of open extensions defined for the message. Nullable.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a CompletableFuture of ExtensionCollectionResponse
     */
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<ExtensionCollectionResponse> get(@javax.annotation.Nullable final java.util.function.Consumer<GetRequestConfiguration> requestConfiguration) {
        try {
            final RequestInformation requestInfo = toGetRequestInformation(requestConfiguration);
            return this.requestAdapter.sendAsync(requestInfo, ExtensionCollectionResponse::createFromDiscriminatorValue, null);
        } catch (URISyntaxException ex) {
            final java.util.concurrent.CompletableFuture<ExtensionCollectionResponse> executionException = new java.util.concurrent.CompletableFuture<ExtensionCollectionResponse>();
            executionException.completeExceptionally(ex);
            return executionException;
        }
    }
    /**
     * Create an open extension (openTypeExtension object) and add custom properties in a new or existing instance of a resource. You can create an open extension in a resource instance and store custom data to it all in the same operation, except for specific resources. See known limitations of open extensions for more information. The table in the Permissions section lists the resources that support open extensions.
     * @param body The request body
     * @return a CompletableFuture of extension
     * @see <a href="https://docs.microsoft.com/graph/api/opentypeextension-post-opentypeextension?view=graph-rest-1.0">Find more info here</a>
     */
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<Extension> post(@javax.annotation.Nonnull final Extension body) {
        try {
            final RequestInformation requestInfo = toPostRequestInformation(body, null);
            return this.requestAdapter.sendAsync(requestInfo, Extension::createFromDiscriminatorValue, null);
        } catch (URISyntaxException ex) {
            final java.util.concurrent.CompletableFuture<Extension> executionException = new java.util.concurrent.CompletableFuture<Extension>();
            executionException.completeExceptionally(ex);
            return executionException;
        }
    }
    /**
     * Create an open extension (openTypeExtension object) and add custom properties in a new or existing instance of a resource. You can create an open extension in a resource instance and store custom data to it all in the same operation, except for specific resources. See known limitations of open extensions for more information. The table in the Permissions section lists the resources that support open extensions.
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a CompletableFuture of extension
     * @see <a href="https://docs.microsoft.com/graph/api/opentypeextension-post-opentypeextension?view=graph-rest-1.0">Find more info here</a>
     */
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<Extension> post(@javax.annotation.Nonnull final Extension body, @javax.annotation.Nullable final java.util.function.Consumer<PostRequestConfiguration> requestConfiguration) {
        Objects.requireNonNull(body);
        try {
            final RequestInformation requestInfo = toPostRequestInformation(body, requestConfiguration);
            return this.requestAdapter.sendAsync(requestInfo, Extension::createFromDiscriminatorValue, null);
        } catch (URISyntaxException ex) {
            final java.util.concurrent.CompletableFuture<Extension> executionException = new java.util.concurrent.CompletableFuture<Extension>();
            executionException.completeExceptionally(ex);
            return executionException;
        }
    }
    /**
     * The collection of open extensions defined for the message. Nullable.
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation toGetRequestInformation() throws URISyntaxException {
        return toGetRequestInformation(null);
    }
    /**
     * The collection of open extensions defined for the message. Nullable.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation toGetRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<GetRequestConfiguration> requestConfiguration) throws URISyntaxException {
        final RequestInformation requestInfo = new RequestInformation();
        requestInfo.httpMethod = HttpMethod.GET;
        requestInfo.urlTemplate = urlTemplate;
        requestInfo.pathParameters = pathParameters;
        requestInfo.headers.add("Accept", "application/json");
        if (requestConfiguration != null) {
            final GetRequestConfiguration requestConfig = new GetRequestConfiguration();
            requestConfiguration.accept(requestConfig);
            requestInfo.addQueryParameters(requestConfig.queryParameters);
            requestInfo.headers.putAll(requestConfig.headers);
            requestInfo.addRequestOptions(requestConfig.options);
        }
        return requestInfo;
    }
    /**
     * Create an open extension (openTypeExtension object) and add custom properties in a new or existing instance of a resource. You can create an open extension in a resource instance and store custom data to it all in the same operation, except for specific resources. See known limitations of open extensions for more information. The table in the Permissions section lists the resources that support open extensions.
     * @param body The request body
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation toPostRequestInformation(@javax.annotation.Nonnull final Extension body) throws URISyntaxException {
        return toPostRequestInformation(body, null);
    }
    /**
     * Create an open extension (openTypeExtension object) and add custom properties in a new or existing instance of a resource. You can create an open extension in a resource instance and store custom data to it all in the same operation, except for specific resources. See known limitations of open extensions for more information. The table in the Permissions section lists the resources that support open extensions.
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation toPostRequestInformation(@javax.annotation.Nonnull final Extension body, @javax.annotation.Nullable final java.util.function.Consumer<PostRequestConfiguration> requestConfiguration) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInformation requestInfo = new RequestInformation();
        requestInfo.httpMethod = HttpMethod.POST;
        requestInfo.urlTemplate = urlTemplate;
        requestInfo.pathParameters = pathParameters;
        requestInfo.headers.add("Accept", "application/json");
        requestInfo.setContentFromParsable(requestAdapter, "application/json", body);
        if (requestConfiguration != null) {
            final PostRequestConfiguration requestConfig = new PostRequestConfiguration();
            requestConfiguration.accept(requestConfig);
            requestInfo.headers.putAll(requestConfig.headers);
            requestInfo.addRequestOptions(requestConfig.options);
        }
        return requestInfo;
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.item.mailFolders.item.messages.item.extensions.item collection
     * @param extensionId Unique identifier of the item
     * @return a ExtensionItemRequestBuilder
     */
    @javax.annotation.Nonnull
    public ExtensionItemRequestBuilder withExtensionId(@javax.annotation.Nonnull final String extensionId) {
        Objects.requireNonNull(extensionId);
        final HashMap<String, Object> urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("extension%2Did", extensionId);
        return new ExtensionItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /**
     * The collection of open extensions defined for the message. Nullable.
     */
    public class GetQueryParameters {
        /** Include count of items */
        @QueryParameter(name = "%24count")
        @javax.annotation.Nullable
        public Boolean count;
        /** Expand related entities */
        @QueryParameter(name = "%24expand")
        @javax.annotation.Nullable
        public String[] expand;
        /** Filter items by property values */
        @QueryParameter(name = "%24filter")
        @javax.annotation.Nullable
        public String filter;
        /** Order items by property values */
        @QueryParameter(name = "%24orderby")
        @javax.annotation.Nullable
        public String[] orderby;
        /** Select properties to be returned */
        @QueryParameter(name = "%24select")
        @javax.annotation.Nullable
        public String[] select;
        /** Skip the first n items */
        @QueryParameter(name = "%24skip")
        @javax.annotation.Nullable
        public Integer skip;
        /** Show only the first n items */
        @QueryParameter(name = "%24top")
        @javax.annotation.Nullable
        public Integer top;
    }
    /**
     * Configuration for the request such as headers, query parameters, and middleware options.
     */
    public class GetRequestConfiguration extends BaseRequestConfiguration {
        /** Request query parameters */
        @javax.annotation.Nullable
        public GetQueryParameters queryParameters = new GetQueryParameters();
    }
    /**
     * Configuration for the request such as headers, query parameters, and middleware options.
     */
    public class PostRequestConfiguration extends BaseRequestConfiguration {
    }
}

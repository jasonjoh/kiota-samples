package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
)

// ItemMessagesRequestBuilder builds and executes requests for operations under \users\{user-id}\messages
type ItemMessagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMessagesRequestBuilderGetQueryParameters get the messages in the signed-in user's mailbox (including the Deleted Items and Clutter folders). Depending on the page size and mailbox data, getting messages from a mailbox can incur multiple requests. The default page size is 10 messages. Use `$top` to customize the page size, within the range of 1 and 1000. To improve the operation response time, use `$select` to specify the exact properties you need; see example 1 below. Fine-tune the values for `$select` and `$top`, especially when you must use a larger page size, as returning a page with hundreds of messages each with a full response payload may trigger the gateway timeout (HTTP 504). To get the next page of messages, simply apply the entire URL returned in `@odata.nextLink` to the next get-messages request. This URL includes any query parameters you may have specified in the initial request.  Do not try to extract the `$skip` value from the `@odata.nextLink` URL to manipulate responses. This API uses the `$skip` value to keep count of all the items it has gone through in the user's mailbox to return a page of message-type items. It's therefore possible that even in the initial response, the `$skip` value is larger than the page size. For more information, see Paging Microsoft Graph data in your app. Currently, this operation returns message bodies in only HTML format. There are two scenarios where an app can get messages in another user's mail folder:
type ItemMessagesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemMessagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMessagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMessagesRequestBuilderGetQueryParameters
}
// ItemMessagesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMessagesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMessagesRequestBuilderInternal instantiates a new MessagesRequestBuilder and sets the default values.
func NewItemMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMessagesRequestBuilder) {
    m := &ItemMessagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/messages{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemMessagesRequestBuilder instantiates a new MessagesRequestBuilder and sets the default values.
func NewItemMessagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMessagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMessagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the messages in the signed-in user's mailbox (including the Deleted Items and Clutter folders). Depending on the page size and mailbox data, getting messages from a mailbox can incur multiple requests. The default page size is 10 messages. Use `$top` to customize the page size, within the range of 1 and 1000. To improve the operation response time, use `$select` to specify the exact properties you need; see example 1 below. Fine-tune the values for `$select` and `$top`, especially when you must use a larger page size, as returning a page with hundreds of messages each with a full response payload may trigger the gateway timeout (HTTP 504). To get the next page of messages, simply apply the entire URL returned in `@odata.nextLink` to the next get-messages request. This URL includes any query parameters you may have specified in the initial request.  Do not try to extract the `$skip` value from the `@odata.nextLink` URL to manipulate responses. This API uses the `$skip` value to keep count of all the items it has gone through in the user's mailbox to return a page of message-type items. It's therefore possible that even in the initial response, the `$skip` value is larger than the page size. For more information, see Paging Microsoft Graph data in your app. Currently, this operation returns message bodies in only HTML format. There are two scenarios where an app can get messages in another user's mail folder:
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/user-list-messages?view=graph-rest-1.0
func (m *ItemMessagesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMessagesRequestBuilderGetRequestConfiguration)(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.CreateMessageCollectionResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageCollectionResponseable), nil
}
// Post create a draft of a new message in either JSON or MIME format. When using JSON format, you can:- Include an attachment to the **message**.- Update the draft later to add content to the **body** or change other message properties. When using MIME format:- Provide the applicable Internet message headers and the MIME content, all encoded in **base64** format in the request body.- /* Add any attachments and S/MIME properties to the MIME content. By default, this operation saves the draft in the Drafts folder. Send the draft message in a subsequent operation. Alternatively, send a new message in a single operation, or create a draft to forward, reply and reply-all to an existing message.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/user-post-messages?view=graph-rest-1.0
func (m *ItemMessagesRequestBuilder) Post(ctx context.Context, body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Messageable, requestConfiguration *ItemMessagesRequestBuilderPostRequestConfiguration)(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Messageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.CreateMessageFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Messageable), nil
}
// ToGetRequestInformation get the messages in the signed-in user's mailbox (including the Deleted Items and Clutter folders). Depending on the page size and mailbox data, getting messages from a mailbox can incur multiple requests. The default page size is 10 messages. Use `$top` to customize the page size, within the range of 1 and 1000. To improve the operation response time, use `$select` to specify the exact properties you need; see example 1 below. Fine-tune the values for `$select` and `$top`, especially when you must use a larger page size, as returning a page with hundreds of messages each with a full response payload may trigger the gateway timeout (HTTP 504). To get the next page of messages, simply apply the entire URL returned in `@odata.nextLink` to the next get-messages request. This URL includes any query parameters you may have specified in the initial request.  Do not try to extract the `$skip` value from the `@odata.nextLink` URL to manipulate responses. This API uses the `$skip` value to keep count of all the items it has gone through in the user's mailbox to return a page of message-type items. It's therefore possible that even in the initial response, the `$skip` value is larger than the page size. For more information, see Paging Microsoft Graph data in your app. Currently, this operation returns message bodies in only HTML format. There are two scenarios where an app can get messages in another user's mail folder:
func (m *ItemMessagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMessagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation create a draft of a new message in either JSON or MIME format. When using JSON format, you can:- Include an attachment to the **message**.- Update the draft later to add content to the **body** or change other message properties. When using MIME format:- Provide the applicable Internet message headers and the MIME content, all encoded in **base64** format in the request body.- /* Add any attachments and S/MIME properties to the MIME content. By default, this operation saves the draft in the Drafts folder. Send the draft message in a subsequent operation. Alternatively, send a new message in a single operation, or create a draft to forward, reply and reply-all to an existing message.
func (m *ItemMessagesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Messageable, requestConfiguration *ItemMessagesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithMessageId gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.messages.item collection
func (m *ItemMessagesRequestBuilder) WithMessageId(messageId string)(*ItemMessagesMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if messageId != "" {
        urlTplParams["message%2Did"] = messageId
    }
    return NewItemMessagesMessageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

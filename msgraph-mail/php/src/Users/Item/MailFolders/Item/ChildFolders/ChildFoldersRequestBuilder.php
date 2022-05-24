<?php

namespace Microsoft\Graph\Users\Item\MailFolders\Item\ChildFolders;

use Exception;
use Http\Promise\Promise;
use Http\Promise\RejectedPromise;
use Microsoft\Graph\Models\Microsoft\Graph\MailFolder;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\QueryParameter;
use Microsoft\Kiota\Abstractions\RequestAdapter;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\RequestOption;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParsableFactory;

class ChildFoldersRequestBuilder 
{
    /**
     * @var array<string, mixed> $pathParameters Path parameters for the request
    */
    private array $pathParameters;
    
    /**
     * @var RequestAdapter $requestAdapter The request adapter to use to execute the requests.
    */
    private RequestAdapter $requestAdapter;
    
    /**
     * @var string $urlTemplate Url template to use to build the URL for the current request builder
    */
    private string $urlTemplate;
    
    /**
     * Instantiates a new ChildFoldersRequestBuilder and sets the default values.
     * @param array<string, mixed> $pathParameters Path parameters for the request
     * @param RequestAdapter $requestAdapter The request adapter to use to execute the requests.
    */
    public function __construct(array $pathParameters, RequestAdapter $requestAdapter) {
        $this->urlTemplate = '{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders{?%24top,%24skip,%24filter,%24count,%24orderby,%24select,%24expand}';
        $this->requestAdapter = $requestAdapter;
        $this->pathParameters = $pathParameters;
    }

    /**
     * The collection of child folders in the mailFolder.
     * @param array<string, mixed>|null $requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return RequestInformation
    */
    public function createGetRequestInformation(?ChildFoldersRequestBuilderGetRequestConfiguration $requestConfiguration = null): RequestInformation {
        $requestInfo = new RequestInformation();
        $requestInfo->urlTemplate = $this->urlTemplate;
        $requestInfo->pathParameters = $this->pathParameters;
        $requestInfo->httpMethod = HttpMethod::GET;
        if ($requestConfiguration !== null) {
            if ($requestConfiguration->headers !== null) {
                $requestInfo->headers = array_merge($requestInfo->headers, $requestConfiguration->headers);
            }
            if ($requestConfiguration->queryParameters !== null) {
                $requestInfo->setQueryParameters($requestConfiguration->queryParameters);
            }
            if ($requestConfiguration->options !== null) {
                $requestInfo->addRequestOptions(...$requestConfiguration->options);
            }
        }
        return $requestInfo;
    }

    /**
     * The collection of child folders in the mailFolder.
     * @param MailFolder $body 
     * @param array<string, mixed>|null $requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return RequestInformation
    */
    public function createPostRequestInformation(MailFolder $body, ?ChildFoldersRequestBuilderPostRequestConfiguration $requestConfiguration = null): RequestInformation {
        $requestInfo = new RequestInformation();
        $requestInfo->urlTemplate = $this->urlTemplate;
        $requestInfo->pathParameters = $this->pathParameters;
        $requestInfo->httpMethod = HttpMethod::POST;
        if ($requestConfiguration !== null) {
            if ($requestConfiguration->headers !== null) {
                $requestInfo->headers = array_merge($requestInfo->headers, $requestConfiguration->headers);
            }
            if ($requestConfiguration->options !== null) {
                $requestInfo->addRequestOptions(...$requestConfiguration->options);
            }
        }
        $requestInfo->setContentFromParsable($this->requestAdapter, "application/json", $body);
        return $requestInfo;
    }

    /**
     * The collection of child folders in the mailFolder.
     * @param array<string, mixed>|null $requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return Promise
    */
    public function get(?ChildFoldersRequestBuilderGetRequestConfiguration $requestConfiguration = null, ?ResponseHandler $responseHandler = null): Promise {
        $requestInfo = $this->createGetRequestInformation($requestConfiguration);
        try {
            return $this->requestAdapter->sendAsync($requestInfo, array(ChildFoldersResponse::class, 'createFromDiscriminatorValue'), $responseHandler, null);
        } catch(Exception $ex) {
            return new RejectedPromise($ex);
        }
    }

    /**
     * The collection of child folders in the mailFolder.
     * @param MailFolder $body 
     * @param array<string, mixed>|null $requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return Promise
    */
    public function post(MailFolder $body, ?ChildFoldersRequestBuilderPostRequestConfiguration $requestConfiguration = null, ?ResponseHandler $responseHandler = null): Promise {
        $requestInfo = $this->createPostRequestInformation($body, $requestConfiguration);
        try {
            return $this->requestAdapter->sendAsync($requestInfo, array(MailFolder::class, 'createFromDiscriminatorValue'), $responseHandler, null);
        } catch(Exception $ex) {
            return new RejectedPromise($ex);
        }
    }

    <?php
    
    class ChildFoldersRequestBuilderGetQueryParameters 
    {
        /**
         * @QueryParameter("%24count")
         * @var bool|null $count Include count of items
        */
        public ?bool $count = null;
        
        /**
         * @QueryParameter("%24expand")
         * @var array<string>|null $expand Expand related entities
        */
        public ?array $expand = null;
        
        /**
         * @QueryParameter("%24filter")
         * @var string|null $filter Filter items by property values
        */
        public ?string $filter = null;
        
        /**
         * @QueryParameter("%24orderby")
         * @var array<string>|null $orderby Order items by property values
        */
        public ?array $orderby = null;
        
        /**
         * @QueryParameter("%24select")
         * @var array<string>|null $select Select properties to be returned
        */
        public ?array $select = null;
        
        /**
         * @QueryParameter("%24skip")
         * @var int|null $skip Skip the first n items
        */
        public ?int $skip = null;
        
        /**
         * @QueryParameter("%24top")
         * @var int|null $top Show only the first n items
        */
        public ?int $top = null;
        
    }
    <?php
    
    class ChildFoldersRequestBuilderGetRequestConfiguration 
    {
        /**
         * @var array|null $headers Request headers
        */
        public ?array $headers = null;
        
        /**
         * @var array|null $options Request options
        */
        public ?array $options = null;
        
        /**
         * @var ChildFoldersRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
        */
        public ?ChildFoldersRequestBuilderGetQueryParameters $queryParameters = null;
        
    }
    <?php
    
    class ChildFoldersRequestBuilderPostRequestConfiguration 
    {
        /**
         * @var array|null $headers Request headers
        */
        public ?array $headers = null;
        
        /**
         * @var array|null $options Request options
        */
        public ?array $options = null;
        
    }
}

require 'microsoft_kiota_abstractions'
require_relative './item/message_rule_request_builder'

module Graphrubyv4::Utilities::Users::MailFolders::MessageRules
    ## 
    # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messageRules
    class MessageRulesRequestBuilder
        
        ## 
        # Current path for the request
        @current_path
        ## 
        # The http core service to use to execute the requests.
        @http_core
        ## 
        # Path segment to use to build the URL for the current request builder
        @path_segment
        ## 
        # Gets an item from the graphrubyv4.utilities.users.mailFolders.messageRules collection
        def [](position)
            return MessageRuleRequestBuilder.new(@current_path + @path_segment  + "/" + position, @http_core)
        end
        ## 
        ## Instantiates a new MessageRulesRequestBuilder and sets the default values.
        ## @param currentPath Current path for the request
        ## @param httpCore The http core service to use to execute the requests.
        ## @return a void
        ## 
        def initialize(current_path, http_core) 
            @path_segment = "/messageRules"
            @http_core = http_core
            @current_path = current_path
        end
        ## 
        ## Get messageRules from users
        ## @param h Request headers
        ## @param q Request query parameters
        ## @return a request_info
        ## 
        def create_get_request_info(q, h) 
            request_info = RequestInfo.new()
            request_info.URI = current_path + path_segment
            request_info.http_method = :GET
            request_info.set_headers_from_raw_object(h)
            request_info.set_query_string_parameters_from_raw_object(q)
            return request_info;
        end
        ## 
        ## Create new navigation property to messageRules for users
        ## @param body 
        ## @param h Request headers
        ## @return a request_info
        ## 
        def create_post_request_info(body, h) 
            request_info = RequestInfo.new()
            request_info.URI = current_path + path_segment
            request_info.http_method = :POST
            request_info.set_headers_from_raw_object(h)
            request_info.set_content_from_parsable(body, self.serializer_factory, "application/json")
            return request_info;
        end
        ## 
        ## Get messageRules from users
        ## @param h Request headers
        ## @param q Request query parameters
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of message_rules_response
        ## 
        def get(q, h, response_handler) 
            request_info = self.create_get_request_info(
                q, h
            )
            return self.http_core.send_async(request_info, response_handler)
        end
        ## 
        ## Create new navigation property to messageRules for users
        ## @param body 
        ## @param h Request headers
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of message_rule
        ## 
        def post(body, h, response_handler) 
            request_info = self.create_post_request_info(
                body, h
            )
            return self.http_core.send_async(request_info, response_handler)
        end
    end
end

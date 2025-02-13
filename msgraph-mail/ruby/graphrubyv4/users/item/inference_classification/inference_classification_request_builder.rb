require 'microsoft_kiota_abstractions'
require_relative '../../../graphrubyv4'
require_relative '../../../models/inference_classification'
require_relative '../../users'
require_relative '../item'
require_relative './inference_classification'
require_relative './overrides/item/inference_classification_override_item_request_builder'
require_relative './overrides/overrides_request_builder'

module Graphrubyv4
    module Users
        module Item
            module InferenceClassification
                ## 
                # Builds and executes requests for operations under \users\{user-id}\inferenceClassification
                class InferenceClassificationRequestBuilder < MicrosoftKiotaAbstractions::BaseRequestBuilder
                    
                    ## 
                    # The overrides property
                    def overrides()
                        return Graphrubyv4::Users::Item::InferenceClassification::Overrides::OverridesRequestBuilder.new(@path_parameters, @request_adapter)
                    end
                    ## 
                    ## Instantiates a new InferenceClassificationRequestBuilder and sets the default values.
                    ## @param path_parameters Path parameters for the request
                    ## @param request_adapter The request adapter to use to execute the requests.
                    ## @return a void
                    ## 
                    def initialize(path_parameters, request_adapter)
                        super(path_parameters, request_adapter, "{+baseurl}/users/{user%2Did}/inferenceClassification{?%24select}")
                    end
                    ## 
                    ## Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
                    ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                    ## @return a Fiber of inference_classification
                    ## 
                    def get(request_configuration=nil)
                        request_info = self.to_get_request_information(
                            request_configuration
                        )
                        return @request_adapter.send_async(request_info, lambda {|pn| Graphrubyv4::Models::InferenceClassification.create_from_discriminator_value(pn) }, nil)
                    end
                    ## 
                    ## Gets an item from the graphrubyv4.users.item.inferenceClassification.overrides.item collection
                    ## @param id Unique identifier of the item
                    ## @return a inference_classification_override_item_request_builder
                    ## 
                    def overrides_by_id(id)
                        raise StandardError, 'id cannot be null' if id.nil?
                        url_tpl_params = @path_parameters.clone
                        url_tpl_params["inferenceClassificationOverride%2Did"] = id
                        return Graphrubyv4::Users::Item::InferenceClassification::Overrides::Item::InferenceClassificationOverrideItemRequestBuilder.new(url_tpl_params, @request_adapter)
                    end
                    ## 
                    ## Update the navigation property inferenceClassification in users
                    ## @param body The request body
                    ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                    ## @return a Fiber of void
                    ## 
                    def patch(body, request_configuration=nil)
                        raise StandardError, 'body cannot be null' if body.nil?
                        request_info = self.to_patch_request_information(
                            body, request_configuration
                        )
                        return @request_adapter.send_async(request_info, nil, nil)
                    end
                    ## 
                    ## Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
                    ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                    ## @return a request_information
                    ## 
                    def to_get_request_information(request_configuration=nil)
                        request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
                        request_info.url_template = @url_template
                        request_info.path_parameters = @path_parameters
                        request_info.http_method = :GET
                        request_info.headers.add('Accept', 'application/json')
                        unless request_configuration.nil?
                            request_info.add_headers_from_raw_object(request_configuration.headers)
                            request_info.set_query_string_parameters_from_raw_object(request_configuration.query_parameters)
                            request_info.add_request_options(request_configuration.options)
                        end
                        return request_info
                    end
                    ## 
                    ## Update the navigation property inferenceClassification in users
                    ## @param body The request body
                    ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                    ## @return a request_information
                    ## 
                    def to_patch_request_information(body, request_configuration=nil)
                        raise StandardError, 'body cannot be null' if body.nil?
                        request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
                        request_info.url_template = @url_template
                        request_info.path_parameters = @path_parameters
                        request_info.http_method = :PATCH
                        unless request_configuration.nil?
                            request_info.add_headers_from_raw_object(request_configuration.headers)
                            request_info.add_request_options(request_configuration.options)
                        end
                        request_info.set_content_from_parsable(self.request_adapter, "application/json", body)
                        return request_info
                    end

                    ## 
                    # Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
                    class InferenceClassificationRequestBuilderGetQueryParameters
                        
                        ## 
                        # Select properties to be returned
                        attr_accessor :select
                        ## 
                        ## Maps the query parameters names to their encoded names for the URI template parsing.
                        ## @param original_name The original query parameter name in the class.
                        ## @return a string
                        ## 
                        def get_query_parameter(original_name)
                            raise StandardError, 'original_name cannot be null' if original_name.nil?
                            case original_name
                                when "select"
                                    return "%24select"
                                else
                                    return original_name
                            end
                        end
                    end
                end
            end
        end
    end
end

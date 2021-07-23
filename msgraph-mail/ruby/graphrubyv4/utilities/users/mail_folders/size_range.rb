require 'microsoft_kiota_abstractions'

module Graphrubyv4::Utilities::Users::MailFolders
    class SizeRange
        include MicrosoftKiotaAbstractions::Parsable
        ## 
        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        @additional_data
        ## 
        # The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        @maximum_size
        ## 
        # The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        @minimum_size
        ## 
        ## Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        ## @return a i_dictionary
        ## 
        def  additional_data
            return @additional_data
        end
        ## 
        ## Gets the maximumSize property value. The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        ## @return a integer
        ## 
        def  maximum_size
            return @maximum_size
        end
        ## 
        ## Gets the minimumSize property value. The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        ## @return a integer
        ## 
        def  minimum_size
            return @minimum_size
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return {
                "maximum_size" => lambda {|o, n| o.maximum_size = n.get_object_value(integer) },
                "minimum_size" => lambda {|o, n| o.minimum_size = n.get_object_value(integer) },
            }
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            writer.write_object_value("maximum_size", self.maximum_size)
            writer.write_object_value("minimum_size", self.minimum_size)
            writer.write_additional_data(self.additional_data)
        end
        ## 
        ## Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        ## @param value Value to set for the AdditionalData property.
        ## @return a void
        ## 
        def  additional_data=(additionalData)
            @additional_data = (additionalData)
        end
        ## 
        ## Sets the maximumSize property value. The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        ## @param value Value to set for the maximumSize property.
        ## @return a void
        ## 
        def  maximum_size=(maximumSize)
            @maximum_size = (maximumSize)
        end
        ## 
        ## Sets the minimumSize property value. The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        ## @param value Value to set for the minimumSize property.
        ## @return a void
        ## 
        def  minimum_size=(minimumSize)
            @minimum_size = (minimumSize)
        end
    end
end

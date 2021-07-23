require 'microsoft_kiota_abstractions'
require_relative './entity'

module Graphrubyv4::Utilities::Users
    class MultiValueLegacyExtendedProperty < Graphrubyv4::Utilities::Users::Entity
        include MicrosoftKiotaAbstractions::Parsable
        ## 
        # A collection of property values.
        @value
        ## 
        ## Gets the value property value. A collection of property values.
        ## @return a string
        ## 
        def  value
            return @value
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return {
                "value" => lambda {|o, n| o.value = n.get_collection_of_primitive_values },
            }
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            super.serialize(writer)
            writer.write_collection_of_primitive_values("value", self.value)
        end
        ## 
        ## Sets the value property value. A collection of property values.
        ## @param value Value to set for the value property.
        ## @return a void
        ## 
        def  value=(value)
            @value = (value)
        end
    end
end

package extensions

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
)

type ExtensionsResponse struct {
    additionalData map[string]interface{};
    nextLink *string;
    value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Extension;
}
func NewExtensionsResponse()(*ExtensionsResponse) {
    m := &ExtensionsResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ExtensionsResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *ExtensionsResponse) GetNextLink()(*string) {
    return m.nextLink
}
func (m *ExtensionsResponse) GetValue()([]ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Extension) {
    return m.value
}
func (m *ExtensionsResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *ExtensionsResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *ExtensionsResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ExtensionsResponse) SetNextLink(value *string)() {
    m.nextLink = value
}
func (m *ExtensionsResponse) SetValue(value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Extension)() {
    m.value = value
}

<?php

namespace Microsoft\Graph\Models;

use Microsoft\Kiota\Abstractions\Serialization\AdditionalDataHolder;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class SingleValueLegacyExtendedPropertyCollectionResponse implements AdditionalDataHolder, Parsable 
{
    /**
     * @var array<string, mixed>|null $additionalData Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    */
    private ?array $additionalData = null;
    
    /**
     * @var string|null $odataNextLink The OdataNextLink property
    */
    private ?string $odataNextLink = null;
    
    /**
     * @var array<SingleValueLegacyExtendedProperty>|null $value The value property
    */
    private ?array $value = null;
    
    /**
     * Instantiates a new SingleValueLegacyExtendedPropertyCollectionResponse and sets the default values.
    */
    public function __construct() {
        $this->setAdditionalData([]);
    }

    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param ParseNode $parseNode The parse node to use to read the discriminator value and create the object
     * @return SingleValueLegacyExtendedPropertyCollectionResponse
    */
    public static function createFromDiscriminatorValue(ParseNode $parseNode): SingleValueLegacyExtendedPropertyCollectionResponse {
        return new SingleValueLegacyExtendedPropertyCollectionResponse();
    }

    /**
     * Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return array<string, mixed>|null
    */
    public function getAdditionalData(): ?array {
        return $this->additionalData;
    }

    /**
     * The deserialization information for the current model
     * @return array<string, callable>
    */
    public function getFieldDeserializers(): array {
        $o = $this;
        return  [
            '@odata.nextLink' => fn(ParseNode $n) => $o->setOdataNextLink($n->getStringValue()),
            'value' => fn(ParseNode $n) => $o->setValue($n->getCollectionOfObjectValues([SingleValueLegacyExtendedProperty::class, 'createFromDiscriminatorValue'])),
        ];
    }

    /**
     * Gets the @odata.nextLink property value. The OdataNextLink property
     * @return string|null
    */
    public function getOdataNextLink(): ?string {
        return $this->odataNextLink;
    }

    /**
     * Gets the value property value. The value property
     * @return array<SingleValueLegacyExtendedProperty>|null
    */
    public function getValue(): ?array {
        return $this->value;
    }

    /**
     * Serializes information the current object
     * @param SerializationWriter $writer Serialization writer to use to serialize this model
    */
    public function serialize(SerializationWriter $writer): void {
        $writer->writeStringValue('@odata.nextLink', $this->getOdataNextLink());
        $writer->writeCollectionOfObjectValues('value', $this->getValue());
        $writer->writeAdditionalData($this->getAdditionalData());
    }

    /**
     * Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param array<string,mixed> $value Value to set for the AdditionalData property.
    */
    public function setAdditionalData(?array $value): void {
        $this->additionalData = $value;
    }

    /**
     * Sets the @odata.nextLink property value. The OdataNextLink property
     * @param string|null $value Value to set for the OdataNextLink property.
    */
    public function setOdataNextLink(?string $value): void {
        $this->odataNextLink = $value;
    }

    /**
     * Sets the value property value. The value property
     * @param array<SingleValueLegacyExtendedProperty>|null $value Value to set for the value property.
    */
    public function setValue(?array $value): void {
        $this->value = $value;
    }

}

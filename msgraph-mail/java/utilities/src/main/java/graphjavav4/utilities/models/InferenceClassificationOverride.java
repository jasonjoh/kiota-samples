package graphjavav4.utilities.models;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class InferenceClassificationOverride extends Entity implements Parsable {
    /** The classifyAs property */
    private InferenceClassificationType classifyAs;
    /** The senderEmailAddress property */
    private EmailAddress senderEmailAddress;
    /**
     * Instantiates a new inferenceClassificationOverride and sets the default values.
     * @return a void
     */
    @javax.annotation.Nullable
    public InferenceClassificationOverride() {
        super();
    }
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a inferenceClassificationOverride
     */
    @javax.annotation.Nonnull
    public static InferenceClassificationOverride createFromDiscriminatorValue(@javax.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new InferenceClassificationOverride();
    }
    /**
     * Gets the classifyAs property value. The classifyAs property
     * @return a inferenceClassificationType
     */
    @javax.annotation.Nullable
    public InferenceClassificationType getClassifyAs() {
        return this.classifyAs;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, java.util.function.Consumer<ParseNode>>
     */
    @javax.annotation.Nonnull
    public Map<String, java.util.function.Consumer<ParseNode>> getFieldDeserializers() {
        final HashMap<String, java.util.function.Consumer<ParseNode>> deserializerMap = new HashMap<String, java.util.function.Consumer<ParseNode>>(super.getFieldDeserializers());
        deserializerMap.put("classifyAs", (n) -> { this.setClassifyAs(n.getEnumValue(InferenceClassificationType.class)); });
        deserializerMap.put("senderEmailAddress", (n) -> { this.setSenderEmailAddress(n.getObjectValue(EmailAddress::createFromDiscriminatorValue)); });
        return deserializerMap;
    }
    /**
     * Gets the senderEmailAddress property value. The senderEmailAddress property
     * @return a emailAddress
     */
    @javax.annotation.Nullable
    public EmailAddress getSenderEmailAddress() {
        return this.senderEmailAddress;
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    @javax.annotation.Nonnull
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        super.serialize(writer);
        writer.writeEnumValue("classifyAs", this.getClassifyAs());
        writer.writeObjectValue("senderEmailAddress", this.getSenderEmailAddress());
    }
    /**
     * Sets the classifyAs property value. The classifyAs property
     * @param value Value to set for the classifyAs property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setClassifyAs(@javax.annotation.Nullable final InferenceClassificationType value) {
        this.classifyAs = value;
    }
    /**
     * Sets the senderEmailAddress property value. The senderEmailAddress property
     * @param value Value to set for the senderEmailAddress property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setSenderEmailAddress(@javax.annotation.Nullable final EmailAddress value) {
        this.senderEmailAddress = value;
    }
}

from __future__ import annotations
from kiota_abstractions.serialization import Parsable, ParseNode, SerializationWriter
from typing import Any, Callable, Dict, List, Optional, TYPE_CHECKING, Union

if TYPE_CHECKING:
    from . import entity, message, message_rule, multi_value_legacy_extended_property, single_value_legacy_extended_property

from . import entity

class MailFolder(entity.Entity):
    def __init__(self,) -> None:
        """
        Instantiates a new mailFolder and sets the default values.
        """
        super().__init__()
        # The number of immediate child mailFolders in the current mailFolder.
        self._child_folder_count: Optional[int] = None
        # The collection of child folders in the mailFolder.
        self._child_folders: Optional[List[MailFolder]] = None
        # The mailFolder's display name.
        self._display_name: Optional[str] = None
        # Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
        self._is_hidden: Optional[bool] = None
        # The collection of rules that apply to the user's Inbox folder.
        self._message_rules: Optional[List[message_rule.MessageRule]] = None
        # The collection of messages in the mailFolder.
        self._messages: Optional[List[message.Message]] = None
        # The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
        self._multi_value_extended_properties: Optional[List[multi_value_legacy_extended_property.MultiValueLegacyExtendedProperty]] = None
        # The unique identifier for the mailFolder's parent mailFolder.
        self._parent_folder_id: Optional[str] = None
        # The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
        self._single_value_extended_properties: Optional[List[single_value_legacy_extended_property.SingleValueLegacyExtendedProperty]] = None
        # The number of items in the mailFolder.
        self._total_item_count: Optional[int] = None
        # The number of items in the mailFolder marked as unread.
        self._unread_item_count: Optional[int] = None
    
    @property
    def child_folder_count(self,) -> Optional[int]:
        """
        Gets the childFolderCount property value. The number of immediate child mailFolders in the current mailFolder.
        Returns: Optional[int]
        """
        return self._child_folder_count
    
    @child_folder_count.setter
    def child_folder_count(self,value: Optional[int] = None) -> None:
        """
        Sets the childFolderCount property value. The number of immediate child mailFolders in the current mailFolder.
        Args:
            value: Value to set for the child_folder_count property.
        """
        self._child_folder_count = value
    
    @property
    def child_folders(self,) -> Optional[List[MailFolder]]:
        """
        Gets the childFolders property value. The collection of child folders in the mailFolder.
        Returns: Optional[List[MailFolder]]
        """
        return self._child_folders
    
    @child_folders.setter
    def child_folders(self,value: Optional[List[MailFolder]] = None) -> None:
        """
        Sets the childFolders property value. The collection of child folders in the mailFolder.
        Args:
            value: Value to set for the child_folders property.
        """
        self._child_folders = value
    
    @staticmethod
    def create_from_discriminator_value(parse_node: Optional[ParseNode] = None) -> MailFolder:
        """
        Creates a new instance of the appropriate class based on discriminator value
        Args:
            parseNode: The parse node to use to read the discriminator value and create the object
        Returns: MailFolder
        """
        if parse_node is None:
            raise Exception("parse_node cannot be undefined")
        return MailFolder()
    
    @property
    def display_name(self,) -> Optional[str]:
        """
        Gets the displayName property value. The mailFolder's display name.
        Returns: Optional[str]
        """
        return self._display_name
    
    @display_name.setter
    def display_name(self,value: Optional[str] = None) -> None:
        """
        Sets the displayName property value. The mailFolder's display name.
        Args:
            value: Value to set for the display_name property.
        """
        self._display_name = value
    
    def get_field_deserializers(self,) -> Dict[str, Callable[[ParseNode], None]]:
        """
        The deserialization information for the current model
        Returns: Dict[str, Callable[[ParseNode], None]]
        """
        from . import entity, message, message_rule, multi_value_legacy_extended_property, single_value_legacy_extended_property

        fields: Dict[str, Callable[[Any], None]] = {
            "childFolders": lambda n : setattr(self, 'child_folders', n.get_collection_of_object_values(MailFolder)),
            "childFolderCount": lambda n : setattr(self, 'child_folder_count', n.get_int_value()),
            "displayName": lambda n : setattr(self, 'display_name', n.get_str_value()),
            "isHidden": lambda n : setattr(self, 'is_hidden', n.get_bool_value()),
            "messages": lambda n : setattr(self, 'messages', n.get_collection_of_object_values(message.Message)),
            "messageRules": lambda n : setattr(self, 'message_rules', n.get_collection_of_object_values(message_rule.MessageRule)),
            "multiValueExtendedProperties": lambda n : setattr(self, 'multi_value_extended_properties', n.get_collection_of_object_values(multi_value_legacy_extended_property.MultiValueLegacyExtendedProperty)),
            "parentFolderId": lambda n : setattr(self, 'parent_folder_id', n.get_str_value()),
            "singleValueExtendedProperties": lambda n : setattr(self, 'single_value_extended_properties', n.get_collection_of_object_values(single_value_legacy_extended_property.SingleValueLegacyExtendedProperty)),
            "totalItemCount": lambda n : setattr(self, 'total_item_count', n.get_int_value()),
            "unreadItemCount": lambda n : setattr(self, 'unread_item_count', n.get_int_value()),
        }
        super_fields = super().get_field_deserializers()
        fields.update(super_fields)
        return fields
    
    @property
    def is_hidden(self,) -> Optional[bool]:
        """
        Gets the isHidden property value. Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
        Returns: Optional[bool]
        """
        return self._is_hidden
    
    @is_hidden.setter
    def is_hidden(self,value: Optional[bool] = None) -> None:
        """
        Sets the isHidden property value. Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
        Args:
            value: Value to set for the is_hidden property.
        """
        self._is_hidden = value
    
    @property
    def message_rules(self,) -> Optional[List[message_rule.MessageRule]]:
        """
        Gets the messageRules property value. The collection of rules that apply to the user's Inbox folder.
        Returns: Optional[List[message_rule.MessageRule]]
        """
        return self._message_rules
    
    @message_rules.setter
    def message_rules(self,value: Optional[List[message_rule.MessageRule]] = None) -> None:
        """
        Sets the messageRules property value. The collection of rules that apply to the user's Inbox folder.
        Args:
            value: Value to set for the message_rules property.
        """
        self._message_rules = value
    
    @property
    def messages(self,) -> Optional[List[message.Message]]:
        """
        Gets the messages property value. The collection of messages in the mailFolder.
        Returns: Optional[List[message.Message]]
        """
        return self._messages
    
    @messages.setter
    def messages(self,value: Optional[List[message.Message]] = None) -> None:
        """
        Sets the messages property value. The collection of messages in the mailFolder.
        Args:
            value: Value to set for the messages property.
        """
        self._messages = value
    
    @property
    def multi_value_extended_properties(self,) -> Optional[List[multi_value_legacy_extended_property.MultiValueLegacyExtendedProperty]]:
        """
        Gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
        Returns: Optional[List[multi_value_legacy_extended_property.MultiValueLegacyExtendedProperty]]
        """
        return self._multi_value_extended_properties
    
    @multi_value_extended_properties.setter
    def multi_value_extended_properties(self,value: Optional[List[multi_value_legacy_extended_property.MultiValueLegacyExtendedProperty]] = None) -> None:
        """
        Sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
        Args:
            value: Value to set for the multi_value_extended_properties property.
        """
        self._multi_value_extended_properties = value
    
    @property
    def parent_folder_id(self,) -> Optional[str]:
        """
        Gets the parentFolderId property value. The unique identifier for the mailFolder's parent mailFolder.
        Returns: Optional[str]
        """
        return self._parent_folder_id
    
    @parent_folder_id.setter
    def parent_folder_id(self,value: Optional[str] = None) -> None:
        """
        Sets the parentFolderId property value. The unique identifier for the mailFolder's parent mailFolder.
        Args:
            value: Value to set for the parent_folder_id property.
        """
        self._parent_folder_id = value
    
    def serialize(self,writer: SerializationWriter) -> None:
        """
        Serializes information the current object
        Args:
            writer: Serialization writer to use to serialize this model
        """
        if writer is None:
            raise Exception("writer cannot be undefined")
        super().serialize(writer)
        writer.write_collection_of_object_values("childFolders", self.child_folders)
        writer.write_int_value("childFolderCount", self.child_folder_count)
        writer.write_str_value("displayName", self.display_name)
        writer.write_bool_value("isHidden", self.is_hidden)
        writer.write_collection_of_object_values("messages", self.messages)
        writer.write_collection_of_object_values("messageRules", self.message_rules)
        writer.write_collection_of_object_values("multiValueExtendedProperties", self.multi_value_extended_properties)
        writer.write_str_value("parentFolderId", self.parent_folder_id)
        writer.write_collection_of_object_values("singleValueExtendedProperties", self.single_value_extended_properties)
        writer.write_int_value("totalItemCount", self.total_item_count)
        writer.write_int_value("unreadItemCount", self.unread_item_count)
    
    @property
    def single_value_extended_properties(self,) -> Optional[List[single_value_legacy_extended_property.SingleValueLegacyExtendedProperty]]:
        """
        Gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
        Returns: Optional[List[single_value_legacy_extended_property.SingleValueLegacyExtendedProperty]]
        """
        return self._single_value_extended_properties
    
    @single_value_extended_properties.setter
    def single_value_extended_properties(self,value: Optional[List[single_value_legacy_extended_property.SingleValueLegacyExtendedProperty]] = None) -> None:
        """
        Sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
        Args:
            value: Value to set for the single_value_extended_properties property.
        """
        self._single_value_extended_properties = value
    
    @property
    def total_item_count(self,) -> Optional[int]:
        """
        Gets the totalItemCount property value. The number of items in the mailFolder.
        Returns: Optional[int]
        """
        return self._total_item_count
    
    @total_item_count.setter
    def total_item_count(self,value: Optional[int] = None) -> None:
        """
        Sets the totalItemCount property value. The number of items in the mailFolder.
        Args:
            value: Value to set for the total_item_count property.
        """
        self._total_item_count = value
    
    @property
    def unread_item_count(self,) -> Optional[int]:
        """
        Gets the unreadItemCount property value. The number of items in the mailFolder marked as unread.
        Returns: Optional[int]
        """
        return self._unread_item_count
    
    @unread_item_count.setter
    def unread_item_count(self,value: Optional[int] = None) -> None:
        """
        Sets the unreadItemCount property value. The number of items in the mailFolder marked as unread.
        Args:
            value: Value to set for the unread_item_count property.
        """
        self._unread_item_count = value
    


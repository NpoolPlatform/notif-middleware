// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AnnouncementsColumns holds the columns for the "announcements" table.
	AnnouncementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "lang_id", Type: field.TypeUUID, Nullable: true},
		{Name: "title", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "content", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "channel", Type: field.TypeString, Nullable: true, Default: "DefaultChannel"},
		{Name: "start_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "end_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "type", Type: field.TypeString, Nullable: true, Default: "DefaultNotifType"},
	}
	// AnnouncementsTable holds the schema information for the "announcements" table.
	AnnouncementsTable = &schema.Table{
		Name:       "announcements",
		Columns:    AnnouncementsColumns,
		PrimaryKey: []*schema.Column{AnnouncementsColumns[0]},
	}
	// ContactsColumns holds the columns for the "contacts" table.
	ContactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "used_for", Type: field.TypeString, Nullable: true, Default: "DefaultUsedFor"},
		{Name: "sender", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "account", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "account_type", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// ContactsTable holds the schema information for the "contacts" table.
	ContactsTable = &schema.Table{
		Name:       "contacts",
		Columns:    ContactsColumns,
		PrimaryKey: []*schema.Column{ContactsColumns[0]},
	}
	// EmailTemplatesColumns holds the columns for the "email_templates" table.
	EmailTemplatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "lang_id", Type: field.TypeUUID},
		{Name: "default_to_username", Type: field.TypeString},
		{Name: "used_for", Type: field.TypeString, Nullable: true, Default: "DefaultUsedFor"},
		{Name: "sender", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "reply_tos", Type: field.TypeJSON, Nullable: true},
		{Name: "cc_tos", Type: field.TypeJSON, Nullable: true},
		{Name: "subject", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "body", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
	}
	// EmailTemplatesTable holds the schema information for the "email_templates" table.
	EmailTemplatesTable = &schema.Table{
		Name:       "email_templates",
		Columns:    EmailTemplatesColumns,
		PrimaryKey: []*schema.Column{EmailTemplatesColumns[0]},
	}
	// FrontendTemplatesColumns holds the columns for the "frontend_templates" table.
	FrontendTemplatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "lang_id", Type: field.TypeUUID},
		{Name: "used_for", Type: field.TypeString, Nullable: true, Default: "DefaultUsedFor"},
		{Name: "title", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "content", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
	}
	// FrontendTemplatesTable holds the schema information for the "frontend_templates" table.
	FrontendTemplatesTable = &schema.Table{
		Name:       "frontend_templates",
		Columns:    FrontendTemplatesColumns,
		PrimaryKey: []*schema.Column{FrontendTemplatesColumns[0]},
	}
	// NotifsColumns holds the columns for the "notifs" table.
	NotifsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "notified", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "lang_id", Type: field.TypeUUID, Nullable: true},
		{Name: "event_id", Type: field.TypeUUID, Nullable: true},
		{Name: "event_type", Type: field.TypeString, Nullable: true, Default: "DefaultUsedFor"},
		{Name: "use_template", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "title", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "content", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "channel", Type: field.TypeString, Nullable: true, Default: "DefaultChannel"},
		{Name: "extra", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "type", Type: field.TypeString, Nullable: true, Default: "DefaultNotifType"},
	}
	// NotifsTable holds the schema information for the "notifs" table.
	NotifsTable = &schema.Table{
		Name:       "notifs",
		Columns:    NotifsColumns,
		PrimaryKey: []*schema.Column{NotifsColumns[0]},
	}
	// NotifChannelsColumns holds the columns for the "notif_channels" table.
	NotifChannelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "event_type", Type: field.TypeString, Nullable: true, Default: "DefaultUsedFor"},
		{Name: "channel", Type: field.TypeString, Nullable: true, Default: "DefaultChannel"},
	}
	// NotifChannelsTable holds the schema information for the "notif_channels" table.
	NotifChannelsTable = &schema.Table{
		Name:       "notif_channels",
		Columns:    NotifChannelsColumns,
		PrimaryKey: []*schema.Column{NotifChannelsColumns[0]},
	}
	// ReadAnnouncementsColumns holds the columns for the "read_announcements" table.
	ReadAnnouncementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "announcement_id", Type: field.TypeUUID, Nullable: true},
	}
	// ReadAnnouncementsTable holds the schema information for the "read_announcements" table.
	ReadAnnouncementsTable = &schema.Table{
		Name:       "read_announcements",
		Columns:    ReadAnnouncementsColumns,
		PrimaryKey: []*schema.Column{ReadAnnouncementsColumns[0]},
	}
	// SmsTemplatesColumns holds the columns for the "sms_templates" table.
	SmsTemplatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "lang_id", Type: field.TypeUUID},
		{Name: "used_for", Type: field.TypeString, Nullable: true, Default: "DefaultUsedFor"},
		{Name: "subject", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "message", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// SmsTemplatesTable holds the schema information for the "sms_templates" table.
	SmsTemplatesTable = &schema.Table{
		Name:       "sms_templates",
		Columns:    SmsTemplatesColumns,
		PrimaryKey: []*schema.Column{SmsTemplatesColumns[0]},
	}
	// SendAnnouncementsColumns holds the columns for the "send_announcements" table.
	SendAnnouncementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "announcement_id", Type: field.TypeUUID, Nullable: true},
		{Name: "channel", Type: field.TypeString, Nullable: true, Default: "DefaultChannel"},
	}
	// SendAnnouncementsTable holds the schema information for the "send_announcements" table.
	SendAnnouncementsTable = &schema.Table{
		Name:       "send_announcements",
		Columns:    SendAnnouncementsColumns,
		PrimaryKey: []*schema.Column{SendAnnouncementsColumns[0]},
	}
	// TxNotifStatesColumns holds the columns for the "tx_notif_states" table.
	TxNotifStatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "tx_id", Type: field.TypeUUID, Nullable: true},
		{Name: "notif_state", Type: field.TypeString, Nullable: true, Default: "DefaultState"},
		{Name: "tx_type", Type: field.TypeString, Nullable: true, Default: "DefaultTxType"},
	}
	// TxNotifStatesTable holds the schema information for the "tx_notif_states" table.
	TxNotifStatesTable = &schema.Table{
		Name:       "tx_notif_states",
		Columns:    TxNotifStatesColumns,
		PrimaryKey: []*schema.Column{TxNotifStatesColumns[0]},
	}
	// UserAnnouncementsColumns holds the columns for the "user_announcements" table.
	UserAnnouncementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "announcement_id", Type: field.TypeUUID, Nullable: true},
	}
	// UserAnnouncementsTable holds the schema information for the "user_announcements" table.
	UserAnnouncementsTable = &schema.Table{
		Name:       "user_announcements",
		Columns:    UserAnnouncementsColumns,
		PrimaryKey: []*schema.Column{UserAnnouncementsColumns[0]},
	}
	// UserNotifsColumns holds the columns for the "user_notifs" table.
	UserNotifsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "event_type", Type: field.TypeString, Nullable: true, Default: "DefaultUsedFor"},
	}
	// UserNotifsTable holds the schema information for the "user_notifs" table.
	UserNotifsTable = &schema.Table{
		Name:       "user_notifs",
		Columns:    UserNotifsColumns,
		PrimaryKey: []*schema.Column{UserNotifsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AnnouncementsTable,
		ContactsTable,
		EmailTemplatesTable,
		FrontendTemplatesTable,
		NotifsTable,
		NotifChannelsTable,
		ReadAnnouncementsTable,
		SmsTemplatesTable,
		SendAnnouncementsTable,
		TxNotifStatesTable,
		UserAnnouncementsTable,
		UserNotifsTable,
	}
)

func init() {
}

// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package generated

// Entries defines model for Entries.
type Entries = []Entry

// Entry defines model for Entry.
type Entry struct {
	// Id エントリID
	Id EntryID `json:"id"`

	// IsDone エントリが既読かどうか。trueが、falseが既読を表す
	IsDone EntryIsDone `json:"is_done"`

	// Url エントリのURL
	Url EntryURL `json:"url"`
}

// EntryID エントリID
type EntryID = int

// EntryIsDone エントリが既読かどうか。trueが、falseが既読を表す
type EntryIsDone = bool

// EntryURL エントリのURL
type EntryURL = string

// Env 実行モード
type Env = string

// MessageResponse defines model for MessageResponse.
type MessageResponse struct {
	// Code HTTP status code
	Code int `json:"code"`

	// Message HTTP status message
	Message string `json:"message"`
}

// Root defines model for Root.
type Root struct {
	// Env 実行モード
	Env Env `json:"env"`

	// Status サーバステータス。固定
	Status Status `json:"status"`

	// Version APIバージョン
	Version Version `json:"version"`
}

// Status サーバステータス。固定
type Status = string

// Version APIバージョン
type Version = string

// EntryIdParam エントリID
type EntryIdParam = EntryID

// IgnoreIds defines model for ignoreIds.
type IgnoreIds = []int

// Size defines model for size.
type Size = int

// N500InternalServerError defines model for 500InternalServerError.
type N500InternalServerError = MessageResponse

// RespEntries defines model for RespEntries.
type RespEntries = Entries

// RespEntry defines model for RespEntry.
type RespEntry = Entry

// RespRoot defines model for RespRoot.
type RespRoot = Root

// GetEntriesParams defines parameters for GetEntries.
type GetEntriesParams struct {
	// Size 取得件数
	Size *Size `form:"size,omitempty" json:"size,omitempty"`

	// IgnoreIds 取得結果から排除するID。取得件数は排除した結果の件数とする
	IgnoreIds *IgnoreIds `form:"ignore_ids,omitempty" json:"ignore_ids,omitempty"`
}

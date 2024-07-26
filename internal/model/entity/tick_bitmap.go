// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TickBitmap is the golang structure for table tick_bitmap.
type TickBitmap struct {
	PoolId       int64  `json:"pool_id"       orm:"pool_id"       ` //
	WordPosition int    `json:"word_position" orm:"word_position" ` //
	Bitmap       string `json:"bitmap"        orm:"bitmap"        ` //
}

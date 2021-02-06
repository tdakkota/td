// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// StatsMegagroupStats represents TL type `stats.megagroupStats#ef7ff916`.
// Supergroup statistics¹
//
// Links:
//  1) https://core.telegram.org/api/stats
//
// See https://core.telegram.org/constructor/stats.megagroupStats for reference.
type StatsMegagroupStats struct {
	// Period in consideration
	Period StatsDateRangeDays
	// Member count change for period in consideration
	Members StatsAbsValueAndPrev
	// Message number change for period in consideration
	Messages StatsAbsValueAndPrev
	// Number of users that viewed messages, for range in consideration
	Viewers StatsAbsValueAndPrev
	// Number of users that posted messages, for range in consideration
	Posters StatsAbsValueAndPrev
	// Supergroup growth graph (absolute subscriber count)
	GrowthGraph StatsGraphClass
	// Members growth (relative subscriber count)
	MembersGraph StatsGraphClass
	// New members by source graph
	NewMembersBySourceGraph StatsGraphClass
	// Subscriber language graph (piechart)
	LanguagesGraph StatsGraphClass
	// Message activity graph (stacked bar graph, message type)
	MessagesGraph StatsGraphClass
	// Group activity graph (deleted, modified messages, blocked users)
	ActionsGraph StatsGraphClass
	// Activity per hour graph (absolute)
	TopHoursGraph StatsGraphClass
	// Activity per day of week graph (absolute)
	WeekdaysGraph StatsGraphClass
	// Info about most active group members
	TopPosters []StatsGroupTopPoster
	// Info about most active group admins
	TopAdmins []StatsGroupTopAdmin
	// Info about most active group inviters
	TopInviters []StatsGroupTopInviter
	// Info about users mentioned in statistics
	Users []UserClass
}

// StatsMegagroupStatsTypeID is TL type id of StatsMegagroupStats.
const StatsMegagroupStatsTypeID = 0xef7ff916

func (m *StatsMegagroupStats) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Period.Zero()) {
		return false
	}
	if !(m.Members.Zero()) {
		return false
	}
	if !(m.Messages.Zero()) {
		return false
	}
	if !(m.Viewers.Zero()) {
		return false
	}
	if !(m.Posters.Zero()) {
		return false
	}
	if !(m.GrowthGraph == nil) {
		return false
	}
	if !(m.MembersGraph == nil) {
		return false
	}
	if !(m.NewMembersBySourceGraph == nil) {
		return false
	}
	if !(m.LanguagesGraph == nil) {
		return false
	}
	if !(m.MessagesGraph == nil) {
		return false
	}
	if !(m.ActionsGraph == nil) {
		return false
	}
	if !(m.TopHoursGraph == nil) {
		return false
	}
	if !(m.WeekdaysGraph == nil) {
		return false
	}
	if !(m.TopPosters == nil) {
		return false
	}
	if !(m.TopAdmins == nil) {
		return false
	}
	if !(m.TopInviters == nil) {
		return false
	}
	if !(m.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *StatsMegagroupStats) String() string {
	if m == nil {
		return "StatsMegagroupStats(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StatsMegagroupStats")
	sb.WriteString("{\n")
	sb.WriteString("\tPeriod: ")
	sb.WriteString(fmt.Sprint(m.Period))
	sb.WriteString(",\n")
	sb.WriteString("\tMembers: ")
	sb.WriteString(fmt.Sprint(m.Members))
	sb.WriteString(",\n")
	sb.WriteString("\tMessages: ")
	sb.WriteString(fmt.Sprint(m.Messages))
	sb.WriteString(",\n")
	sb.WriteString("\tViewers: ")
	sb.WriteString(fmt.Sprint(m.Viewers))
	sb.WriteString(",\n")
	sb.WriteString("\tPosters: ")
	sb.WriteString(fmt.Sprint(m.Posters))
	sb.WriteString(",\n")
	sb.WriteString("\tGrowthGraph: ")
	sb.WriteString(fmt.Sprint(m.GrowthGraph))
	sb.WriteString(",\n")
	sb.WriteString("\tMembersGraph: ")
	sb.WriteString(fmt.Sprint(m.MembersGraph))
	sb.WriteString(",\n")
	sb.WriteString("\tNewMembersBySourceGraph: ")
	sb.WriteString(fmt.Sprint(m.NewMembersBySourceGraph))
	sb.WriteString(",\n")
	sb.WriteString("\tLanguagesGraph: ")
	sb.WriteString(fmt.Sprint(m.LanguagesGraph))
	sb.WriteString(",\n")
	sb.WriteString("\tMessagesGraph: ")
	sb.WriteString(fmt.Sprint(m.MessagesGraph))
	sb.WriteString(",\n")
	sb.WriteString("\tActionsGraph: ")
	sb.WriteString(fmt.Sprint(m.ActionsGraph))
	sb.WriteString(",\n")
	sb.WriteString("\tTopHoursGraph: ")
	sb.WriteString(fmt.Sprint(m.TopHoursGraph))
	sb.WriteString(",\n")
	sb.WriteString("\tWeekdaysGraph: ")
	sb.WriteString(fmt.Sprint(m.WeekdaysGraph))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range m.TopPosters {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range m.TopAdmins {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range m.TopInviters {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range m.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (m *StatsMegagroupStats) TypeID() uint32 {
	return StatsMegagroupStatsTypeID
}

// Encode implements bin.Encoder.
func (m *StatsMegagroupStats) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode stats.megagroupStats#ef7ff916 as nil")
	}
	b.PutID(StatsMegagroupStatsTypeID)
	if err := m.Period.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field period: %w", err)
	}
	if err := m.Members.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field members: %w", err)
	}
	if err := m.Messages.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field messages: %w", err)
	}
	if err := m.Viewers.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field viewers: %w", err)
	}
	if err := m.Posters.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field posters: %w", err)
	}
	if m.GrowthGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field growth_graph is nil")
	}
	if err := m.GrowthGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field growth_graph: %w", err)
	}
	if m.MembersGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field members_graph is nil")
	}
	if err := m.MembersGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field members_graph: %w", err)
	}
	if m.NewMembersBySourceGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field new_members_by_source_graph is nil")
	}
	if err := m.NewMembersBySourceGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field new_members_by_source_graph: %w", err)
	}
	if m.LanguagesGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field languages_graph is nil")
	}
	if err := m.LanguagesGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field languages_graph: %w", err)
	}
	if m.MessagesGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field messages_graph is nil")
	}
	if err := m.MessagesGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field messages_graph: %w", err)
	}
	if m.ActionsGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field actions_graph is nil")
	}
	if err := m.ActionsGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field actions_graph: %w", err)
	}
	if m.TopHoursGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_hours_graph is nil")
	}
	if err := m.TopHoursGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_hours_graph: %w", err)
	}
	if m.WeekdaysGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field weekdays_graph is nil")
	}
	if err := m.WeekdaysGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field weekdays_graph: %w", err)
	}
	b.PutVectorHeader(len(m.TopPosters))
	for idx, v := range m.TopPosters {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_posters element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.TopAdmins))
	for idx, v := range m.TopAdmins {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_admins element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.TopInviters))
	for idx, v := range m.TopInviters {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_inviters element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.Users))
	for idx, v := range m.Users {
		if v == nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetPeriod returns value of Period field.
func (m *StatsMegagroupStats) GetPeriod() (value StatsDateRangeDays) {
	return m.Period
}

// GetMembers returns value of Members field.
func (m *StatsMegagroupStats) GetMembers() (value StatsAbsValueAndPrev) {
	return m.Members
}

// GetMessages returns value of Messages field.
func (m *StatsMegagroupStats) GetMessages() (value StatsAbsValueAndPrev) {
	return m.Messages
}

// GetViewers returns value of Viewers field.
func (m *StatsMegagroupStats) GetViewers() (value StatsAbsValueAndPrev) {
	return m.Viewers
}

// GetPosters returns value of Posters field.
func (m *StatsMegagroupStats) GetPosters() (value StatsAbsValueAndPrev) {
	return m.Posters
}

// GetGrowthGraph returns value of GrowthGraph field.
func (m *StatsMegagroupStats) GetGrowthGraph() (value StatsGraphClass) {
	return m.GrowthGraph
}

// GetMembersGraph returns value of MembersGraph field.
func (m *StatsMegagroupStats) GetMembersGraph() (value StatsGraphClass) {
	return m.MembersGraph
}

// GetNewMembersBySourceGraph returns value of NewMembersBySourceGraph field.
func (m *StatsMegagroupStats) GetNewMembersBySourceGraph() (value StatsGraphClass) {
	return m.NewMembersBySourceGraph
}

// GetLanguagesGraph returns value of LanguagesGraph field.
func (m *StatsMegagroupStats) GetLanguagesGraph() (value StatsGraphClass) {
	return m.LanguagesGraph
}

// GetMessagesGraph returns value of MessagesGraph field.
func (m *StatsMegagroupStats) GetMessagesGraph() (value StatsGraphClass) {
	return m.MessagesGraph
}

// GetActionsGraph returns value of ActionsGraph field.
func (m *StatsMegagroupStats) GetActionsGraph() (value StatsGraphClass) {
	return m.ActionsGraph
}

// GetTopHoursGraph returns value of TopHoursGraph field.
func (m *StatsMegagroupStats) GetTopHoursGraph() (value StatsGraphClass) {
	return m.TopHoursGraph
}

// GetWeekdaysGraph returns value of WeekdaysGraph field.
func (m *StatsMegagroupStats) GetWeekdaysGraph() (value StatsGraphClass) {
	return m.WeekdaysGraph
}

// GetTopPosters returns value of TopPosters field.
func (m *StatsMegagroupStats) GetTopPosters() (value []StatsGroupTopPoster) {
	return m.TopPosters
}

// GetTopAdmins returns value of TopAdmins field.
func (m *StatsMegagroupStats) GetTopAdmins() (value []StatsGroupTopAdmin) {
	return m.TopAdmins
}

// GetTopInviters returns value of TopInviters field.
func (m *StatsMegagroupStats) GetTopInviters() (value []StatsGroupTopInviter) {
	return m.TopInviters
}

// GetUsers returns value of Users field.
func (m *StatsMegagroupStats) GetUsers() (value []UserClass) {
	return m.Users
}

// Decode implements bin.Decoder.
func (m *StatsMegagroupStats) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode stats.megagroupStats#ef7ff916 to nil")
	}
	if err := b.ConsumeID(StatsMegagroupStatsTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: %w", err)
	}
	{
		if err := m.Period.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field period: %w", err)
		}
	}
	{
		if err := m.Members.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field members: %w", err)
		}
	}
	{
		if err := m.Messages.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field messages: %w", err)
		}
	}
	{
		if err := m.Viewers.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field viewers: %w", err)
		}
	}
	{
		if err := m.Posters.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field posters: %w", err)
		}
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field growth_graph: %w", err)
		}
		m.GrowthGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field members_graph: %w", err)
		}
		m.MembersGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field new_members_by_source_graph: %w", err)
		}
		m.NewMembersBySourceGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field languages_graph: %w", err)
		}
		m.LanguagesGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field messages_graph: %w", err)
		}
		m.MessagesGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field actions_graph: %w", err)
		}
		m.ActionsGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_hours_graph: %w", err)
		}
		m.TopHoursGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field weekdays_graph: %w", err)
		}
		m.WeekdaysGraph = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_posters: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StatsGroupTopPoster
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_posters: %w", err)
			}
			m.TopPosters = append(m.TopPosters, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_admins: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StatsGroupTopAdmin
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_admins: %w", err)
			}
			m.TopAdmins = append(m.TopAdmins, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_inviters: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StatsGroupTopInviter
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_inviters: %w", err)
			}
			m.TopInviters = append(m.TopInviters, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field users: %w", err)
			}
			m.Users = append(m.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsMegagroupStats.
var (
	_ bin.Encoder = &StatsMegagroupStats{}
	_ bin.Decoder = &StatsMegagroupStats{}
)

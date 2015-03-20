// Code generated by protoc-gen-go.
// source: kythe/proto/serving.proto
// DO NOT EDIT!

/*
Package serving_proto is a generated protocol buffer package.

It is generated from these files:
	kythe/proto/serving.proto

It has these top-level messages:
	Node
	EdgeSet
	PagedEdgeSet
	PageIndex
	EdgePage
	FileDirectory
	CorpusRoots
	FileDecorations
	PageToken
*/
package serving_proto

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// A derivative of xref.NodeInfo for serving.
type Node struct {
	Ticket string       `protobuf:"bytes,1,opt,name=ticket" json:"ticket,omitempty"`
	Fact   []*Node_Fact `protobuf:"bytes,2,rep,name=fact" json:"fact,omitempty"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}

func (m *Node) GetFact() []*Node_Fact {
	if m != nil {
		return m.Fact
	}
	return nil
}

type Node_Fact struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Node_Fact) Reset()         { *m = Node_Fact{} }
func (m *Node_Fact) String() string { return proto.CompactTextString(m) }
func (*Node_Fact) ProtoMessage()    {}

// An EdgeSet represents a collection of edges from a single node.  The edges
// are organized into groups, each sharing a common edge kind.
//
// The number of edges represented by an EdgeSet es, denoted len(es), is the
// sum of the lengths of the repeated target_ticket fields for all the groups
// in the EdgeSet.  This count is used to determine page size in a request.
//
// Note: this is a derivative of xref.EdgeSet
type EdgeSet struct {
	// The ticket of the source node for all the edges in the edge set.
	SourceTicket string `protobuf:"bytes,1,opt,name=source_ticket" json:"source_ticket,omitempty"`
	// Each group is a collection of outbound edges from source node sharing a
	// given kind.  In a given EdgeSet, the server will not send more than one
	// group with the same kind label.
	Group []*EdgeSet_Group `protobuf:"bytes,2,rep,name=group" json:"group,omitempty"`
}

func (m *EdgeSet) Reset()         { *m = EdgeSet{} }
func (m *EdgeSet) String() string { return proto.CompactTextString(m) }
func (*EdgeSet) ProtoMessage()    {}

func (m *EdgeSet) GetGroup() []*EdgeSet_Group {
	if m != nil {
		return m.Group
	}
	return nil
}

type EdgeSet_Group struct {
	Kind         string   `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	TargetTicket []string `protobuf:"bytes,2,rep,name=target_ticket" json:"target_ticket,omitempty"`
}

func (m *EdgeSet_Group) Reset()         { *m = EdgeSet_Group{} }
func (m *EdgeSet_Group) String() string { return proto.CompactTextString(m) }
func (*EdgeSet_Group) ProtoMessage()    {}

// PagedEdgeSets are used for efficiently storing EdgeSets, all originating from
// the same source ticket, in order to handle pagination requests.
type PagedEdgeSet struct {
	// Collection of edges on the first page.  If the number of edges for a
	// particular source_ticket is small, this may contain all known edges and
	// no page_index will exist.
	EdgeSet *EdgeSet `protobuf:"bytes,1,opt,name=edge_set" json:"edge_set,omitempty"`
	// Total number of edges in all of the EdgePages, including this one.
	TotalEdges int32 `protobuf:"varint,2,opt,name=total_edges" json:"total_edges,omitempty"`
	// Page indices for other EdgePages, sorted by edge kind.
	PageIndex []*PageIndex `protobuf:"bytes,3,rep,name=page_index" json:"page_index,omitempty"`
}

func (m *PagedEdgeSet) Reset()         { *m = PagedEdgeSet{} }
func (m *PagedEdgeSet) String() string { return proto.CompactTextString(m) }
func (*PagedEdgeSet) ProtoMessage()    {}

func (m *PagedEdgeSet) GetEdgeSet() *EdgeSet {
	if m != nil {
		return m.EdgeSet
	}
	return nil
}

func (m *PagedEdgeSet) GetPageIndex() []*PageIndex {
	if m != nil {
		return m.PageIndex
	}
	return nil
}

// PageIndex is a pointer to an EdgePage.  In order to keep the PagedEdgeSet
// small, we don't store edges here.  We just store a key for looking up an
// EdgePage and the type of edge.
type PageIndex struct {
	// The kind of all edges on the referred EdgePage.
	EdgeKind string `protobuf:"bytes,1,opt,name=edge_kind" json:"edge_kind,omitempty"`
	// Total number of edges on the referred EdgePage.
	EdgeCount int32 `protobuf:"varint,2,opt,name=edge_count" json:"edge_count,omitempty"`
	// Key that can be used to lookup the referred EdgePage.
	PageKey string `protobuf:"bytes,3,opt,name=page_key" json:"page_key,omitempty"`
}

func (m *PageIndex) Reset()         { *m = PageIndex{} }
func (m *PageIndex) String() string { return proto.CompactTextString(m) }
func (*PageIndex) ProtoMessage()    {}

// EdgePages are a group of edges for a particular edge kind and source ticket.
type EdgePage struct {
	// Corresponding PageIndex key that can be used to lookup this page.
	PageKey      string         `protobuf:"bytes,1,opt,name=page_key" json:"page_key,omitempty"`
	SourceTicket string         `protobuf:"bytes,2,opt,name=source_ticket" json:"source_ticket,omitempty"`
	EdgesGroup   *EdgeSet_Group `protobuf:"bytes,3,opt,name=edges_group" json:"edges_group,omitempty"`
}

func (m *EdgePage) Reset()         { *m = EdgePage{} }
func (m *EdgePage) String() string { return proto.CompactTextString(m) }
func (*EdgePage) ProtoMessage()    {}

func (m *EdgePage) GetEdgesGroup() *EdgeSet_Group {
	if m != nil {
		return m.EdgesGroup
	}
	return nil
}

// FileDirectory describes a virtual directory of file nodes.
type FileDirectory struct {
	// Set of URIs for each contained sub-directory's corpus, root, and full path.
	Subdirectory []string `protobuf:"bytes,1,rep,name=subdirectory" json:"subdirectory,omitempty"`
	// Set of file node tickets contained within this directory.
	FileTicket []string `protobuf:"bytes,2,rep,name=file_ticket" json:"file_ticket,omitempty"`
}

func (m *FileDirectory) Reset()         { *m = FileDirectory{} }
func (m *FileDirectory) String() string { return proto.CompactTextString(m) }
func (*FileDirectory) ProtoMessage()    {}

// CorpusRoots describes all of the known corpus/root pairs that contain file
// nodes.
type CorpusRoots struct {
	Corpus []*CorpusRoots_Corpus `protobuf:"bytes,1,rep,name=corpus" json:"corpus,omitempty"`
}

func (m *CorpusRoots) Reset()         { *m = CorpusRoots{} }
func (m *CorpusRoots) String() string { return proto.CompactTextString(m) }
func (*CorpusRoots) ProtoMessage()    {}

func (m *CorpusRoots) GetCorpus() []*CorpusRoots_Corpus {
	if m != nil {
		return m.Corpus
	}
	return nil
}

type CorpusRoots_Corpus struct {
	Corpus string   `protobuf:"bytes,1,opt,name=corpus" json:"corpus,omitempty"`
	Root   []string `protobuf:"bytes,2,rep,name=root" json:"root,omitempty"`
}

func (m *CorpusRoots_Corpus) Reset()         { *m = CorpusRoots_Corpus{} }
func (m *CorpusRoots_Corpus) String() string { return proto.CompactTextString(m) }
func (*CorpusRoots_Corpus) ProtoMessage()    {}

// FileDecorations stores a file's contents and all contained anchor edges.
type FileDecorations struct {
	FileTicket string `protobuf:"bytes,1,opt,name=file_ticket" json:"file_ticket,omitempty"`
	SourceText []byte `protobuf:"bytes,2,opt,name=source_text,proto3" json:"source_text,omitempty"`
	Encoding   string `protobuf:"bytes,3,opt,name=encoding" json:"encoding,omitempty"`
	// The decorations located in the file, sorted by starting offset.
	Decoration []*FileDecorations_Decoration `protobuf:"bytes,4,rep,name=decoration" json:"decoration,omitempty"`
}

func (m *FileDecorations) Reset()         { *m = FileDecorations{} }
func (m *FileDecorations) String() string { return proto.CompactTextString(m) }
func (*FileDecorations) ProtoMessage()    {}

func (m *FileDecorations) GetDecoration() []*FileDecorations_Decoration {
	if m != nil {
		return m.Decoration
	}
	return nil
}

// Represents an edge from an anchor contained within the file to some target.
type FileDecorations_Decoration struct {
	Anchor       *FileDecorations_Decoration_Anchor `protobuf:"bytes,1,opt,name=anchor" json:"anchor,omitempty"`
	TargetTicket string                             `protobuf:"bytes,2,opt,name=target_ticket" json:"target_ticket,omitempty"`
	Kind         string                             `protobuf:"bytes,3,opt,name=kind" json:"kind,omitempty"`
}

func (m *FileDecorations_Decoration) Reset()         { *m = FileDecorations_Decoration{} }
func (m *FileDecorations_Decoration) String() string { return proto.CompactTextString(m) }
func (*FileDecorations_Decoration) ProtoMessage()    {}

func (m *FileDecorations_Decoration) GetAnchor() *FileDecorations_Decoration_Anchor {
	if m != nil {
		return m.Anchor
	}
	return nil
}

type FileDecorations_Decoration_Anchor struct {
	Ticket      string `protobuf:"bytes,1,opt,name=ticket" json:"ticket,omitempty"`
	StartOffset int32  `protobuf:"varint,2,opt,name=start_offset" json:"start_offset,omitempty"`
	EndOffset   int32  `protobuf:"varint,3,opt,name=end_offset" json:"end_offset,omitempty"`
}

func (m *FileDecorations_Decoration_Anchor) Reset()         { *m = FileDecorations_Decoration_Anchor{} }
func (m *FileDecorations_Decoration_Anchor) String() string { return proto.CompactTextString(m) }
func (*FileDecorations_Decoration_Anchor) ProtoMessage()    {}

// Internal encoding for an EdgesReply page_token
type PageToken struct {
	// Index into sequence of edges to return in EdgesReply.
	Index int32 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
}

func (m *PageToken) Reset()         { *m = PageToken{} }
func (m *PageToken) String() string { return proto.CompactTextString(m) }
func (*PageToken) ProtoMessage()    {}

func init() {
}
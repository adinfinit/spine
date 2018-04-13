package spine

import "fmt"

type Attachment interface {
	GetName() string
}

type RegionAttachment struct {
	Name  string
	Path  string
	Size  Vector
	Local Transform
	Color Color
}

func (attachment *RegionAttachment) GetName() string { return attachment.Name }

type PointAttachment struct {
	Name  string
	Local Transform
	Color Color
}

func (attachment *PointAttachment) GetName() string { return attachment.Name }

type MeshAttachment struct {
	Name         string
	Path         string
	Local        Transform
	Size         Vector
	Hull         int
	UV           []Vector
	Vertices     []Vertex
	Triangles    [][3]int
	Edges        [][2]int
	Color        Color
	Weighted     bool
	BindingCount int
}

func (attachment *MeshAttachment) GetName() string { return attachment.Name }

func (mesh *MeshAttachment) CalculateWorldVertices(skel *Skeleton, slot *Slot) []Vector {
	vertices := make([]Vector, len(mesh.Vertices))
	if !mesh.Weighted {
		final := slot.Bone.World.Mul(mesh.Local.Affine())
		if len(slot.Deform) == 0 {
			for i := range mesh.Vertices {
				p := mesh.Vertices[i].Position
				vertices[i] = final.Transform(p)
			}
		} else {
			if len(vertices) != len(slot.Deform) {
				fmt.Println(len(vertices), len(slot.Deform))
				panic("invalid deform")
			}
			for i := range mesh.Vertices {
				p := mesh.Vertices[i].Position
				vertices[i] = final.Transform(p.Add(slot.Deform[i]))
			}
		}
	} else {
		bones := skel.Bones
		if len(slot.Deform) == 0 {
			for i := range mesh.Vertices {
				v := &mesh.Vertices[i]
				var w Vector
				for _, binding := range v.Bindings {
					bone := bones[binding.Bone]
					p := binding.Position
					w = w.Add(bone.World.WeightedTransform(binding.Weight, p))
				}
				vertices[i] = w
			}
		} else {
			deform, di := slot.Deform, 0
			for i := range mesh.Vertices {
				v := &mesh.Vertices[i]
				var w Vector
				for _, binding := range v.Bindings {
					bone := bones[binding.Bone]
					p := binding.Position.Add(deform[di])
					di++
					w = w.Add(bone.World.WeightedTransform(binding.Weight, p))
				}
				vertices[i] = w
			}
		}
	}
	return vertices
}

type Vertex struct {
	Position Vector
	Bindings []VertexBinding
}

type VertexBinding struct {
	Bone     int
	Position Vector
	Weight   float32
}

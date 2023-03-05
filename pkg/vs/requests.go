package vs

type VirtualServiceUri struct {
	Name      string `uri:"name" binding:"required"`
	Namespace string `uri:"ns" binding:"required"`
}

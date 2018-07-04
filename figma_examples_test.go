package figma_test

import (
	"os"

	"github.com/tmc/figma"
)

func ExampleNewClient() {
	c, _ := figma.NewClient(os.Getenv("FIGMA_TOKEN"))
	projects, _ := c.GetProjectsForTeam(os.Getenv("FIGMA_TEAM_ID"))
	_ = projects
	// do something with projects, don't ignore errors.
}

func ExampleClient_GetFilesForProject() {
	c, _ := figma.NewClient(os.Getenv("FIGMA_TOKEN"))
	files, _ := c.GetFilesForProject(os.Getenv("FIGMA_PROJECT_ID"))
	_ = files
	// do something with projects, don't ignore errors.
}

func ExampleClient_GetFile() {
	c, _ := figma.NewClient(os.Getenv("FIGMA_TOKEN"))
	files, _ := c.GetFilesForProject(os.Getenv("FIGMA_PROJECT_ID"))
	_, _ = c.GetProjectsForTeam(os.Getenv("FIGMA_TEAM_ID"))
	for _, f := range files {
		c.GetFile(f.Key)
	}
	// don't ignore errors.
}

func ExampleClient_GetImage() {
	c, _ := figma.NewClient(os.Getenv("FIGMA_TOKEN"))
	i, _ := c.GetImage(os.Getenv("FIGMA_FILE_ID"), figma.ImageOptions{
		IDs:    figma.DefaultID,
		Format: "png",
		Scale:  1.0,
	})
	_ = i
	// i.Images has images URLs, keyed by ID.
	// don't ignore errors.
}

func ExampleClient_GetFileVersions() {
	c, _ := figma.NewClient(os.Getenv("FIGMA_TOKEN"))
	v, _ := c.GetFileVersions(os.Getenv("FIGMA_FILE_ID"))
	_ = v
	// don't ignore errors.
}

func ExampleClient_GetFileComments() {
	c, _ := figma.NewClient(os.Getenv("FIGMA_TOKEN"))
	comments, _ := c.GetFileComments(os.Getenv("FIGMA_FILE_ID"))
	_ = comments
	// don't ignore errors.
}

func ExampleClient_CreateFileComment() {
	c, _ := figma.NewClient(os.Getenv("FIGMA_TOKEN"))
	file, _ := c.GetFile(os.Getenv("FIGMA_FILE_ID"))
	comment, err := c.CreateFileComment(os.Getenv("FIGMA_FILE_ID"), figma.CreateCommentOptions{
		Message: "Beep Boop! I am a bot! Can you make it pop more? 🤖",
		ClientMeta: figma.VectorOrFrameOffset{
			NodeID:     file.Document.Children[0].GetID(),
			NodeOffset: &figma.Vector{X: 0, Y: 0},
		},
	})
	_, _ = comment, err
	// don't ignore errors.
}

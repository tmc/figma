package figma

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const defaultBaseURL = "https://api.figma.com/v1/"

// Client is the primary type that implements an interface to the Figma.com API.
type Client struct {
	client  *http.Client
	baseURL string
	token   string
}

// NewClient initializes a new Client.
func NewClient(token string, opts ...ClientOption) (*Client, error) {
	c := &Client{
		token:   token,
		baseURL: defaultBaseURL,
	}

	for _, o := range opts {
		o(c)
	}
	if c.client == nil {
		c.client = http.DefaultClient
	}
	return c, nil
}

func (c *Client) url(path string) string {
	return fmt.Sprintf("%s%s", c.baseURL, path)
}

func (c *Client) get(pattern string, args ...interface{}) ([]byte, error) {
	return c.do("GET", nil, pattern, args...)
}

func (c *Client) post(payload interface{}, pattern string, args ...interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(payload); err != nil {
		return nil, err
	}
	fmt.Println(string(buf.Bytes()))
	return c.do("POST", buf, pattern, args...)
}

func (c *Client) do(method string, body io.Reader, pattern string, args ...interface{}) ([]byte, error) {
	path := c.url(fmt.Sprintf(pattern, args...))
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}
	req.Header.Set("X-Figma-Token", c.token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "performing request")
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return buf, &Error{
			URL:        path,
			StatusCode: resp.StatusCode,
			Body:       string(buf),
		}
	}
	return buf, nil
}

type filesForProject struct {
	Files []FileMeta `json:"files,omitempty"`
}

// GetFilesForProject returns a list of FileMetas for a given project id.
func (c *Client) GetFilesForProject(projectID string) ([]FileMeta, error) {
	b, err := c.get("projects/%s/files", projectID)
	if err != nil {
		return nil, err
	}
	result := &filesForProject{}
	return result.Files, json.Unmarshal(b, result)
}

type projectsForTeam struct {
	Projects []Project `json:"projects,omitempty"`
}

// GetProjectsForTeam returns a list of Projects given a team id.
func (c *Client) GetProjectsForTeam(teamID string) ([]Project, error) {
	b, err := c.get("teams/%s/projects", teamID)
	if err != nil {
		return nil, err
	}
	result := &projectsForTeam{}
	return result.Projects, json.Unmarshal(b, result)
}

// GetFile returns details for a given file key.
func (c *Client) GetFile(fileKey string) (*File, error) {
	b, err := c.getFile(fileKey)
	if err != nil {
		return nil, err
	}
	result := &File{}
	return result, json.Unmarshal(b, result)
}

func (c *Client) getFile(fileKey string) ([]byte, error) {
	return c.get("files/%s", fileKey)
}

// FileOptions allows configuration of the Get File request.
type FileOptions struct {
	GeometryPaths bool
	Version       string
}

// GetFileWithOptions is similar to GetFile but allows more specific requests to be made.
func (c *Client) GetFileWithOptions(fileKey string, opts FileOptions) (*File, error) {
	b, err := c.getFileWithOptions(fileKey, opts)
	if err != nil {
		return nil, err
	}
	result := &File{}
	return result, json.Unmarshal(b, result)
}

func (c *Client) getFileWithOptions(fileKey string, opts FileOptions) ([]byte, error) {
	o := url.Values{}
	o.Set("version", opts.Version)
	if opts.GeometryPaths {
		o.Set("geometry", "paths")
	}
	return c.get("files/%s?%s", fileKey, o.Encode())
}

// ImageFormat encodes the possible values for an image.
type ImageFormat string

const (
	ImageFormatJPG ImageFormat = "jpg"
	ImageFormatPNG             = "png"
	ImageFormatSVG             = "svg"
)

// ImageOptions allows configuration of the Get Image request.
type ImageOptions struct {
	// A comma separated list of node IDs to render.
	IDs string `json:"ids"`
	// A number between 0.01 and 4, the image scaling factor.
	Scale float64 `json:"scale"`
	// A string enum for the image output format, can be "jpg", "png", or "svg".
	Format string `json:"format"`
	// Whether to include id attributes for all SVG elements. Default: false.
	SVGIncludeID bool `json:"svg_include_id,omitempty"`
	// Whether to skip simplifying inside/outside strokes and use stroke attribute if possible instead of <mask>. Default: false.
	SkipSVGSimplifyStroke bool `json:"svg_simplify_stroke,omitempty"`
	// A specific version ID to use. Omitting this will use the current version of the file.
	Version string `json:"version,omitempty"`
}

// GetImage gets an image from the Figma API.
func (c *Client) GetImage(fileKey string, opts ImageOptions) (*Image, error) {
	b, err := c.getImage(fileKey, opts)
	if err != nil {
		return nil, err
	}
	result := &Image{}
	return result, json.Unmarshal(b, result)
}

func (c *Client) getImage(fileKey string, opts ImageOptions) ([]byte, error) {
	o := url.Values{}
	o.Set("ids", opts.IDs)
	o.Set("scale", fmt.Sprint(opts.Scale))
	o.Set("version", opts.Version)
	o.Set("format", opts.Format)
	o.Set("svg_include_id", fmt.Sprint(opts.SVGIncludeID))
	if opts.SkipSVGSimplifyStroke {
		o.Set("svg_simplify_stroke", "false")
	}
	return c.get("images/%s?%s", fileKey, o.Encode())
}

// GetFileVersions returns a list of versions for a file.
func (c *Client) GetFileVersions(fileKey string) ([]Version, error) {
	b, err := c.getFileVersions(fileKey)
	if err != nil {
		return nil, err
	}
	result := struct {
		Versions []Version `json:"versions"`
	}{}
	return result.Versions, json.Unmarshal(b, &result)
}

func (c *Client) getFileVersions(fileKey string) ([]byte, error) {
	return c.get("files/%s/versions", fileKey)
}

// GetFileComments gets the list of comments associated with the given file.
func (c *Client) GetFileComments(fileKey string) ([]Comment, error) {
	b, err := c.getFileComments(fileKey)
	if err != nil {
		return nil, err
	}
	result := struct {
		Comments []Comment `json:"comments"`
	}{}
	return result.Comments, json.Unmarshal(b, &result)
}

func (c *Client) getFileComments(fileKey string) ([]byte, error) {
	return c.get("files/%s/comments", fileKey)
}

// CreateCommentOptions describes a comment to be made on a file.
type CreateCommentOptions struct {
	// The text contents of the comment to post.
	Message string `json:"message"`
	// The position of where to place the comment. This can either be an absolute canvas position or the relative position within a frame..
	ClientMeta VectorOrFrameOffset `json:"client_meta"`
}

// CreateFileComment creates a comment on a file.
func (c *Client) CreateFileComment(fileKey string, opts CreateCommentOptions) (*Comment, error) {
	b, err := c.postFileComment(fileKey, opts)
	if err != nil {
		return nil, err
	}
	result := struct {
		Comment *Comment `json:"comment"`
	}{}
	return result.Comment, json.Unmarshal(b, &result)
}

func (c *Client) postFileComment(fileKey string, opts CreateCommentOptions) ([]byte, error) {
	return c.post(opts, "files/%s/comments", fileKey)
}

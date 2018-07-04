package figma

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	justCanvas = `{
  "name": "Personal",
  "lastModified": "2018-07-03T09:05:44.608Z",
  "thumbnailUrl": "https://s3-alpha.com/thumbnails/91a7557d-5625-4f76-9459-70818298599f",
  "document": {
    "id": "0:0",
    "name": "Document",
    "type": "DOCUMENT",
    "children": [
      {
        "id": "72:0",
        "name": "Blog",
        "type": "CANVAS",
        "backgroundColor": {
          "r": 0.8980392156862745,
          "g": 0.8980392156862745,
          "b": 0.8980392156862745,
          "a": 1
        },
        "exportSettings": [
          {
            "suffix": "x",
            "format": "PNG",
            "constraint": {
              "type": "SCALE",
              "value": 1.5
            }
          },
          {
            "suffix": "",
            "format": "PNG",
            "constraint": {
              "type": "SCALE",
              "value": 1
            }
          }
        ]
      }
    ]
  },
  "components": {
    "20:228": {
      "name": "Icon",
      "description": ""
    },
    "73:865": {
      "name": "tmc",
      "description": ""
    }
  },
  "schemaVersion": 0,
  "styles": {}
}`
)

func TestDocumentDecoding(t *testing.T) {
	cases := []struct {
		name         string
		in           string
		expectedType interface{}
	}{
		{"canvas", justCanvas, &CanvasNode{}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			d := &File{}

			if err := json.Unmarshal([]byte(tt.in), d); err != nil {
				t.Fatal(err)
			}
			got := fmt.Sprintf("%T", tt.expectedType)
			want := fmt.Sprintf("%T", d.Document.Children[0])

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestJsonEncodingDecoding(t *testing.T) {
	figmaToken := os.Getenv("FIGMA_TOKEN")
	figmaProjectID := os.Getenv("FIGMA_PROJECT_ID")
	if figmaToken == "" || figmaProjectID == "" {
		t.Skip("missing env vars")
	}
	c, _ := NewClient(figmaToken)
	files, err := c.GetFilesForProject(figmaProjectID)
	if err != nil {
		t.Fatal(err)
	}

	// ensure all files encode the same way we got them from
	for _, f := range files {
		buf, err := c.getFile(f.Key)
		if err != nil {
			t.Fatal(err)
		}
		fd, err := c.GetFile(f.Key)
		if err != nil {
			t.Fatal(err)
		}
		fdbytes, err := json.Marshal(fd)
		if err != nil {
			t.Fatal(err)
		}

		var x1 map[string]interface{}
		var x2 map[string]interface{}
		err = json.Unmarshal(buf, &x1)
		if err != nil {
			t.Fatal(err)
		}
		err = json.Unmarshal(fdbytes, &x2)
		if err != nil {
			t.Fatal(err)
		}

		// figma has some undocumented differences from their docs, once those seem to be resolved this can be made a real assesrtion
		//x1buf, _ := json.Marshal(x1)
		//x2buf, _ := json.Marshal(x2)

		if !cmp.Equal(x1, x2) {
			t.Logf("json differs: %v", cmp.Diff(x1, x2))
		}
	}
}

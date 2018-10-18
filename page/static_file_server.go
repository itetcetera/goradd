package page

import (
	"github.com/spekary/goradd/util"
	"goradd-project/config"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"path"
)

var assetDirectories =  map[string]string{}


// RenderAssetTag will render a tag that points to a static file asset that should be served by the MUX. filePath points
// to the file on the development server, and it will be copied to the appropriate subdirectory in the local assets directory
// for easy deployment. tag is the tag label to put in the tag, and attributes are additional attributes to include in the tag.
// The copied location of the file and structure of the tag will be deduced from the type of tag and the label of the file.
// The type of tag will also be used to automatically insert the location of the file into the correct tag attribute.
/*
func RenderAssetTag(filePath string, tag string, attributes *html.Attributes, content string) string {
	var typ string

	_,fileName := filepath.Split(filePath)
	ext := path.Ext(fileName)


	switch ext {
	case "js": fallthrough
	case "javascript":
		typ = "js"
	case "css":
		typ = "css"
	case "jpg": fallthrough
	case "jpeg":fallthrough
	case "png": fallthrough
	case "gif": fallthrough
	case "bmp": fallthrough
	case "ico":
		typ = "img"
	default:
		switch tag {
		case "script":
			typ = "js"
		case "a":
			typ = "file" // a download file type likely, if we haven't already recognized it as something else.
		default:
			panic ("Unknown file type")
		}
	}

	url := "/" + typ + "/" + fileName

	url = RegisterAssetFile(url, filePath)

	switch tag {
	case script
	}
}
*/


func GetAssetLocation(url string) string {
	// If we have an AssetDirectory, either we are in release mode, or we are locally testing the release process
	if config.AssetDirectory != "" {
		if !util.StartsWith(url, config.AssetPrefix) {
			panic("Assets must start with the asset prefix.")
		}
		fPath := strings.TrimPrefix(url, config.AssetPrefix)
		return filepath.Join(config.AssetDirectory, filepath.FromSlash(fPath))
	}
	for dirUrl,dir := range assetDirectories {
		if util.StartsWith(url, dirUrl) {
			fPath := strings.TrimPrefix(url, dirUrl)
			return filepath.Join(dir, filepath.FromSlash(fPath))
		}
	}
	return ""
}

func GetAssetUrl(location string) string {
	if config.Release {
		if !util.StartsWith(location, config.AssetPrefix) {
			panic("In the release build, asset locations should be the same as asset paths.")
		}
		return location
	}
	for url,dir := range assetDirectories {
		if util.StartsWith(location, dir) {
			fPath := strings.TrimPrefix(location, dir)
			return path.Join(url, filepath.ToSlash(fPath))
		}
	}
	return ""
}


func ServeAsset(w http.ResponseWriter, r *http.Request) {
	localpath := GetAssetLocation(r.URL.Path)
	if localpath == "" {
		log.Printf("Invalid asset %s", r.URL.Path)
		return
	}
	//log.Printf("Served %s", localpath)

	if !config.Release {
		// TODO: Set up per file cache control

		// During development, tell the browser not to cache our assets so that if we change an asset, we don't have to deal with
		// trying to get the browser to refresh
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	} else {
		// TODO: Set up a validating cache control
	}

	http.ServeFile(w, r, localpath)
}


// RegisterAssetDirectory registers the given directory as a static file server. The files are served by the normal
// go static file process. This must happen during application initialization, as the static file directories
// are added to the MUX at startup time.
func RegisterAssetDirectory(dir string, pattern string) {
	if _,ok := assetDirectories[pattern]; ok {
		panic(pattern + " is already registered as an asset directory.")
	}
	assetDirectories[pattern] = dir
}

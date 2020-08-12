// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/strfmt"
)

// GetPresignedForClusterFilesURL generates an URL for the get presigned for cluster files operation
type GetPresignedForClusterFilesURL struct {
	ClusterID strfmt.UUID

	FileName string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetPresignedForClusterFilesURL) WithBasePath(bp string) *GetPresignedForClusterFilesURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetPresignedForClusterFilesURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetPresignedForClusterFilesURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/clusters/{cluster_id}/downloads/files-presigned"

	clusterID := o.ClusterID.String()
	if clusterID != "" {
		_path = strings.Replace(_path, "{cluster_id}", clusterID, -1)
	} else {
		return nil, errors.New("clusterId is required on GetPresignedForClusterFilesURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/assisted-install/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	fileNameQ := o.FileName
	if fileNameQ != "" {
		qs.Set("file_name", fileNameQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetPresignedForClusterFilesURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetPresignedForClusterFilesURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetPresignedForClusterFilesURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetPresignedForClusterFilesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetPresignedForClusterFilesURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetPresignedForClusterFilesURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}

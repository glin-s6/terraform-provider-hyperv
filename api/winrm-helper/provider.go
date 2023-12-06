package winrm_helper

import (
	"context"
	"text/template"
)

type Client interface {
	RunFireAndForgetScript(ctx context.Context, script *template.Template, args interface{}) error
	RunScriptWithResult(ctx context.Context, script *template.Template, args interface{}, result interface{}) (err error)
	UploadFile(ctx context.Context, filePath string) (remoteRootPath string, err error)
	UploadDirectory(ctx context.Context, rootPath string, excludeList []string) (remoteRootPath string, remoteAbsolutePaths []string, err error)
	DeleteFileOrDirectory(ctx context.Context, filePath string) (err error)
}

type Provider struct {
	Client Client
}

// frontend 包用于嵌入前端静态文件。
// 必须放入 dist 文件夹，否则无法访问 index.html
package frontend

import "embed"

//go:embed *
var FS embed.FS

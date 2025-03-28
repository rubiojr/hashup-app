// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.856
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func NotConfigured() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>HashUp Search</title><script src=\"https://unpkg.com/htmx.org@1.9.6\"></script><style>\n\t\t\tbody {\n\t\t\t\tfont-family: system-ui, -apple-system, sans-serif;\n\t\t\t\tmax-width: 600px;\n\t\t\t\tmargin: 0 auto;\n\t\t\t\tpadding: 2rem;\n\t\t\t\tbackground-color: #0f0f0f;\n\t\t\t\tcolor: #d96f28;\n\t\t\t}\n\t\t\t.search-title {\n\t\t\t    color: #d96f28;\n\t\t\t\tfont-weight: bold;\n\t\t\t\tfont-size: 2.5rem;\n\t\t\t}\n\t\t\t.search-container {\n\t\t\t\tmargin: 1rem 0;\n\t\t\t\ttext-align: center;\n\t\t\t}\n\t\t\t.search-input {\n\t\t\t\twidth: 100%;\n\t\t\t\t/* max-width: 600px; */\n\t\t\t\tpadding: 12px;\n\t\t\t\tfont-size: 16px;\n\t\t\t\tborder: 1px solid #ddd;\n\t\t\t\tborder-radius: 24px;\n\t\t\t\toutline: none;\n\t\t\t}\n\t\t\t.search-input:focus {\n\t\t\t\tbox-shadow: 0 0 0 2px rgba(0, 0, 255, 0.2);\n\t\t\t}\n\t\t\t.results-container {\n\t\t\t\tmargin-top: 1px;\n\t\t\t\t/* max-width: 500px; */\n\t\t\t}\n\t\t\t.missconfigured {\n\t\t\t\tfont-weight: bold;\n\t\t\t}\n\t\t\t.missconfigured a{\n\t\t\t\tcolor: #d96f28;\n\t\t\t}\n\t\t</style></head><body><div class=\"search-container\"><h1 class=\"search-title\">HashUp Search</h1><input type=\"text\" name=\"q\" class=\"search-input\" placeholder=\"Search for files...\"><div class=\"missconfigured\"><h2>HashUp needs to be configured first.</h2><p>Read the quickstart guide at:</p><a href=\"https://github.com/rubiojr/hashub/tree/main/docs/quickstart.md\" target=\"_blank\">https://github.com/rubiojr/hashub/tree/main/docs/quickstart.md</a></div></div><div id=\"search-results\" class=\"results-container\"></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate

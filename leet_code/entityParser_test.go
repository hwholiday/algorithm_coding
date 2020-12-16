package leet_code

import (
	"strings"
	"testing"
)

//1410. HTML 实体解析器
func entityParser(text string) string {
	return strings.NewReplacer("&quot;", `"`, "&apos;", `'`, "&amp;", `&`, "&gt;", `>`, "&lt;", `<`, "&frasl;", `/`).Replace(text)
}
func Test_entityParser(t *testing.T) {
	t.Log(entityParser("x &gt; y &amp;&amp; x &lt; y is always false"))
}

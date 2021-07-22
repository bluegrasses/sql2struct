package sql2struct

import (
	"fmt"
	"os"
	"text/template"

	"github.com/bluegrasses/sql2struct/internal/tools"
	"github.com/bluegrasses/sql2struct/internal/word"
)

// type 大写驼峰的表名称 struct {
// 	//注释
// 	字段名 字段类型
// 	//注释
// 	字段名 字段类型
// }
const strcutTpl = `package modles

type {{.TableName | ToCamelCase}} struct {
	{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
		{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
	{{end}}}
	func (model {{.TableName | ToCamelCase}}) TableName() string {
		return "{{.TableName}}"
	}`

// 定义模型文件的路径
const dir = "./internal/models/"

type StructTemplate struct {
	strcutTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{strcutTpl: strcutTpl}
}

// 将tbColumns转换为StructColum 结构体
func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

//将数据解析到模板中
func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUppercamelCase,
	}).Parse(t.strcutTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}

	// err := tpl.Execute(os.Stdout, tplDB)
	exist, err := tools.PathExists(dir)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return err
	}
	if !exist {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
	fname := word.UnderscoreToUppercamelCase(tplDB.TableName)
	filename := dir + fname + ".go"

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = tpl.Execute(f, tplDB)
	if err != nil {
		return err
	}
	return nil
}

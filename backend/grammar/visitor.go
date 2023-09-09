package grammar

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"proyecto2/parser"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// Funcion Principal del objeto Visitor
func (v *Visitor) Visit(tree antlr.ParseTree) Value {
	switch val := tree.(type) {
	case *parser.InitContext:
		return v.VisitInit(val)
	case *parser.BlockContext:
		return v.VisitBlock(val)
	case *parser.StmtContext:
		return v.VisitStmt(val)
	case *parser.IntExprContext:
		return v.VisitIntExpr(val)
	case *parser.DoubleExprContext:
		return v.VisitDoubleExpr(val)
	case *parser.NilExprContext:
		return v.VisitNilExpr(val)
	case *parser.IdExprContext:
		return v.VisitIdExpr(val)
	case *parser.StrExprContext:
		return v.VisitStrExpr(val)
	case *parser.CharExprContext:
		return v.VisitCharExpr(val)
	case *parser.BoolExprContext:
		return v.VisitBoolExpr(val)
	case *parser.NotExprContext:
		return v.VisitNotExpr(val)
	case *parser.UnaryExprContext:
		return v.VisitUnaryExpr(val)
	case *parser.ParExprContext:
		return v.VisitParExpr(val)
	case *parser.OpExprContext:
		return v.VisitOpExpr(val)
	case *parser.ExprWithParamContext:
		return v.VisitExprWithParam(val)
	case *parser.ExprWithParamsContext:
		return v.VisitExprWithParams(val)
	case *parser.PrimitiveTypeContext:
		return v.VisitPrimitiveType(val)
	case *parser.DeclstmtWithTypeAndExprContext:
		return v.VisitDeclstmtWithTypeAndExpr(val)
	case *parser.DeclstmtWithExprContext:
		return v.VisitDeclstmtWithExpr(val)
	case *parser.DeclstmtWithTypeContext:
		return v.VisitDeclstmtWithType(val)
	case *parser.AsignstmtContext:
		return v.VisitAsignstmt(val)
	case *parser.IncstmtContext:
		return v.VisitIncstmt(val)
	case *parser.DecstmtContext:
		return v.VisitDecstmt(val)
	case *parser.IfSimpleContext:
		return v.VisitIfSimple(val)
	case *parser.IfWithElseContext:
		return v.VisitIfWithElse(val)
	case *parser.IfWithElseIfContext:
		return v.VisitIfWithElseIf(val)
	case *parser.SwitchstmtContext:
		return v.VisitSwitchstmt(val)
	case *parser.WhilestmtContext:
		return v.VisitWhilestmt(val)
	case *parser.PrintlnstmtContext:
		return v.VisitPrintlnstmt(val)
	case *parser.ForWithRangeContext:
		return v.VisitForWithRange(val)
	case *parser.GuardstmtContext:
		return v.VisitGuardstmt(val)
	case *parser.BreakstmtContext:
		return v.VisitBreakstmt(val)
	case *parser.ContinuestmtContext:
		return v.VisitContinuestmt(val)
	case *parser.ReturnstmtContext:
		return v.VisitReturnstmt(val)
	case *parser.IntstmtContext:
		return v.VisitIntstmt(val)
	case *parser.IntConvExprContext:
		return v.VisitIntConvExpr(val)
	case *parser.FloatstmtContext:
		return v.VisitFloatstmt(val)
	case *parser.FloatConvExprContext:
		return v.VisitFloatConvExpr(val)
	case *parser.StringstmtContext:
		return v.VisitStringstmt(val)
	case *parser.StringConvExprContext:
		return v.VisitStringConvExpr(val)
	case *parser.ParametroContext:
		return v.VisitParametro(val)
	case *parser.FuncParametersContext:
		return v.VisitFuncParameters(val)
	case *parser.FuncParameterContext:
		return v.VisitFuncParameter(val)
	case *parser.FuncstmtContext:
		return v.VisitFuncstmt(val)
	case *parser.CallstmtContext:
		return v.VisitCallstmt(val)
	case *parser.CallExprContext:
		return v.VisitCallExpr(val)
	}
	return Value{value: "la funcion todavia no se ha implementado"}
}

func (v *Visitor) VisitInit(ctx *parser.InitContext) Value {
	return v.Visit(ctx.Block())
}

func CreateCST(tree antlr.ParseTree) string {
	var graph strings.Builder
	nodeCount := 0
	graph.WriteString("digraph {\n")
	generateBody(tree, &graph, &nodeCount)
	graph.WriteString("}\n")
	return graph.String()
}
func generateBody(node antlr.ParseTree, graph *strings.Builder, nodeCount *int) int {
	currentNode := *nodeCount
	*nodeCount++
	nodeName := fmt.Sprintf("Node%d", currentNode)
	nodeLabel := convertNodeToStr(node)
	if nodeLabel != "nil" {
		graph.WriteString(fmt.Sprintf("  %s [label=\"%s\"];\n", nodeName, nodeLabel))
		for i := 0; i < node.GetChildCount(); i++ {
			child := node.GetChild(i).(antlr.ParseTree)
			childNode := generateBody(child, graph, nodeCount)
			if childNode != -1 {
				graph.WriteString(fmt.Sprintf("  %s -> Node%d;\n", nodeName, childNode))
			}
		}
	}

	return currentNode
}

func convertNodeToStr(node antlr.ParseTree) string {
	switch node.(type) {
	case antlr.TerminalNode:
		text := node.GetText()
		if strings.Contains(text, "\"") {
			newText := strings.ReplaceAll(text, "\"", "\\\"")
			newText = strings.ReplaceAll(newText, "\n", "\\n")
			newText = strings.ReplaceAll(newText, "\t", "\\t")
			return newText
		}
		_, err := strconv.ParseFloat(text, 64)
		if err == nil {
			return text
		}
		return text
	default:
		return fmt.Sprintf("%T", node)
	}
}

func GenerateHTMLERROR(v *Visitor) {
	var graph strings.Builder
	graph.WriteString("<!DOCTYPE html>\n")
	graph.WriteString("<html lang=\"en\">\n")
	graph.WriteString("<head>\n")
	graph.WriteString("<meta charset=\"UTF-8\">\n")
	graph.WriteString("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n")
	graph.WriteString("<title>Tabla de Errores</title>\n")
	graph.WriteString("<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css\">\n")
	graph.WriteString("</head>\n")
	graph.WriteString("<body>\n")
	graph.WriteString("<div class=\"container mt-4\">\n")
	graph.WriteString("<h1>Tabla de Errores</h1>\n")
	graph.WriteString("<table class=\"table table-striped\">\n")
	graph.WriteString("<thead><tr>\n")
	graph.WriteString("<th>descripcion</th><th>linea</th><th>columna</th>\n")
	graph.WriteString("</tr></head>\n")
	graph.WriteString("<tbody>\n")
	for _, element := range v.Errores {
		graph.WriteString("<tr>")
		graph.WriteString("<th>")
		graph.WriteString(element.desc)
		graph.WriteString("</th>")
		graph.WriteString("<th>")
		graph.WriteString(fmt.Sprintf("%v", element.line))
		graph.WriteString("</th>")
		graph.WriteString("<th>")
		graph.WriteString(fmt.Sprintf("%v", element.column))
		graph.WriteString("</th></tr>\n")
	}
	graph.WriteString("</tbody>\n</table>\n</div>\n")
	graph.WriteString("<script src=\"https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js\"></script>\n")
	graph.WriteString("<script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.16.0/umd/popper.min.js\"></script>\n")
	graph.WriteString("<script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.5.2/js/bootstrap.min.js\"></script>\n")
	graph.WriteString("</body>\n")
	graph.WriteString("</html>\n")
	err := os.WriteFile("tablaDeErrores.html", []byte(graph.String()), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateHTML(v *Visitor) {
	for v.Environment.padre != nil {
		v.Environment = v.Environment.padre
	}
	var graph strings.Builder
	graph.WriteString("<!DOCTYPE html>\n")
	graph.WriteString("<html lang=\"en\">\n")
	graph.WriteString("<head>\n")
	graph.WriteString("<meta charset=\"UTF-8\">\n")
	graph.WriteString("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n")
	graph.WriteString("<title>Tabla de Simbolos</title>\n")
	graph.WriteString("<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css\">\n")
	graph.WriteString("</head>\n")
	graph.WriteString("<body>\n")
	graph.WriteString("<div class=\"container mt-4\">\n")
	graph.WriteString("<h1>Tabla de Simbolos</h1>\n")
	graph.WriteString("<table class=\"table table-striped\">\n")
	graph.WriteString("<thead><tr>\n")
	graph.WriteString("<th>Id</th><th>Tipo de dato</th><th>Editable</th><th>Valor</th>\n")
	graph.WriteString("</tr></head>\n")
	graph.WriteString("<tbody>\n")
	graph.WriteString(v.Environment.generateSymbolTable())
	graph.WriteString("</tbody>\n</table>\n</div>\n")
	graph.WriteString("<script src=\"https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js\"></script>\n")
	graph.WriteString("<script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.16.0/umd/popper.min.js\"></script>\n")
	graph.WriteString("<script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.5.2/js/bootstrap.min.js\"></script>\n")
	graph.WriteString("</body>\n")
	graph.WriteString("</html>\n")
	err := os.WriteFile("tablaDeSimbolos.html", []byte(graph.String()), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func Execute() {
	input := `
	`
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewGrammarLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGrammarParser(tokenStream)
	p.BuildParseTrees = true
	tree := p.Init()
	eval := Visitor{
		Environment: NewEnvironment(nil),
	}
	eval.Visit(tree)
	cst := CreateCST(tree)
	err := os.WriteFile("cstree.dot", []byte(cst), 0644)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("dot", "-Tpdf", "cstree.dot", "-o", "cstree.pdf")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error al convertir el archivo .dot a .pdf", err.Error())
		return
	}
	GenerateHTML(&eval)
	GenerateHTMLERROR(&eval)
	fmt.Println(eval.Errores)
}

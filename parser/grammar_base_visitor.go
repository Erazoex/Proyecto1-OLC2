// Code generated from Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr4-go/antlr/v4"

type BaseGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGrammarVisitor) VisitInit(ctx *InitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitDeclstmt(ctx *DeclstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitPrintstmt(ctx *PrintstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitIfstmt(ctx *IfstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitForstmt(ctx *ForstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitDoubleExpr(ctx *DoubleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitParExpr(ctx *ParExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStrExpr(ctx *StrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitIntExpr(ctx *IntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitNilExpr(ctx *NilExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitOpExpr(ctx *OpExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}
